package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"gitlab.crja72.ru/gospec/go16/easydeploy/web_app/internal/views" 
	"gitlab.crja72.ru/gospec/go16/easydeploy/web_app/internal/views/components"
)

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