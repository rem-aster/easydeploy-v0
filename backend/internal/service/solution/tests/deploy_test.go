package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/client/runnerService"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/client/runnerService/mocks"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository"
	repositoryMocks "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/mocks"
	solution2 "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service/solution"
)

func TestSolution_Deploy(t *testing.T) {
	//t.Parallel()
	type deployRepositoryMockFunc func(mc *minimock.Controller) repository.DeployRepository
	type runnerServiceMockFunc func(mc *minimock.Controller) runnerService.RunnerService
	type solutionServiceMockFunc func(mc *minimock.Controller) repository.SolutionRepository

	type args struct {
		ctx context.Context
		req *model.Deploy
	}

	logger.TestInit()

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		repositoryErr = fmt.Errorf("repository error")
		serviceErr    = fmt.Errorf("service error")

		req = &model.Deploy{
			DeployInfo: model.DeployInfo{
				SolutionId: 1,
				SSHAddress: "root@127.0.0.1",
				SSHKey:     "ssh key",
				Extra:      map[string]string{"key": "value", "key2": "value2"},
			},
		}
		solution = &model.Solution{
			ID: 1,
			Info: model.SolutionInfo{
				Name:        "Test",
				Description: "Test desc",
				Playbook:    "1.yml",
			},
			CreatedAt: time.Now(),
			UpdatedAt: sql.NullTime{},
		}
		idPlaybook = "abcdeeradfsdaf"

		res = int64(1)
	)

	tests := []struct {
		name                   string
		args                   args
		want                   int64
		err                    error
		deployRepositoryMock   deployRepositoryMockFunc
		runnerServiceMock      runnerServiceMockFunc
		solutionRepositoryMock solutionServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			deployRepositoryMock: func(mc *minimock.Controller) repository.DeployRepository {
				mock := repositoryMocks.NewDeployRepositoryMock(mc)
				mock.DeployMock.Expect(ctx, req).Return(1, nil)
				return mock
			},
			solutionRepositoryMock: func(mc *minimock.Controller) repository.SolutionRepository {
				mock := repositoryMocks.NewSolutionRepositoryMock(mc)
				mock.GetMock.Expect(ctx, req.DeployInfo.SolutionId).Return(solution, nil)
				return mock
			},
			runnerServiceMock: func(mc *minimock.Controller) runnerService.RunnerService {
				mock := mocks.NewRunnerServiceMock(mc)
				mock.RunPlaybookMock.Expect(ctx, *req, *solution).Return(idPlaybook, "")
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  serviceErr,
			deployRepositoryMock: func(mc *minimock.Controller) repository.DeployRepository {
				mock := repositoryMocks.NewDeployRepositoryMock(mc)
				return mock
			},
			solutionRepositoryMock: func(mc *minimock.Controller) repository.SolutionRepository {
				mock := repositoryMocks.NewSolutionRepositoryMock(mc)
				mock.GetMock.Expect(ctx, req.DeployInfo.SolutionId).Return(nil, serviceErr)
				return mock
			},
			runnerServiceMock: func(mc *minimock.Controller) runnerService.RunnerService {
				mock := mocks.NewRunnerServiceMock(mc)
				return mock
			},
		},
		{
			name: "repository error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repositoryErr,
			deployRepositoryMock: func(mc *minimock.Controller) repository.DeployRepository {
				mock := repositoryMocks.NewDeployRepositoryMock(mc)
				mock.DeployMock.Expect(ctx, req).Return(0, repositoryErr)
				return mock
			},
			solutionRepositoryMock: func(mc *minimock.Controller) repository.SolutionRepository {
				mock := repositoryMocks.NewSolutionRepositoryMock(mc)
				mock.GetMock.Expect(ctx, req.DeployInfo.SolutionId).Return(solution, nil)
				return mock
			},
			runnerServiceMock: func(mc *minimock.Controller) runnerService.RunnerService {
				mock := mocks.NewRunnerServiceMock(mc)
				req.DeployInfo.SolutionId = 1
				mock.RunPlaybookMock.Expect(ctx, *req, *solution).Return(idPlaybook, "")
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			createDeployRepositoryMock := tt.deployRepositoryMock(mc)
			createSolutionRepositoryMock := tt.solutionRepositoryMock(mc)
			createRunnerServiceMock := tt.runnerServiceMock(mc)
			serv := solution2.NewService(createSolutionRepositoryMock, createDeployRepositoryMock, createRunnerServiceMock, nil)

			newID, err := serv.Deploy(tt.args.ctx, tt.args.req)
			if err != nil {
				require.Equal(t, err, tt.err)
			}
			require.Equal(t, tt.want, newID)
		})
	}
}
