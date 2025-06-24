package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rajaabdullahnasir/Custom-Payload-Generator/modules"
	"github.com/rajaabdullahnasir/Custom-Payload-Generator/utils"
)

var helpText = `
██████╗ ██████╗ ██████╗ ███████╗██████╗  ██████╗  ██████╗  ██████╗ 
██╔══██╗██╔═══██╗██╔══██╗██╔════╝██╔══██╗██╔═══██╗██╔═══██╗██╔════╝ 
██████╔╝██║   ██║██████╔╝█████╗  ██████╔╝██║   ██║██║   ██║██║  ███╗
██╔═══╝ ██║   ██║██╔═══╝ ██╔══╝  ██╔═══╝ ██║   ██║██║   ██║██║   ██║
██║     ╚██████╔╝██║     ███████╗██║     ╚██████╔╝╚██████╔╝╚██████╔╝
╚═╝      ╚═════╝ ╚═╝     ╚══════╝╚═╝      ╚═════╝  ╚═════╝  ╚═════╝ 

Modular Payload Generator Tool by @rajaabdullahnasir

USAGE:
  ./payloadgen [--xss | --sqli | --cmdi] [--output=json|txt|console] [--save] [--clipboard]

FLAGS:
  --xss           Generate XSS payloads
  --sqli          Generate SQL Injection payloads
  --cmdi          Generate Command Injection payloads
  --output        Set output format: console (default), json, txt
  --save          Save output to ./reports/
  --clipboard     Copy output to clipboard
  --help          Show this help menu

EXAMPLES:
  ./payloadgen --xss --output=json 
  ./payloadgen --cmdi --output=txt 
  ./payloadgen --sqli

  Enjoy hacking ethically! 🔐
`

func main() {
	// Define flags
	xss := flag.Bool("xss", false, "Generate XSS payloads")
	sqli := flag.Bool("sqli", false, "Generate SQLi payloads")
	cmdi := flag.Bool("cmdi", false, "Generate CMDi payloads")

	output := flag.String("output", "console", "Output format: json, txt, console")
	save := flag.Bool("save", false, "Save to reports/")
	clip := flag.Bool("clipboard", false, "Copy payloads to clipboard")
	help := flag.Bool("help", false, "Show help menu")

	flag.Parse()

	// Show help if requested
	if *help || (!*xss && !*sqli && !*cmdi) {
		fmt.Println(helpText)
		return
	}

	if *xss {
		payloads, err := modules.GenerateXSSPayloads()
		if err != nil {
			log.Fatalf("❌ Failed to generate XSS payloads: %v", err)
		}
		handleOutput("xss_payloads", payloads, *output, *save, *clip)
	}

	if *sqli {
		payloads, err := modules.GenerateSQLiPayloads()
		if err != nil {
			log.Fatalf("❌ Failed to generate SQLi payloads: %v", err)
		}
		handleOutput("sqli_payloads", payloads, *output, *save, *clip)
	}

	if *cmdi {
		payloads := modules.GenerateCMDiPayloads()
		handleOutput("cmdi_payloads", payloads, *output, *save, *clip)
	}
}

func handleOutput(name string, payloads interface{}, format string, save bool, clip bool) {
	switch format {
	case "json":
		if save {
			err := utils.SaveAsJSON(payloads, name)
			if err != nil {
				log.Printf("⚠️ Could not save JSON: %v", err)
			} else {
				fmt.Printf("✅ Saved %s.json in /reports/\n", name)
			}
		} else {
			utils.PrintToConsole(name, payloads)
		}
	case "txt":
		lines := flattenPayloads(payloads)
		if save {
			err := utils.SaveAsTXT(lines, name)
			if err != nil {
				log.Printf("⚠️ Could not save TXT: %v", err)
			} else {
				fmt.Printf("✅ Saved %s.txt in /reports/\n", name)
			}
		} else {
			utils.PrintToConsole(name, lines)
		}
	case "console":
		utils.PrintToConsole(name, payloads)
	default:
		fmt.Println("❌ Invalid output format. Use json, txt, or console.")
		os.Exit(1)
	}

	if clip {
		lines := flattenPayloads(payloads)
		if len(lines) > 0 {
			err := utils.CopyToClipboard(lines[0])
			if err != nil {
				log.Printf("⚠️ Could not copy to clipboard: %v", err)
			} else {
				fmt.Println("📋 First payload copied to clipboard!")
			}
		}
	}
}

func flattenPayloads(data interface{}) []string {
	var lines []string
	switch v := data.(type) {
	case []modules.XSSPayload:
		for _, p := range v {
			lines = append(lines, p.Payload)
		}
	case []modules.SQLiPayload:
		for _, p := range v {
			lines = append(lines, p.Payload)
		}
	case []modules.CMDPayload:
		for _, p := range v {
			lines = append(lines, p.Original)
		}
	default:
		lines = append(lines, fmt.Sprintf("%v", v))
	}
	return lines
}
