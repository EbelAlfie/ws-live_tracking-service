package route

import (
	"live-track/main/controller"
	"net/http"
)

func LocationRoute() http.HandlerFunc {
	controller := controller.NewLocationController()

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		http.HandleFunc("/send/", controller.Publish)
	})
}
