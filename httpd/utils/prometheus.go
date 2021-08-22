package utils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/alecthomas/template"
	"github.com/hashicorp/consul/api"
	"net/http"
	"portal/global/consul"
)

func PrometheusReload(url string) error {
	reloadUrl := fmt.Sprintf("%s%s", url, "/-/reload")
	resp, err := http.Post(reloadUrl,"application/json",nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("reload prometheus error")
	}
	return nil
}

const AlertingRuleTemplate = `  - alert: '{{ .Alert }}'
    expr: {{.Expr}}
    for: {{.For}}
    labels:
      severity: {{.Severity}}
    annotations:
      summary: '{{.Summary}}'
      description: '{{.Description}}'
`

const RecordingRuleTemplate = `  - record: {{ .Record }}
    expr: {{ .Expr }}
`

type RecordingRule struct {
	Record  string        `json:"record"`
	Expr    string        `json:"expr"`
}

type AlertingRule struct {
	Alert     string `json:"alert"`
	Expr  string `json:"expr"`
	For  string `json:"for"`
	Severity    string `json:"severity"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

func GetRecordingRuleTemplate(rule *RecordingRule) (string, error) {
	t := template.New("recording_rule")
	t, err := t.Parse(RecordingRuleTemplate)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	err = t.Execute(&b, rule)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func GetAlertingRuleTemplate(rule *AlertingRule) (string, error) {
	t := template.New("alerting_rule")
	t, err := t.Parse(AlertingRuleTemplate)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	err = t.Execute(&b, rule)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func PutConsul(key string, value string) error {
	client := consul.GetClient()
	kv := client.KV()
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err := kv.Put(p, nil)
	if err != nil {
		return err
	}
	return nil
}