import React from "react";

interface IProps {
    crypto: {
      code: string
      name: string
      price: number
    }[]
  }

const List: React.FC<IProps> = ({ crypto }) => {

    const renderList =(): JSX.Element[] => {
        return crypto.map(eachcrypto => {
            return (
                <li className="List">
                    <div className="List-header">
                        <h2>{eachcrypto.name}</h2>
                    </div>
                    <p>{eachcrypto.code}</p>
                    <p className="List-note">AUD <b>${eachcrypto.price}</b></p>
                    <button className="button">Trade</button>
                </li>
            )
        })
    }
    return (
        <ul className="List-container">
            {renderList()}
        </ul>
    )
}

export default List