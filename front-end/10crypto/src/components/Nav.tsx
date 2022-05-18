import React from "react";
import {Link, Navigate} from 'react-router-dom';
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
                        <li><Link to="/">Home</Link></li>
                    </ul>
                </nav>
                <div className="user-link"><Link to="/userDetail">{store.initUser.Username}</Link></div>
                <Link to="/login">
                    <button onClick={()=>{
                        store.setisLoggedin(false);
                        alert("Log out successfully!")
                        }}>
                        Log Out
                    </button>
                </Link>
                
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

