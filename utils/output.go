package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// GenericPayload is a common struct to export any payload
type GenericPayload struct {
	Type    string `json:"type,omitempty"`
	Payload string `json:"payload"`
	Encoded string `json:"encoded,omitempty"`
	Bypass  bool   `json:"bypass,omitempty"`
}

// ExportToJSON writes the list of payloads to a JSON file
func ExportToJSON(filename string, payloads []GenericPayload) error {
	data, err := json.MarshalIndent(payloads, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal payloads: %v", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write JSON file: %v", err)
	}

	fmt.Printf("✅ Payloads saved to %s\n", filename)
	return nil
}

// ExportToText writes only raw payloads to a text file (one per line)
func ExportToText(filename string, payloads []GenericPayload) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create text file: %v", err)
	}
	defer file.Close()

	for _, p := range payloads {
		_, err := file.WriteString(p.Payload + "\n")
		if err != nil {
			return fmt.Errorf("failed to write payload to file: %v", err)
		}
	}

	fmt.Printf("✅ Payloads saved to %s\n", filename)
	return nil
}
