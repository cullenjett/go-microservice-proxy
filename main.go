package main

import (
	"fmt"
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	loggerMiddleware(e)

	microservices := GetMicroservices()

	for _, microservice := range microservices {
		proxy(e, microservice.PublicPath, microservice.URL)
	}

	e.Logger.Fatal(e.Start(":1323"))
}

func loggerMiddleware(e *echo.Echo) {
	logFormat := `{
		"method":"${method}",
		"uri":"${uri}",
		"status":${status},
		"latency":"${latency_human}"
	}` + "\n"

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: logFormat,
	}))
}

func proxy(e *echo.Echo, path string, upstreamURL string) {
	g := e.Group(path)

	parsedURL, err := url.Parse(upstreamURL)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	g.Use(middleware.Proxy(&middleware.RoundRobinBalancer{
		Targets: []*middleware.ProxyTarget{
			{
				URL: parsedURL,
			},
		},
	}))
}
