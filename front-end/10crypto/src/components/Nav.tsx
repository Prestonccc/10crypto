import React from "react";
import {Link} from 'react-router-dom';
import { useContainer } from "unstated-next";
import { StoreContainer } from "./others/globalState";



export default function Nav(){
    const store = useContainer(StoreContainer)

    if(store.isLoggedin){
        return(
            <div className="nav-container">
                <h1 className="App-name">CRYPTO PRES</h1>
                <nav>
                    <ul className="nav-links">
                        <Link to="/">Home</Link>
                    </ul>
                </nav>
                <button>Log Out</button>
                <p>{store.usernameReg}</p>
        </div>
        )
    }

    return(
        <div className="nav-container">
                <h1 className="App-name">CRYPTO PRES</h1>
                <nav>
                    <ul className="nav-links">
                        <Link to="/">Home</Link>
                    </ul>
                </nav>
                <Link to="/login"><button>Login</button></Link>
        </div>
    )

}

