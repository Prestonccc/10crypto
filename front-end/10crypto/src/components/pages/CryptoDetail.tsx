import React from "react"
import { useParams } from "react-router"
import useFetch from '../useFetch'
import { useNavigate } from 'react-router-dom';


export const CryptoDetail = () => {
    const navigate = useNavigate();

    const { code } = useParams();
    const {data, error, isFetching} = useFetch('http://localhost:8080/api/crypto/' + code);

    const renderList =(): JSX.Element[] => {
        return data.map(eachcrypto => {
            return (
                <div key={eachcrypto.code} className="detail-container">
                    {isFetching && <div>Loading...</div>}
                    {error && <div>{ error }</div>}
                    {crypto && (
                        <div className="cryptodetail">
                            <h2>{eachcrypto.code}</h2>
                            <p>{eachcrypto.name}</p>
                            <p>AUD ${eachcrypto.price}</p>
                            <form method="post">
                                <div className="txt_field">
                                    <input type="number" required />
                                    <label>Price AUD$</label>
                                </div>
                                <input type="submit" value="BUY" />
                            </form>

                            <form method="post">
                                <div className="txt_field">
                                    <input type="number" required />
                                    <label>Price AUD$</label>
                                </div>
                                <input type="submit" value="SELL" />
                            </form>
                        </div>
                        
                    )}
            </div>
            )
        })}

    return (
        <div>
            <div className="back">
                <button onClick={() => navigate('/')}>&#x3C; Back</button>
            </div>
            <h1>
                {renderList()}
            </h1>

        </div>
    )

}