package modules

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/rajaabdullahnasir/Custom-Payload-Generator/utils"
)

// CMDiPayload represents a command injection payload with various encodings.
type CMDiPayload struct {
	OS         string `json:"os"`
	Command    string `json:"command"`
	Operator   string `json:"operator"`
	Original   string `json:"original"`
	Base64     string `json:"base64"`
	URLEncoded string `json:"url_encoded"`
	HexEncoded string `json:"hex_encoded"`
	Unicode    string `json:"unicode_escaped"`
	Obfuscated string `json:"obfuscated"`
}

// GenerateCMDiPayloads builds encoded & obfuscated payloads based on OS, command, and operators.
func GenerateCMDiPayloads(osType, command string, operators []string) []CMDiPayload {
	payloads := make([]CMDiPayload, 0, len(operators))

	for _, op := range operators {
		full := op + command
		payloads = append(payloads, CMDiPayload{
			OS:         osType,
			Command:    command,
			Operator:   op,
			Original:   full,
			Base64:     encodeBase64(full),
			URLEncoded: encodeURL(full),
			HexEncoded: utils.EncodeHex(full),
			Unicode:    utils.EncodeUnicode(full),
			Obfuscated: utils.Obfuscate(full),
		})
	}

	return payloads
}

// ExportCMDiPayloadsToJSON saves the payloads as a pretty-printed JSON file.
func ExportCMDiPayloadsToJSON(payloads []CMDiPayload) error {
	data, err := json.MarshalIndent(payloads, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal CMDi payloads: %w", err)
	}

	outputPath := filepath.Clean("reports/payloads_cmdi.json")
	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write CMDi payloads to file: %w", err)
	}

	return nil
}

// Internal encoding helpers
func encodeBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func encodeURL(input string) string {
	return url.QueryEscape(input)
}
