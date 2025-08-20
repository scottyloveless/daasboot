package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type MachineCatalogGetInfo struct {
	AgentVersion         string `json:"AgentVersion"`
	HypervisorConnection any    `json:"HypervisorConnection"`
	OSType               string `json:"OSType"`
	OSVersion            string `json:"OSVersion"`
	DeliveryGroups       []any  `json:"DeliveryGroups"`
	UpgradeDetail        struct {
		ScheduleStatus           string `json:"ScheduleStatus"`
		LastStateChangeTimeUtc   string `json:"LastStateChangeTimeUtc"`
		TotalMachines            int    `json:"TotalMachines"`
		SuccessCount             int    `json:"SuccessCount"`
		ValidationFailureCount   int    `json:"ValidationFailureCount"`
		InProgressCount          int    `json:"InProgressCount"`
		UpgradeFailureCount      int    `json:"UpgradeFailureCount"`
		ScheduledTimeUtc         string `json:"ScheduledTimeUtc"`
		DurationInHours          int    `json:"DurationInHours"`
		TargetPackageVersion     any    `json:"TargetPackageVersion"`
		CancelledUpgradeCount    int    `json:"CancelledUpgradeCount"`
		WaitingToUpgradeCount    int    `json:"WaitingToUpgradeCount"`
		AvailableForUpgradeCount int    `json:"AvailableForUpgradeCount"`
	} `json:"UpgradeDetail"`
	Name                     string `json:"Name"`
	FullName                 string `json:"FullName"`
	ID                       string `json:"Id"`
	UID                      int    `json:"Uid"`
	AllocationType           string `json:"AllocationType"`
	AssignedCount            int    `json:"AssignedCount"`
	AvailableAssignedCount   int    `json:"AvailableAssignedCount"`
	AvailableCount           int    `json:"AvailableCount"`
	AvailableUnassignedCount int    `json:"AvailableUnassignedCount"`
	Description              string `json:"Description"`
	IsPowerManaged           bool   `json:"IsPowerManaged"`
	IsRemotePC               bool   `json:"IsRemotePC"`
	JobsInProgress           any    `json:"JobsInProgress"`
	MachineType              string `json:"MachineType"`
	Metadata                 []any  `json:"Metadata"`
	MinimumFunctionalLevel   string `json:"MinimumFunctionalLevel"`
	HasBeenPromoted          bool   `json:"HasBeenPromoted"`
	HasBeenPromotedFrom      string `json:"HasBeenPromotedFrom"`
	CanRollbackVMImage       bool   `json:"CanRollbackVMImage"`
	CanRecreateCatalog       bool   `json:"CanRecreateCatalog"`
	PersistChanges           string `json:"PersistChanges"`
	ProvisioningScheme       any    `json:"ProvisioningScheme"`
	ProvisioningType         string `json:"ProvisioningType"`
	ProvisioningProgress     any    `json:"ProvisioningProgress"`
	PvsAddress               any    `json:"PvsAddress"`
	PvsDomain                any    `json:"PvsDomain"`
	RemotePCEnrollmentScopes any    `json:"RemotePCEnrollmentScopes"`
	Scopes                   []struct {
		ID            string `json:"Id"`
		UID           any    `json:"Uid"`
		Name          string `json:"Name"`
		Description   any    `json:"Description"`
		IsBuiltIn     bool   `json:"IsBuiltIn"`
		IsAllScope    bool   `json:"IsAllScope"`
		IsTenantScope bool   `json:"IsTenantScope"`
		TenantID      any    `json:"TenantId"`
		TenantName    any    `json:"TenantName"`
	} `json:"Scopes"`
	Tenants                         any    `json:"Tenants"`
	SessionSupport                  string `json:"SessionSupport"`
	SharingKind                     string `json:"SharingKind"`
	TotalCount                      int    `json:"TotalCount"`
	IsBroken                        bool   `json:"IsBroken"`
	IsMasterImageAssociated         any    `json:"IsMasterImageAssociated"`
	ImageUpdateStatus               string `json:"ImageUpdateStatus"`
	Errors                          []any  `json:"Errors"`
	Warnings                        []any  `json:"Warnings"`
	UnassignedCount                 int    `json:"UnassignedCount"`
	UsedCount                       int    `json:"UsedCount"`
	AvailableCountOfSuspend         any    `json:"AvailableCountOfSuspend"`
	AvailableAssignedCountOfSuspend any    `json:"AvailableAssignedCountOfSuspend"`
	UpgradeInfo                     struct {
		UpgradeType                 string `json:"UpgradeType"`
		UpgradeState                string `json:"UpgradeState"`
		UpgradeScheduleStatus       string `json:"UpgradeScheduleStatus"`
		UpgradeOngoingMachinesCount int    `json:"UpgradeOngoingMachinesCount"`
		UpgradeFailedMachinesCount  int    `json:"UpgradeFailedMachinesCount"`
	} `json:"UpgradeInfo"`
	Zone struct {
		ID   string `json:"Id"`
		UID  any    `json:"Uid"`
		Name string `json:"Name"`
	} `json:"Zone"`
	AdminFolder struct {
		ID   string `json:"Id"`
		UID  int    `json:"Uid"`
		Name string `json:"Name"`
	} `json:"AdminFolder"`
	HypervisorVMTagging bool `json:"HypervisorVMTagging"`
}

func (cfg *config) getMachineCatalog() (MachineCatalogGetInfo, error) {
	url := "https://api.cloud.com/cvad/manage/MachineCatalogs/" + "MC_" + cfg.Clinic

	req, err := http.NewRequest("GET", url, nil)
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

	if res.StatusCode >= 400 && res.StatusCode <= 499 {
		return MachineCatalogGetInfo{}, fmt.Errorf("citrix error checking mc exists: status=%d body=%s", res.StatusCode, strings.TrimSpace(string(data)))
	}
	if res.StatusCode > 500 {
		return MachineCatalogGetInfo{}, fmt.Errorf("server error checking mc exists: status=%d body=%s", res.StatusCode, strings.TrimSpace(string(data)))
	}

	var rules MachineCatalogGetInfo
	if err := json.Unmarshal(data, &rules); err != nil {
		return MachineCatalogGetInfo{}, fmt.Errorf("parse JSON: %w; raw=%s", err, string(data))
	}

	return rules, nil
}
