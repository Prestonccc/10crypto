import React from "react"
import { useContainer } from "unstated-next"
import { StoreContainer } from "../others/globalState"
import { Link, Navigate } from "react-router-dom"

export const UserDetail = () => {
    const {initUser, setTopup, handleSubmit, topup, isToppedup, setisToppedup} = useContainer(StoreContainer)
    if(isToppedup){
        setisToppedup(false)
        return(
            <Navigate to="/" />
        )
    }
    return(
        <div className="user-detail">
            <h2><b>USERNAME:</b>  {initUser.Username}</h2>
            <h2><b>ID:</b>  {initUser.Id}</h2>
            <h2><b>EMAIL:</b>  {initUser.Email}</h2>
            <h2><b>BALANCE(AUD$):</b>  {initUser.Balance}</h2>
            <form onSubmit={handleSubmit(topup)}>
                <div className="txt_field">
                    <input type="number" 
                        onChange={(e) => {
                            setTopup(e.target.value)
                        }}
                    />
                    <label>Price AUD$</label>
                </div>
                <input type="submit" value="TOP UP"/>
            </form>
            <Link to="/"><button className="start-trading">Start Trading</button></Link>
        </div>
    )
}