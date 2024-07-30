package domain

import "net/http"

type LocationRepository interface {
	Publish() http.HandlerFunc
}

type LocationController interface {
	Publish(writer http.ResponseWriter, request http.Request)
}
