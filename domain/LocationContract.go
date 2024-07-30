package domain

import "net/http"

type LocationRepository interface {
	Publish(writer http.ResponseWriter, req *http.Request)
}

type LocationController interface {
	Publish(writer http.ResponseWriter, req *http.Request)
}
