package main

// TODO: Finish this func
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type DeliveryGroupCheckParams struct {
	Name string `json:"Name"`
}

type DeliveryGroupCheckResponse struct {
	ErrorMessage string `json:"ErrorMessage"`
}

func (cfg *config) checkDeliveryGroupExists() (bool, error) {
	const url = "https://api.cloud.com/cvad/manage/DeliveryGroups/$checkName"

	body := CheckCatalogNameRequest{
		Name: "DG_" + cfg.Clinic,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return false, fmt.Errorf("error marshalling json: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return false, fmt.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set("Citrix-InstanceId", cfg.SiteID)
	req.Header.Set("Citrix-CustomerId", cfg.CCID)
	req.Header.Set("Authorization", fmt.Sprintf("CWSAuth Bearer=%v", cfg.BearerToken))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("error calling Citrix API for MC list: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return false, fmt.Errorf("error reading response body: %v", err)
	}

	if res.StatusCode >= 400 && res.StatusCode < 404 {
		return false, fmt.Errorf("citrix error checking mc exists: status=%d body=%s", res.StatusCode, strings.TrimSpace(string(data)))
	}
	if res.StatusCode > 404 {
		return false, fmt.Errorf("server error checking mc exists: status=%d body=%s", res.StatusCode, strings.TrimSpace(string(data)))
	}
	if res.StatusCode == 404 {
		return false, nil
	}
	return true, nil
}
