import React, { Component, useState} from 'react';
import logo from './logo.svg';
import './App.css';
import useSWR from "swr"
import {Box} from "@mantine/core"
import List from './components/List'

export const ENDPOINT = 'htttp://localhost:8080'
const fetcher =(url: string) => 
fetch(`${ENDPOINT}/home`).then((r) => r.json());


interface IState {
  crypto: {
    code: string
    name: string
    price: number
    url: string
  }[]
}

function App() {
  // const {data,mutate} = useSWR('home',fetcher) 



  const [crypto, setCrypto] = useState<IState["crypto"]>([
    {
      code: "BTC",
      name: "Bitcoin",
      price: 55853.679688,
      url: "https://bitcoin.org/img/icons/opengraph.png?1648897668"
    }
  ])

  return (
    <div className="App">
      <h1>Ten Crypto</h1>
      {/* <List crypto={crypto}/> */}
      {/* <Box>{JSON.stringify(data)}</Box> */}
    </div>
  );
}

export default App;
