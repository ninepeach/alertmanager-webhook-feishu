package main

import (
	"encoding/json"
	"fmt"
	alog "github.com/apex/log"
	"github.com/gin-gonic/gin"
	"github.com/magicst0ne/alertmanager-webhook-feishu/feishu"
	"github.com/magicst0ne/alertmanager-webhook-feishu/model"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
)

var (
	Version       string
	BuildRevision string
	BuildBranch   string
	BuildTime     string
	BuildHost     string
	rootLoggerCtx *alog.Entry

	configFile = kingpin.Flag(
		"config.file",
		"configuration file path.",
	).Short('c').Default("config.yml").String()

	serverPort = kingpin.Flag(
		"web.listen-address",
		"Address to listen on",
	).Short('p').Default(":8086").String()

	sc = &SafeConfig{
		C: &Config{},
	}
	reloadCh chan chan error
)

func init() {
	rootLoggerCtx = alog.WithFields(alog.Fields{
		"app": "prometheus-webhook-feish",
	})

	hostname, _ := os.Hostname()
	rootLoggerCtx.Infof("version %s, build reversion %s, build branch %s, build at %s on host %s", Version, BuildRevision, BuildBranch, BuildTime, hostname)
}

func main() {
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	configLoggerCtx := rootLoggerCtx.WithField("config", *configFile)
	configLoggerCtx.Info("starting app")

	// load config  first time
	if err := sc.ReloadConfig(*configFile); err != nil {
		configLoggerCtx.WithError(err).Error("error parsing config file")
		panic(err)
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/webhook", func(c *gin.Context) {

		body_raw, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			rootLoggerCtx.Infof("Read Request Body Error: %v", err.Error())
		}

		var alertMsg model.AlertMessage
		err = json.Unmarshal([]byte(body_raw), &alertMsg)
		if err != nil {
			rootLoggerCtx.Infof("Parser WebhookMessage Error: %v", err.Error())
			c.JSON(200, gin.H{"ret": "-1", "msg": "invalid data"})
		} else {

			access_token := c.Query("access_token")
			receiver := c.Query("receiver")

			receiverConfig, err := sc.GetConfigByName(receiver)
			if err != nil {
				c.JSON(200, gin.H{"ret": "-1", "msg": "receiver not exists"})
				return
			}

			if access_token != receiverConfig.AccessToken {
				c.JSON(200, gin.H{"ret": "-1", "msg": "invaild access_token"})
				return
			}

			fmt.Println(receiverConfig.Fsurl, "https://open.feishu.cn/open-apis/bot/v2/hook/de30801b-f747-48e2-99ce-39d8dd8d2eaf")
			fshu, _ := feishu.NewFeishu(receiverConfig.Fsurl)

			webhookMessage := model.WebhookMessage{AlertMessage: alertMsg}
			webhookMessage.AlertHosts = make(map[string]string)
			err = fshu.Send(&webhookMessage)

			c.JSON(200, gin.H{"ret": "0", "msg": "ok"})
		}

		//From Alertmanager
		//jobName := alerts.Alerts[0].Labels["job"]

		/*
				//From config
				receiverTarget := config.Target{}
			OuterLoop:
				for _, value := range conf.Targets {
					if value.JOB == jobName {
						receiverTarget = value
						break OuterLoop
					} else {
						receiverTarget = conf.Targets["receiver_default"]
					}
				}
		*/

		//bot, _ := feishu.NewBot(receiverTarget)
		//err = bot.Send(&alerts)
		//if err == nil {
		//	c.JSON(200, gin.H{"ret": "0", "msg": "ok"})
		//} else {
		//	c.JSON(200, gin.H{"ret": "-1", "msg": err.Error()})
		//}
	})

	r.Run(*serverPort)
}
