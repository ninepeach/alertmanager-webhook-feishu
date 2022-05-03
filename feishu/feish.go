package feishu

import (
	"bytes"
	"fmt"
	"github.com/icza/gox/stringsx"
	"github.com/magicst0ne/alertmanager-webhook-feishu/model"
	"github.com/magicst0ne/alertmanager-webhook-feishu/tmpl"
	"strings"
	"text/template"
)

type Feishu struct {
	webhook  string
	sdk      *Sdk
	tpl      *template.Template
	alertTpl *template.Template
}

func NewFeishu(fsurl string) (*Feishu, error) {

	// template
	tpl, err := tmpl.GetEmbedTemplate("default.tmpl")
	alertTpl, err := tmpl.GetEmbedTemplate("default_alert.tmpl")

	if err != nil {
		return nil, err
	}

	return &Feishu{
		webhook:  fsurl,
		sdk:      NewSDK("", ""),
		tpl:      tpl,
		alertTpl: alertTpl,
	}, nil
}

func (b Feishu) Send(alerts *model.WebhookMessage) error {

	// prepare data
	err := b.preprocessAlerts(alerts)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = b.tpl.Execute(&buf, alerts)
	if err != nil {
		return err
	}
	//fmt.Println(buf.String())

	return b.sdk.WebhookV2(b.webhook, &buf)
}

// right is immutable
func mergeMap(left, right map[string]string) map[string]string {
	if len(right) == 0 {
		return left
	}
	if left == nil {
		left = make(map[string]string)
	}
	for k, v := range right {
		if _, ok := left[k]; !ok {
			left[k] = v
		}
	}
	return left
}

// field description may contain double quote, non printable chars
func fixDescription(s string) string {
	// feishu fix: clean non printable char
	s = stringsx.Clean(s)
	// feishu fix: unescape a string
	s = fmt.Sprintf("%#v", s)
	// remove prefix and suffix double quote, means we just unescape inner text
	s = strings.TrimPrefix(s, "\"")
	s = strings.TrimSuffix(s, "\"")
	return s
}

func (b Feishu) preprocessAlerts(webhookMsg *model.WebhookMessage) error {

	// preprocess using alert template
	webhookMsg.Severity = webhookMsg.Alerts.Severity()

	n := 0
	for _, alert := range webhookMsg.Alerts.Firing() {
		webhookMsg.FiringAlerts = append(webhookMsg.FiringAlerts, alert)

		if _, ok1 := alert.Labels["hostname"]; ok1 {
			if _, ok2 := webhookMsg.AlertHosts[alert.Labels["hostname"]]; ok2 {
				if _, ok3 := alert.Labels["instance"]; ok3 {
					webhookMsg.AlertHosts[alert.Labels["hostname"]] = alert.Labels["instance"]
				}
			}
		}

		n++
	}
	webhookMsg.FiringNum = n

	for _, alert := range webhookMsg.Alerts.Resolved() {
		webhookMsg.ResolvedAlerts = append(webhookMsg.ResolvedAlerts, alert)

	}

	return nil
}
