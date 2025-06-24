// main.go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rajaabdullahnasir/Custom-Payload-Generator/modules"
)

func main() {
	// Create report files
	jsonFile, err := os.Create("reports/results.json")
	if err != nil {
		fmt.Println("❌ Failed to create results.json:", err)
		return
	}
	defer jsonFile.Close()

	txtFile, err := os.Create("reports/payloads.txt")
	if err != nil {
		fmt.Println("❌ Failed to create payloads.txt:", err)
		return
	}
	defer txtFile.Close()

	allResults := make(map[string]interface{})

	// =======================
	// SQLi Payloads
	// =======================
	sqlPayloads, err := modules.GenerateSQLiPayloads()
	writeSectionHeader(txtFile, "🔍 SQL Injection Payloads")
	if err != nil {
		fmt.Fprintf(txtFile, "❌ Failed to load SQLi payloads: %v\n", err)
	} else {
		allResults["sqli"] = sqlPayloads
		for _, p := range sqlPayloads {
			writePayloadToReport(txtFile, map[string]string{
				"Type":    p.Type,
				"Bypass":  fmt.Sprint(p.Bypass),
				"Payload": p.Payload,
				"Encoded": p.Encoded,
			})
		}
	}

	// =======================
	// XSS Payloads
	// =======================
	xssPayloads, err := modules.GenerateXSSPayloads()
	writeSectionHeader(txtFile, "🚨 XSS Payloads")
	if err != nil {
		fmt.Fprintf(txtFile, "❌ Failed to load XSS payloads: %v\n", err)
	} else {
		allResults["xss"] = xssPayloads
		for _, p := range xssPayloads {
			writePayloadToReport(txtFile, map[string]string{
				"Type":    p.Type,
				"Bypass":  fmt.Sprint(p.Bypass),
				"Payload": p.Payload,
				"Encoded": p.Encoded,
			})
		}
	}

	// =======================
	// CMD Injection Payloads
	// =======================
	cmdPayloads := modules.GenerateCMDiPayloads("linux", "whoami", []string{";", "&&", "|"})
	writeSectionHeader(txtFile, "💣 Command Injection Payloads (Linux)")
	allResults["cmd"] = cmdPayloads
	for _, p := range cmdPayloads {
		writePayloadToReport(txtFile, map[string]string{
			"OS":         p.OS,
			"Command":    p.Command,
			"Operator":   p.Operator,
			"Original":   p.Original,
			"Base64":     p.Base64,
			"URLEncoded": p.URLEncoded,
			"HexEncoded": p.HexEncoded,
			"Unicode":    p.Unicode,
			"Obfuscated": p.Obfuscated,
		})
	}

	// Write JSON Report
	jsonData, err := json.MarshalIndent(allResults, "", "  ")
	if err != nil {
		fmt.Println("❌ Failed to marshal results to JSON:", err)
		return
	}
	if _, err := jsonFile.Write(jsonData); err != nil {
		fmt.Println("❌ Failed to write to results.json:", err)
		return
	}

	fmt.Println("\n✅ All payloads written to reports/results.json and reports/payloads.txt")
}

func writeSectionHeader(f *os.File, title string) {
	fmt.Fprintln(f, "\n====================================")
	fmt.Fprintln(f, title)
	fmt.Fprintln(f, "====================================")
}

func writePayloadToReport(f *os.File, data map[string]string) {
	for k, v := range data {
		fmt.Fprintf(f, "%s: %s\n", k, v)
	}
	fmt.Fprintln(f, "------------------------------------")
}
