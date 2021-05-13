import React, {Component} from 'react';
import './App.css';
import {BrowserRouter as Router, Link, Route, Switch} from 'react-router-dom'
import {Layouts} from "./layout";
import {Login} from "./components/user/login";
import {Register} from "./components/user/register";



class App extends Component {

    render() {
        return (
            <Router>
                <Switch>
                    <Route path={"/login"} component={Login}/>
                    <Route path={"/register"} component={Register}/>
                    <Route path="/" component={Layouts}/>

                </Switch>
            </Router>
        );
    }
}

export default App;
