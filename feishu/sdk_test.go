package feishu

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func ProcessWebhookMessage() {

}

func TestSdk_WebhookV2(t *testing.T) {
	sdk := NewSDK("", "")

	var buf bytes.Buffer

	data := `{
	"card": {
		"header": {
			"template": "red",
			"title": {
				"i18n": {
					"en_us": "title",
					"zh_cn": "title"
				},
				"tag": "plain_text"
			}
		},
		"i18n_elements": {
			"en_us": [
				{
					"tag": "div",
					"text": {
						"content": "test\n",
						"tag": "lark_md"
					}
				},
				{
					"tag": "hr"
				},
				{
					"elements": [
						{
							"content": "**Please fix it asap**\t",
							"tag": "lark_md"
						}
					],
					"tag": "note"
				}
			],
			"zh_cn": [
				{
					"tag": "div",
					"text": {
						"content": "test\n",
						"tag": "lark_md"
					}
				},
				{
					"tag": "hr"
				},
				{
					"elements": [
						{
							"content": "test\t",
							"tag": "lark_md"
						}
					],
					"tag": "note"
				}
			]
		}
	},
	"msg_type": "interactive"
}`

	buf.WriteString(data)
	err := sdk.WebhookV2("https://open.feishu.cn/open-apis/bot/v2/hook/de30801b-f747-48e2-99ce-39d8dd8d2eaf", &buf)
	require.Nil(t, err)
}
