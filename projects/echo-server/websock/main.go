package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/net/websocket"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./public")
	e.GET("/ws", wsHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func wsHandler(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close() // Close the WS connection after they disconnect (aka lets handle disconnection later)
		for {
			err := websocket.Message.Send(ws, "THe sky is black!")
			if err != nil {
				c.Logger().Error(err)
			}

			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%v\n", msg)
		}

	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
