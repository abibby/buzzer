package room

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Player struct {
	conn     *websocket.Conn
	readMtx  *sync.Mutex
	writeMtx *sync.Mutex

	Name string
}

func NewPlayer(conn *websocket.Conn, name string) *Player {
	return &Player{
		conn:     conn,
		readMtx:  &sync.Mutex{},
		writeMtx: &sync.Mutex{},
		Name:     name,
	}
}

func (r *Player) Read() (*RoomRequest, error) {
	r.readMtx.Lock()
	defer r.readMtx.Unlock()

	msg := &RoomRequest{}
	err := r.conn.ReadJSON(msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (r *Player) Write(resp *RoomResponse) error {
	r.writeMtx.Lock()
	defer r.writeMtx.Unlock()

	err := r.conn.WriteJSON(resp)
	if err != nil {
		return err
	}

	return nil
}
