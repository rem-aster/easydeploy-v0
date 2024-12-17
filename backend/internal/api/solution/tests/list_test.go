package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/api/solution"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/model"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service"
	serviceMocks "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service/mocks"
	desc "gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/solution_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestImplementation_List(t *testing.T) {
	t.Parallel()
	type solutionServiceMockFunc func(mc *minimock.Controller) service.SolutionService

	type args struct {
		ctx context.Context
		req *desc.ListRequest
	}

	logger.TestInit()

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)
		now = time.Now()

		serviceErr = fmt.Errorf("service error")

		req = &desc.ListRequest{}

		solutions = []*desc.Solution{
			{
				Id: 1,
				Info: &desc.SolutionInfo{
					Name:        "First",
					Description: "First description",
				},
				CreatedAt: timestamppb.New(now),
				UpdatedAt: nil,
			},
		}

		serviceSolutions = []*model.Solution{
			{
				ID: 1,
				Info: model.SolutionInfo{
					Name:        "First",
					Description: "First description",
				},
				CreatedAt: now,
				UpdatedAt: sql.NullTime{},
			},
		}

		res = &desc.ListResponse{
			Solutions: solutions,
		}
	)

	tests := []struct {
		name                string
		args                args
		want                *desc.ListResponse
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
				mock.ListMock.Expect(ctx).Return(serviceSolutions, nil)
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
				mock.ListMock.Expect(ctx).Return(nil, serviceErr)
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

			newID, err := api.List(tt.args.ctx, tt.args.req)
			if err != nil {
				require.Equal(t, err, solution.InternalError)
			}
			require.Equal(t, tt.want, newID)
		})
	}
}
