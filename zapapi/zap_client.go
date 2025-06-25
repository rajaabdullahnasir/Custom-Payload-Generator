package zapapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type ZAPClient struct {
	BaseURL string
	APIKey  string
}

// NewClient returns a new initialized ZAPClient
func NewClient(host, port, apiKey string) *ZAPClient {
	base := fmt.Sprintf("http://%s:%s", host, port)
	return &ZAPClient{BaseURL: base, APIKey: apiKey}
}

// StartActiveScan initiates an active scan using ZAP.
func (z *ZAPClient) StartActiveScan(target string) (string, error) {
	apiURL := fmt.Sprintf("%s/JSON/ascan/action/scan/?apikey=%s&url=%s",
		z.BaseURL, z.APIKey, url.QueryEscape(target))

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("start scan error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ZAP returned status %d: %s", resp.StatusCode, string(body))
	}

	var result map[string]string
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("invalid JSON from scan: %v", err)
	}
	return result["scan"], nil
}

// CheckScanStatus returns the scan status.
func (z *ZAPClient) CheckScanStatus(scanID string) (string, error) {
	apiURL := fmt.Sprintf("%s/JSON/ascan/view/status/?apikey=%s&scanId=%s",
		z.BaseURL, z.APIKey, scanID)

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("status check error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ZAP status error %d: %s", resp.StatusCode, string(body))
	}

	var result map[string]string
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("invalid status JSON: %v", err)
	}
	return result["status"], nil
}

// WaitForScanCompletion blocks until scan reaches 100%
func (z *ZAPClient) WaitForScanCompletion(scanID string) error {
	for {
		status, err := z.CheckScanStatus(scanID)
		if err != nil {
			return err
		}
		if status == "100" {
			return nil
		}
		time.Sleep(2 * time.Second)
	}
}

// GetAlerts fetches all alerts for a base URL
func (z *ZAPClient) GetAlerts(target string) ([]map[string]interface{}, error) {
	apiURL := fmt.Sprintf("%s/JSON/alert/view/alerts/?apikey=%s&baseurl=%s",
		z.BaseURL, z.APIKey, url.QueryEscape(target))

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("get alerts error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ZAP alert error %d: %s", resp.StatusCode, string(body))
	}

	var result map[string][]map[string]interface{}
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("invalid alert JSON: %v", err)
	}
	return result["alerts"], nil
}
