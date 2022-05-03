package tmpl

import (
	"fmt"
	"github.com/magicst0ne/alertmanager-webhook-feishu/model"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestFeishuCard(t *testing.T) {
	alerts := model.WebhookMessage{Data: newAlerts()}

	tpl := embedTemplates["default.tmpl"]
	tpl.Execute(os.Stdout, alerts)

	err := tpl.Execute(os.Stdout, alerts)
	fmt.Println(err)
	require.Nil(t, err)
}

func newAlerts() types.Data {
	return types.Data{
		Alerts: types.Alerts{
			types.Alert{
				Status:       "firing",
				Annotations:  map[string]string{"a_key": "a_value"},
				Labels:       map[string]string{"l_key": "l_value"},
				StartsAt:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				EndsAt:       time.Date(2000, 1, 1, 0, 0, 1, 0, time.UTC),
				GeneratorURL: "file://generatorUrl",
			},
			types.Alert{
				Annotations: map[string]string{"a_key_warn": "a_value_warn"},
				Labels:      map[string]string{"l_key_warn": "l_value_warn"},
				Status:      "warning",
			},
		},
		CommonAnnotations: map[string]string{"ca_key": "ca_value"},
		CommonLabels:      map[string]string{"cl_key": "cl_value"},
		GroupLabels:       map[string]string{"gl_key": "gl_value"},
		ExternalURL:       "file://externalUrl",
		Receiver:          "test-receiver",
	}
}
