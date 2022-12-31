export type RoomRequest = {
    action: 'buzz' | 'reset'
}

export type RoomResponse = {
    active_player: string
}

export class Room {
    private readonly ws: WebSocket
    constructor(name: string, roomCode: string) {
        const url = new URL(location.origin)
        url.protocol = 'ws'
        if (location.protocol === 'https:') {
            url.protocol = 'wss'
        }
        url.pathname = 'room'
        url.searchParams.set('name', name)
        url.searchParams.set('room', roomCode)
        console.log(url.toString())

        this.ws = new WebSocket(url.toString())
    }

    public send(data: RoomRequest) {
        this.ws.send(JSON.stringify(data))
    }

    public onMessage(cb: (resp: RoomResponse) => void): () => void {
        const message = (e: MessageEvent) => {
            cb(JSON.parse(e.data))
        }
        this.ws.addEventListener('message', message)
        return () => {
            this.ws.removeEventListener('message', message)
        }
    }
}
