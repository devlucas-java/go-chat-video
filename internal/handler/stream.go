package handler

import (
	"flag"
	"net/http"
	"os"

	"gihub.com/devlucas-java/go-chat-video/internal/handler"
)

var (
	addr = flag.String("addr", os.Getenv("PORT"), "")
	cert = flag.String("cert", "", "")
	key  = flag.String("key", "", "")
)

func Run() error {

	flag.Parse()

	if *addr == "" {
		*addr = ":8080"
	}

	http.HandleFunc("/", handler.Welcome)

	http.HandlerFunc("/room/create", handler.RoomCreate)
	http.HandlerFunc("/room/:uuid", handler.Room)
	app.Get("/room/:uuid/websocket")
	http.HandlerFunc("/room/:uuid/chat", handler.RooChat) // room chat
	app.Get("/room/:uuid/chat", websocket.New(handler.RoomChatWebSocket))
	app.Get("/room/uuid/viewer/websocket", websocket.New(handler.RoomViewerWebSocket))

	app.Get("/stream/:ssuid", handler.Stream)
	app.Get("/stream/:ssuid/websocket")
	app.Get("/stream/ssuid/chat/websocket")
	app.Get("/stream/ssuid/viewer/websocket")

	return nil
}
