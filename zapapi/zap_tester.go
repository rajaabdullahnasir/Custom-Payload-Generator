package zapapi

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rajaabdullahnasir/Custom-Payload-Generator/reports"
)

type ScanResult struct {
	TargetURL string                   `json:"target_url"`
	ScanID    string                   `json:"scan_id"`
	Timestamp string                   `json:"timestamp"`
	Alerts    []map[string]interface{} `json:"alerts"`
}

// RunZAPScan performs the full scan and report generation process
func RunZAPScan(targetURL, host, port, apiKey string) error {
	fmt.Println("🚀 Starting ZAP Scan on:", targetURL)

	client := NewClient(host, port, apiKey)

	scanID, err := client.StartActiveScan(targetURL)
	if err != nil {
		return fmt.Errorf("❌ Failed to start scan: %v", err)
	}
	fmt.Println("🔍 Scan ID:", scanID)

	fmt.Println("⏳ Waiting for scan to complete...")
	if err := client.WaitForScanCompletion(scanID); err != nil {
		return fmt.Errorf("❌ Scan wait error: %v", err)
	}
	fmt.Println("✅ Scan complete!")

	alerts, err := client.GetAlerts(targetURL)
	if err != nil {
		return fmt.Errorf("❌ Could not retrieve alerts: %v", err)
	}
	fmt.Printf("📦 %d alerts retrieved\n", len(alerts))

	// Step 1: Save alerts to results.json
	result := ScanResult{
		TargetURL: targetURL,
		ScanID:    scanID,
		Timestamp: time.Now().Format(time.RFC3339),
		Alerts:    alerts,
	}
	if err := saveResults(result); err != nil {
		return fmt.Errorf("❌ Failed to save results.json: %v", err)
	}
	fmt.Println("📝 Results saved to reports/results.json")

	// Step 2: Generate HTML report
	fmt.Println("📄 Generating HTML report...")
	if err := reports.GenerateHTMLReport("reports/results.json"); err != nil {
		return fmt.Errorf("❌ HTML report generation failed: %v", err)
	}

	fmt.Println("✅ HTML report generated successfully.")
	return nil
}

// saveResults stores the JSON output of the scan
func saveResults(result ScanResult) error {
	if err := os.MkdirAll("reports", 0755); err != nil {
		return err
	}
	path := filepath.Join("reports", "results.json")
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}
