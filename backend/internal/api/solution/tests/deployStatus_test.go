package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/api/solution"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/logger"
	"gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service"
	serviceMocks "gitlab.crja72.ru/gospec/go16/easydeploy/backend/internal/service/mocks"
	desc "gitlab.crja72.ru/gospec/go16/easydeploy/backend/pkg/solution_v1"
)

func TestImplementation_DeployTest(t *testing.T) {
	t.Parallel()
	type solutionServiceMockFunc func(mc *minimock.Controller) service.SolutionService

	type args struct {
		ctx context.Context
		req *desc.GetDeployStatusRequest
	}

	logger.TestInit()

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		serviceErr = fmt.Errorf("service error")

		req = &desc.GetDeployStatusRequest{}

		res = &desc.GetDeployStatusResponse{}
	)

	tests := []struct {
		name                string
		args                args
		want                *desc.GetDeployStatusResponse
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
				mock.DeployStatusMock.Expect(ctx).Return(nil)
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
				mock.DeployStatusMock.Expect(ctx).Return(serviceErr)
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

			newID, err := api.DeployStatus(tt.args.ctx, tt.args.req)
			if err != nil {
				require.Equal(t, err, solution.InternalError)
			}
			require.Equal(t, tt.want, newID)
		})
	}
}
