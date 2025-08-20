package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type MachineCatalogs struct {
	Items []struct {
		Name                     string `json:"Name"`
		FullName                 string `json:"FullName"`
		ID                       string `json:"Id"`
		AllocationType           string `json:"AllocationType"`
		AssignedCount            int    `json:"AssignedCount"`
		AvailableAssignedCount   int    `json:"AvailableAssignedCount"`
		AvailableCount           int    `json:"AvailableCount"`
		AvailableUnassignedCount int    `json:"AvailableUnassignedCount"`
		Description              string `json:"Description"`
		IsPowerManaged           bool   `json:"IsPowerManaged"`
		IsRemotePC               bool   `json:"IsRemotePC"`
		JobsInProgress           []struct {
			ID   string `json:"Id"`
			UID  int    `json:"Uid"`
			Name string `json:"Name"`
		} `json:"JobsInProgress"`
		MachineType string `json:"MachineType"`
		Metadata    []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"Metadata"`
		MinimumFunctionalLevel string `json:"MinimumFunctionalLevel"`
		HasBeenPromoted        bool   `json:"HasBeenPromoted"`
		HasBeenPromotedFrom    string `json:"HasBeenPromotedFrom"`
		CanRollbackVMImage     bool   `json:"CanRollbackVMImage"`
		CanRecreateCatalog     bool   `json:"CanRecreateCatalog"`
		PersistChanges         string `json:"PersistChanges"`
		ProvisioningScheme     struct {
			ID                  string   `json:"Id"`
			Name                string   `json:"Name"`
			CleanOnBoot         bool     `json:"CleanOnBoot"`
			ControllerAddresses []string `json:"ControllerAddresses"`
			CPUCount            int      `json:"CpuCount"`
			CoresPerCPUCount    int      `json:"CoresPerCpuCount"`
			CurrentDiskImage    struct {
				FunctionalLevel string `json:"FunctionalLevel"`
				Image           struct {
					ID                 string `json:"Id"`
					Name               string `json:"Name"`
					XDPath             string `json:"XDPath"`
					RelativePath       string `json:"RelativePath"`
					FullRelativePath   string `json:"FullRelativePath"`
					ResourceType       string `json:"ResourceType"`
					ObjectTypeName     string `json:"ObjectTypeName"`
					ResourcePoolXDPath string `json:"ResourcePoolXDPath"`
				} `json:"Image"`
				ImageStatus     string `json:"ImageStatus"`
				Date            string `json:"Date"`
				MasterImageNote string `json:"MasterImageNote"`
			} `json:"CurrentDiskImage"`
			HistoricalDiskImages []struct {
				FunctionalLevel string `json:"FunctionalLevel"`
				Image           struct {
					ID                 string `json:"Id"`
					Name               string `json:"Name"`
					XDPath             string `json:"XDPath"`
					RelativePath       string `json:"RelativePath"`
					FullRelativePath   string `json:"FullRelativePath"`
					ResourceType       string `json:"ResourceType"`
					ObjectTypeName     string `json:"ObjectTypeName"`
					ResourcePoolXDPath string `json:"ResourcePoolXDPath"`
				} `json:"Image"`
				ImageStatus     string `json:"ImageStatus"`
				Date            string `json:"Date"`
				MasterImageNote string `json:"MasterImageNote"`
			} `json:"HistoricalDiskImages"`
			CustomProperties []struct {
				Name  string `json:"Name"`
				Value string `json:"Value"`
			} `json:"CustomProperties"`
			ImageRuntimeEnvironment struct {
				OperatingSystem struct {
					Type     string `json:"Type"`
					FullName string `json:"FullName"`
				} `json:"OperatingSystem"`
				VDASessionSupport      string `json:"VDASessionSupport"`
				IdentityJoinStatus     string `json:"IdentityJoinStatus"`
				DeviceEnrollmentStatus string `json:"DeviceEnrollmentStatus"`
				PagingFileSettings     []struct {
					Path            string `json:"Path"`
					InitialSizeInMb int    `json:"InitialSizeInMb"`
					MaxSizeInMb     int    `json:"MaxSizeInMb"`
					IsSystemManaged bool   `json:"IsSystemManaged"`
				} `json:"PagingFileSettings"`
				Capabilities struct {
					AdditionalProp1 string `json:"additionalProp1"`
					AdditionalProp2 string `json:"additionalProp2"`
					AdditionalProp3 string `json:"additionalProp3"`
				} `json:"Capabilities"`
				VDAComponents struct {
					AdditionalProp1 string `json:"additionalProp1"`
					AdditionalProp2 string `json:"additionalProp2"`
					AdditionalProp3 string `json:"additionalProp3"`
				} `json:"VDAComponents"`
			} `json:"ImageRuntimeEnvironment"`
			Warning  string `json:"Warning"`
			Warnings []struct {
				Type    string `json:"Type"`
				Message string `json:"Message"`
			} `json:"Warnings"`
			CustomPropertiesInString string   `json:"CustomPropertiesInString"`
			DedicatedTenancy         bool     `json:"DedicatedTenancy"`
			TenancyType              string   `json:"TenancyType"`
			AzureAdJoinType          string   `json:"AzureAdJoinType"`
			AzureADTenantID          string   `json:"AzureADTenantId"`
			ServiceAccountUID        []string `json:"ServiceAccountUid"`
			AzureAdSecurityGroupName string   `json:"AzureAdSecurityGroupName"`
			IdentityType             string   `json:"IdentityType"`
			DeviceManagementType     string   `json:"DeviceManagementType"`
			IdentityContent          string   `json:"IdentityContent"`
			DiskSizeGB               int      `json:"DiskSizeGB"`
			GpuTypeID                string   `json:"GpuTypeId"`
			ResourcePool             struct {
				ID               string `json:"Id"`
				Name             string `json:"Name"`
				XDPath           string `json:"XDPath"`
				FullRelativePath string `json:"FullRelativePath"`
				Hypervisor       struct {
					ID                string `json:"Id"`
					UID               int    `json:"Uid"`
					Name              string `json:"Name"`
					ConnectionType    string `json:"ConnectionType"`
					PluginFactoryName string `json:"PluginFactoryName"`
				} `json:"Hypervisor"`
			} `json:"ResourcePool"`
			IdentityPool struct {
				ID   string `json:"Id"`
				UID  int    `json:"Uid"`
				Name string `json:"Name"`
			} `json:"IdentityPool"`
			MachineCount int `json:"MachineCount"`
			MasterImage  struct {
				ID                 string `json:"Id"`
				Name               string `json:"Name"`
				XDPath             string `json:"XDPath"`
				RelativePath       string `json:"RelativePath"`
				FullRelativePath   string `json:"FullRelativePath"`
				ResourceType       string `json:"ResourceType"`
				ObjectTypeName     string `json:"ObjectTypeName"`
				ResourcePoolXDPath string `json:"ResourcePoolXDPath"`
			} `json:"MasterImage"`
			MachineProfile struct {
				ID                 string `json:"Id"`
				Name               string `json:"Name"`
				XDPath             string `json:"XDPath"`
				RelativePath       string `json:"RelativePath"`
				FullRelativePath   string `json:"FullRelativePath"`
				ResourceType       string `json:"ResourceType"`
				ObjectTypeName     string `json:"ObjectTypeName"`
				ResourcePoolXDPath string `json:"ResourcePoolXDPath"`
			} `json:"MachineProfile"`
			MemoryMB int `json:"MemoryMB"`
			Metadata []struct {
				Name  string `json:"Name"`
				Value string `json:"Value"`
			} `json:"Metadata"`
			NetworkMaps []struct {
				DeviceName string `json:"DeviceName"`
				DeviceID   string `json:"DeviceId"`
				Network    struct {
					ID                 string `json:"Id"`
					Name               string `json:"Name"`
					XDPath             string `json:"XDPath"`
					RelativePath       string `json:"RelativePath"`
					FullRelativePath   string `json:"FullRelativePath"`
					ResourceType       string `json:"ResourceType"`
					ObjectTypeName     string `json:"ObjectTypeName"`
					ResourcePoolXDPath string `json:"ResourcePoolXDPath"`
				} `json:"Network"`
				IsCardEnabled bool `json:"IsCardEnabled"`
			} `json:"NetworkMaps"`
			NutanixContainer             string   `json:"NutanixContainer"`
			ResetAdministratorPasswords  bool     `json:"ResetAdministratorPasswords"`
			SecurityGroups               []string `json:"SecurityGroups"`
			ServiceOffering              string   `json:"ServiceOffering"`
			UseFullDiskCloneProvisioning bool     `json:"UseFullDiskCloneProvisioning"`
			UseWriteBackCache            bool     `json:"UseWriteBackCache"`
			VMMetadata                   struct {
				MachineSize                       string `json:"MachineSize"`
				OsDiskCaching                     string `json:"OsDiskCaching"`
				Tags                              string `json:"Tags"`
				BootDiagnostics                   bool   `json:"BootDiagnostics"`
				AcceleratedNetwork                bool   `json:"AcceleratedNetwork"`
				SupportsHibernation               bool   `json:"SupportsHibernation"`
				SecurityType                      string `json:"SecurityType"`
				DiskSecurityType                  string `json:"DiskSecurityType"`
				ConfidentialVMDiskEncryptionSetID string `json:"ConfidentialVmDiskEncryptionSetId"`
				EnableSecureBoot                  bool   `json:"EnableSecureBoot"`
				EnableVTPM                        bool   `json:"EnableVTPM"`
				EncryptionAtHost                  bool   `json:"EncryptionAtHost"`
				Labels                            string `json:"Labels"`
				ZoneName                          string `json:"ZoneName"`
				StorageType                       string `json:"StorageType"`
				EncryptionKeyID                   string `json:"EncryptionKeyId"`
			} `json:"VMMetadata"`
			WindowsActivationType       string `json:"WindowsActivationType"`
			WorkgroupMachines           bool   `json:"WorkgroupMachines"`
			WriteBackCacheDiskIndex     int    `json:"WriteBackCacheDiskIndex"`
			WriteBackCacheDiskSizeGB    int    `json:"WriteBackCacheDiskSizeGB"`
			WriteBackCacheMemorySizeMB  int    `json:"WriteBackCacheMemorySizeMB"`
			WriteBackCacheDriveLetter   string `json:"WriteBackCacheDriveLetter"`
			IsPreviousImageLegacyVda    bool   `json:"IsPreviousImageLegacyVda"`
			MachineAccountCreationRules struct {
				NamingScheme     string `json:"NamingScheme"`
				NamingSchemeType string `json:"NamingSchemeType"`
				Ou               string `json:"OU"`
				Domain           struct {
					Forest                               string   `json:"Forest"`
					Parent                               string   `json:"Parent"`
					Name                                 string   `json:"Name"`
					Children                             []string `json:"Children"`
					Sid                                  string   `json:"Sid"`
					GUID                                 string   `json:"Guid"`
					NetBiosName                          string   `json:"NetBiosName"`
					DistinguishedName                    string   `json:"DistinguishedName"`
					Controllers                          []string `json:"Controllers"`
					DefaultController                    string   `json:"DefaultController"`
					TrustedDomains                       []string `json:"TrustedDomains"`
					UpnSuffixes                          []string `json:"UpnSuffixes"`
					ServiceConnectionPointConfigurations []struct {
						Name     string   `json:"Name"`
						Keywords []string `json:"Keywords"`
					} `json:"ServiceConnectionPointConfigurations"`
					PossibleLookupFailure bool `json:"PossibleLookupFailure"`
					PropertiesFetched     int  `json:"PropertiesFetched"`
				} `json:"Domain"`
				NextValue string `json:"NextValue"`
			} `json:"MachineAccountCreationRules"`
			NumAvailableMachineAccounts int    `json:"NumAvailableMachineAccounts"`
			PVSSite                     string `json:"PVSSite"`
			PVSVDisk                    string `json:"PVSVDisk"`
			ProvisioningSchemeType      string `json:"ProvisioningSchemeType"`
			CurrentImageVersion         struct {
				ImageVersion struct {
					ID              string `json:"Id"`
					Number          int    `json:"Number"`
					ImageDefinition struct {
						ID   string `json:"Id"`
						UID  int    `json:"Uid"`
						Name string `json:"Name"`
					} `json:"ImageDefinition"`
					Description       string `json:"Description"`
					ImageVersionSpecs []struct {
						ID              string `json:"Id"`
						PreparationType string `json:"PreparationType"`
						ResourcePool    struct {
							ID               string `json:"Id"`
							Name             string `json:"Name"`
							XDPath           string `json:"XDPath"`
							FullRelativePath string `json:"FullRelativePath"`
							Hypervisor       struct {
								ID                string `json:"Id"`
								UID               int    `json:"Uid"`
								Name              string `json:"Name"`
								ConnectionType    string `json:"ConnectionType"`
								PluginFactoryName string `json:"PluginFactoryName"`
							} `json:"Hypervisor"`
						} `json:"ResourcePool"`
					} `json:"ImageVersionSpecs"`
				} `json:"ImageVersion"`
				Date                string `json:"Date"`
				ImageAssignmentNote string `json:"ImageAssignmentNote"`
				IsImageAvailable    bool   `json:"IsImageAvailable"`
			} `json:"CurrentImageVersion"`
			HistoricalImageVersions []struct {
				ImageVersion struct {
					ID              string `json:"Id"`
					Number          int    `json:"Number"`
					ImageDefinition struct {
						ID   string `json:"Id"`
						UID  int    `json:"Uid"`
						Name string `json:"Name"`
					} `json:"ImageDefinition"`
					Description       string `json:"Description"`
					ImageVersionSpecs []struct {
						ID              string `json:"Id"`
						PreparationType string `json:"PreparationType"`
						ResourcePool    struct {
							ID               string `json:"Id"`
							Name             string `json:"Name"`
							XDPath           string `json:"XDPath"`
							FullRelativePath string `json:"FullRelativePath"`
							Hypervisor       struct {
								ID                string `json:"Id"`
								UID               int    `json:"Uid"`
								Name              string `json:"Name"`
								ConnectionType    string `json:"ConnectionType"`
								PluginFactoryName string `json:"PluginFactoryName"`
							} `json:"Hypervisor"`
						} `json:"ResourcePool"`
					} `json:"ImageVersionSpecs"`
				} `json:"ImageVersion"`
				Date                string `json:"Date"`
				ImageAssignmentNote string `json:"ImageAssignmentNote"`
				IsImageAvailable    bool   `json:"IsImageAvailable"`
			} `json:"HistoricalImageVersions"`
		} `json:"ProvisioningScheme"`
		ProvisioningType     string `json:"ProvisioningType"`
		ProvisioningProgress struct {
			IsRunning           bool       `json:"IsRunning"`
			Progress            int        `json:"Progress"`
			ProgressMessage     string     `json:"ProgressMessage"`
			ProgressMessageList [][]string `json:"ProgressMessageList"`
		} `json:"ProvisioningProgress"`
		PvsAddress               string `json:"PvsAddress"`
		PvsDomain                string `json:"PvsDomain"`
		RemotePCEnrollmentScopes []struct {
			Ou                   string   `json:"OU"`
			IncludeSubfolders    bool     `json:"IncludeSubfolders"`
			IsOrganizationalUnit bool     `json:"IsOrganizationalUnit"`
			MachinesExcluded     []string `json:"MachinesExcluded"`
			MachinesIncluded     []string `json:"MachinesIncluded"`
		} `json:"RemotePCEnrollmentScopes"`
		Scopes []struct {
			ID            string `json:"Id"`
			Name          string `json:"Name"`
			Description   string `json:"Description"`
			IsBuiltIn     bool   `json:"IsBuiltIn"`
			IsAllScope    bool   `json:"IsAllScope"`
			IsTenantScope bool   `json:"IsTenantScope"`
			TenantID      string `json:"TenantId"`
			TenantName    string `json:"TenantName"`
		} `json:"Scopes"`
		Tenants []struct {
			ID   string `json:"Id"`
			UID  int    `json:"Uid"`
			Name string `json:"Name"`
		} `json:"Tenants"`
		SessionSupport          string   `json:"SessionSupport"`
		SharingKind             string   `json:"SharingKind"`
		TotalCount              int      `json:"TotalCount"`
		IsBroken                bool     `json:"IsBroken"`
		IsMasterImageAssociated bool     `json:"IsMasterImageAssociated"`
		ImageUpdateStatus       string   `json:"ImageUpdateStatus"`
		Errors                  []string `json:"Errors"`
		Warnings                []struct {
			Type    string `json:"Type"`
			Subtype string `json:"Subtype"`
			Message string `json:"Message"`
		} `json:"Warnings"`
		UnassignedCount                 int `json:"UnassignedCount"`
		UsedCount                       int `json:"UsedCount"`
		AvailableCountOfSuspend         int `json:"AvailableCountOfSuspend"`
		AvailableAssignedCountOfSuspend int `json:"AvailableAssignedCountOfSuspend"`
		UpgradeInfo                     struct {
			UpgradeType                 string `json:"UpgradeType"`
			UpgradeState                string `json:"UpgradeState"`
			UpgradeScheduleStatus       string `json:"UpgradeScheduleStatus"`
			UpgradeOngoingMachinesCount int    `json:"UpgradeOngoingMachinesCount"`
			UpgradeFailedMachinesCount  int    `json:"UpgradeFailedMachinesCount"`
		} `json:"UpgradeInfo"`
		Zone struct {
			ID   string `json:"Id"`
			UID  int    `json:"Uid"`
			Name string `json:"Name"`
		} `json:"Zone"`
		AdminFolder struct {
			ID   string `json:"Id"`
			UID  int    `json:"Uid"`
			Name string `json:"Name"`
		} `json:"AdminFolder"`
		HypervisorVMTagging bool `json:"HypervisorVMTagging"`
	} `json:"Items"`
	ContinuationToken string `json:"ContinuationToken"`
	TotalItems        int    `json:"TotalItems"`
}

func (cfg *config) getMachineCatalogs() (MachineCatalogs, error) {
	const url = "https://api.cloud.com/cvad/manage/MachineCatalogs"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return MachineCatalogs{}, fmt.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set("Citrix-InstanceId", cfg.SiteID)
	req.Header.Set("Citrix-CustomerId", cfg.CCID)
	req.Header.Set("Authorization", fmt.Sprintf("CWSAuth Bearer=%v", cfg.BearerToken))
	req.Header.Set("Accept", "application/json")

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return MachineCatalogs{}, fmt.Errorf("error calling Citrix API for MC list: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return MachineCatalogs{}, fmt.Errorf("error reading response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return MachineCatalogs{}, fmt.Errorf("citrix failure: status=%d body=%s", res.StatusCode, strings.TrimSpace(string(data)))
	}

	var out MachineCatalogs
	if err := json.Unmarshal(data, &out); err != nil {
		return MachineCatalogs{}, fmt.Errorf("error unmarshalling JSON for get MCs: %v", err)
	}

	return out, nil
}
