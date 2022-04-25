import React, { Component, useState, useEffect} from 'react';
import logo from './logo.svg';
import './App.css';
import { useQuery } from 'react-fetching-library';
// import { List } from '@mantine/core';
import List from './components/List'
import { render } from '@testing-library/react';
import Nav from './components/Nav';
import Footer from './components/Footer'

interface IState {
  crypto: {
    id: number
    code: string
    name: string
    price: number
  }[]
}

function App() {
  const [data, setData] = useState<IState["crypto"]>([])
  const [isFetching, setFetch] = useState(false)
  useEffect(() => {
    async function fetchData() {
    setFetch(true)
    await fetch("http://localhost:8080/api/crypto")
    .then(res => res.json()
    .then(setData))
    setFetch(false)
    }
    fetchData()
  },[])

  if (isFetching) {
    return(
      <div className="App">
        <Nav />
        <h1 className="App-load">...Data Loading...<br/>
        It takes a while, please wait...</h1>
      </div>
    )
  }
   
  return (
    <div className="App">      
      <Nav />    
      <List crypto={data}/>
      <Footer />
    </div>
  );
}

export default App;
