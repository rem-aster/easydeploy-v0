package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository"
	repositoryMocks "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/repository/mocks"
	solution2 "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service/solution"
)

func TestSolution_List(t *testing.T) {
	t.Parallel()
	type solutionRepositoryMockFunc func(mc *minimock.Controller) repository.SolutionRepository

	type args struct {
		ctx context.Context
	}

	logger.TestInit()

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		repositoryErr = fmt.Errorf("repository error")

		res = []*model.Solution{
			{
				ID: 1,
				Info: model.SolutionInfo{
					Name:        "First",
					Description: "First description",
				},
				CreatedAt: time.Now(),
				UpdatedAt: sql.NullTime{},
			},
			{
				ID: 2,
				Info: model.SolutionInfo{
					Name:        "Second",
					Description: "Second description",
				},
				CreatedAt: time.Now(),
				UpdatedAt: sql.NullTime{},
			},
		}
	)

	tests := []struct {
		name                   string
		args                   args
		want                   []*model.Solution
		err                    error
		solutionRepositoryMock solutionRepositoryMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
			},
			want: res,
			err:  nil,
			solutionRepositoryMock: func(mc *minimock.Controller) repository.SolutionRepository {
				mock := repositoryMocks.NewSolutionRepositoryMock(mc)
				mock.ListMock.Expect(ctx).Return(res, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
			},
			want: nil,
			err:  repositoryErr,
			solutionRepositoryMock: func(mc *minimock.Controller) repository.SolutionRepository {
				mock := repositoryMocks.NewSolutionRepositoryMock(mc)
				mock.ListMock.Expect(ctx).Return(nil, repositoryErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			createRepositoryMock := tt.solutionRepositoryMock(mc)
			serv := solution2.NewService(createRepositoryMock, nil, nil)

			newID, err := serv.List(tt.args.ctx)
			if err != nil {
				require.Equal(t, err, tt.err)
			}
			require.Equal(t, tt.want, newID)
		})
	}
}
