package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/tiptophelmet/nomess-core/logger"
	mw "github.com/tiptophelmet/nomess-core/middleware"
	"github.com/tiptophelmet/nomess-core/responder"
)

var router *Router

type Router struct {
	mux *mux.Router
}

func Init() *Router {
	router = &Router{mux: mux.NewRouter()}

	return router
}

func GetMuxRouter() *mux.Router {
	return router.mux
}

func Handle(pattern string, handlr func(http.ResponseWriter, *http.Request)) *mux.Route {
	PeekRouteLock(pattern)

	return router.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		responder.Init(w, r)

		mw.WithMiddleware(w, r)

		handlr(w, r)
	})
}

// enhanced websocket handler
func WebSocket(pattern string, upgrader *websocket.Upgrader, handlr func(*websocket.Conn)) *mux.Route {
	PeekRouteLock(pattern)

	return router.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		mw.WithMiddleware(w, r)

		ws, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			logger.Error("failed to upgrade HTTP connection for %v with error: %v", pattern, err.Error())
			return
		}

		handlr(ws)
	})
}
