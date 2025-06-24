package utils

import (
	"fmt"

	"github.com/atotto/clipboard"
)

// CopyToClipboard copies the provided payload to the system clipboard
func CopyToClipboard(payload string) error {
	err := clipboard.WriteAll(payload)
	if err != nil {
		return fmt.Errorf("clipboard error: %v", err)
	}
	return nil
}
