import React from "react";
import useFetch from "./useFetch";
import CircularProgress from '@mui/material/CircularProgress';
import {Link} from 'react-router-dom'

const List = () => {
// const List: React.FC<IProps> = () => {
    const {data, isFetching} = useFetch('http://localhost:8080/api/crypto');
    const renderList =(): JSX.Element[] => {

        return data.map(eachcrypto => {
            return (
                <li key={eachcrypto.code} className="List">
                    <div className="List-header">
                        <h2>{eachcrypto.code}</h2>
                    </div>
                    <p>{eachcrypto.name}</p>
                    <p className="List-note">AUD <b>${eachcrypto.price}</b></p>
                    <Link to={`/crypto/${eachcrypto.code}`}><button className="button">Trade</button></Link>
                </li>
            )
        })
    }
    return (
        <ul className="List-container">
            {isFetching && <div className="circle"><p>Don't switch to other pages.<br/>Real-time rate, please wait...</p><CircularProgress/></div>}
            {renderList()}
        </ul>
    )
}

export default List