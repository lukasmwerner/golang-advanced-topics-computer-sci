package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		hostParts := strings.Split(req.Host, ".")
		if len(hostParts) > 2 {
			dockerhostname := hostParts[0]
			client := &http.Client{}
			req, err := http.NewRequest(req.Method, fmt.Sprintf("http://%s/%s", dockerhostname, req.URL.RawPath), nil)
			if err != nil {
				return err
			}
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			return c.String(resp.StatusCode, string(body))
		}
		return fmt.Errorf("Unable to parse subdomain")
	})
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":1323"))
}
