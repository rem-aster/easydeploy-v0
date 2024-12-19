//nolint:all
package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"context"
	"time"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gitlab.crja72.ru/gospec/go16/easydeploy/web_app/internal/model"
	"gitlab.crja72.ru/gospec/go16/easydeploy/web_app/internal/views" 
	"gitlab.crja72.ru/gospec/go16/easydeploy/web_app/internal/views/components"

	pb "gitlab.crja72.ru/gospec/go16/easydeploy/web_app/pkg/solution_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)



// renderHTML renders a templ.Component into an HTTP response.
func renderHTML(component templ.Component) func(echo.Context) error {
	return func(c echo.Context) error {
		buf := templ.GetBuffer()
		defer templ.ReleaseBuffer(buf)

		// Render the component into the buffer
		if err := component.Render(c.Request().Context(), buf); err != nil {
			return err
		}

		// Send the rendered HTML to the client
		return c.HTML(http.StatusOK, buf.String())
	}
}

// viewIndex serves the index page.
func viewIndex(c echo.Context) error {
	return renderHTML(views.Page([]templ.Component{components.Index()}))(c)
}

// viewSolutions serves the page displaying all solutions.
func viewSolutions(c echo.Context) error {
	solutions := GetSolutions() // Fetch solutions (mocked for now)
	return renderHTML(components.AllSolutions(solutions))(c)
}

// viewSolution serves a page for a specific solution.
func viewSolution(c echo.Context) error {
	id := c.Param("id") // Extract solution ID from the request
	solution, err := GetSolution(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound) // Return 404 if solution is not found
	}

	return renderHTML(views.Page([]templ.Component{
		components.Solution(solution),
	}))(c)
}

// viewDeployStatus serves a page showing the deployment status of a solution.
func viewDeployStatus(c echo.Context) error {
	id := c.Param("id") // Extract deployment ID from the request
	deploy, err := GetDeployStatus(id)
	if err != nil {
		return c.NoContent(http.StatusNotFound) // Return 404 if deployment status is not found
	}

	return renderHTML(views.Page([]templ.Component{
		components.Deploy(deploy),
	}))(c)
}

// handleDeploy handles the form submission for deployment.
func handleDeploy(c echo.Context) error {
	id := c.Param("id") // Extract solution ID from the URL parameter

	// Parse form data
	user := c.FormValue("user")
	ip := c.FormValue("ip")
	password := c.FormValue("password")

	// Print the form data to the console
	fmt.Printf("Received deployment request for solution ID: %s\n", id)
	fmt.Printf("User: %s, IP: %s, Password: %s\n", user, ip, password)

	// Generate a random UUID for the deployment job
	deploymentID := uuid.New().String()

	// Respond with Hx-Redirect header for htmx redirect
	c.Response().Header().Set("Hx-Redirect", fmt.Sprintf("/deploy/%s", deploymentID))
	return c.NoContent(http.StatusOK)
}

// apiGetSolutions handles the API request for fetching all solutions.
func apiGetSolutions(c echo.Context) error {
	return c.JSON(http.StatusOK, GetSolutions())
}

// apiGetSolution handles the API request for fetching a specific solution.
func apiGetSolution(c echo.Context) error {
	id := c.Param("id") // Extract solution ID from the request
	solution, err := GetSolution(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Solution not found"}) // Return error in JSON format
	}
	return c.JSON(http.StatusOK, solution)
}

// apiGetDeployStatus handles the API request for fetching deployment status.
func apiGetDeployStatus(c echo.Context) error {
	id := c.Param("id") // Extract deployment ID from the request
	deploy, err := GetDeployStatus(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Deploy status not found"}) // Return error in JSON format
	}
	return c.JSON(http.StatusOK, deploy)
}

// gRPC client connection
var grpcClient pb.SolutionV1Client

// main is the entry point for the application.
func main() {

	e := echo.New()

	// Initialize gRPC connection
	conn, err := grpc.NewClient("solution:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		e.Logger.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	grpcClient = pb.NewSolutionV1Client(conn)

	// Middleware setup
	e.Use(middleware.Logger()) // Log HTTP requests
	e.Use(middleware.Recover()) // Recover from panics and return a 500 error
	// e.Use(middleware.SecureWithConfig(middleware.DefaultSecureConfig))

	// API routes
	api := e.Group("/api")
	api.GET("/solutions", apiGetSolutions)
	api.GET("/solution/:id", apiGetSolution)
	api.GET("/deploy/:id", apiGetDeployStatus)

	// Web application routes
	webapp := e.Group("")
	webapp.Static("/static", "public/static") // Serve static files
	webapp.GET("/", viewIndex).Name = "index" // Index page
	webapp.GET("/solutions", viewSolutions).Name = "solutions" // Solutions page
	webapp.GET("/solution/:id", viewSolution).Name = "solution" // Specific solution page
	webapp.GET("/deploy/:id", viewDeployStatus).Name = "status" // Deployment status page
	webapp.POST("/deploy/:id", handleDeploy) // Deployment form handler
	webapp.Any("/*", echo.NotFoundHandler) // Catch-all for undefined routes

	// Start the server
	e.Logger.Fatal(e.Start(":80"))
}

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
			State: model.StateConnectionError,
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


func mapGrpcStatusToState(status string) model.DeployState {
    switch status {
    case "ready":
        return model.StateReady
    case "deploying":
        return model.StateDeploying
    case "unknown_error":
        return model.StateUnknownError
    case "connection_error":
        return model.StateConnectionError
    default:
        return model.StateUnknown
    }
}

func parseInt64OrZero(s string) int64 {
    id, err := strconv.ParseInt(s, 10, 64)
    if err != nil {
        return 0
    }
    return id
}