package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kakebon/backend/domain/repository"
)

type slackNotifier struct {
	webhookURL string
}

func NewSlackNotifier (url string) repository.Notifier {
	return &slackNotifier{webhookURL: url}
}

func (s *slackNotifier) Send(message string) error {
	payload := map[string]string{"text": message}
	data, _ := json.Marshal(payload)

	res, err := http.Post(s.webhookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
			buf := new(bytes.Buffer)
			buf.ReadFrom(res.Body)
			return fmt.Errorf("slack error: status %d, body: %s", res.StatusCode, buf.String())
	}
    
	fmt.Println("Slack says: OK!")

	return nil
}