import React from "react";

function Nav(){
    return(
        <div className="nav-container">
            <h1 className="App-name">CRYPTO PRES</h1>
            <nav>
                <ul className="nav-links">
                    <li><a href="#">Home</a></li>
                    <li><a href="#">Crypto</a></li>
                </ul>
            </nav>
            <a className="login" href="#"><button>Login</button></a>
        </div>
    )

}

export default Nav