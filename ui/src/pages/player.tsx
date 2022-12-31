import { bind } from '@zwzn/spicy'
import { h } from 'preact'
import { Link } from 'preact-router'
import { useCallback, useEffect, useState } from 'preact/hooks'
import { Room } from '../services/room'
import styles from './player.module.css'

h

export interface PlayerProps {
    host?: boolean
}

export function Player({ host }: PlayerProps) {
    const url = new URL(location.href)
    const name = url.searchParams.get('name') ?? ''
    const roomCode = url.searchParams.get('room') ?? ''

    const [activePlayer, setActivePlayer] = useState<string | null>(null)
    const [room, setRoom] = useState<Room | null>(null)
    useEffect(() => {
        const r = new Room(name, roomCode)
        setRoom(r)

        return r.onMessage(msg => {
            setActivePlayer(msg.active_player)
        })
    }, [name, roomCode])

    const send = useCallback(
        (action: 'buzz' | 'reset') => {
            room?.send({ action: action })
        },
        [room],
    )

    return (
        <div class={styles.player}>
            <div class={styles.room}>{roomCode}</div>

            {activePlayer !== null && (
                <div class={styles.activePlayerPopup}>
                    <div class={styles.activePlayer}>{activePlayer}</div>

                    {host && (
                        <button
                            class={styles.reset}
                            onClick={bind('reset', send)}
                            disabled={activePlayer === null}
                        >
                            Reset
                        </button>
                    )}
                </div>
            )}

            <button
                class={styles.buzz}
                onClick={bind('buzz', send)}
                disabled={activePlayer !== null}
            >
                Buzz
            </button>

            <div class={styles.name}>Name: {name}</div>

            <Link class={styles.home} href='/'>
                Home
            </Link>
        </div>
    )
}
