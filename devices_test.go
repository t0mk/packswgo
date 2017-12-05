package packswgo

import (
	"fmt"
	"testing"
	"time"
)

func waitDeviceActive(id string, c *Client) (*Device, error) {
	// 15 minutes = 180 * 5sec-retry
	for i := 0; i < 180; i++ {
		<-time.After(5 * time.Second)
		d, resp, err := c.FindDeviceById(id, "facility")
		if err != nil {
			return nil, err
		}
		err = CheckForAppError(resp)
		if err != nil {
			return nil, err
		}
		if d.State == "active" {
			return d, nil
		}
	}
	return nil, fmt.Errorf("device %s is still not active after timeout", id)
}

func deleteDevice(t *testing.T, c *Client, id string) {
	_, err := c.DeleteDevice(id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAccDeviceBasic(t *testing.T) {
	skipUnlessAcceptanceTestsAllowed(t)
	t.Parallel()

	c, projectId, teardown := setupWithProject(t)
	defer teardown()

	hn := randString8()

	cr := DeviceCreateInput{
		Hostname:        hn,
		Facility:        testFacility(),
		Plan:            "baremetal_0",
		OperatingSystem: "ubuntu_16_04",
		BillingCycle:    "hourly",
	}

	d, resp, err := c.CreateDevice(projectId, cr)
	if err != nil {
		t.Fatal(err)
	}
	err = CheckForAppError(resp)
	if err != nil {
		t.Fatal(err)
	}

	defer deleteDevice(t, c, d.Id)

	dId := d.Id

	d, err = waitDeviceActive(dId, c)
	if err != nil {
		t.Fatal(err)
	}

	if len(d.RootPassword) == 0 {
		t.Fatal("root_password is empty or non-existent")
	}

	newHN := randString8()
	ur := DeviceUpdateInput{Hostname: newHN}

	newD, _, err := c.UpdateDevice(dId, ur)
	if err != nil {
		t.Fatal(err)
	}

	if newD.Hostname != newHN {
		t.Fatalf("hostname of test device should be %s, but is %s", newHN, newD.Hostname)
	}
	for _, ipa := range newD.IpAddresses {
		if !ipa.Management {
			t.Fatalf("management flag for all the IP addresses in a new device should be True: was %s", ipa)
		}
	}
}

/*

func TestAccDevicePXE(t *testing.T) {
	skipUnlessAcceptanceTestsAllowed(t)
	t.Parallel()

	c, projectId, teardown := setupWithProject(t)
	defer teardown()
	hn := randString8()
	pxeURL := "https://boot.netboot.xyz"

	cr := DeviceCreateRequest{
		Hostname:      "pxe-" + hn,
		Facility:      testFacility(),
		Plan:          "baremetal_0",
		ProjectId:     projectId,
		BillingCycle:  "hourly",
		OS:            "custom_ipxe",
		IPXEScriptURL: "https://boot.netboot.xyz",
		AlwaysPXE:     true,
	}

	d, _, err := c.Devices.Create(&cr)
	if err != nil {
		t.Fatal(err)
	}

	defer deleteDevice(t, c, d.Id)

	d, err = waitDeviceActive(d.Id, c)
	if err != nil {
		t.Fatal(err)
	}

	if !d.AlwaysPXE {
		t.Fatal("always_pxe should be set")
	}
	if d.IPXEScriptURL != pxeURL {
		t.Fatalf("ipxe_script_url should be %s", pxeURL)
	}
}

func TestAccDeviceAssignIP(t *testing.T) {
	skipUnlessAcceptanceTestsAllowed(t)
	t.Parallel()

	c, projectId, teardown := setupWithProject(t)
	defer teardown()
	hn := randString8()

	testFac := testFacility()

	cr := DeviceCreateRequest{
		Hostname:     hn,
		Facility:     testFac,
		Plan:         "baremetal_0",
		ProjectId:    projectId,
		BillingCycle: "hourly",
		OS:           "ubuntu_16_04",
	}

	d, _, err := c.Devices.Create(&cr)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteDevice(t, c, d.Id)

	d, err = waitDeviceActive(d.Id, c)
	if err != nil {
		t.Fatal(err)
	}

	req := IPReservationRequest{
		Type:     "public_ipv4",
		Quantity: 1,
		Comments: "packngo test",
		Facility: testFac,
	}

	reservation, _, err := c.ProjectIPs.Request(projectId, &req)
	if err != nil {
		t.Fatal(err)
	}

	af := AddressStruct{Address: fmt.Sprintf("%s/%d", reservation.Address, reservation.CIdR)}

	assignment, _, err := c.DeviceIPs.Assign(d.Id, &af)
	if err != nil {
		t.Fatal(err)
	}

	if assignment.Management {
		t.Error("Management flag for assignment resource must be False")
	}

	d, _, err = c.Devices.Get(d.Id)
	if err != nil {
		t.Fatal(err)
	}

	// If the Quantity in the IPReservationRequest is >1, this test won't work.
	// The assignment CIdR would then have to be extracted from the reserved
	// block.
	reservation, _, err = c.ProjectIPs.Get(reservation.Id)
	if err != nil {
		t.Fatal(err)
	}

	if len(reservation.Assignments) != 1 {
		t.Fatalf("reservation %s should have exactly 1 assignment", reservation)
	}

	if reservation.Assignments[0].Href != assignment.Href {
		t.Fatalf("assignment %s should be listed in reservation resource %s",
			assignment.Href, reservation)

	}

	func() {
		for _, ipa := range d.Network {
			if ipa.Href == assignment.Href {
				return
			}
		}
		t.Fatalf("assignment %s should be listed in device %s", assignment, d)
	}()

	if assignment.AssignedTo.Href != d.Href {
		t.Fatalf("device %s should be listed in assignment %s",
			d, assignment)
	}

	_, err = c.DeviceIPs.Unassign(assignment.Id)
	if err != nil {
		t.Fatal(err)
	}

	// reload reservation, now without any assignment
	reservation, _, err = c.ProjectIPs.Get(reservation.Id)
	if err != nil {
		t.Fatal(err)
	}

	if len(reservation.Assignments) != 0 {
		t.Fatalf("reservation %s shoud be without assignments. Was %v",
			reservation, reservation.Assignments)
	}

	// reload device, now without the assigned floating IP
	d, _, err = c.Devices.Get(d.Id)
	if err != nil {
		t.Fatal(err)
	}

	for _, ipa := range d.Network {
		if ipa.Href == assignment.Href {
			t.Fatalf("assignment %s shoud be not listed in device %s anymore",
				assignment, d)
		}
	}
}

func TestAccDeviceAttachVolume(t *testing.T) {
	skipUnlessAcceptanceTestsAllowed(t)
	t.Parallel()

	c, projectId, teardown := setupWithProject(t)
	defer teardown()
	hn := randString8()

	cr := DeviceCreateRequest{
		Hostname:     hn,
		Facility:     testFacility(),
		Plan:         "baremetal_0",
		ProjectId:    projectId,
		BillingCycle: "hourly",
		OS:           "ubuntu_16_04",
	}

	d, _, err := c.Devices.Create(&cr)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteDevice(t, c, d.Id)

	d, err = waitDeviceActive(d.Id, c)
	if err != nil {
		t.Fatal(err)
	}

	vcr := VolumeCreateRequest{
		Size:         10,
		BillingCycle: "hourly",
		PlanId:       "storage_1",
		FacilityId:   testFacility(),
	}

	v, _, err := c.Volumes.Create(&vcr, projectId)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Volumes.Delete(v.Id)

	v, err = waitVolumeActive(v.Id, c)
	if err != nil {
		t.Fatal(err)
	}

	a, _, err := c.VolumeAttachments.Create(v.Id, d.Id)
	if err != nil {
		t.Fatal(err)
	}

	if path.Base(a.Volume.Href) != v.Id {
		t.Fatalf("wrong volume href in the attachment: %s, should be %s", a.Volume.Href, v.Id)
	}

	if path.Base(a.Device.Href) != d.Id {
		t.Fatalf("wrong device href in the attachment: %s, should be %s", a.Device.Href, d.Id)
	}

	v, _, err = c.Volumes.Get(v.Id)
	if err != nil {
		t.Fatal(err)
	}

	d, _, err = c.Devices.Get(d.Id)
	if err != nil {
		t.Fatal(err)
	}

	if v.Attachments[0].Device.Id != d.Id {
		t.Fatalf("wrong device linked in volume attachment: %s, should be %s", v.Attachments[0].Device.Id, d.Id)
	}
	if path.Base(d.Volumes[0].Href) != v.Id {
		t.Fatalf("wrong volume linked in device.volumes: %s, should be %s", d.Volumes[0].Href, v.Id)
	}

	_, err = c.VolumeAttachments.Delete(a.Id)
	if err != nil {
		t.Fatal(err)
	}

}

func TestAccDeviceSpotInstance(t *testing.T) {
	skipUnlessAcceptanceTestsAllowed(t)
	t.Parallel()

	c, projectId, teardown := setupWithProject(t)
	defer teardown()
	hn := randString8()

	testFac := "nrt1"
	testSPM := 0.04
	testTerm := &Timestamp{Time: time.Now().Add(time.Hour - (time.Minute * 10))}

	cr := DeviceCreateRequest{
		Hostname:        hn,
		Facility:        testFac,
		Plan:            "baremetal_0",
		OS:              "coreos_stable",
		ProjectId:       projectId,
		BillingCycle:    "hourly",
		SpotInstance:    true,
		SpotPriceMax:    testSPM,
		TerminationTime: testTerm,
	}

	d, _, err := c.Devices.Create(&cr)
	if err != nil {
		t.Fatal(err)
	}
	defer deleteDevice(t, c, d.Id)

	d, err = waitDeviceActive(d.Id, c)
	if err != nil {
		t.Fatal(err)
	}

	if !d.SpotInstance {
		t.Fatal("spot_instance is false, should be true")
	}

	if d.SpotPriceMax != testSPM {
		t.Fatalf("spot_price_max is %f, should be %f", d.SpotPriceMax, testSPM)
	}

	if !d.TerminationTime.Time.Truncate(time.Minute).Equal(testTerm.Time.Truncate(time.Minute)) {
		t.Fatalf("termination_time is %s, should be %s",
			d.TerminationTime.Time.Local(), testTerm.Time.Local())
	}
}

*/
