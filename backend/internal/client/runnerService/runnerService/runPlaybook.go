package runnerService

import (
	"context"

	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/runner_v1"
)

// RunPlaybook реализует метод интерфейса RunnerService
func (c *Client) RunPlaybook(ctx context.Context, deploy model.Deploy, solution model.Solution) (string, string) {
	client := runner_v1.NewPlaybookServiceV1Client(c.conn)
	resp, err := client.RunPlaybook(ctx, &runner_v1.RunPlaybookRequest{
		PlaybookName: solution.Info.Playbook,
		SshAddress:   deploy.DeployInfo.SSHAddress,
		SshPassword:  string(deploy.DeployInfo.SSHKey),
		ExtraVars:    nil,
	})
	if err != nil {
		return "", err.Error()
	}
	return resp.GetId(), resp.GetStatus()
}
