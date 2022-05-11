import React from "react"
import {Link, Navigate} from 'react-router-dom'
import { useContainer } from "unstated-next"
import { StoreContainer } from "../others/globalState"

export const Login = () => {
    const {setEmailLog, setPasswordLog, isLoggedin, handleSubmit, Signin} = useContainer(StoreContainer)

    if(isLoggedin){
        return (
            <Navigate to="/"/>)
    }
    return(
        <div className="login">
            <h1>Login</h1>
            <form onSubmit={handleSubmit(Signin)}>
                <div className="txt_field">
                    <input type="email" required
                     onChange={(e) => {
                        setEmailLog(e.target.value)
                    }}
                    />
                    <span></span>
                    <label>Email</label>
                </div>

                <div className="txt_field">
                    <input type="password" required
                     onChange={(e) => {
                        setPasswordLog(e.target.value)
                    }}
                    />
                    <span></span>
                    <label>Password</label>
                </div>

                <div className="pass">Forget Password?</div>

                <input type="submit" value="Login" />

                <div className="signup_link">
                    Not a user?<Link to="/signup"><button>Signup</button></Link>
                </div>
            </form>
        </div>
        )
    
}