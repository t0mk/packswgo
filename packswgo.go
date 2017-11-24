package packswgo

import (
    "os"
    "fmt"
)

// APITokenVar is the name of teh envvar which holds auth token to Packet API
const (
    APITokenVar = "PACKET_AUTH_TOKEN"
    APIURLVar = "PACKET_API_URL"
)

type Client struct {
	VirtualNetworksApi
	NotificationsApi
	LicensesApi
	InstancesBatchApi
	InternetGatewaysApi
	MembershipsApi
	EventsApi
	EmailsApi
	InstanceBandwidthApi
	FacilitiesApi
	UsersApi
	PlansApi
	VPNApi
	SpotMarketPricesApi
	TrafficApi
	VolumesApi
	SSHKeysApi
	BgpConfigsApi
	TransferRequestsApi
	VolumeSnapshotPoliciesApi
	HardwareReservationsApi
	ProjectsApi
	InvitationsApi
	BatchesApi
	VolumeSnapshotsApi
	VolumeAttachmentsApi
	BgpSessionsApi
	IPAddressesApi
	CapacityApi
	OperatingSystemsApi
	DevicesApi
	ActionsApi
}

func getCfg(token, apiURL string) *Configuration {
	cfg := NewConfiguration()
        cfg.APIKey["X-Auth-Token"] = token
        cfg.BasePath = apiURL
	return cfg
}

// NewClient returns Client
func NewClient() (*Client, error) {
        return NewClientWithToken(os.Getenv(APITokenVar))
}

// NewClientWithToken returns API client
func NewClientWithToken(token string) (*Client, error) {
        if token == "" {
            return nil, fmt.Errorf("You must pass Packet API Token, for example via env var %s", APITokenVar)
        }
        apiURL := "https://api.packet.net"
        if os.Getenv(APIURLVar) != "" {
            apiURL = os.Getenv(APIURLVar)
        }

	cfg := getCfg(token, apiURL)
	return &Client{
		VirtualNetworksApi{Configuration: cfg},
		NotificationsApi{Configuration: cfg},
		LicensesApi{Configuration: cfg},
		InstancesBatchApi{Configuration: cfg},
		InternetGatewaysApi{Configuration: cfg},
		MembershipsApi{Configuration: cfg},
		EventsApi{Configuration: cfg},
		EmailsApi{Configuration: cfg},
		InstanceBandwidthApi{Configuration: cfg},
		FacilitiesApi{Configuration: cfg},
		UsersApi{Configuration: cfg},
		PlansApi{Configuration: cfg},
		VPNApi{Configuration: cfg},
		SpotMarketPricesApi{Configuration: cfg},
		TrafficApi{Configuration: cfg},
		VolumesApi{Configuration: cfg},
		SSHKeysApi{Configuration: cfg},
		BgpConfigsApi{Configuration: cfg},
		TransferRequestsApi{Configuration: cfg},
		VolumeSnapshotPoliciesApi{Configuration: cfg},
		HardwareReservationsApi{Configuration: cfg},
		ProjectsApi{Configuration: cfg},
		InvitationsApi{Configuration: cfg},
		BatchesApi{Configuration: cfg},
		VolumeSnapshotsApi{Configuration: cfg},
		VolumeAttachmentsApi{Configuration: cfg},
		BgpSessionsApi{Configuration: cfg},
		IPAddressesApi{Configuration: cfg},
		CapacityApi{Configuration: cfg},
		OperatingSystemsApi{Configuration: cfg},
		DevicesApi{Configuration: cfg},
		ActionsApi{Configuration: cfg},
	}, nil
}


