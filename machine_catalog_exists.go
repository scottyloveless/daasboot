package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type CheckCatalogNameRequest struct {
	Name string `json:"Name"`
}

type CheckCatalogNameResponse struct {
	Exists bool   `json:"Exists"`
	Error  string `json:"Error,omitempty"`
}

func (cfg *config) checkMachineCatalogName() (CheckCatalogNameResponse, error) {
	const url = "https://api.cloud.com/cvad/manage/MachineCatalogs/MachineCatalogs/$checkCatalogName"

	body := CheckCatalogNameRequest{
		Name: "MC_" + cfg.Clinic,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return CheckCatalogNameResponse{}, fmt.Errorf("error marshalling json: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return CheckCatalogNameResponse{}, fmt.Errorf("error creating HTTP request: %v", err)
	}
	req.Header.Set("Citrix-InstanceId", cfg.SiteID)
	req.Header.Set("Citrix-CustomerId", cfg.CCID)
	req.Header.Set("Authorization", fmt.Sprintf("CWSAuth Bearer=%v", cfg.BearerToken))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return CheckCatalogNameResponse{}, fmt.Errorf("error calling Citrix API for MC list: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return CheckCatalogNameResponse{}, fmt.Errorf("error reading response body: %v", err)
	}

	if res.StatusCode >= 400 && res.StatusCode <= 499 {
		return CheckCatalogNameResponse{}, fmt.Errorf("citrix error checking mc exists: status=%d body=%s", res.StatusCode, strings.TrimSpace(string(data)))
	}
	if res.StatusCode >= 500 {
		return CheckCatalogNameResponse{}, fmt.Errorf("server error checking mc exists: status=%d body=%s", res.StatusCode, strings.TrimSpace(string(data)))
	}

	var result CheckCatalogNameResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return CheckCatalogNameResponse{}, fmt.Errorf("parse JSON: %w (raw=%s)", err, string(data))
	}

	return result, nil
}
