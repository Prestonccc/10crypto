import React, { Component, useState, useEffect} from 'react';
import './App.css';
import List from './components/List'
import Nav from './components/Nav';
import Footer from './components/Footer'
import {Login} from './components/pages/Login';
import { CryptoDetail } from './components/pages/CryptoDetail';
import {BrowserRouter as Router, Routes, Route, Link} from 'react-router-dom';
import { Signup } from './components/pages/Signup';
import { NotFound } from './components/pages/Notfound';
import { StoreContainer } from './components/others/globalState';
export default class App extends Component{
  constructor(props) {
    super(props);

    this.state = {
      loggedInstatus: "NOT_LOGGED_IN",
      user: {}
    }
  }
  render(){  
    return (
    <StoreContainer.Provider>
    <div className="App">
      <Router>
        <Nav />
        <div className="content">
          <Routes>
            <Route path="/" element={<List />} />
            <Route path="/login" element={<Login />} />
            <Route path="/signup" element={<Signup />} />
            <Route path="/crypto/:code" element={<CryptoDetail />} />
            <Route path="*" element={<NotFound />} />
          </Routes>  
        </div>
        <Footer />
      </Router>
    </div>
    </StoreContainer.Provider>
);}

}

// export default App;
