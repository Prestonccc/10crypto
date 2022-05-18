import React from "react"
import { useState } from "react"
import { useForm } from "react-hook-form"
import { Md5 } from "ts-md5"
import { log_user, reg_user } from "./interface"
import { createContainer } from "unstated-next"
import { formHelperTextClasses, SliderTrack } from "@mui/material"
import { IState } from './interface'

export const UserState = () =>{
    const [usernameReg, setUsernameReg] = useState("")
    const [passwordReg, setPasswordReg] = useState("")
    const [emailReg, setEmailReg] = useState("")
    const [balanceReg, setBalance] = useState()
    

    const [emailLog, setEmailLog] = useState("")
    const [passwordLog, setPasswordLog] = useState("")
    const [initTopup, setTopup] = useState("")
    const [isRegistered, setisRegistered] = useState<boolean>(false)
    const [isLoggedin, setisLoggedin] = useState<boolean>(false)
    const [isToppedup, setisToppedup] = useState<boolean>(false)
    const [initUser, setInitUser] = useState<log_user>()
    

    const [initCrypto, setCrypto]= useState<IState["crypto"]>()
    // Register Form handler
    const { register, handleSubmit, formState: { errors } } = useForm();
    const onSubmit = () => {
            fetch('http://localhost:8080/api/crypto/signup', {
                method: "post",
                body: JSON.stringify({
                    username: usernameReg, 
                    password: Md5.hashStr(passwordReg), 
                    email: emailReg,
                    balance: 0})
               
            })
            setisRegistered(true)
            alert("Successfully registered! Please Login.")
            // setisLoggedin(true)
        }


    const Signin = () =>{
        fetch('http://localhost:8080/api/crypto/signin', {
                method: "POST",
                headers: {"Contect-Type":"application/json"},
                body: JSON.stringify({
                    password: Md5.hashStr(passwordLog), 
                    email: emailLog,
                })             
            })
            .then(res => res.json())
            .then(data => {
                if(data[0].user_id != undefined){
                    setisLoggedin(true)
                }
                else{
                    alert("Username or Password is incorrect!")
                }
                const currUser: log_user = {
                    Id: data[0].user_id,
                    Username: data[0].username,
                    Password:data[0].password,
                    Email: data[0].email,
                    Balance: data[0].balance,
                    isLoggedin: true,
                }
                setInitUser(currUser)
                alert("Login successfully!")
            })      
    }

    const topup = () => {
        console.log("topup", initTopup)
        fetch('http://localhost:8080/api/crypto/topup', {
            method: "POST",
            headers: {"Contect-Type":"application/json"},
            body: JSON.stringify({
                username: initUser.Username,
                email: initUser.Email,
                balance: initTopup,
            })
        })
        .then(res => res.json())
        .then(data => {
            initUser.Balance =  data[0].balance
        })
        setisToppedup(true)
        alert("Topup successfully!")
    }

    return {
        onSubmit, 
        Signin, 
        usernameReg, 
        passwordReg, 
        emailReg, 
        isRegistered, 
        balanceReg, 
        initUser, 
        handleSubmit, 
        setUsernameReg, 
        setEmailReg, 
        setPasswordReg, 
        setInitUser, 
        isLoggedin, 
        setisLoggedin,
        emailLog, 
        passwordLog,
        setEmailLog,
        setPasswordLog,
        initCrypto, 
        setCrypto,
        initTopup, 
        setTopup,
        topup,
        isToppedup,
        setisToppedup
    }
}

export const StoreContainer = createContainer(UserState)