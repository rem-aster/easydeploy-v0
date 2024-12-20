package main


import (
	"fmt"
	"net/http"
	"context"
	"time"

	"github.com/labstack/echo/v4"

	"gitlab.crja72.ru/gospec/go16/easydeploy/web_app/internal/model"

	pb "gitlab.crja72.ru/gospec/go16/easydeploy/web_app/pkg/solution_v1"
)

// GetSolutions returns a list of available solutions via gRPC.
func GetSolutions() []model.Solution {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the gRPC List method
	response, err := grpcClient.List(ctx, &pb.ListRequest{})
	if err != nil {
		return nil // Return empty list in case of error
	}

	// Map gRPC response to model.Solution
	solutions := make([]model.Solution, 0, len(response.Solutions))
	for _, sol := range response.Solutions {
		solutions = append(solutions, model.Solution{
			ID:          fmt.Sprintf("%d", sol.Id),
			Label:       sol.Info.Name,
			Description: sol.Info.Description,
		})
	}

	return solutions
}

// GetSolution fetches a specific solution by ID.
func GetSolution(id string) (model.Solution, error) {
	// Get all solutions
	solutions := GetSolutions()
	if solutions == nil {
		return model.Solution{}, echo.NewHTTPError(http.StatusInternalServerError, "Unable to fetch solutions")
	}

	// Find the solution by ID
	for _, s := range solutions {
		if s.ID == id {
			return s, nil
		}
	}

	return model.Solution{}, echo.ErrNotFound // Return error if solution is not found
}

// GetDeployStatus fetches the deployment status for a specific ID.
func GetDeployStatus(id string) (model.DeployStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the gRPC GetDeployStatus method
	response, err := grpcClient.GetDeployStatus(ctx, &pb.GetDeployStatusRequest{
		Id: parseInt64OrZero(id),
	})
	if err != nil {
		return model.DeployStatus{
			ID:    id,
			State: model.StateUnknown,
		}, err
	}

	// Map gRPC response status to DeployState
	state := mapGrpcStatusToState(response.Status)
	return model.DeployStatus{
		ID:    id,
		State: state,
	}, nil
}

// Deploy starts a deploy task.
func Deploy(id string, sshUser string, sshPassword string, sshIP string, extraVars map[string]string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Construct gRPC DeployRequest
	req := &pb.DeployRequest{
		SolutionId: parseInt64OrZero(id),
		SshAddress: fmt.Sprintf("%s@%s", sshUser, sshIP),
		SshKey:     sshPassword, // Assuming SSH key for simplicity
		ExtraVars:  extraVars,
	}
	// Call the gRPC Deploy method
	response, err := grpcClient.Deploy(ctx, req)
	if err != nil {
		return -1, err
	}
	return int(response.Id), nil
}

