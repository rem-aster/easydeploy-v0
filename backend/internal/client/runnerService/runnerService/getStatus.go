package runnerService

import (
	"context"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/runner_v1"
)

// GetTaskStatus реализует метод интерфейса RunnerService
func (c *Client) GetTaskStatus(ctx context.Context, id string) (string, string) {
	client := runner_v1.NewPlaybookServiceV1Client(c.conn)
	resp, err := client.GetTaskStatus(ctx, &runner_v1.PlaybookStatusRequest{
		Id: id,
	})
	if err != nil {
		return "", err.Error()
	}
	return resp.GetStatus(), resp.Error
}
