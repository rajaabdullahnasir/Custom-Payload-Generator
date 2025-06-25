package reports

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"time"
) 

// Alert defines the structure of a ZAP scan alert
type Alert struct {
	Alert   string `json:"alert"`
	Name    string `json:"name"`
	Risk    string `json:"risk"`
	Desc    string `json:"description"`
	Solution string `json:"solution"`
	Param   string `json:"param"`
	Evidence string `json:"evidence"`
	URL     string `json:"url"`
}

// ScanResult is the full structure of scan result
type ScanResult struct {
	TargetURL string                   `json:"target_url"`
	ScanID    string                   `json:"scan_id"`
	Timestamp string                   `json:"timestamp"`
	Alerts    []map[string]interface{} `json:"alerts"` // raw alert objects
}

// HTMLReportData defines the structure for templating
type HTMLReportData struct {
	Title      string
	Date       string
	TargetURL  string
	ScanID     string
	AlertCount int
	Alerts     []Alert
}

// GenerateHTMLReport converts scan result into a beautiful HTML report
func GenerateHTMLReport(scanPath string) error {
	file, err := os.ReadFile(scanPath)
	if err != nil {
		return fmt.Errorf("failed to read scan result: %v", err)
	}

	var result ScanResult
	if err := json.Unmarshal(file, &result); err != nil {
		return fmt.Errorf("failed to parse scan result: %v", err)
	}

	// Parse raw alerts into structured list
	var alerts []Alert
	for _, a := range result.Alerts {
		alert := Alert{
			Alert:    toString(a["alert"]),
			Name:     toString(a["name"]),
			Risk:     toString(a["risk"]),
			Desc:     toString(a["description"]),
			Solution: toString(a["solution"]),
			Param:    toString(a["param"]),
			Evidence: toString(a["evidence"]),
			URL:      toString(a["url"]),
		}
		alerts = append(alerts, alert)
	}

	data := HTMLReportData{
		Title:      "Automated Vulnerability Assessment Report",
		Date:       time.Now().Format("02-Jan-2006 15:04:05"),
		TargetURL:  result.TargetURL,
		ScanID:     result.ScanID,
		AlertCount: len(alerts),
		Alerts:     alerts,
	}

	tmpl, err := template.New("report").Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse HTML template: %v", err)
	}

	// Create reports directory if not exists
	_ = os.MkdirAll("reports", 0755)
	fileName := fmt.Sprintf("reports/report_%s.html", time.Now().Format("20060102_150405"))
	fileOut, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create report file: %v", err)
	}
	defer fileOut.Close()

	err = tmpl.Execute(fileOut, data)
	if err != nil {
		return fmt.Errorf("failed to write HTML report: %v", err)
	}

	fmt.Println("✅ HTML report generated:", fileName)
	return nil
}

// toString safely casts interface to string
func toString(i interface{}) string {
	if str, ok := i.(string); ok {
		return str
	}
	return ""
}

// HTML template string
const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>{{ .Title }}</title>
  <style>
    body {
      font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
      margin: 40px;
      color: #333;
      background-color: #f4f4f4;
    }
    h1, h2, h3 { color: #1a1a1a; }
    .header {
      text-align: center;
      background-color: #222;
      color: #fff;
      padding: 40px 0;
    }
    .section {
      margin-top: 40px;
      background: #fff;
      padding: 30px;
      border-radius: 8px;
      box-shadow: 0 0 10px rgba(0,0,0,0.1);
    }
    .alert {
      border-left: 5px solid #c0392b;
      padding-left: 15px;
      margin-bottom: 20px;
    }
    .risk-High { color: red; }
    .risk-Medium { color: orange; }
    .risk-Low { color: green; }
    .risk-Informational { color: blue; }
    .chart {
      width: 100%;
      max-width: 600px;
      height: 300px;
      margin: auto;
    }
    table {
      width: 100%;
      border-collapse: collapse;
      margin-top: 10px;
    }
    th, td {
      border: 1px solid #ccc;
      padding: 12px;
      text-align: left;
    }
    th {
      background: #eaeaea;
    }
  </style>
</head>
<body>

  <div class="header">
    <h1>{{ .Title }}</h1>
    <p><strong>Date:</strong> {{ .Date }}</p>
    <p><strong>Target:</strong> {{ .TargetURL }}</p>
    <p><strong>Scan ID:</strong> {{ .ScanID }}</p>
  </div>

  <div class="section">
    <h2>Table of Contents</h2>
    <ul>
      <li>1. Executive Summary</li>
      <li>2. Methodology</li>
      <li>3. Vulnerability Findings</li>
      <li>4. Graphical Summary</li>
    </ul>
  </div>

  <div class="section">
    <h2>1. Executive Summary</h2>
    <p>This report presents the results of an automated security assessment conducted using the Modular Payload Generator Tool and OWASP ZAP. The purpose of this assessment was to identify potential vulnerabilities such as Cross-Site Scripting (XSS), SQL Injection (SQLi), and Command Injection in the specified web target.</p>
  </div>

  <div class="section">
    <h2>2. Methodology</h2>
    <p>The assessment process includes:</p>
    <ul>
      <li>Payload generation (XSS, SQLi, CMDi)</li>
      <li>Encoding and WAF bypass techniques</li>
      <li>Automatic injection of payloads into URL/query/form parameters</li>
      <li>Active scan using OWASP ZAP’s daemon API</li>
      <li>Collection and analysis of vulnerabilities</li>
    </ul>
  </div>

  <div class="section">
    <h2>3. Vulnerability Findings ({{ .AlertCount }} Total)</h2>
    {{range .Alerts}}
    <div class="alert">
      <h3>{{.Alert}}</h3>
      <p><strong>Risk:</strong> <span class="risk-{{.Risk}}">{{.Risk}}</span></p>
      <p><strong>Parameter:</strong> {{.Param}}</p>
      <p><strong>Evidence:</strong> {{.Evidence}}</p>
      <p><strong>Description:</strong> {{.Desc}}</p>
      <p><strong>Solution:</strong> {{.Solution}}</p>
      <p><strong>URL:</strong> {{.URL}}</p>
    </div>
    {{end}}
  </div>

  <div class="section">
    <h2>4. Graphical Summary</h2>
    <canvas id="vulnChart" class="chart"></canvas>
  </div>

  <script>
    const ctx = document.getElementById('vulnChart').getContext('2d');
    const data = {
      labels: ['High', 'Medium', 'Low', 'Informational'],
      datasets: [{
        label: 'Vulnerabilities by Risk Level',
        backgroundColor: ['#e74c3c','#f39c12','#2ecc71','#3498db'],
        data: [
          {{countRisk .Alerts "High"}},
          {{countRisk .Alerts "Medium"}},
          {{countRisk .Alerts "Low"}},
          {{countRisk .Alerts "Informational"}}
        ]
      }]
    };
    const config = {
      type: 'bar',
      data: data,
    };
    new Chart(ctx, config);
  </script>
  <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>

</body>
</html>
`
