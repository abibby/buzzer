package room

import (
	"sync"

	"github.com/abibby/nulls"
)

var rooms = map[string]*Room{}
var roomMtx = &sync.RWMutex{}

type Action string

const (
	ActionBuzz  = Action("buzz")
	ActionReset = Action("reset")
)

type RoomRequest struct {
	Action Action `json:"action"`
}
type RoomResponse struct {
	ActivePlayer *nulls.String `json:"active_player"`
}

// func JoinRoom

type Room struct {
	Players      []*Player
	ActivePlayer string
}

func newRoom() *Room {
	return &Room{
		Players: []*Player{},
	}
}

func Join(code string, p *Player) {
	roomMtx.Lock()
	defer roomMtx.Unlock()

	room, ok := rooms[code]
	if !ok {
		room = newRoom()
	}
	room.Players = append(room.Players, p)
	rooms[code] = room
}

func Get(code string) *Room {
	roomMtx.RLock()
	defer roomMtx.RUnlock()

	room, ok := rooms[code]
	if !ok {
		return newRoom()
	}
	return room
}

func Set(code string, cb func(r *Room) *Room) *Room {
	roomMtx.Lock()
	defer roomMtx.Unlock()

	room, ok := rooms[code]
	if !ok {
		room = newRoom()
	}
	newRoom := cb(room)
	rooms[code] = newRoom
	return newRoom
}

func Delete(code string) {
	roomMtx.Lock()
	defer roomMtx.Unlock()

	delete(rooms, code)
}
