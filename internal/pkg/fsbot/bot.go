package fsbot

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	"github.com/daniel-hutao/crbot/internal/pkg/ferry"
	"github.com/daniel-hutao/crbot/internal/pkg/log"
)

var webhook = "https://open.feishu.cn/open-apis/bot/v2/hook/39413a88-045b-4f47-85ab-f70b1913343c"

//var msgTpl = `{"msg_type":"text","content":{"text":"%s"}}`

var msgTpl = `
{
	"msg_type": "post",
	"content": {
		"post": {
			"zh_cn": {
				"title": ">>> Code Review Reminder <<<",
				"content": [
					[{{range .Datas}}
						{
							"tag": "text",
							"text": "{{.CreatedAt}} "
						},
						{
							"tag": "a",
							"text": "#{{.Number}}",
							"href": "{{.URL}}"
						},
						{
							"tag": "text",
							"text": " {{.Title}}\n\n"
						},{{end}}
						{
							"tag": "text",
							"text": "{{.Now}} now is."
						}
					]
				]
			}
		}
	}
}
`

//,
//{
//"tag": "at",
//"user_id": ""
//}

type Bot struct {
	Webhook string
	MsgTpl  string
	client  *http.Client
}

func NewBot() *Bot {
	return &Bot{
		Webhook: webhook,
		MsgTpl:  msgTpl,
	}
}

func (b *Bot) Run() {
	for {
		select {
		case message := <-ferry.GlobalMessageChan:
			if err := b.run(message); err != nil {
				log.Error(err)
				continue
			}
		}
	}
}

func (b *Bot) run(messages ferry.Message) error {
	t, err := template.New("default").Parse(msgTpl)
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	err = t.Execute(&buff, messages)
	if err != nil {
		return err
	}

	log.Infof(buff.String())

	client := &http.Client{}
	req, err := http.NewRequest("POST", webhook,
		strings.NewReader(buff.String()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Infof(string(body))
	return nil
}
