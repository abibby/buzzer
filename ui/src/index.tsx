import { h, render } from 'preact'
import Router, { Route } from 'preact-router'
import './main.css'
import { Home } from './pages/home'
import { Player } from './pages/player'

h

function Main() {
    return (
        <Router>
            <Route component={Home} path='/' />
            <Route component={Player} path='/player' />
            <Route component={Player} path='/host' host={true} />
        </Router>
    )
}

render(<Main />, document.getElementById('app')!)
