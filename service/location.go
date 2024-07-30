package service

import (
	"context"
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

func (repo locationRepository) Publish(writer http.ResponseWriter, req *http.Request) {
	c, err := websocket.Accept(writer, req, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.CloseNow()

	ctx, cancel := context.WithTimeout(req.Context(), time.Second*10)
	defer cancel()

	var v interface{}
	err = wsjson.Read(ctx, c, &v)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("received: %v", v)

	c.Close(websocket.StatusNormalClosure, "")
}
