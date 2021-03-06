// Code generated by swagger_to. DO NOT EDIT.
package foo

// Automatically generated file by swagger_to. DO NOT EDIT OR APPEND ANYTHING!

import (
	"github.com/gorilla/mux"
	"net/http"
)

// SetupRouter sets up a router. If you don't use any middleware, you are good to go.
// Otherwise, you need to maually re-implement this function with your middlewares.
func SetupRouter(h Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc(`/`,
		func(w http.ResponseWriter, r *http.Request) {
			WrapGetFoo(h, w, r)
		}).Methods("get")

	return r
}

// WrapGetFoo wraps the path `/` with the method "get".
func WrapGetFoo(h Handler, w http.ResponseWriter, r *http.Request) {
	h.GetFoo(w, r)
}

// Automatically generated file by swagger_to. DO NOT EDIT OR APPEND ANYTHING!
