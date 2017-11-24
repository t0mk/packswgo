package packswgo

import "os"

// APITokenVar is the name of teh envvar which holds auth token to Packet API
const APITokenVar = "PACKET_AUTH_TOKEN"

type Client struct {
	VirtualNetworks        VirtualNetworksApi
	Notifications          NotificationsApi
	Licenses               LicensesApi
	InstancesBatch         InstancesBatchApi
	InternetGateways       InternetGatewaysApi
	Memberships            MembershipsApi
	Events                 EventsApi
	Emails                 EmailsApi
	InstanceBandwidth      InstanceBandwidthApi
	Facilities             FacilitiesApi
	Users                  UsersApi
	Plans                  PlansApi
	VPN                    VPNApi
	SpotMarketPrices       SpotMarketPricesApi
	Traffic                TrafficApi
	Volumes                VolumesApi
	SSHKeys                SSHKeysApi
	BgpConfigs             BgpConfigsApi
	TransferRequests       TransferRequestsApi
	VolumeSnapshotPolicies VolumeSnapshotPoliciesApi
	HardwareReservations   HardwareReservationsApi
	Projects               ProjectsApi
	Invitations            InvitationsApi
	Batches                BatchesApi
	VolumeSnapshots        VolumeSnapshotsApi
	VolumeAttachments      VolumeAttachmentsApi
	BgpSessions            BgpSessionsApi
	IPAddresses            IPAddressesApi
	Capacity               CapacityApi
	OperatingSystems       OperatingSystemsApi
	Devices                DevicesApi
	Actions                ActionsApi
}

func getCfg() *Configuration {
	cfg := NewConfiguration()
	cfg.APIKey["X-Auth-Token"] = os.Getenv(APITokenVar)
	return cfg
}

// NewClient returns API client
func NewClient() Client {
	cfg := getCfg()
	return Client{
		VirtualNetworks:        VirtualNetworksApi{Configuration: cfg},
		Notifications:          NotificationsApi{Configuration: cfg},
		Licenses:               LicensesApi{Configuration: cfg},
		InstancesBatch:         InstancesBatchApi{Configuration: cfg},
		InternetGateways:       InternetGatewaysApi{Configuration: cfg},
		Memberships:            MembershipsApi{Configuration: cfg},
		Events:                 EventsApi{Configuration: cfg},
		Emails:                 EmailsApi{Configuration: cfg},
		InstanceBandwidth:      InstanceBandwidthApi{Configuration: cfg},
		Facilities:             FacilitiesApi{Configuration: cfg},
		Users:                  UsersApi{Configuration: cfg},
		Plans:                  PlansApi{Configuration: cfg},
		VPN:                    VPNApi{Configuration: cfg},
		SpotMarketPrices:       SpotMarketPricesApi{Configuration: cfg},
		Traffic:                TrafficApi{Configuration: cfg},
		Volumes:                VolumesApi{Configuration: cfg},
		SSHKeys:                SSHKeysApi{Configuration: cfg},
		BgpConfigs:             BgpConfigsApi{Configuration: cfg},
		TransferRequests:       TransferRequestsApi{Configuration: cfg},
		VolumeSnapshotPolicies: VolumeSnapshotPoliciesApi{Configuration: cfg},
		HardwareReservations:   HardwareReservationsApi{Configuration: cfg},
		Projects:               ProjectsApi{Configuration: cfg},
		Invitations:            InvitationsApi{Configuration: cfg},
		Batches:                BatchesApi{Configuration: cfg},
		VolumeSnapshots:        VolumeSnapshotsApi{Configuration: cfg},
		VolumeAttachments:      VolumeAttachmentsApi{Configuration: cfg},
		BgpSessions:            BgpSessionsApi{Configuration: cfg},
		IPAddresses:            IPAddressesApi{Configuration: cfg},
		Capacity:               CapacityApi{Configuration: cfg},
		OperatingSystems:       OperatingSystemsApi{Configuration: cfg},
		Devices:                DevicesApi{Configuration: cfg},
		Actions:                ActionsApi{Configuration: cfg},
	}
}

func (c *Client) AssignPort(id string, vnid string) (*Port, *APIResponse, error) {
	return c.VirtualNetworks.AssignPort(id, vnid)
}

func (c *Client) BondPort(id string, bulkEnable bool) (*Port, *APIResponse, error) {
	return c.VirtualNetworks.BondPort(id, bulkEnable)
}

func (c *Client) CreateVirtualNetwork(id string, virtualNetwork VirtualNetworkCreateInput) (*VirtualNetwork, *APIResponse, error) {
	return c.VirtualNetworks.CreateVirtualNetwork(id, virtualNetwork)
}

func (c *Client) DeleteVirtualNetwork(id string) (*VirtualNetwork, *APIResponse, error) {
	return c.VirtualNetworks.DeleteVirtualNetwork(id)
}

func (c *Client) DisbondPort(id string, bulkDisable bool) (*Port, *APIResponse, error) {
	return c.VirtualNetworks.DisbondPort(id, bulkDisable)
}

func (c *Client) FindVirtualNetworks(id string, include string) (*VirtualNetworkList, *APIResponse, error) {
	return c.VirtualNetworks.FindVirtualNetworks(id, include)
}

func (c *Client) UnassignPort(id string, vnid string) (*Port, *APIResponse, error) {
	return c.VirtualNetworks.UnassignPort(id, vnid)
}

func (c *Client) FindNotificationById(id string, include string) (*Notification, *APIResponse, error) {
	return c.Notifications.FindNotificationById(id, include)
}

func (c *Client) FindNotifications(all bool, since string, include string, page int32, perPage int32) (*NotificationList, *APIResponse, error) {
	return c.Notifications.FindNotifications(all, since, include, page, perPage)
}

func (c *Client) UpdateNotification(id string) (*Notification, *APIResponse, error) {
	return c.Notifications.UpdateNotification(id)
}

func (c *Client) CreateLicense(id string, license LicenseCreateInput) (*License, *APIResponse, error) {
	return c.Licenses.CreateLicense(id, license)
}

func (c *Client) DeleteLicense(id string) (*APIResponse, error) {
	return c.Licenses.DeleteLicense(id)
}

func (c *Client) FindLicenseById(id string, include string) (*License, *APIResponse, error) {
	return c.Licenses.FindLicenseById(id, include)
}

func (c *Client) FindLicenses(id string, include string) (*LicenseList, *APIResponse, error) {
	return c.Licenses.FindLicenses(id, include)
}

func (c *Client) UpdateLicense(id string, license LicenseUpdateInput) (*License, *APIResponse, error) {
	return c.Licenses.UpdateLicense(id, license)
}

func (c *Client) CreateDeviceBatch(id string, batch InstancesBatchCreateInput) (*Batch, *APIResponse, error) {
	return c.InstancesBatch.CreateDeviceBatch(id, batch)
}

func (c *Client) CreateInternetGateway(id string, length string) (*InternetGateway, *APIResponse, error) {
	return c.InternetGateways.CreateInternetGateway(id, length)
}

func (c *Client) DeleteMembership(id string) (*APIResponse, error) {
	return c.Memberships.DeleteMembership(id)
}

func (c *Client) FindMembershipById(id string, include string) (*Membership, *APIResponse, error) {
	return c.Memberships.FindMembershipById(id, include)
}

func (c *Client) UpdateMembership(id string, membership MembershipInput) (*Membership, *APIResponse, error) {
	return c.Memberships.UpdateMembership(id, membership)
}

func (c *Client) FindDeviceEvents(id string, include string, page int32, perPage int32) (*EventList, *APIResponse, error) {
	return c.Events.FindDeviceEvents(id, include, page, perPage)
}

func (c *Client) FindEventById(id string, include string) (*Event, *APIResponse, error) {
	return c.Events.FindEventById(id, include)
}

func (c *Client) FindEvents(include string, page int32, perPage int32) (*EventList, *APIResponse, error) {
	return c.Events.FindEvents(include, page, perPage)
}

func (c *Client) FindProjectEvents(id string, include string, page int32, perPage int32) (*EventList, *APIResponse, error) {
	return c.Events.FindProjectEvents(id, include, page, perPage)
}

func (c *Client) FindVolumeEvents(id string, include string, page int32, perPage int32) (*EventList, *APIResponse, error) {
	return c.Events.FindVolumeEvents(id, include, page, perPage)
}

func (c *Client) CreateEmail(email EmailInput) (*Email, *APIResponse, error) {
	return c.Emails.CreateEmail(email)
}

func (c *Client) DeleteEmail(id string) (*APIResponse, error) {
	return c.Emails.DeleteEmail(id)
}

func (c *Client) FindEmailById(id string, include string) (*Email, *APIResponse, error) {
	return c.Emails.FindEmailById(id, include)
}

func (c *Client) UpdateEmail(id string, email EmailInput) (*Email, *APIResponse, error) {
	return c.Emails.UpdateEmail(id, email)
}

func (c *Client) FindInstanceBandwidth(id string, from string, until string) (*APIResponse, error) {
	return c.InstanceBandwidth.FindInstanceBandwidth(id, from, until)
}

func (c *Client) FindFacilities() (*FacilityList, *APIResponse, error) {
	return c.Facilities.FindFacilities()
}

func (c *Client) FindFacilitiesByProject(id string, include string, page int32, perPage int32) (*FacilityList, *APIResponse, error) {
	return c.Facilities.FindFacilitiesByProject(id, include, page, perPage)
}

func (c *Client) FindCurrentUser(include string) (*User, *APIResponse, error) {
	return c.Users.FindCurrentUser(include)
}

func (c *Client) FindUserById(id string, include string) (*User, *APIResponse, error) {
	return c.Users.FindUserById(id, include)
}

func (c *Client) FindUsers(include string, page int32, perPage int32) (*UserList, *APIResponse, error) {
	return c.Users.FindUsers(include, page, perPage)
}

func (c *Client) UpdateCurrentUser(user UserUpdateInput) (*User, *APIResponse, error) {
	return c.Users.UpdateCurrentUser(user)
}

func (c *Client) FindPlans(include string) (*PlanList, *APIResponse, error) {
	return c.Plans.FindPlans(include)
}

func (c *Client) FindPlansByProject(id string, include string, page int32, perPage int32) (*PlanList, *APIResponse, error) {
	return c.Plans.FindPlansByProject(id, include, page, perPage)
}

func (c *Client) FindCurrentUserVpnConfig(code string) (*VpnConfig, *APIResponse, error) {
	return c.VPN.FindCurrentUserVpnConfig(code)
}

func (c *Client) TurnOffCurrentUserVpn() (*APIResponse, error) {
	return c.VPN.TurnOffCurrentUserVpn()
}

func (c *Client) TurnOnCurrentUserVpn() (*APIResponse, error) {
	return c.VPN.TurnOnCurrentUserVpn()
}

func (c *Client) FindSpotMarketPrices(facility string, plan string) (*APIResponse, error) {
	return c.SpotMarketPrices.FindSpotMarketPrices(facility, plan)
}

func (c *Client) FindSpotMarketPricesHistory(facility string, plan string) (*APIResponse, error) {
	return c.SpotMarketPrices.FindSpotMarketPricesHistory(facility, plan)
}

func (c *Client) FindTraffic(id string, direction string, timeframe Timeframe, interval string, bucket string) (*APIResponse, error) {
	return c.Traffic.FindTraffic(id, direction, timeframe, interval, bucket)
}

func (c *Client) CloneVolume(id string, snapshotTimestamp string) (*Volume, *APIResponse, error) {
	return c.Volumes.CloneVolume(id, snapshotTimestamp)
}

func (c *Client) CreateVolume(id string, volume VolumeCreateInput) (*Volume, *APIResponse, error) {
	return c.Volumes.CreateVolume(id, volume)
}

func (c *Client) DeleteVolume(id string) (*APIResponse, error) {
	return c.Volumes.DeleteVolume(id)
}

func (c *Client) FindVolumeById(id string, include string) (*Volume, *APIResponse, error) {
	return c.Volumes.FindVolumeById(id, include)
}

func (c *Client) FindVolumes(id string, include string, page int32, perPage int32) (*VolumeList, *APIResponse, error) {
	return c.Volumes.FindVolumes(id, include, page, perPage)
}

func (c *Client) RestoreVolume(id string, restorePoint string) (*Volume, *APIResponse, error) {
	return c.Volumes.RestoreVolume(id, restorePoint)
}

func (c *Client) UpdateVolume(id string, volume VolumeUpdateInput) (*Volume, *APIResponse, error) {
	return c.Volumes.UpdateVolume(id, volume)
}

func (c *Client) CreateSSHKey(sshKey SshKeyInput) (*SshKey, *APIResponse, error) {
	return c.SSHKeys.CreateSSHKey(sshKey)
}

func (c *Client) CreateSSHKey_1(id string, sshKey SshKeyInput) (*SshKey, *APIResponse, error) {
	return c.SSHKeys.CreateSSHKey_1(id, sshKey)
}

func (c *Client) DeleteSSHKey(id string) (*APIResponse, error) {
	return c.SSHKeys.DeleteSSHKey(id)
}

func (c *Client) FindSSHKeyById(id string, include string) (*SshKey, *APIResponse, error) {
	return c.SSHKeys.FindSSHKeyById(id, include)
}

func (c *Client) FindSSHKeys(include string) (*SshKeyList, *APIResponse, error) {
	return c.SSHKeys.FindSSHKeys(include)
}

func (c *Client) FindSSHKeys_2(id string, include string) (*SshKeyList, *APIResponse, error) {
	return c.SSHKeys.FindSSHKeys_2(id, include)
}

func (c *Client) UpdateSSHKey(id string, sshKey SshKeyInput) (*SshKey, *APIResponse, error) {
	return c.SSHKeys.UpdateSSHKey(id, sshKey)
}

func (c *Client) FindBgpConfigByProject(id string, include string) (*BgpConfig, *APIResponse, error) {
	return c.BgpConfigs.FindBgpConfigByProject(id, include)
}

func (c *Client) RequestBgpConfig(id string, bgpConfigRequest BgpConfigRequestInput) (*APIResponse, error) {
	return c.BgpConfigs.RequestBgpConfig(id, bgpConfigRequest)
}

func (c *Client) AcceptTransferRequest(id string) (*APIResponse, error) {
	return c.TransferRequests.AcceptTransferRequest(id)
}

func (c *Client) CreateTransferRequest(id string, transferRequest TransferRequestInput) (*TransferRequest, *APIResponse, error) {
	return c.TransferRequests.CreateTransferRequest(id, transferRequest)
}

func (c *Client) DeclineTransferRequest(id string) (*APIResponse, error) {
	return c.TransferRequests.DeclineTransferRequest(id)
}

func (c *Client) FindTransferRequestById(id string, include string) (*TransferRequest, *APIResponse, error) {
	return c.TransferRequests.FindTransferRequestById(id, include)
}

func (c *Client) CreateVolumeSnapshotPolicy(id string, snapshotFrequency string, snapshotCount int32) (*SnapshotPolicy, *APIResponse, error) {
	return c.VolumeSnapshotPolicies.CreateVolumeSnapshotPolicy(id, snapshotFrequency, snapshotCount)
}

func (c *Client) DeleteVolumeSnapshotPolicy(id string) (*APIResponse, error) {
	return c.VolumeSnapshotPolicies.DeleteVolumeSnapshotPolicy(id)
}

func (c *Client) UpdateVolumeSnapshotPolicy(id string, snapshotFrequency string, snapshotCount int32) (*SnapshotPolicy, *APIResponse, error) {
	return c.VolumeSnapshotPolicies.UpdateVolumeSnapshotPolicy(id, snapshotFrequency, snapshotCount)
}

func (c *Client) FindHardwareReservationById(id string, include string) (*Device, *APIResponse, error) {
	return c.HardwareReservations.FindHardwareReservationById(id, include)
}

func (c *Client) FindHardwareReservations(id string, include string, page int32, perPage int32) (*HardwareReservationList, *APIResponse, error) {
	return c.HardwareReservations.FindHardwareReservations(id, include, page, perPage)
}

func (c *Client) CreateProject(project ProjectCreateInput) (*Project, *APIResponse, error) {
	return c.Projects.CreateProject(project)
}

func (c *Client) DeleteProject(id string) (*APIResponse, error) {
	return c.Projects.DeleteProject(id)
}

func (c *Client) FindProjectById(id string, include string) (*Project, *APIResponse, error) {
	return c.Projects.FindProjectById(id, include)
}

func (c *Client) FindProjects(include string, page int32, perPage int32) (*ProjectList, *APIResponse, error) {
	return c.Projects.FindProjects(include, page, perPage)
}

func (c *Client) UpdateProject(id string, project ProjectUpdateInput) (*Project, *APIResponse, error) {
	return c.Projects.UpdateProject(id, project)
}

func (c *Client) AcceptInvitation(id string) (*Membership, *APIResponse, error) {
	return c.Invitations.AcceptInvitation(id)
}

func (c *Client) CreateInvitation(id string, invitation InvitationInput) (*Invitation, *APIResponse, error) {
	return c.Invitations.CreateInvitation(id, invitation)
}

func (c *Client) DeclineInvitation(id string) (*APIResponse, error) {
	return c.Invitations.DeclineInvitation(id)
}

func (c *Client) FindInvitationById(id string, include string) (*Invitation, *APIResponse, error) {
	return c.Invitations.FindInvitationById(id, include)
}

func (c *Client) FindBatchById(id string, include string) (*Batch, *APIResponse, error) {
	return c.Batches.FindBatchById(id, include)
}

func (c *Client) FindBatchesByProject(id string, include string) (*BatchesList, *APIResponse, error) {
	return c.Batches.FindBatchesByProject(id, include)
}

func (c *Client) CreateVolumeSnapshot(id string) (*APIResponse, error) {
	return c.VolumeSnapshots.CreateVolumeSnapshot(id)
}

func (c *Client) DeleteVolumeSnapshot(volumeId string, id string) (*APIResponse, error) {
	return c.VolumeSnapshots.DeleteVolumeSnapshot(volumeId, id)
}

func (c *Client) FindVolumeSnapshots(id string, include string) (*VolumeSnapshotList, *APIResponse, error) {
	return c.VolumeSnapshots.FindVolumeSnapshots(id, include)
}

func (c *Client) CreateVolumeAttachment(id string, attachment VolumeAttachmentInput) (*VolumeAttachment, *APIResponse, error) {
	return c.VolumeAttachments.CreateVolumeAttachment(id, attachment)
}

func (c *Client) DeleteVolumeAttachment(id string) (*APIResponse, error) {
	return c.VolumeAttachments.DeleteVolumeAttachment(id)
}

func (c *Client) FindVolumeAttachmentById(id string, include string) (*VolumeAttachment, *APIResponse, error) {
	return c.VolumeAttachments.FindVolumeAttachmentById(id, include)
}

func (c *Client) FindVolumeAttachments(id string, include string) (*VolumeAttachmentList, *APIResponse, error) {
	return c.VolumeAttachments.FindVolumeAttachments(id, include)
}

func (c *Client) CreateBgpSession(id string, addressFamily string) (*BgpSession, *APIResponse, error) {
	return c.BgpSessions.CreateBgpSession(id, addressFamily)
}

func (c *Client) DeleteBgpSession(id string) (*APIResponse, error) {
	return c.BgpSessions.DeleteBgpSession(id)
}

func (c *Client) FindBgpSessionById(id string, include string) (*BgpSession, *APIResponse, error) {
	return c.BgpSessions.FindBgpSessionById(id, include)
}

func (c *Client) FindBgpSessions(id string) (*BgpSessionList, *APIResponse, error) {
	return c.BgpSessions.FindBgpSessions(id)
}

func (c *Client) CreateIPAssignment(id string, ipAssignment IpAssignmentInput) (*IpAssignment, *APIResponse, error) {
	return c.IPAddresses.CreateIPAssignment(id, ipAssignment)
}

func (c *Client) DeleteIPAddress(id string) (*APIResponse, error) {
	return c.IPAddresses.DeleteIPAddress(id)
}

func (c *Client) FindIPAddressById(id string, include string) (*IpAssignment, *APIResponse, error) {
	return c.IPAddresses.FindIPAddressById(id, include)
}

func (c *Client) FindIPAssignments(id string, include string) (*IpAssignmentList, *APIResponse, error) {
	return c.IPAddresses.FindIPAssignments(id, include)
}

func (c *Client) FindIPAvailabilities(id string, cidr string) (*IpAvailabilitiesList, *APIResponse, error) {
	return c.IPAddresses.FindIPAvailabilities(id, cidr)
}

func (c *Client) FindIPReservations(id string, include string) (*IpReservationList, *APIResponse, error) {
	return c.IPAddresses.FindIPReservations(id, include)
}

func (c *Client) RequestIPReservation(id string, ipReservationRequest IpReservationRequestInput) (*IpReservationRequest, *APIResponse, error) {
	return c.IPAddresses.RequestIPReservation(id, ipReservationRequest)
}

func (c *Client) CheckCapacity(facility CapacityInput) (*APIResponse, error) {
	return c.Capacity.CheckCapacity(facility)
}

func (c *Client) FindCapacity() (*CapacityList, *APIResponse, error) {
	return c.Capacity.FindCapacity()
}

func (c *Client) FindOperatingSystems() (*OperatingSystemList, *APIResponse, error) {
	return c.OperatingSystems.FindOperatingSystems()
}

func (c *Client) CreateDevice(id string, device DeviceCreateInput) (*Device, *APIResponse, error) {
	return c.Devices.CreateDevice(id, device)
}

func (c *Client) DeleteDevice(id string) (*APIResponse, error) {
	return c.Devices.DeleteDevice(id)
}

func (c *Client) FindDeviceById(id string, include string) (*Device, *APIResponse, error) {
	return c.Devices.FindDeviceById(id, include)
}

func (c *Client) FindDevices(id string, include string, page int32, perPage int32) (*DeviceList, *APIResponse, error) {
	return c.Devices.FindDevices(id, include, page, perPage)
}

func (c *Client) UpdateDevice(id string, device DeviceUpdateInput) (*Device, *APIResponse, error) {
	return c.Devices.UpdateDevice(id, device)
}

func (c *Client) PerformAction(id string, type_ string) (*APIResponse, error) {
	return c.Actions.PerformAction(id, type_)
}
