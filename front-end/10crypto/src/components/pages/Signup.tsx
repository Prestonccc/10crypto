import React from "react"
import {Link, Navigate} from 'react-router-dom'
import {useState, useEffect} from 'react'
import { useForm } from "react-hook-form"
import { Md5 } from "ts-md5"
import { reg_user } from "../others/interface"
import { UserState } from "../others/globalState"
import { useContainer } from "unstated-next"
import { StoreContainer } from "../others/globalState"

export const Signup = () => {
    const {isRegistered, setUsernameReg, setEmailReg, setPasswordReg, handleSubmit, onSubmit} = useContainer(StoreContainer)
    if(isRegistered){
        return (
        <Navigate to="/"/>)
    }
    return(
        <div className="login">
            <h1>Sign Up</h1>
            <form onSubmit={handleSubmit(onSubmit)}>
                <div className="txt_field">
                    <input type="text" required 
                    onChange={(e) => {
                        setUsernameReg(e.target.value)
                    }} 
                    />
                    <span></span>
                    <label>Username</label>
                </div>

                <div className="txt_field">
                    <input type="email" required
                    onChange={(e) => {
                        setEmailReg(e.target.value)
                    }} 
                    />
                    <span></span>
                    <label>Email</label>
                </div>

                <div className="txt_field">
                    <input type="password" required
                    onChange={(e) => {
                        setPasswordReg(e.target.value)
                    }}
                    />
                    <span></span>
                    <label>Password</label>
                </div>

                <input type="submit" value="Sign Up"/>

                <div className="signup_link">
                    Already have an account?<Link to="/login"><button>Login here</button></Link>
                </div>
            </form>
        </div>
        )
    
}