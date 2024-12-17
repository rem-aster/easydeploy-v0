//nolint:all
package main

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"gitlab.crja72.ru/gospec/go16/easydeploy/web_app/internal/model"
	"gitlab.crja72.ru/gospec/go16/easydeploy/web_app/internal/views" 
	"gitlab.crja72.ru/gospec/go16/easydeploy/web_app/internal/views/components"
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

// main is the entry point for the application.
func main() {
	e := echo.New()

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
	e.Logger.Fatal(e.Start(":8080"))
}

// GetSolutions returns a list of available solutions (mocked for now).
func GetSolutions() []model.Solution {
	return []model.Solution{
		{ID: "factorio", Label: "Factorio", Description: "Your own server for an amazing game about production and defense. Factory must grow!"},
		{ID: "monica", Label: "Monica", Description: "Great personal relationship management system to organize interactions with your loved ones."},
	}
}

// GetSolution fetches a specific solution by ID.
func GetSolution(id string) (model.Solution, error) {
	for _, s := range GetSolutions() {
		if s.ID == id {
			return s, nil // Return the solution if found
		}
	}
	return model.Solution{}, echo.ErrNotFound // Return error if solution is not found
}

// GetDeployStatus fetches the deployment status for a specific ID
func GetDeployStatus(id string) (model.DeployStatus, error) {
	return model.DeployStatus{ID: id, State: model.StateUnknown}, nil // Default to unknown state if not found
}
