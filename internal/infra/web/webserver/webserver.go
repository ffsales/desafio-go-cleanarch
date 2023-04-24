package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router         chi.Router
	Handlers       map[string]http.HandlerFunc
	MethodHandlers []MethodHandler
	WebServerPort  string
}

type MethodHandler struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) AddMethod(method string, path string, handler http.HandlerFunc) {
	s.MethodHandlers = append(s.MethodHandlers, MethodHandler{Method: method, Path: path, Handler: handler})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		s.Router.Handle(path, handler)
	}
	for _, methodHandle := range s.MethodHandlers {
		s.Router.Method(methodHandle.Method, methodHandle.Path, methodHandle.Handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
