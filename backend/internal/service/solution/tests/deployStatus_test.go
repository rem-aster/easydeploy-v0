package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository"
	repositoryMocks "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/mocks"
	solution2 "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service/solution"
)

func TestSolution_DeployStatus(t *testing.T) {
	t.Parallel()
	type deployRepositoryMockFunc func(mc *minimock.Controller) repository.DeployRepository

	type args struct {
		ctx context.Context
	}

	logger.TestInit()

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		repositoryErr = fmt.Errorf("repository error")
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
			},
			err: nil,
			deployRepositoryMock: func(mc *minimock.Controller) repository.DeployRepository {
				mock := repositoryMocks.NewDeployRepositoryMock(mc)
				mock.DeployStatusMock.Expect(ctx).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
			},
			want: 0,
			err:  repositoryErr,
			deployRepositoryMock: func(mc *minimock.Controller) repository.DeployRepository {
				mock := repositoryMocks.NewDeployRepositoryMock(mc)
				mock.DeployStatusMock.Expect(ctx).Return(repositoryErr)
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

			err := serv.DeployStatus(tt.args.ctx)
			if err != nil {
				require.Equal(t, err, tt.err)
			}
		})
	}
}
