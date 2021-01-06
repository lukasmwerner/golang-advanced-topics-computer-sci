package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	Host struct {
		Echo *echo.Echo
	}
	ReqSubDomain struct {
		Domain string `json:"domain"`
		Resp   string `json:"body"`
	}
)

func main() {

	hosts := map[string]*Host{}

	test := echo.New()
	test.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("request: %v\n", c.Request()))
	})
	test.Use(middleware.Recover())

	hosts["req.localhost:1323"] = &Host{test}

	allocr := echo.New()
	allocr.POST("/", func(c echo.Context) error {
		sub := new(ReqSubDomain)

		if err := c.Bind(sub); err != nil {
			return err
		}
		newServer := echo.New()
		newServer.GET("/", func(c echo.Context) error {
			return c.HTML(http.StatusOK, fmt.Sprintf("%s", sub.Resp))
		})
		hosts[sub.Domain] = &Host{newServer}

		return c.JSON(http.StatusOK, sub)
	})

	hosts["alloc.localhost:1323"] = &Host{allocr}

	e := echo.New()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":1323"))
}
