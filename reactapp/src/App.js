// App.js

import React, { Component } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Download from './components/Download';

import '../node_modules/bootstrap/dist/css/bootstrap.min.css';

class App extends Component {
    render() {
        return (
            <Router>
                <div className="container">
                    <h2>React & Golang file download</h2>
                    <hr />
                    <Switch>
                        <Route path='/' component={ Download } />
                    </Switch>
                </div>
            </Router>
        );
    }
}

export default App;
