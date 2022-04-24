import React, { Component, useState, useEffect} from 'react';
import logo from './logo.svg';
import './App.css';
import { useQuery } from 'react-fetching-library';
// import { List } from '@mantine/core';
import List from './components/List'
import { render } from '@testing-library/react';

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
  // if (!isFetching) {
  //   <div>
  //     <h1>LOADING...PLEASE WAIT...</h1>
  //   </div>
  //   }
  useEffect(() => {
    async function fetchData() {
    setFetch(true)
    await fetch("http://localhost:8080/home")
    .then(res => res.json()
    .then(setData))
    setFetch(false)
    }
    fetchData()
  },[])

  if (isFetching) {
    return(
      <div className="App">
        <h1>...Data Loading...</h1>
        <h1>It takes a while, please wait...</h1>
      </div>
    )
  }
   
  return (
    <div className="App">
      <h1>Top Ten Crypto</h1>
      <List crypto={data}/>
    </div>
  );
}

export default App;
