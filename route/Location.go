package route

import (
	"live-track/main/controller"
	"net/http"
)

func LocationRoute(mux *http.ServeMux) {
	controller := controller.NewLocationController()

	mux.HandleFunc("/send/", controller.Publish)
}
