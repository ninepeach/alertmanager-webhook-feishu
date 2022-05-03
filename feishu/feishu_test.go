package feishu

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/magicst0ne/alertmanager-webhook-feishu/model"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestFeishu_Send(t *testing.T) {
	//return
	feishu, err := NewFeishu("https://open.feishu.cn/open-apis/bot/v2/hook/de30801b-f747-48e2-99ce-39d8dd8d2eaf")
	require.Nil(t, err)

	//alerts := model.WebhookMessage{
	//	Data: newAlerts(),
	//	Meta: map[string]string{"group": "hello", "url": "www.baidu.com"},
	//}
	webhookMessage := model.WebhookMessage{AlertMessage: newAlerts()}
	webhookMessage.AlertHosts = make(map[string]string)

	err = feishu.Send(&webhookMessage)
	spew.Dump(err)
	require.Nil(t, err)
}

func newAlerts() model.AlertMessage {
	return model.AlertMessage{
		Alerts: model.Alerts{
			model.Alert{
				Status: "firing",
				Annotations: map[string]string{
					"alertmsg":    "Disk free < 15%",
					"description": "Disk is almost full (< 10% left)\n  VALUE = 31.05208958154261\n  LABELS = map[device:/dev/mapper/centos-root env:Dev fstype:xfs hostname:db-mysql-martintest02 instance:192.168.150.123:9100 job:MartinTest mountpoint:/ servicename:MySQL]",
					"summary":     "Free disk space is less than 5%  on volume /",
				},
				Labels: map[string]string{
					"alertname":   "HostOutOfDiskSpace",
					"device":      "/dev/mapper/centos-root",
					"env":         "Dev",
					"fstype":      "xfs",
					"hostname":    "db-mysql-martintest02",
					"instance":    "192.168.150.123:9100",
					"job":         "MartinTest",
					"mountpoint":  "/",
					"servicename": "MySQL",
					"severity":    "warning",
				},
				StartsAt:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				EndsAt:       time.Date(2000, 1, 1, 0, 0, 1, 0, time.UTC),
				GeneratorURL: "file://generatorUrl",
			},
			model.Alert{
				Status: "firing",
				Annotations: map[string]string{
					"alertmsg":    "Disk free < 15%",
					"description": "Disk is almost full (< 10% left)\n  VALUE = 87.7288276627219\n  LABELS = map[device:/dev/vda1 env:Dev fstype:xfs hostname:db-mysql-martintest02 instance:192.168.150.123:9100 job:MartinTest mountpoint:/boot servicename:MySQL]",
					"summary":     "Free disk space is less than 5%  on volume /boot",
				},
				Labels: map[string]string{
					"alertname":  "HostOutOfMemory3",
					"env":        "Dev",
					"hostname":   "db-mysql-martintest01",
					"instance":   "192.168.150.232:9100",
					"job":        "MartinTest",
					"servername": "MySQL",
					"severity":   "error",
				},
				StartsAt:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				EndsAt:       time.Date(2000, 1, 1, 0, 0, 1, 0, time.UTC),
				GeneratorURL: "file://generatorUrl",
			},
			model.Alert{
				Status: "firing",
				Annotations: map[string]string{
					"alertmsg":    "Disk free < 15%",
					"description": "Disk is almost full (< 10% left)\n  VALUE = 31.05208958154261\n  LABELS = map[device:rootfs env:Dev fstype:rootfs hostname:db-mysql-martintest02 instance:192.168.150.123:9100 job:MartinTest mountpoint:/ servicename:MySQL]",
					"summary":     "Free disk space is less than 5%  on volume /",
				},
				Labels: map[string]string{
					"alertname":   "HostOutOfDiskSpace",
					"device":      "rootfs",
					"env":         "Dev",
					"fstype":      "rootfs",
					"hostname":    "db-mysql-martintest02",
					"instance":    "192.168.150.123:9100",
					"job":         "MartinTest",
					"mountpoint":  "/",
					"servicename": "MySQL",
					"severity":    "warning",
				},
				StartsAt:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				EndsAt:       time.Date(2000, 1, 1, 0, 0, 1, 0, time.UTC),
				GeneratorURL: "file://generatorUrl",
			},
			model.Alert{
				Status: "firing",
				Annotations: map[string]string{
					"alertmsg":    "Have Slow Query",
					"description": "192.168.150.232:9104 of job MartinTest Connection More Than 4",
					"summary":     " 192.168.150.232:9104 MySQL Connection More Than 4",
				},
				Labels: map[string]string{
					"alertname":   "MySQLConnectedMoreThan4",
					"env":         "Dev",
					"hostname":    "db-mysql-martintest03",
					"instance":    "192.168.150.232:9104",
					"job":         "MartinTest",
					"servicename": "MySQL",
					"severity":    "error",
				},
				StartsAt:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				EndsAt:       time.Date(2000, 1, 1, 0, 0, 1, 0, time.UTC),
				GeneratorURL: "file://generatorUrl",
			},
			model.Alert{
				Status: "firing",
				Annotations: map[string]string{
					"alertmsg":    "Disk free < 15%",
					"description": "Disk is almost full (< 10% left)\n  VALUE = 31.05208958154261\n  LABELS = map[device:/dev/mapper/centos-root env:Dev fstype:xfs hostname:db-mysql-martintest02 instance:192.168.150.123:9100 job:MartinTest mountpoint:/ servicename:MySQL]",
					"summary":     "Free disk space is less than 15%  on volume /",
				},
				Labels: map[string]string{
					"alertname":   "HostOutOfDiskSpace",
					"device":      "/dev/mapper/centos-root",
					"env":         "Dev",
					"fstype":      "xfs",
					"hostname":    "db-mysql-martintest03",
					"instance":    "192.168.21.126:9100",
					"job":         "MartinTest",
					"mountpoint":  "/",
					"servicename": "MySQL",
					"severity":    "warning",
				},
				StartsAt:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				EndsAt:       time.Date(2000, 1, 1, 0, 0, 1, 0, time.UTC),
				GeneratorURL: "file://generatorUrl",
			},
			model.Alert{
				Status: "firing",
				Annotations: map[string]string{
					"alertmsg":    "Disk free < 15%",
					"description": "Disk is almost full (< 10% left)\n  VALUE = 31.05208958154261\n  LABELS = map[device:/dev/mapper/centos-root env:Dev fstype:xfs hostname:db-mysql-martintest02 instance:192.168.150.123:9100 job:MartinTest mountpoint:/ servicename:MySQL]",
					"summary":     "Free disk space is less than 15%  on volume /",
				},
				Labels: map[string]string{
					"alertname":   "HostOutOfDiskSpace",
					"device":      "/dev/vda1",
					"env":         "Dev",
					"fstype":      "xfs",
					"hostname":    "db-mysql-martintest03",
					"instance":    "192.168.21.126:9100",
					"job":         "MartinTest",
					"mountpoint":  "/boot",
					"servicename": "MySQL",
					"severity":    "warning",
				},
				StartsAt:     time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				EndsAt:       time.Date(2000, 1, 1, 0, 0, 1, 0, time.UTC),
				GeneratorURL: "file://generatorUrl",
			},
		},
		CommonAnnotations: map[string]string{"ca_key": "ca_value"},
		CommonLabels: map[string]string{
			"env":         "Dev",
			"job":         "MartinTest",
			"servicename": "MySQL",
		},
		GroupLabels: map[string]string{
			"job": "MartinTest",
		},
		ExternalURL: "file://externalUrl",
		Receiver:    "test-receiver",
	}
}
