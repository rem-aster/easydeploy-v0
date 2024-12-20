package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)


// handleDeploy handles the form submission for deployment.
func handleDeploy(c echo.Context) error {
	id := c.Param("id") // Extract solution ID from the URL parameter
	c.Echo().Logger.Debugf("Extracted ID: %s", id)

	// Parse form data
	user := c.FormValue("user")
	ip := c.FormValue("ip")
	password := c.FormValue("password")

	// Print the form data to the console
	c.Echo().Logger.Debug("Received deployment request for solution ID: %s\n", id)

	deploymentID := c.Param("id")
	resId, err := Deploy(deploymentID, user, password, ip, map[string]string{})

	if err != nil {
		c.Echo().Logger.Fatalf("Deploy error!!!")
	}

	// Respond with Hx-Redirect header for htmx redirect
	c.Response().Header().Set("Hx-Redirect", fmt.Sprintf("/deploy/%s", resId))
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

// apiDeploy handles the API request for initiating a deployment with credentials.
func apiDeploy(c echo.Context) error {
	var req struct {
		SSHUser    string            `json:"ssh_user"`
		SSHPassword string           `json:"ssh_password"`
		SSHIP      string            `json:"ssh_ip"`
		ExtraVars  map[string]string `json:"extra_vars"`
	}

	id := c.Param("id") // Extract deployment ID from the request

	// Parse JSON body
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}

	// Initiate deployment
	deploymentID, err := Deploy(id, req.SSHUser, req.SSHPassword, req.SSHIP, req.ExtraVars)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to initiate deployment", "details": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"deployment_id": deploymentID})
}