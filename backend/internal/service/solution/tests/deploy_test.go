package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository"
	repositoryMocks "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/mocks"
	solution2 "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service/solution"
)

func TestSolution_Deploy(t *testing.T) {
	t.Parallel()
	type deployRepositoryMockFunc func(mc *minimock.Controller) repository.DeployRepository

	type args struct {
		ctx context.Context
		req *model.Deploy
	}

	logger.TestInit()

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		repositoryErr = fmt.Errorf("repository error")

		req = &model.Deploy{
			DeployInfo: model.DeployInfo{
				SolutionId: 1,
				SSHAddress: "root@127.0.0.1",
				SSHKey:     []byte("ssh key"),
				Extra:      map[string]string{"key": "value", "key2": "value2"},
			},
		}

		res = int64(1)
	)

	tests := []struct {
		name                 string
		args                 args
		want                 int64
		err                  error
		deployRepositoryMock deployRepositoryMockFunc
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
		},
		{
			name: "service error case",
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
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			createRepositoryMock := tt.deployRepositoryMock(mc)
			serv := solution2.NewService(nil, createRepositoryMock, nil)

			newID, err := serv.Deploy(tt.args.ctx, tt.args.req)
			if err != nil {
				require.Equal(t, err, tt.err)
			}
			require.Equal(t, tt.want, newID)
		})
	}
}
