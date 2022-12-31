package controllers

import (
	"log"
	"net/http"

	"github.com/abibby/buzzer/room"
	"github.com/abibby/nulls"
	"github.com/abibby/validate"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type RoomRequest struct {
	Room string `query:"room" validate:"required"`
	Name string `query:"name" validate:"required"`
}

func JoinRoom(w http.ResponseWriter, r *http.Request) {
	req := &RoomRequest{}
	err := validate.Run(r, req)
	if err != nil {
		log.Print(err)
		errorResponse(w, err)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
		errorResponse(w, err)
		return
	}
	defer func() {
		conn.Close()
		r := room.Set(req.Room, func(r *room.Room) *room.Room {
			players := make([]*room.Player, 0, len(r.Players)-1)
			for _, p := range r.Players {
				if p.Name != req.Name {
					players = append(players, p)
				}
			}
			r.Players = players
			return r
		})
		if len(r.Players) == 0 {
			room.Delete(req.Room)
		}
	}()

	player := room.NewPlayer(conn, req.Name)
	room.Join(req.Room, player)

	pushUpdate(room.Get(req.Room))

	for {
		msg, err := player.Read()
		if err != nil {
			log.Print(err)
			return
		}

		switch msg.Action {
		case room.ActionBuzz:
			buzz(req.Room, req.Name)
		case room.ActionReset:
			reset(req.Room)
		}
	}
}

func buzz(code, name string) error {
	r := room.Get(code)
	if r.ActivePlayer != "" {
		return nil
	}
	room.Set(code, func(r *room.Room) *room.Room {
		r.ActivePlayer = name
		return r
	})

	return pushUpdate(r)
}
func reset(code string) error {
	r := room.Set(code, func(r *room.Room) *room.Room {
		r.ActivePlayer = ""
		return r
	})

	return pushUpdate(r)
}

func pushUpdate(r *room.Room) error {
	for _, p := range r.Players {
		activePlayer := nulls.NewString(r.ActivePlayer)
		if activePlayer.String() == "" {
			activePlayer = nil
		}
		err := p.Write(&room.RoomResponse{
			ActivePlayer: activePlayer,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
