package modules

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rajaabdullahnasir/Custom-Payload-Generator/utils"
)

// SQLiPayload represents a SQL injection payload and its encoded version
type SQLiPayload struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Bypass  bool   `json:"bypass"`
	Encoded string `json:"encoded,omitempty"`
}

// GenerateSQLiPayloads reads payloads from file and encodes them
func GenerateSQLiPayloads() ([]SQLiPayload, error) {
	payloads, err := loadSQLiPayloads("payloads/sqli.json")
	if err != nil {
		return nil, fmt.Errorf("load error: %w", err)
	}

	for i := range payloads {
		payloads[i].Encoded = utils.EncodeURL(payloads[i].Payload)
	}

	return payloads, nil
}

// loadSQLiPayloads loads payloads from a given JSON file path
func loadSQLiPayloads(path string) ([]SQLiPayload, error) {
	fullPath := filepath.Clean(path)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", fullPath, err)
	}

	var payloads []SQLiPayload
	if err := json.Unmarshal(data, &payloads); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return payloads, nil
}
