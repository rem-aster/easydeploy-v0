package main

import (
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"gitlab.crja72.ru/gospec/go16/easydeploy/web_app/internal/model"
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

func mapGrpcStatusToState(status string) model.DeployState {
    switch status {
    case "running":
        return model.StateRunning
    case "failed":
        return model.StateFailed
    case "success":
        return model.StateSuccess
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