import React from "react"
import { useState } from "react"
import { useForm } from "react-hook-form"
import { Md5 } from "ts-md5"
import { log_user, reg_user } from "./interface"
import { createContainer } from "unstated-next"
import { formHelperTextClasses, SliderTrack } from "@mui/material"

export const UserState = () =>{
    const [usernameReg, setUsernameReg] = useState("")
    const [passwordReg, setPasswordReg] = useState("")
    const [emailReg, setEmailReg] = useState("")
    const [balanceReg, setBalance] = useState()
    const [passwordLog, setPasswordLog] = useState("")

    const [emailLog, setEmailLog] = useState("")
    const [isRegistered, setisRegistered] = useState<boolean>(false)
    const [isLoggedin, setisLoggedin] = useState<boolean>(false)
    const [initUser, setInitUser] = useState<log_user>()

    // Register Form handler
    const { register, handleSubmit, formState: { errors } } = useForm();
    const onSubmit = () => {
            fetch('http://localhost:8080/api/crypto/signup', {
                method: "post",
                body: JSON.stringify({
                    "username": usernameReg, 
                    "password": Md5.hashStr(passwordReg), 
                    "email": emailReg,
                    "balance": 0})
               
            })
            
            // const currUser: reg_user = {
            //         Username: usernameReg,
            //         Password:passwordReg,
            //         Email: emailReg,
            //         Balance: balanceReg,
            //         isRegistered: true,
            //         isLoggedin: true,
            // }
            setisRegistered(true)
            setisLoggedin(true)
            // console.log(currUser)
            // setInitUser(currUser)
            // console.log(initUser)        
        }


    const Signin = () =>{
        fetch('http://localhost:8080/api/crypto/signin', {
                method: "POST",
                headers: {"Contect-Type":"application/json"},
                body: JSON.stringify({
                    "Password": Md5.hashStr(passwordLog), 
                    "Email": emailLog,
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
            })
                
            
   
            // const currUser: reg_user = {
            //         Username: usernameReg,
            //         Password:passwordReg,
            //         Email: emailReg,
            //         Balance: balanceReg,
            //         isRegistered: true,
            //         isLoggedin: true,
            // }
            
            console.log(isLoggedin)
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
        emailLog, 
        passwordLog,
        setEmailLog,
        setPasswordLog,
    }
}

export const StoreContainer = createContainer(UserState)