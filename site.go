package main

//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"strings"
// )
//
// type Sites struct {
// 	UserID                string `json:"UserId"`
// 	DisplayName           string `json:"DisplayName"`
// 	ExpiryTime            string `json:"ExpiryTime"`
// 	RefreshExpirationTime string `json:"RefreshExpirationTime"`
// 	VerifiedEmail         string `json:"VerifiedEmail"`
// 	Customers             []struct {
// 		ID    string `json:"Id"`
// 		Name  any    `json:"Name"`
// 		Sites []struct {
// 			ID   string `json:"Id"`
// 			Name string `json:"Name"`
// 		} `json:"Sites"`
// 	} `json:"Customers"`
// }
//
// func (cfg *config) getSites() (Sites, error) {
// 	const url = "https://api.cloud.com/cvad/manage/me"
//
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return Sites{}, fmt.Errorf("error creating HTTP request: %v", err)
// 	}
// 	req.Header.Set("Citrix-CustomerId", cfg.CCID)
// 	req.Header.Set("Authorization", fmt.Sprintf("CWSAuth Bearer=%v", cfg.BearerToken))
// 	req.Header.Set("Accept", "application/json")
// 	req.Header.Set("Content-Type", "application/json")
//
// 	client := http.DefaultClient
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return Sites{}, fmt.Errorf("error calling Citrix API for MC list: %v", err)
// 	}
// 	defer res.Body.Close()
//
// 	data, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return Sites{}, fmt.Errorf("error reading response body: %v", err)
// 	}
//
// 	if res.StatusCode != http.StatusOK {
// 		return Sites{}, fmt.Errorf("citrix failure: status=%d body=%s", res.StatusCode, strings.TrimSpace(string(data)))
// 	}
//
// 	var out Sites
// 	if err := json.Unmarshal(data, &out); err != nil {
// 		return Sites{}, fmt.Errorf("error unmarshalling JSON for get MCs: %v", err)
// 	}
//
// 	return out, nil
// }
