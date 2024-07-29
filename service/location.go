package service

import (
	"context"
	"fmt"
	"live-track/main/domain"
	"log"
	"net/http"
	"time"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type locationRepository struct {
}

func NewLocationRepository() domain.LocationRepository {
	return locationRepository{}
}

func (repo locationRepository) Publish() http.HandlerFunc {
	fmt.Print("Hello world")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, nil)
		if err != nil {
			// ...
		}
		defer c.CloseNow()

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		var v interface{}
		err = wsjson.Read(ctx, c, &v)
		if err != nil {
			// ...
		}

		log.Printf("received: %v", v)

		c.Close(websocket.StatusNormalClosure, "")
	})
}
