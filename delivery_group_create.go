package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type DeliveryGroupParams struct {
	// AdminFolder  string `json:"AdminFolder"`
	Applications ApplicationsStruct `json:"Applications"`
	// UserManagement string `json:"UserManagement"`
	DeliveryType string `json:"DeliveryType"`
	Description  string `json:"Description"`
	//	Desktops       []struct {
	//		ID                        string   `json:"Id"`
	//		ColorDepth                string   `json:"ColorDepth"`
	//		Description               string   `json:"Description"`
	//		Enabled                   bool     `json:"Enabled"`
	//		ExcludedUserFilterEnabled bool     `json:"ExcludedUserFilterEnabled"`
	//		ExcludedUsers             []string `json:"ExcludedUsers"`
	//		Icon                      string   `json:"Icon"`
	//		IncludedUserFilterEnabled bool     `json:"IncludedUserFilterEnabled"`
	//		IncludedUsers             []string `json:"IncludedUsers"`
	//		LeasingBehavior           string   `json:"LeasingBehavior"`
	//		MaxDesktops               int      `json:"MaxDesktops"`
	//		Name                      string   `json:"Name"`
	//		PublishedName             string   `json:"PublishedName"`
	//		RestrictToTag             string   `json:"RestrictToTag"`
	//		SecureIcaRequired         bool     `json:"SecureIcaRequired"`
	//		SessionReconnection       string   `json:"SessionReconnection"`
	//	} `json:"Desktops"`
	//
	// Enabled           bool `json:"Enabled"`
	// InMaintenanceMode bool `json:"InMaintenanceMode"`
	//
	MachineCatalogs MachineCatalogStruct `json:"MachineCatalogs"`
	//
	SessionSupport string `json:"SessionSupport"`
	SharingKind    string `json:"SharingKind"`
	// IsRemotePC                         bool     `json:"IsRemotePC"`
	MinimumFunctionalLevel string `json:"MinimumFunctionalLevel"`
	Name                   string `json:"Name"`
	// RequireUserHomeZone                bool     `json:"RequireUserHomeZone"`
	// Scopes                             []string `json:"Scopes"`
	// Tenants                            []string `json:"Tenants"`
	// AppProtectionKeyLoggingRequired    bool     `json:"AppProtectionKeyLoggingRequired"`
	// AppProtectionScreenCaptureRequired bool     `json:"AppProtectionScreenCaptureRequired"`
	//
	AppAccessPolicy AppAccessPolicyStruct `json:"AppAccessPolicy"`
	//
	// AutomaticPowerOnForAssigned           bool   `json:"AutomaticPowerOnForAssigned"`
	// AutomaticPowerOnForAssignedDuringPeak bool   `json:"AutomaticPowerOnForAssignedDuringPeak"`
	// AutoScaleEnabled                      bool   `json:"AutoScaleEnabled"`
	// RestrictAutoscaleTag                  string `json:"RestrictAutoscaleTag"`
	// ColorDepth                            string `json:"ColorDepth"`
	// DefaultDesktopIcon                    string `json:"DefaultDesktopIcon"`
	// DefaultDesktopPublishedName           string `json:"DefaultDesktopPublishedName"`
	// HdxSslEnabled                         bool   `json:"HdxSslEnabled"`
	//
	LingerSettings LingerSettingsStruct `json:"LingerSettings"`
	//
	// OffPeakBufferSizePercent                             int    `json:"OffPeakBufferSizePercent"`
	// OffPeakDisconnectAction                              string `json:"OffPeakDisconnectAction"`
	// OffPeakDisconnectTimeoutMinutes                      int    `json:"OffPeakDisconnectTimeoutMinutes"`
	// OffPeakExtendedDisconnectAction                      string `json:"OffPeakExtendedDisconnectAction"`
	// OffPeakExtendedDisconnectTimeoutMinutes              int    `json:"OffPeakExtendedDisconnectTimeoutMinutes"`
	// OffPeakLogOffAction                                  string `json:"OffPeakLogOffAction"`
	// OffPeakLogOffTimeoutMinutes                          int    `json:"OffPeakLogOffTimeoutMinutes"`
	// LimitSecondsToForceLogOffUserDuringPeak              int    `json:"LimitSecondsToForceLogOffUserDuringPeak"`
	// LimitSecondsToForceLogOffUserDuringOffPeak           int    `json:"LimitSecondsToForceLogOffUserDuringOffPeak"`
	// LogOffWarningMessage                                 string `json:"LogOffWarningMessage"`
	// LogOffWarningTitle                                   string `json:"LogOffWarningTitle"`
	// LoadBalanceType                                      string `json:"LoadBalanceType"`
	// RestrictAutoscaleMinIdleUntaggedPercentDuringPeak    int    `json:"RestrictAutoscaleMinIdleUntaggedPercentDuringPeak"`
	// RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak int    `json:"RestrictAutoscaleMinIdleUntaggedPercentDuringOffPeak"`
	// AutoscaleLogOffReminderEnabled                       bool   `json:"AutoscaleLogOffReminderEnabled"`
	// AutoscaleLogOffReminderIntervalSecondsOffPeak        int    `json:"AutoscaleLogOffReminderIntervalSecondsOffPeak"`
	// AutoscaleLogOffReminderIntervalSecondsPeak           int    `json:"AutoscaleLogOffReminderIntervalSecondsPeak"`
	// PeakAutoscaleAssignedPowerOnIdleTimeoutMinutes       int    `json:"PeakAutoscaleAssignedPowerOnIdleTimeoutMinutes"`
	// PeakAutoscaleAssignedPowerOnIdleAction               string `json:"PeakAutoscaleAssignedPowerOnIdleAction"`
	// AutoscaleLogOffReminderMessage                       string `json:"AutoscaleLogOffReminderMessage"`
	// AutoscaleLogOffReminderTitle                         string `json:"AutoscaleLogOffReminderTitle"`
	// PeakBufferSizePercent                                int    `json:"PeakBufferSizePercent"`
	// PeakDisconnectAction                                 string `json:"PeakDisconnectAction"`
	// PeakDisconnectTimeoutMinutes                         int    `json:"PeakDisconnectTimeoutMinutes"`
	// PeakExtendedDisconnectAction                         string `json:"PeakExtendedDisconnectAction"`
	// PeakExtendedDisconnectTimeoutMinutes                 int    `json:"PeakExtendedDisconnectTimeoutMinutes"`
	// PeakLogOffAction                                     string `json:"PeakLogOffAction"`
	// PeakLogOffTimeoutMinutes                             int    `json:"PeakLogOffTimeoutMinutes"`
	// DisconnectPeakIdleSessionAfterSeconds                int    `json:"DisconnectPeakIdleSessionAfterSeconds"`
	// DisconnectOffPeakIdleSessionAfterSeconds             int    `json:"DisconnectOffPeakIdleSessionAfterSeconds"`
	// LogoffPeakDisconnectedSessionAfterSeconds            int    `json:"LogoffPeakDisconnectedSessionAfterSeconds"`
	// LogoffOffPeakDisconnectedSessionAfterSeconds         int    `json:"LogoffOffPeakDisconnectedSessionAfterSeconds"`
	// AutoscaleScaleDownActionDuringPeak                   string `json:"AutoscaleScaleDownActionDuringPeak"`
	// AutoscaleScaleDownActionDuringOffPeak                string `json:"AutoscaleScaleDownActionDuringOffPeak"`
	//
	//	PrelaunchSettings                                    struct {
	//		Enabled                        bool     `json:"Enabled"`
	//		MaxAverageLoadThreshold        int      `json:"MaxAverageLoadThreshold"`
	//		MaxLoadPerMachineThreshold     int      `json:"MaxLoadPerMachineThreshold"`
	//		MaxTimeBeforeDisconnectMinutes int      `json:"MaxTimeBeforeDisconnectMinutes"`
	//		MaxTimeBeforeTerminateMinutes  int      `json:"MaxTimeBeforeTerminateMinutes"`
	//		IncludedUserFilterEnabled      bool     `json:"IncludedUserFilterEnabled"`
	//		IncludedUsers                  []string `json:"IncludedUsers"`
	//	} `json:"PrelaunchSettings"`
	//
	//	PowerTimeSchemes []struct {
	//		DaysOfWeek       []string `json:"DaysOfWeek"`
	//		Name             string   `json:"Name"`
	//		DisplayName      string   `json:"DisplayName"`
	//		PeakTimeRanges   []string `json:"PeakTimeRanges"`
	//		PoolSizeSchedule []struct {
	//			TimeRange string `json:"TimeRange"`
	//			PoolSize  int    `json:"PoolSize"`
	//		} `json:"PoolSizeSchedule"`
	//		PoolUsingPercentage bool `json:"PoolUsingPercentage"`
	//	} `json:"PowerTimeSchemes"`
	//
	// PowerOffDelayMinutes                 int      `json:"PowerOffDelayMinutes"`
	// MachineCost                          int      `json:"MachineCost"`
	// MachineLogOnType                     string   `json:"MachineLogOnType"`
	// ProductCode                          string   `json:"ProductCode"`
	// ReuseMachinesWithoutShutdownInOutage bool     `json:"ReuseMachinesWithoutShutdownInOutage"`
	// LicenseModel                         string   `json:"LicenseModel"`
	// ProtocolPriority                     []string `json:"ProtocolPriority"`
	//
	//	RebootSchedules                      []struct {
	//		ID                           string    `json:"Id"`
	//		Name                         string    `json:"Name"`
	//		RebootDurationMinutes        int       `json:"RebootDurationMinutes"`
	//		Day                          string    `json:"Day"`
	//		DaysInWeek                   []string  `json:"DaysInWeek"`
	//		WeekInMonth                  string    `json:"WeekInMonth"`
	//		DayInMonth                   string    `json:"DayInMonth"`
	//		Description                  string    `json:"Description"`
	//		Enabled                      bool      `json:"Enabled"`
	//		IgnoreMaintenanceMode        bool      `json:"IgnoreMaintenanceMode"`
	//		UseNaturalReboot             bool      `json:"UseNaturalReboot"`
	//		Frequency                    string    `json:"Frequency"`
	//		FrequencyFactor              int       `json:"FrequencyFactor"`
	//		RestrictToTag                string    `json:"RestrictToTag"`
	//		StartDate                    string    `json:"StartDate"`
	//		StartTime                    time.Time `json:"StartTime"`
	//		WarningDurationMinutes       int       `json:"WarningDurationMinutes"`
	//		WarningMessage               string    `json:"WarningMessage"`
	//		WarningRepeatIntervalMinutes int       `json:"WarningRepeatIntervalMinutes"`
	//		WarningTitle                 string    `json:"WarningTitle"`
	//	} `json:"RebootSchedules"`
	//
	// SecureIcaRequired                         bool `json:"SecureIcaRequired"`
	// SettlementPeriodBeforeAutoShutdownSeconds int  `json:"SettlementPeriodBeforeAutoShutdownSeconds"`
	// SettlementPeriodBeforeUseSeconds          int  `json:"SettlementPeriodBeforeUseSeconds"`
	// ShutdownDesktopsAfterUse                  bool `json:"ShutdownDesktopsAfterUse"`
	//
	//	SimpleAccessPolicy                        struct {
	//		AllowAnonymous                       bool `json:"AllowAnonymous"`
	//		AllowHdxAccess                       bool `json:"AllowHdxAccess"`
	//		AllowMachineRestart                  bool `json:"AllowMachineRestart"`
	//		AllowRdpAccess                       bool `json:"AllowRdpAccess"`
	//		ConnectNotViaNetScalerGatewayAllowed bool `json:"ConnectNotViaNetScalerGatewayAllowed"`
	//		ConnectViaNetScalerGatewayAllowed    bool `json:"ConnectViaNetScalerGatewayAllowed"`
	//		IncludedSmartAccessFilterEnabled     bool `json:"IncludedSmartAccessFilterEnabled"`
	//		IncludedSmartAccessTags              []struct {
	//			Farm   string `json:"Farm"`
	//			Filter string `json:"Filter"`
	//		} `json:"IncludedSmartAccessTags"`
	//		IncludedUserFilterEnabled bool     `json:"IncludedUserFilterEnabled"`
	//		IncludedUsers             []string `json:"IncludedUsers"`
	//		ExcludedUserFilterEnabled bool     `json:"ExcludedUserFilterEnabled"`
	//		ExcludedUsers             []string `json:"ExcludedUsers"`
	//	} `json:"SimpleAccessPolicy"`
	//
	//	StoreFrontServersForHostedReceiver []struct {
	//		ID          string `json:"Id"`
	//		Name        string `json:"Name"`
	//		Description string `json:"Description"`
	//		URL         string `json:"Url"`
	//		Enabled     bool   `json:"Enabled"`
	//	} `json:"StoreFrontServersForHostedReceiver"`
	//
	// TimeZone            string   `json:"TimeZone"`
	// TurnOnAddedMachines bool     `json:"TurnOnAddedMachines"`
	// ZonePreferences     []string `json:"ZonePreferences"`
	//
	//	Metadata            []struct {
	//		Name  string `json:"Name"`
	//		Value string `json:"Value"`
	//	} `json:"Metadata"`
	//
	// PolicySetGUID           string `json:"PolicySetGuid"`
	// RequiredSleepCapability string `json:"RequiredSleepCapability"`
}

type MachineCatalogStruct []struct {
	MachineCatalog string `json:"MachineCatalog"`
	// Count                 int    `json:"Count"`
	// AssignMachinesToUsers []struct {
	// 	Machine              string   `json:"Machine"`
	// 	MachinePublishedName string   `json:"MachinePublishedName"`
	// 	Users                []string `json:"Users"`
	// 	Icon                 string   `json:"Icon"`
	// } `json:"AssignMachinesToUsers"`
}
type AppAccessPolicyStruct struct {
	Enabled bool `json:"Enabled"`
	//		ExcludedUserFilterEnabled bool     `json:"ExcludedUserFilterEnabled"`
	//		ExcludedUsers             []string `json:"ExcludedUsers"`
	//		IncludedUserFilterEnabled bool     `json:"IncludedUserFilterEnabled"`
	IncludedUsers []string `json:"IncludedUsers"`
	//		LeasingBehavior           string   `json:"LeasingBehavior"`
	//		SessionReconnection       string   `json:"SessionReconnection"`
}
type LingerSettingsStruct struct {
	Enabled bool `json:"Enabled"`
	//		MaxAverageLoadThreshold        int      `json:"MaxAverageLoadThreshold"`
	//		MaxLoadPerMachineThreshold     int      `json:"MaxLoadPerMachineThreshold"`
	//		MaxTimeBeforeDisconnectMinutes int      `json:"MaxTimeBeforeDisconnectMinutes"`
	MaxTimeBeforeTerminateMinutes int `json:"MaxTimeBeforeTerminateMinutes"`
	//		IncludedUserFilterEnabled      bool     `json:"IncludedUserFilterEnabled"`
	//		IncludedUsers                  []string `json:"IncludedUsers"`
}
type ApplicationsStruct struct {
	NewApplications NewApplicationsStruct `json:"NewApplications"`
}

type NewApplicationsStruct []struct {
	// ApplicationFolder             string   `json:"ApplicationFolder"`
	// ApplicationType               string   `json:"ApplicationType"`
	// PackagedApplicationType       string   `json:"PackagedApplicationType"`
	// PackagedApplicationVisibility string   `json:"PackagedApplicationVisibility"`
	// BrowserName                   string   `json:"BrowserName"`
	// ClientFolder                  string   `json:"ClientFolder"`
	// CPUPriorityLevel              string   `json:"CpuPriorityLevel"`
	// ApplicationGroups             []string `json:"ApplicationGroups"`
	Description string `json:"Description"`
	Enabled     bool   `json:"Enabled"`
	// DoNotEnumerate                bool     `json:"DoNotEnumerate"`
	// HomeZone                      string   `json:"HomeZone"`
	// HomeZoneMode                  string   `json:"HomeZoneMode"`
	// Icon string `json:"Icon"`
	IconFromClient bool `json:"IconFromClient"`
	// IncludedUserFilterEnabled     bool     `json:"IncludedUserFilterEnabled"`
	// IncludedUsers                 []string `json:"IncludedUsers"`
	InstalledAppProperties InstalledAppPropertiesStruct `json:"InstalledAppProperties"`
	// PackagedAppProperties struct {
	// 	ID                            string `json:"Id"`
	// 	Identifier                    string `json:"Identifier"`
	// 	PackageID                     string `json:"PackageId"`
	// 	PackageName                   string `json:"PackageName"`
	// 	PackageVersion                string `json:"PackageVersion"`
	// 	PackageVersionID              string `json:"PackageVersionId"`
	// 	PublishingServer              string `json:"PublishingServer"`
	// 	SequenceLocation              string `json:"SequenceLocation"`
	// 	ServerMachineConfigurationUID string `json:"ServerMachineConfigurationUid"`
	// 	TargetInPackage               bool   `json:"TargetInPackage"`
	// } `json:"PackagedAppProperties"`
	// AppVAppProperties struct {
	// 	ID                            string `json:"Id"`
	// 	Identifier                    string `json:"Identifier"`
	// 	PackageID                     string `json:"PackageId"`
	// 	PackageName                   string `json:"PackageName"`
	// 	PackageVersion                string `json:"PackageVersion"`
	// 	PackageVersionID              string `json:"PackageVersionId"`
	// 	PublishingServer              string `json:"PublishingServer"`
	// 	SequenceLocation              string `json:"SequenceLocation"`
	// 	ServerMachineConfigurationUID string `json:"ServerMachineConfigurationUid"`
	// 	TargetInPackage               bool   `json:"TargetInPackage"`
	// } `json:"AppVAppProperties"`
	// ContentLocation          string `json:"ContentLocation"`
	// MaxPerUserInstances      int    `json:"MaxPerUserInstances"`
	// MaxTotalInstances        int    `json:"MaxTotalInstances"`
	Name string `json:"Name"`
	// PublishedName            string `json:"PublishedName"`
	// ShortcutAddedToDesktop   bool   `json:"ShortcutAddedToDesktop"`
	// ShortcutAddedToStartMenu bool   `json:"ShortcutAddedToStartMenu"`
	// StartMenuFolder          string `json:"StartMenuFolder"`
	// Visible                  bool   `json:"Visible"`
	// WaitForPrinterCreation   bool   `json:"WaitForPrinterCreation"`
	// Metadata                 []struct {
	// 	Name  string `json:"Name"`
	// 	Value string `json:"Value"`
	// } `json:"Metadata"`
	// FileTypes []struct {
	// 	ContentType   string `json:"ContentType"`
	// 	ExtensionName string `json:"ExtensionName"`
	// 	Description   string `json:"Description"`
	// 	HandlerName   string `json:"HandlerName"`
	// 	OpenArguments string `json:"OpenArguments"`
	// } `json:"FileTypes"`
	// 		ExistingApplicationGroups []struct {
	// Item     string `json:"Item"`
	// Priority int    `json:"Priority"`
}

type InstalledAppPropertiesStruct struct {
	// 	CommandLineArguments  string `json:"CommandLineArguments"`
	CommandLineExecutable string `json:"CommandLineExecutable"`
	WorkingDirectory      string `json:"WorkingDirectory"`
}

func (cfg *config) createDeliveryGroup() error {
	const url = "https://api.cloud.com/cvad/manage/DeliveryGroups"

	body := DeliveryGroupParams{
		Name:                   "DG_" + cfg.Clinic,
		SessionSupport:         "MultiSession",
		SharingKind:            "Shared",
		DeliveryType:           "DesktopsAndApps",
		MinimumFunctionalLevel: "L7_20",
		MachineCatalogs: MachineCatalogStruct{
			{MachineCatalog: "MC_" + cfg.Clinic},
		},
		// AppAccessPolicy: AppAccessPolicyStruct{
		// 	// Enabled: true,
		// 	// IncludedUsers: []string{fmt.Sprintf("SVP\\grp_citrix_%s", cfg.Clinic), "SVP\\grp_citrix_admins"},
		// },
		LingerSettings: LingerSettingsStruct{
			Enabled:                       true,
			MaxTimeBeforeTerminateMinutes: 75,
		},
		Applications: ApplicationsStruct{
			NewApplications: NewApplicationsStruct{
				{
					Name:           cfg.Clinic + "-" + cfg.PIMS,
					Enabled:        true,
					Description:    "Keywords:Mandatory",
					IconFromClient: true,
					InstalledAppProperties: InstalledAppPropertiesStruct{
						CommandLineExecutable: "C:\\" + cfg.PIMS + "\\" + cfg.PIMS + ".exe",
						WorkingDirectory:      "C:\\" + cfg.PIMS + "\\",
					},
				},
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set("Citrix-InstanceId", cfg.SiteID)
	req.Header.Set("Citrix-CustomerId", cfg.CCID)
	req.Header.Set("Authorization", fmt.Sprintf("CWSAuth Bearer=%v", cfg.BearerToken))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error calling Citrix API for MC list: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("citrix failure: status=%d body=%s", res.StatusCode, strings.TrimSpace(string(data)))
	}
	return nil
}
