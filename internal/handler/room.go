package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/google/uuid"
)

func RoomCreate(w http.ResponseWriter, r *http.Request) {
	id := uuid.New().String()

	w.Header().Set("Location", fmt.Sprintf("/room/%s", id))
	w.WriteHeader(http.StatusFound)
}

func Room(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id is required"))
		return
	}

	uuid, suuid, _ := createOrGetRoom(id)
}

func RoomWebSocket(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id is required"))
		return
	}

	conn, err := websocket.Accept(w, r,
		&websocket.AcceptOptions{
			OriginPatterns:     []string{"*"},
			InsecureSkipVerify: true,
		})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error in acept conection web socket: ", err)
	}

	_, _, room := createOrGetRoom(id)
}

func createOrGetRoom(id string) (uuid.UUID, uuid.UUID, Room) {
	return uuid.UUID{}, uuid.UUID{}, nil
}
