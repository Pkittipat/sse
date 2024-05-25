package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	e.Static("/", "./")
	e.GET("/sse", sse)

	e.Logger.Fatal(e.Start(":8080"))
}

type Message struct {
	Time string `json:"time"`
}

func sse(c echo.Context) error {
	w := c.Response()
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-c.Request().Context().Done():
			return nil
		case s := <-ticker.C:
			msg := Message{Time: s.Format(time.RFC3339)}

			data, err := json.Marshal(msg)
			if err != nil {
				return err
			}
			log.Info(data)
			if _, err := fmt.Fprintf(w, "data: %s\n\n", data); err != nil {
				return err
			}
			w.Flush()
		}
	}
}
