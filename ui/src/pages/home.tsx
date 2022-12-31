import { bindValue } from '@zwzn/spicy'
import { h } from 'preact'
import { Link } from 'preact-router'
import { useState } from 'preact/hooks'
import styles from './home.module.css'

h

export function Home() {
    const [roomCode, setRoomCode] = useState('')
    const [name, setName] = useState('')

    return (
        <div>
            <h1>Buzz</h1>
            <label>
                <div>Name</div>
                <input value={name} onInput={bindValue(setName)} />
            </label>
            <section>
                <h2>Join</h2>
                <label>
                    <div>Room Code</div>
                    <input
                        value={roomCode}
                        onInput={bindValue(toUpperCase(setRoomCode))}
                    />
                </label>
                <div>
                    <Link
                        class={styles.button}
                        href={`/player?name=${name}&room=${roomCode}`}
                        disabled={name === '' || roomCode === ''}
                    >
                        Join Room
                    </Link>
                </div>
            </section>

            <section>
                <h2>Create</h2>
                <Link
                    class={styles.button}
                    href={`/host?name=${name}&room=${randomString(4)}`}
                    disabled={name === ''}
                >
                    Create New Room
                </Link>
            </section>
        </div>
    )
}

function randomString(length: number): string {
    let str = ''
    const opts = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
    for (let i = 0; i < length; i++) {
        str += opts[Math.floor(Math.random() * opts.length)]
    }
    return str
}

function toUpperCase(cb: (v: string) => void): (v: string) => void {
    return (v: string) => {
        cb(v.toUpperCase())
    }
}
