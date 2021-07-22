package utils

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"github.com/alecthomas/template"
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

const ruleTemplate = `  - alert: '{{ Key . }}'
    expr: {{.Expr}}
    for: {{.For}}
    labels:
      severity: {{.Severity}}
    annotations:
      group: '{{.Group}}'
      desc: '{{.Desc}}'
      summary: '{{.Summary}}'
      value: {{"'{{$value}}'"}}
`

const RecordingTemplate = `  - record: {{ .Record }}
    expr: {{ .Expr }}
`

type APIRule struct {
	Id     string `json:"id" bson:"_id" yaml:"id"`
	IdStr  string        `json:"id_str" yaml:"id_str"`
	Group  string        `json:"group" yaml:"group"`
	Record string        `json:"record" yaml:"record"`
	// Alert          string              `json:"alert"`
	Metric   string `json:"metric" yaml:"metric"`
	Expr     string `json:"expr" yaml:"expr"`
	For      string `json:"for" yaml:"for"`
	Severity string `json:"severity" yaml:"severity"`
	Desc     string `json:"desc" yaml:"desc"`
	Summary  string `json:"summary" yaml:"summary"`
	// Object         string              `json:"object" yaml:"object"`
	// Left           string                 `json:"left" yaml:"left"`
	Right          string                 `json:"right" yaml:"right"`
	Operator       string                 `json:"operator" yaml:"operator"`
	Repeat         bool                   `json:"repeat" yaml:"repeat"`
	RepeatInterval string                 `json:"repeatInterval" yaml:"repeat_interval"`
	TagValues      []map[string]string    `json:"tagValues" yaml:"tag_values"`
	EnvType        string                 `json:"env_type" yaml:"env_type"`
	Envs           []string               `json:"envs" yaml:"envs"`
	Start          string                 `json:"start" yaml:"start"`
	End            string                 `json:"end" yaml:"end"`
	Disable        bool                   `json:"disable" yaml:"disable"`
	UnixTime       int64                  `json:"unix_time" yaml:"unix_time" bson:"unix_time"`
	SendResolve    bool                   `json:"send_resolve" yaml:"send_resolve" bson:"send_resolve"`
	Extra          map[string]interface{} `json:"extra" yaml:"extra"`
	AlertGroupIDs  []string               `json:"alert_group_ids" yaml:"alert_group_ids" bson:"alert_group_ids"`
	AlertGroups    []AlertGroup           `json:"alert_groups" yaml:"alert_groups" bson:"-"`
	CreateTime     string                 `json:"create_time" yaml:"create_time" bson:"create_time"`
	UpdateTime     string                 `json:"update_time" yaml:"update_time" bson:"update_time"`
	ResolveTimeout string                 `json:"resolve_timeout" yaml:"resolve_timeout" bson:"resolve_timeout"`
}

func (a *APIRule) GetKey() string {
	return a.Id + "_" + a.Metric
}

type Recording struct {
	Id      string `json:"id" bson:"_id"`
	IdStr   string        `json:"id_str"`
	Envs    []string      `json:"envs"`
	Desc    string        `json:"desc"`
	Record  string        `json:"record"`
	Expr    string        `json:"expr"`
	Disable bool          `json:"disable"`
}

func (r *Recording) GetKey() string {
	return r.Id + "_" + r.Record
}

// AlertGroup 告警组
type AlertGroup struct {
	ID         string          `json:"id" bson:"_id" yaml:"id"`
	Name       string                 `json:"name" yaml:"name"`
	UserIDS    []string               `json:"user_ids" yaml:"user_ids" bson:"user_ids"`
	Users      []AlertUser            `json:"users" yaml:"users" bson:"-"`
	Extra      map[string]interface{} `json:"extra" yaml:"extra"`
	Preview    map[string]string      `json:"preview" yaml:"preview" bson:"-"`
	CreateTime string                 `json:"create_time" yaml:"create_time" bson:"create_time"`
	UpdateTime string                 `json:"update_time" yaml:"update_time" bson:"update_time"`
}

// AlertUser 告警联系人
type AlertUser struct {
	ID           string          `json:"id" bson:"_id" yaml:"id"`
	Name         string                 `json:"name" yaml:"name"`
	MailEnable   bool                   `json:"mail_enable" yaml:"mail_enable" bson:"mail_enable"`
	Mail         string                 `json:"mail" yaml:"mail"`
	WechatEnable bool                   `json:"wechat_enable" yaml:"wechat_enable" bson:"wechat_enable"`
	Wechat       string                 `json:"wechat" yaml:"wechat"`
	WechatType   string                 `json:"wechat_type" yaml:"wechat_type"`
	SmsEnable    bool                   `json:"sms_enable" yaml:"sms_enable" bson:"sms_enable"`
	Mobile       string                 `json:"mobile" yaml:"mobile"`
	Extra        map[string]interface{} `json:"extra" yaml:"extra"`
	CreateTime   string                 `json:"create_time" yaml:"create_time" bson:"create_time"`
	UpdateTime   string                 `json:"update_time" yaml:"update_time" bson:"update_time"`
}

func getRuleTemplate(rule *APIRule) (string, error) {
	t := template.New("rule").Funcs(template.FuncMap{
		"Key": func(rule APIRule) string {
			return rule.GetKey()
		},
		// "Object": func(rule apptypes.APIRule) string {
		// 	return objectCache[rule.Object].Template
		// },
	})
	t, err := t.Parse(ruleTemplate)
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

func getRecordingTemplate(r *Recording) (string, error) {
	t := template.New("recording")
	t, err := t.Parse(RecordingTemplate)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	err = t.Execute(&b, r)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}