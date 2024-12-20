package main

import (
	"golang.org/x/time/rate"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/gommon/log"

	pb "gitlab.crja72.ru/gospec/go16/easydeploy/web_app/pkg/solution_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// gRPC client connection
var grpcClient pb.SolutionV1Client

// main is the entry point for the application.
func main() {

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	c := jaegertracing.New(e, nil)
    defer c.Close()

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
	e.Use(echoprometheus.NewMiddleware("webapp")) // adds middleware to gather metrics
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(10))))

	// Util routes
	e.GET("/metrics", echoprometheus.NewHandler()) // adds route to serve gathered metrics

	// API routes
	api := e.Group("/api")
	api.GET("/solutions", apiGetSolutions)
	api.GET("/solution/:id", apiGetSolution)
	api.GET("/deploy/:id", apiGetDeployStatus)
	api.POST("/deploy/:id", apiDeploy)

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
	e.Logger.Fatal(e.Start(":8082"))
}
