package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type CreateMachineCatalogRules struct {
	AdminFolder          string `json:"AdminFolder"`
	Name                 string `json:"Name"`
	AllocationType       string `json:"AllocationType"`
	Description          string `json:"Description"`
	HypervisorConnection string `json:"HypervisorConnection"`
	IsRemotePC           bool   `json:"IsRemotePC"`
	// RemotePCEnrollmentScopes []struct {
	// 	Ou                   string   `json:"OU"`
	// 	IncludeSubfolders    bool     `json:"IncludeSubfolders"`
	// 	IsOrganizationalUnit bool     `json:"IsOrganizationalUnit"`
	// 	MachinesExcluded     []string `json:"MachinesExcluded"`
	// 	MachinesIncluded     []string `json:"MachinesIncluded"`
	// 	AssignedUsers        []string `json:"AssignedUsers"`
	// } `json:"RemotePCEnrollmentScopes"`
	MachineType            string `json:"MachineType"`
	MinimumFunctionalLevel string `json:"MinimumFunctionalLevel"`
	PersistUserChanges     string `json:"PersistUserChanges"`
	ProvisioningType       string `json:"ProvisioningType"`
	// ProvisioningScheme     struct {
	// 	// 	ResourcePool                           string `json:"ResourcePool"`
	// 	MasterImagePath string `json:"MasterImagePath"`
	// 	// 	AssignImageVersionToProvisioningScheme struct {
	// 	// 		ImageDefinition     string `json:"ImageDefinition"`
	// 	// 		ImageVersion        string `json:"ImageVersion"`
	// 	// 		ImageAssignmentNote string `json:"ImageAssignmentNote"`
	// 	// 	} `json:"AssignImageVersionToProvisioningScheme"`
	// 	// MachineProfilePath         string `json:"MachineProfilePath"`
	// 	// MasterImageNote            string `json:"MasterImageNote"`
	// 	// CPUCount                   int    `json:"CpuCount"`
	// 	// CoresPerCPUCount           int    `json:"CoresPerCpuCount"`
	// 	// MemoryMB                   int    `json:"MemoryMB"`
	// 	// UseWriteBackCache          bool   `json:"UseWriteBackCache"`
	// 	// WriteBackCacheDiskSizeGB   int    `json:"WriteBackCacheDiskSizeGB"`
	// 	// WriteBackCacheMemorySizeMB int    `json:"WriteBackCacheMemorySizeMB"`
	// 	// WriteBackCacheDriveLetter  string `json:"WriteBackCacheDriveLetter"`
	// 	// PrepareImage               bool   `json:"PrepareImage"`
	// 	// NetworkMapping             []struct {
	// 	// 	NetworkDeviceNameOrID string `json:"NetworkDeviceNameOrId"`
	// 	// 	NetworkPath           string `json:"NetworkPath"`
	// 	// } `json:"NetworkMapping"`
	// 	// Metadata []struct {
	// 	// 	Name  string `json:"Name"`
	// 	// 	Value string `json:"Value"`
	// 	// } `json:"Metadata"`
	// 	// ServiceOfferingPath          string   `json:"ServiceOfferingPath"`
	// 	// SecurityGroups               []string `json:"SecurityGroups"`
	// 	// DedicatedTenancy             bool     `json:"DedicatedTenancy"`
	// 	// TenancyType                  string   `json:"TenancyType"`
	// 	// AzureAdJoinType              string   `json:"AzureAdJoinType"`
	// 	// IdentityType                 string   `json:"IdentityType"`
	// 	// DeviceManagementType         string   `json:"DeviceManagementType"`
	// 	// AzureADSecurityGroupName     string   `json:"AzureADSecurityGroupName"`
	// 	// AzureADSecurityGroupNestID   string   `json:"AzureADSecurityGroupNestId"`
	// 	// AzureADSecurityGroupNestName string   `json:"AzureADSecurityGroupNestName"`
	// 	// AzureADTenantID              string   `json:"AzureADTenantId"`
	// 	// CustomProperties             []struct {
	// 	// 	Name  string `json:"Name"`
	// 	// 	Value string `json:"Value"`
	// 	// } `json:"CustomProperties"`
	// 	// CustomPropertiesInString     string `json:"CustomPropertiesInString"`
	// 	// ResetAdministratorPasswords  bool   `json:"ResetAdministratorPasswords"`
	// 	// UseFullDiskCloneProvisioning bool   `json:"UseFullDiskCloneProvisioning"`
	// 	// WorkGroupMachines            bool   `json:"WorkGroupMachines"`
	// 	// NumTotalMachines             int    `json:"NumTotalMachines"`
	// 	MachineAccountCreationRules struct {
	// 		NamingScheme     string `json:"NamingScheme"`
	// 		NamingSchemeType string `json:"NamingSchemeType"`
	// 		// 	Ou               string `json:"OU"`
	// 		Domain string `json:"Domain"`
	// 		// 	NextValue        string `json:"NextValue"`
	// 		// 	IdentityPoolID   string `json:"IdentityPoolId"`
	// 	} `json:"MachineAccountCreationRules"`
	// 	// AvailableMachineAccounts []struct {
	// 	// 	ADAccountName  string `json:"ADAccountName"`
	// 	// 	ResetPassword  bool   `json:"ResetPassword"`
	// 	// 	Password       string `json:"Password"`
	// 	// 	PasswordFormat string `json:"PasswordFormat"`
	// 	// } `json:"AvailableMachineAccounts"`
	// 	// PVSSite           string   `json:"PVSSite"`
	// 	// PVSVDisk          string   `json:"PVSVDisk"`
	// 	// ServiceAccountUID []string `json:"ServiceAccountUid"`
	// } `json:"ProvisioningScheme,omitempty"`
	// ProvisioningSchemeImportData struct {
	// 	IdentityPool       string `json:"IdentityPool"`
	// 	ProvisioningScheme string `json:"ProvisioningScheme"`
	// } `json:"ProvisioningSchemeImportData"`
	// Machines []struct {
	// 	MachineName          string   `json:"MachineName"`
	// 	AssignedClientName   string   `json:"AssignedClientName"`
	// 	AssignedIPAddress    string   `json:"AssignedIPAddress"`
	// 	AssignedUsers        []string `json:"AssignedUsers"`
	// 	HostedMachineID      string   `json:"HostedMachineId"`
	// 	HypervisorConnection string   `json:"HypervisorConnection"`
	// 	InMaintenanceMode    bool     `json:"InMaintenanceMode"`
	// 	Metadata             []struct {
	// 		Name  string `json:"Name"`
	// 		Value string `json:"Value"`
	// 	} `json:"Metadata"`
	// } `json:"Machines"`
	// PvsAddress          string   `json:"PvsAddress"`
	// PvsDomain           string   `json:"PvsDomain"`
	// PvsCollectionIds    []string `json:"PvsCollectionIds"`
	// Scopes              []string `json:"Scopes"`
	// Tenants             []string `json:"Tenants"`
	SessionSupport string `json:"SessionSupport"`
	// VdaUpgradeType      string   `json:"VdaUpgradeType"`
	Zone string `json:"Zone"`
	// HypervisorVMTagging bool     `json:"HypervisorVMTagging"`
	// Metadata            []struct {
	// 	Name  string `json:"Name"`
	// 	Value string `json:"Value"`
	// } `json:"Metadata"`
}

func (cfg *config) createMachineCatalog() (MachineCatalogGetInfo, error) {
	const url = "https://api.cloud.com/cvad/manage/MachineCatalogs"

	body := CreateMachineCatalogRules{
		Name:                   "MC_" + cfg.Clinic,
		Description:            "MC_" + cfg.Clinic,
		AllocationType:         "Random",
		IsRemotePC:             false,
		MachineType:            "Physical",
		MinimumFunctionalLevel: "L7_20",
		PersistUserChanges:     "OnLocal",
		ProvisioningType:       "Manual",
		SessionSupport:         "MultiSession",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return MachineCatalogGetInfo{}, fmt.Errorf("error marshalling JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return MachineCatalogGetInfo{}, fmt.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set("Citrix-InstanceId", cfg.SiteID)
	req.Header.Set("Citrix-CustomerId", cfg.CCID)
	req.Header.Set("Authorization", fmt.Sprintf("CWSAuth Bearer=%v", cfg.BearerToken))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return MachineCatalogGetInfo{}, fmt.Errorf("error calling Citrix API for MC list: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return MachineCatalogGetInfo{}, fmt.Errorf("error reading response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return MachineCatalogGetInfo{}, fmt.Errorf("citrix failure: status=%d body=%s", res.StatusCode, strings.TrimSpace(string(data)))
	}

	var out MachineCatalogGetInfo
	if err := json.Unmarshal(data, &out); err != nil {
		return MachineCatalogGetInfo{}, fmt.Errorf("error unmarshalling JSON for get MCs: %v", err)
	}

	return out, nil
}
