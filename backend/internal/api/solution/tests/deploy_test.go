package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/api/solution"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service"
	serviceMocks "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service/mocks"
	desc "gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/solution_v1"
)

func TestImplementation_Deploy(t *testing.T) {
	t.Parallel()
	type solutionServiceMockFunc func(mc *minimock.Controller) service.SolutionService

	type args struct {
		ctx context.Context
		req *desc.DeployRequest
	}

	logger.TestInit()

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		serviceErr = fmt.Errorf("service error")

		req = &desc.DeployRequest{
			SolutionId: 1,
			SshAddress: "root@127.0.0.1",
			SshKey:     []byte("ssh key"),
			ExtraVars:  map[string]string{"key": "value"},
		}

		deploy = &model.Deploy{
			DeployInfo: model.DeployInfo{
				SolutionId: 1,
				SSHAddress: "root@127.0.0.1",
				SSHKey:     []byte("ssh key"),
				Extra:      map[string]string{"key": "value"},
			},
		}

		res = &desc.DeployResponse{
			Id: 1,
		}
	)

	tests := []struct {
		name                string
		args                args
		want                *desc.DeployResponse
		err                 error
		solutionServiceMock solutionServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			solutionServiceMock: func(mc *minimock.Controller) service.SolutionService {
				mock := serviceMocks.NewSolutionServiceMock(mc)
				mock.DeployMock.Expect(ctx, deploy).Return(1, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			solutionServiceMock: func(mc *minimock.Controller) service.SolutionService {
				mock := serviceMocks.NewSolutionServiceMock(mc)
				mock.DeployMock.Expect(ctx, deploy).Return(0, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			createServiceMock := tt.solutionServiceMock(mc)
			api := solution.NewImplementation(createServiceMock)

			newID, err := api.Deploy(tt.args.ctx, tt.args.req)
			if err != nil {
				require.Equal(t, err, solution.InternalError)
			}
			require.Equal(t, tt.want, newID)
		})
	}
}
