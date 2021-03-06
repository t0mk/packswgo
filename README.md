# Go API client for packswgo

This is the API for Packet. Interact with your devices, user account, and projects.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./packswgo"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.packet.net*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActionsApi* | [**PerformAction**](docs/ActionsApi.md#performaction) | **Post** /devices/{id}/actions | Perform an action
*BatchesApi* | [**FindBatchById**](docs/BatchesApi.md#findbatchbyid) | **Get** /batches/{id} | Retrieve a Batch
*BatchesApi* | [**FindBatchesByProject**](docs/BatchesApi.md#findbatchesbyproject) | **Get** /projects/{id}/batches | Retrieve all batches by project
*BgpConfigsApi* | [**FindBgpConfigByProject**](docs/BgpConfigsApi.md#findbgpconfigbyproject) | **Get** /projects/{id}/bgp-config | Retrieve a bgp config
*BgpConfigsApi* | [**RequestBgpConfig**](docs/BgpConfigsApi.md#requestbgpconfig) | **Post** /projects/{id}/bgp-configs | Requesting bgp config
*BgpSessionsApi* | [**CreateBgpSession**](docs/BgpSessionsApi.md#createbgpsession) | **Post** /devices/{id}/bgp/sessions | Create a BGP session
*BgpSessionsApi* | [**DeleteBgpSession**](docs/BgpSessionsApi.md#deletebgpsession) | **Delete** /bgp/sessions/{id} | Delete the BGP session
*BgpSessionsApi* | [**FindBgpSessionById**](docs/BgpSessionsApi.md#findbgpsessionbyid) | **Get** /bgp/sessions/{id} | Retrieve a BGP session
*BgpSessionsApi* | [**FindBgpSessions**](docs/BgpSessionsApi.md#findbgpsessions) | **Get** /devices/{id}/bgp/sessions | Retrieve all BGP sessions
*CapacityApi* | [**CheckCapacity**](docs/CapacityApi.md#checkcapacity) | **Post** /capacity | Check capacity
*CapacityApi* | [**FindCapacity**](docs/CapacityApi.md#findcapacity) | **Get** /capacity | View capacity
*DevicesApi* | [**CreateDevice**](docs/DevicesApi.md#createdevice) | **Post** /projects/{id}/devices | Create a device
*DevicesApi* | [**DeleteDevice**](docs/DevicesApi.md#deletedevice) | **Delete** /devices/{id} | Delete the device
*DevicesApi* | [**FindDeviceById**](docs/DevicesApi.md#finddevicebyid) | **Get** /devices/{id} | Retrieve a device
*DevicesApi* | [**FindDevices**](docs/DevicesApi.md#finddevices) | **Get** /projects/{id}/devices | Retrieve all devices
*DevicesApi* | [**UpdateDevice**](docs/DevicesApi.md#updatedevice) | **Put** /devices/{id} | Update the device
*EmailsApi* | [**CreateEmail**](docs/EmailsApi.md#createemail) | **Post** /emails | Create an email
*EmailsApi* | [**DeleteEmail**](docs/EmailsApi.md#deleteemail) | **Delete** /emails/{id} | Delete the email
*EmailsApi* | [**FindEmailById**](docs/EmailsApi.md#findemailbyid) | **Get** /emails/{id} | Retrieve an email
*EmailsApi* | [**UpdateEmail**](docs/EmailsApi.md#updateemail) | **Put** /emails/{id} | Update the email
*EventsApi* | [**FindDeviceEvents**](docs/EventsApi.md#finddeviceevents) | **Get** /devices/{id}/events | Retrieve device&#39;s events
*EventsApi* | [**FindEventById**](docs/EventsApi.md#findeventbyid) | **Get** /events/{id} | Retrieve an event
*EventsApi* | [**FindEvents**](docs/EventsApi.md#findevents) | **Get** /events | Retrieve current user&#39;s events
*EventsApi* | [**FindProjectEvents**](docs/EventsApi.md#findprojectevents) | **Get** /projects/{id}/events | Retrieve project&#39;s events
*EventsApi* | [**FindVolumeEvents**](docs/EventsApi.md#findvolumeevents) | **Get** /volumes/{id}/events | Retrieve volume&#39;s events
*FacilitiesApi* | [**FindFacilities**](docs/FacilitiesApi.md#findfacilities) | **Get** /facilities | Retrieve all facilities
*FacilitiesApi* | [**FindFacilitiesByProject**](docs/FacilitiesApi.md#findfacilitiesbyproject) | **Get** /projects/{id}/facilities | Retrieve all facilities visible by the project
*HardwareReservationsApi* | [**FindHardwareReservationById**](docs/HardwareReservationsApi.md#findhardwarereservationbyid) | **Get** /hardware-reservations/{id} | Retrieve a hardware reservation
*HardwareReservationsApi* | [**FindHardwareReservations**](docs/HardwareReservationsApi.md#findhardwarereservations) | **Get** /projects/{id}/hardware-reservations | Retrieve all hardware reservations
*IPAddressesApi* | [**CreateIPAssignment**](docs/IPAddressesApi.md#createipassignment) | **Post** /devices/{id}/ips | Create a ip assignment
*IPAddressesApi* | [**DeleteIPAddress**](docs/IPAddressesApi.md#deleteipaddress) | **Delete** /ips/{id} | Unassign an ip address
*IPAddressesApi* | [**FindIPAddressById**](docs/IPAddressesApi.md#findipaddressbyid) | **Get** /ips/{id} | Retrieve an ip address
*IPAddressesApi* | [**FindIPAssignments**](docs/IPAddressesApi.md#findipassignments) | **Get** /devices/{id}/ips | Retrieve all ip assignments
*IPAddressesApi* | [**FindIPAvailabilities**](docs/IPAddressesApi.md#findipavailabilities) | **Get** /ips/{id}/available | Retrieve all available subnets of a particular reservation
*IPAddressesApi* | [**FindIPReservations**](docs/IPAddressesApi.md#findipreservations) | **Get** /projects/{id}/ips | Retrieve all ip reservations
*IPAddressesApi* | [**RequestIPReservation**](docs/IPAddressesApi.md#requestipreservation) | **Post** /projects/{id}/ips | Requesting ip reservations
*InstanceBandwidthApi* | [**FindInstanceBandwidth**](docs/InstanceBandwidthApi.md#findinstancebandwidth) | **Get** /devices/{id}/bandwidth | Retrieve an instance bandwidth
*InstancesBatchApi* | [**CreateDeviceBatch**](docs/InstancesBatchApi.md#createdevicebatch) | **Post** /projects/{id}/devices/batch | Create a devices batch
*InternetGatewaysApi* | [**CreateInternetGateway**](docs/InternetGatewaysApi.md#createinternetgateway) | **Post** /virtual-networks/{id}/internet-gateways | Create an internet gateway
*InvitationsApi* | [**AcceptInvitation**](docs/InvitationsApi.md#acceptinvitation) | **Put** /invitations/{id} | Accept an invitation
*InvitationsApi* | [**CreateInvitation**](docs/InvitationsApi.md#createinvitation) | **Post** /projects/{id}/invitations | Create an invitation
*InvitationsApi* | [**DeclineInvitation**](docs/InvitationsApi.md#declineinvitation) | **Delete** /invitations/{id} | Decline an invitation
*InvitationsApi* | [**FindInvitationById**](docs/InvitationsApi.md#findinvitationbyid) | **Get** /invitations/{id} | View an invitation
*LicensesApi* | [**CreateLicense**](docs/LicensesApi.md#createlicense) | **Post** /projects/{id}/licenses | Create a License
*LicensesApi* | [**DeleteLicense**](docs/LicensesApi.md#deletelicense) | **Delete** /licenses/{id} | Delete the license
*LicensesApi* | [**FindLicenseById**](docs/LicensesApi.md#findlicensebyid) | **Get** /licenses/{id} | Retrieve a license
*LicensesApi* | [**FindLicenses**](docs/LicensesApi.md#findlicenses) | **Get** /projects/{id}/licenses | Retrieve all licenses
*LicensesApi* | [**UpdateLicense**](docs/LicensesApi.md#updatelicense) | **Put** /licenses/{id} | Update the license
*MembershipsApi* | [**DeleteMembership**](docs/MembershipsApi.md#deletemembership) | **Delete** /memberships/{id} | Delete the membership
*MembershipsApi* | [**FindMembershipById**](docs/MembershipsApi.md#findmembershipbyid) | **Get** /memberships/{id} | Retrieve a membership
*MembershipsApi* | [**UpdateMembership**](docs/MembershipsApi.md#updatemembership) | **Put** /memberships/{id} | Update the membership
*NotificationsApi* | [**FindNotificationById**](docs/NotificationsApi.md#findnotificationbyid) | **Get** /notifications/{id} | Retrieve a notification
*NotificationsApi* | [**FindNotifications**](docs/NotificationsApi.md#findnotifications) | **Get** /notifications | Retrieve all notifications
*NotificationsApi* | [**UpdateNotification**](docs/NotificationsApi.md#updatenotification) | **Put** /notifications/{id} | Update the notification
*OperatingSystemsApi* | [**FindOperatingSystems**](docs/OperatingSystemsApi.md#findoperatingsystems) | **Get** /operating-systems | Retrieve all operating systems
*PlansApi* | [**FindPlans**](docs/PlansApi.md#findplans) | **Get** /plans | Retrieve all plans
*PlansApi* | [**FindPlansByProject**](docs/PlansApi.md#findplansbyproject) | **Get** /projects/{id}/plans | Retrieve all plans visible by the project
*ProjectsApi* | [**CreateProject**](docs/ProjectsApi.md#createproject) | **Post** /projects | Create a project
*ProjectsApi* | [**DeleteProject**](docs/ProjectsApi.md#deleteproject) | **Delete** /projects/{id} | Delete the project
*ProjectsApi* | [**FindProjectById**](docs/ProjectsApi.md#findprojectbyid) | **Get** /projects/{id} | Retrieve a project
*ProjectsApi* | [**FindProjects**](docs/ProjectsApi.md#findprojects) | **Get** /projects | Retrieve all projects
*ProjectsApi* | [**UpdateProject**](docs/ProjectsApi.md#updateproject) | **Put** /projects/{id} | Update the project
*SSHKeysApi* | [**CreateSSHKey**](docs/SSHKeysApi.md#createsshkey) | **Post** /ssh-keys | Create a ssh key for the current user
*SSHKeysApi* | [**CreateSSHKey_0**](docs/SSHKeysApi.md#createsshkey_0) | **Post** /projects/{id}/ssh-keys | Create a ssh key for the given project
*SSHKeysApi* | [**DeleteSSHKey**](docs/SSHKeysApi.md#deletesshkey) | **Delete** /ssh-keys/{id} | Delete the ssh key
*SSHKeysApi* | [**FindSSHKeyById**](docs/SSHKeysApi.md#findsshkeybyid) | **Get** /ssh-keys/{id} | Retrieve a ssh key
*SSHKeysApi* | [**FindSSHKeys**](docs/SSHKeysApi.md#findsshkeys) | **Get** /ssh-keys | Retrieve all ssk keys
*SSHKeysApi* | [**FindSSHKeys_0**](docs/SSHKeysApi.md#findsshkeys_0) | **Get** /projects/{id}/ssh-keys | Retrieve a project&#39;s ssk keys
*SSHKeysApi* | [**UpdateSSHKey**](docs/SSHKeysApi.md#updatesshkey) | **Put** /ssh-keys/{id} | Update the ssh key
*SpotMarketPricesApi* | [**FindSpotMarketPrices**](docs/SpotMarketPricesApi.md#findspotmarketprices) | **Get** /market/spot/prices | Get current spot market prices
*SpotMarketPricesApi* | [**FindSpotMarketPricesHistory**](docs/SpotMarketPricesApi.md#findspotmarketpriceshistory) | **Get** /market/spot/prices/history | Get spot market prices for a given period of time
*TrafficApi* | [**FindTraffic**](docs/TrafficApi.md#findtraffic) | **Get** /devices/{id}/traffic | Retrieve device traffic
*TransferRequestsApi* | [**AcceptTransferRequest**](docs/TransferRequestsApi.md#accepttransferrequest) | **Put** /transfers/{id} | Accept a transfer request
*TransferRequestsApi* | [**CreateTransferRequest**](docs/TransferRequestsApi.md#createtransferrequest) | **Post** /projects/{id}/transfers | Create a transfer request
*TransferRequestsApi* | [**DeclineTransferRequest**](docs/TransferRequestsApi.md#declinetransferrequest) | **Delete** /transfers/{id} | Decline a transfer request
*TransferRequestsApi* | [**FindTransferRequestById**](docs/TransferRequestsApi.md#findtransferrequestbyid) | **Get** /transfers/{id} | View a transfer request
*UsersApi* | [**FindCurrentUser**](docs/UsersApi.md#findcurrentuser) | **Get** /user | Retrieve the current user
*UsersApi* | [**FindUserById**](docs/UsersApi.md#finduserbyid) | **Get** /users/{id} | Retrieve a user
*UsersApi* | [**FindUsers**](docs/UsersApi.md#findusers) | **Get** /users | Retrieve all users
*UsersApi* | [**UpdateCurrentUser**](docs/UsersApi.md#updatecurrentuser) | **Put** /user | Update the current user
*VPNApi* | [**FindCurrentUserVpnConfig**](docs/VPNApi.md#findcurrentuservpnconfig) | **Get** /user/vpn | Retrieve the client vpn config for current user
*VPNApi* | [**TurnOffCurrentUserVpn**](docs/VPNApi.md#turnoffcurrentuservpn) | **Delete** /user/vpn | Turn off vpn for the current user
*VPNApi* | [**TurnOnCurrentUserVpn**](docs/VPNApi.md#turnoncurrentuservpn) | **Post** /user/vpn | Turn on vpn for the current user
*VirtualNetworksApi* | [**AssignPort**](docs/VirtualNetworksApi.md#assignport) | **Post** /ports/{id}/assign | Assign a port to virtual network
*VirtualNetworksApi* | [**BondPort**](docs/VirtualNetworksApi.md#bondport) | **Post** /ports/{id}/bond | Enabling bonding
*VirtualNetworksApi* | [**CreateVirtualNetwork**](docs/VirtualNetworksApi.md#createvirtualnetwork) | **Post** /projects/{id}/virtual-networks | Create an virtual network
*VirtualNetworksApi* | [**DeleteVirtualNetwork**](docs/VirtualNetworksApi.md#deletevirtualnetwork) | **Delete** /virtual-networks/{id} | Delete a virtual network
*VirtualNetworksApi* | [**DisbondPort**](docs/VirtualNetworksApi.md#disbondport) | **Post** /ports/{id}/disbond | Disabling bonding
*VirtualNetworksApi* | [**FindVirtualNetworks**](docs/VirtualNetworksApi.md#findvirtualnetworks) | **Get** /projects/{id}/virtual-networks | Retrieve all virtual networks
*VirtualNetworksApi* | [**UnassignPort**](docs/VirtualNetworksApi.md#unassignport) | **Post** /ports/{id}/unassign | Unassign a port
*VolumeAttachmentsApi* | [**CreateVolumeAttachment**](docs/VolumeAttachmentsApi.md#createvolumeattachment) | **Post** /storage/{id}/attachments | Attach your volume
*VolumeAttachmentsApi* | [**DeleteVolumeAttachment**](docs/VolumeAttachmentsApi.md#deletevolumeattachment) | **Delete** /storage/attachments/{id} | Detach volume
*VolumeAttachmentsApi* | [**FindVolumeAttachmentById**](docs/VolumeAttachmentsApi.md#findvolumeattachmentbyid) | **Get** /storage/attachments/{id} | Retrieve an attachment
*VolumeAttachmentsApi* | [**FindVolumeAttachments**](docs/VolumeAttachmentsApi.md#findvolumeattachments) | **Get** /storage/{id}/attachments | Retrieve all volume attachment
*VolumeSnapshotPoliciesApi* | [**CreateVolumeSnapshotPolicy**](docs/VolumeSnapshotPoliciesApi.md#createvolumesnapshotpolicy) | **Post** /storage/{id}/snapshot-policies | Create a volume snapshot policy
*VolumeSnapshotPoliciesApi* | [**DeleteVolumeSnapshotPolicy**](docs/VolumeSnapshotPoliciesApi.md#deletevolumesnapshotpolicy) | **Delete** /storage/snapshot-policies/{id} | Delete the volume snapshot policy
*VolumeSnapshotPoliciesApi* | [**UpdateVolumeSnapshotPolicy**](docs/VolumeSnapshotPoliciesApi.md#updatevolumesnapshotpolicy) | **Put** /storage/snapshot-policies/{id} | Update the volume snapshot policy
*VolumeSnapshotsApi* | [**CreateVolumeSnapshot**](docs/VolumeSnapshotsApi.md#createvolumesnapshot) | **Post** /storage/{id}/snapshots | Create a volume snapshot
*VolumeSnapshotsApi* | [**DeleteVolumeSnapshot**](docs/VolumeSnapshotsApi.md#deletevolumesnapshot) | **Delete** /storage/{volume_id}/snapshots/{id} | Delete volume snapshot
*VolumeSnapshotsApi* | [**FindVolumeSnapshots**](docs/VolumeSnapshotsApi.md#findvolumesnapshots) | **Get** /storage/{id}/snapshots | Retrieve all volume snapshot
*VolumesApi* | [**CloneVolume**](docs/VolumesApi.md#clonevolume) | **Post** /storage/{id}/clone | Clone volume/snapshot
*VolumesApi* | [**CreateVolume**](docs/VolumesApi.md#createvolume) | **Post** /projects/{id}/storage | Create a volume
*VolumesApi* | [**DeleteVolume**](docs/VolumesApi.md#deletevolume) | **Delete** /storage/{id} | Delete the volume
*VolumesApi* | [**FindVolumeById**](docs/VolumesApi.md#findvolumebyid) | **Get** /storage/{id} | Retrieve a volume
*VolumesApi* | [**FindVolumes**](docs/VolumesApi.md#findvolumes) | **Get** /projects/{id}/storage | Retrieve all volumes
*VolumesApi* | [**RestoreVolume**](docs/VolumesApi.md#restorevolume) | **Post** /storage/{id}/restore | Restore volume
*VolumesApi* | [**UpdateVolume**](docs/VolumesApi.md#updatevolume) | **Put** /storage/{id} | Update the volume


## Documentation For Models

 - [Address](docs/Address.md)
 - [Batch](docs/Batch.md)
 - [BatchesList](docs/BatchesList.md)
 - [BgpConfig](docs/BgpConfig.md)
 - [BgpConfigRequestInput](docs/BgpConfigRequestInput.md)
 - [BgpSession](docs/BgpSession.md)
 - [BgpSessionList](docs/BgpSessionList.md)
 - [CapacityInput](docs/CapacityInput.md)
 - [CapacityList](docs/CapacityList.md)
 - [CapacityPerBaremetal](docs/CapacityPerBaremetal.md)
 - [CapacityPerFacility](docs/CapacityPerFacility.md)
 - [CapacityReport](docs/CapacityReport.md)
 - [Coordinates](docs/Coordinates.md)
 - [Device](docs/Device.md)
 - [DeviceCreateInput](docs/DeviceCreateInput.md)
 - [DeviceList](docs/DeviceList.md)
 - [DeviceUpdateInput](docs/DeviceUpdateInput.md)
 - [Email](docs/Email.md)
 - [EmailInput](docs/EmailInput.md)
 - [Event](docs/Event.md)
 - [EventInput](docs/EventInput.md)
 - [EventList](docs/EventList.md)
 - [EventType](docs/EventType.md)
 - [EventTypeList](docs/EventTypeList.md)
 - [Facility](docs/Facility.md)
 - [FacilityList](docs/FacilityList.md)
 - [HardwareLocation](docs/HardwareLocation.md)
 - [HardwareReservation](docs/HardwareReservation.md)
 - [HardwareReservationList](docs/HardwareReservationList.md)
 - [Href](docs/Href.md)
 - [InstancesBatchCreateInput](docs/InstancesBatchCreateInput.md)
 - [InstancesBatchCreateInputBatches](docs/InstancesBatchCreateInputBatches.md)
 - [InternetGateway](docs/InternetGateway.md)
 - [Invitation](docs/Invitation.md)
 - [InvitationInput](docs/InvitationInput.md)
 - [IpAssignment](docs/IpAssignment.md)
 - [IpAssignmentInput](docs/IpAssignmentInput.md)
 - [IpAssignmentList](docs/IpAssignmentList.md)
 - [IpAvailabilitiesList](docs/IpAvailabilitiesList.md)
 - [IpReservation](docs/IpReservation.md)
 - [IpReservationList](docs/IpReservationList.md)
 - [IpReservationMessage](docs/IpReservationMessage.md)
 - [IpReservationRequest](docs/IpReservationRequest.md)
 - [IpReservationRequestInput](docs/IpReservationRequestInput.md)
 - [License](docs/License.md)
 - [LicenseCreateInput](docs/LicenseCreateInput.md)
 - [LicenseList](docs/LicenseList.md)
 - [LicenseUpdateInput](docs/LicenseUpdateInput.md)
 - [Membership](docs/Membership.md)
 - [MembershipInput](docs/MembershipInput.md)
 - [Meta](docs/Meta.md)
 - [Notification](docs/Notification.md)
 - [NotificationList](docs/NotificationList.md)
 - [OperatingSystem](docs/OperatingSystem.md)
 - [OperatingSystemList](docs/OperatingSystemList.md)
 - [ParentBlock](docs/ParentBlock.md)
 - [Plan](docs/Plan.md)
 - [PlanList](docs/PlanList.md)
 - [PlanVersion](docs/PlanVersion.md)
 - [Port](docs/Port.md)
 - [Project](docs/Project.md)
 - [ProjectCreateInput](docs/ProjectCreateInput.md)
 - [ProjectList](docs/ProjectList.md)
 - [ProjectUpdateInput](docs/ProjectUpdateInput.md)
 - [ServerInfo](docs/ServerInfo.md)
 - [SnapshotPolicy](docs/SnapshotPolicy.md)
 - [SnapshotPolicyInput](docs/SnapshotPolicyInput.md)
 - [SshKey](docs/SshKey.md)
 - [SshKeyInput](docs/SshKeyInput.md)
 - [SshKeyList](docs/SshKeyList.md)
 - [SupportRequestInput](docs/SupportRequestInput.md)
 - [Timeframe](docs/Timeframe.md)
 - [TransferRequest](docs/TransferRequest.md)
 - [TransferRequestInput](docs/TransferRequestInput.md)
 - [User](docs/User.md)
 - [UserList](docs/UserList.md)
 - [UserUpdateInput](docs/UserUpdateInput.md)
 - [VirtualNetwork](docs/VirtualNetwork.md)
 - [VirtualNetworkCreateInput](docs/VirtualNetworkCreateInput.md)
 - [VirtualNetworkList](docs/VirtualNetworkList.md)
 - [Volume](docs/Volume.md)
 - [VolumeAttachment](docs/VolumeAttachment.md)
 - [VolumeAttachmentInput](docs/VolumeAttachmentInput.md)
 - [VolumeAttachmentList](docs/VolumeAttachmentList.md)
 - [VolumeCreateInput](docs/VolumeCreateInput.md)
 - [VolumeList](docs/VolumeList.md)
 - [VolumeSnapshot](docs/VolumeSnapshot.md)
 - [VolumeSnapshotInput](docs/VolumeSnapshotInput.md)
 - [VolumeSnapshotList](docs/VolumeSnapshotList.md)
 - [VolumeUpdateInput](docs/VolumeUpdateInput.md)
 - [VpnConfig](docs/VpnConfig.md)


## Documentation For Authorization


## x_auth_token

- **Type**: API key 
- **API key parameter name**: X-Auth-Token
- **Location**: HTTP header


## Author

help@packet.net

