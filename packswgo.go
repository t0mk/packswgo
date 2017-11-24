package packswgo

import "os"

// APITokenVar is the name of teh envvar which holds auth token to Packet API
const APITokenVar = "PACKET_AUTH_TOKEN"

// Client is Packet API Client aggregating generated interfaces
type Client struct {
	Actions                ActionsApi
	Batches                BatchesApi
	BgpConfigs             BgpConfigsApi
	BgpSessions            BgpSessionsApi
	Capacity               CapacityApi
	Devices                DevicesApi
	Emails                 EmailsApi
	Events                 EventsApi
	Facilities             FacilitiesApi
	HardwareReservations   HardwareReservationsApi
	InstanceBandwidth      InstanceBandwidthApi
	InstancesBatch         InstancesBatchApi
	InternetGateways       InternetGatewaysApi
	Invitations            InvitationsApi
	IPAddresses            IPAddressesApi
	Licenses               LicensesApi
	Memberships            MembershipsApi
	Notifications          NotificationsApi
	OperatingSystems       OperatingSystemsApi
	Plans                  PlansApi
	Projects               ProjectsApi
	SpotMarketPrices       SpotMarketPricesApi
	SSHKeys                SSHKeysApi
	Traffic                TrafficApi
	TransferRequests       TransferRequestsApi
	Users                  UsersApi
	VirtualNetworks        VirtualNetworksApi
	VolumeAttachments      VolumeAttachmentsApi
	Volumes                VolumesApi
	VolumeSnapshotPolicies VolumeSnapshotPoliciesApi
	VolumeSnapshots        VolumeSnapshotsApi
	VPN                    VPNApi
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
		Actions:                ActionsApi{Configuration: cfg},
		Batches:                BatchesApi{Configuration: cfg},
		BgpConfigs:             BgpConfigsApi{Configuration: cfg},
		BgpSessions:            BgpSessionsApi{Configuration: cfg},
		Capacity:               CapacityApi{Configuration: cfg},
		Devices:                DevicesApi{Configuration: cfg},
		Emails:                 EmailsApi{Configuration: cfg},
		Events:                 EventsApi{Configuration: cfg},
		Facilities:             FacilitiesApi{Configuration: cfg},
		HardwareReservations:   HardwareReservationsApi{Configuration: cfg},
		InstanceBandwidth:      InstanceBandwidthApi{Configuration: cfg},
		InstancesBatch:         InstancesBatchApi{Configuration: cfg},
		InternetGateways:       InternetGatewaysApi{Configuration: cfg},
		Invitations:            InvitationsApi{Configuration: cfg},
		IPAddresses:            IPAddressesApi{Configuration: cfg},
		Licenses:               LicensesApi{Configuration: cfg},
		Memberships:            MembershipsApi{Configuration: cfg},
		Notifications:          NotificationsApi{Configuration: cfg},
		OperatingSystems:       OperatingSystemsApi{Configuration: cfg},
		Plans:                  PlansApi{Configuration: cfg},
		Projects:               ProjectsApi{Configuration: cfg},
		SpotMarketPrices:       SpotMarketPricesApi{Configuration: cfg},
		SSHKeys:                SSHKeysApi{Configuration: cfg},
		Traffic:                TrafficApi{Configuration: cfg},
		TransferRequests:       TransferRequestsApi{Configuration: cfg},
		Users:                  UsersApi{Configuration: cfg},
		VirtualNetworks:        VirtualNetworksApi{Configuration: cfg},
		VolumeAttachments:      VolumeAttachmentsApi{Configuration: cfg},
		Volumes:                VolumesApi{Configuration: cfg},
		VolumeSnapshotPolicies: VolumeSnapshotPoliciesApi{Configuration: cfg},
		VolumeSnapshots:        VolumeSnapshotsApi{Configuration: cfg},
		VPN:                    VPNApi{Configuration: cfg},
	}

}
