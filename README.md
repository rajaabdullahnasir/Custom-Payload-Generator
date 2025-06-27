# 🚀 Modular Payload Generator Tool (Go)

A **custom payload generation toolkit** for Ethical Hackers, Penetration Testers, and Security Researchers. This tool automates the creation of **XSS**, **SQL Injection**, and **Command Injection** payloads, supporting advanced **encoding**, **obfuscation**, and **reporting** techniques — all from the command line.

---

## 🧠 What This Tool Can Do

- ✅ Generate payloads for **XSS**, **SQLi**, and **Command Injection**
- ✅ Apply **Base64, URL, Hex, Unicode** encodings
- ✅ Add **WAF bypass**, random case, comments, homoglyphs, and more
- ✅ Save payloads as `.json`, `.txt`, or display in console
- ✅ Copy payloads to clipboard
- ✅ Modular codebase for easy expansion
- ✅ ZAP integration ready

---

## 📁 Project Structure

```
payload-generator-go/
│
├── main.go                      # Entry point with CLI logic
├── go.mod                       # Go module definition
│
├── /modules/                    # Core payload logic
│   ├── xss.go                   # Reflected, Stored, DOM XSS
│   ├── sqli.go                  # Error-based, Union, Blind SQLi
│   └── cmdinj.go                # Linux/Windows command injection
│
├── /utils/                      # Reusable helper tools
│   ├── encoder.go               # Base64, URL, Hex, Unicode encoders
│   ├── obfuscator.go            # Obfuscation strategies
│   ├── clipboard.go             # Clipboard handler (xclip/pbcopy)
│   └── output.go                # JSON, TXT, console output
│
├── /zapapi/                     # (Optional) OWASP ZAP integration
│   ├── zap_client.go
│   └── zap_tester.go
│
├── /payloads/                   # Input templates
│   ├── xss.json
│   ├── sqli.json
│   └── cmd.json
│
├── /reports/                    # Generated payloads & results
│   ├── results.json
│   ├── payloads.txt
│   └── report.go
│
└── README.md
```

---

## ⚙️ Installation & Setup

### 📦 Requirements

- Go 1.22+
- For Linux: `xclip` for clipboard functionality
- OWASP ZAP running in **daemon mode**

```bash
sudo apt install xclip
```

---

### 🔧 Build

```bash
git clone https://github.com/rajaabdullahnasir/Custom-Payload-Generator.git
cd Custom-Payload-Generator
go mod tidy
go build -o payloadgen
```

---

## 🚀 Usage

### Basic Help

```bash
./payloadgen --help
```

### Generate XSS Payloads

```bash
./payloadgen --xss
./payloadgen --xss --output=json --save
./payloadgen --xss --output=txt --clipboard
```

### Generate SQLi Payloads

```bash
./payloadgen --sqli
./payloadgen --sqli --output=txt --save
```

### Generate Command Injection Payloads

```bash
./payloadgen --cmdi
./payloadgen --cmdi --output=json --save
```

## Perform ZAP Scan (Active Vulnerability Assessment)

```bash
./payloadgen --zapscan --target=http://testasp.vulnweb.com --zap-key=YOUR_ZAP_API_KEY
```

---

## ✨ Features

| Module        | Details |
|---------------|---------|
| 🔍 XSS         | Reflected, Stored, DOM | 
| 💉 SQLi        | Error, Union, Blind, WAF bypass |
| 💣 CMDi        | Linux/Windows OS commands |
| 🔐 Encoding    | Base64, URL, Hex, Unicode |
| 🎭 Obfuscation | Spacing, comments, homoglyphs |
| 📋 Clipboard   | Copy directly for instant testing |
| 📤 Output      | Console, JSON, TXT |
| 🔎 ZAP Ready   | Integrate with ZAP |

---

## 💡 Examples

```bash
./payloadgen --xss --output=json --save
./payloadgen --sqli --output=txt
./payloadgen --cmdi --clipboard
./payloadgen --zapscan --target=http://testphp.vulnweb.com --zap-key=your_zap_key
```

---

## 🔐 Disclaimer

This tool is for **educational** and **ethical** use only. Do **not** use it against systems you don’t own or have permission to test.

---

## 📜 License

MIT License — feel free to use and contribute.
