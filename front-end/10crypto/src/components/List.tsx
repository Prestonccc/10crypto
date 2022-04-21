import React from "react";

interface IProps {
    crypto: {
      code: string
      name: string
      price: number
      url: string
    }[]
  }

const List: React.FC<IProps> = ({ crypto }) => {

    const renderList =(): JSX.Element[] => {
        return crypto.map(eachcrypto => {
            return (
                <li className="List">
                    <div className="List-header">
                        <img className="List-img" src={eachcrypto.url}/>
                        <h2>{eachcrypto.code}</h2>
                    </div>
                    <p>{eachcrypto.price}</p>
                    <p className="List-note">{eachcrypto.name}</p>
                </li>
            )
        })
    }
    return (
        <ul>
            {renderList()}
        </ul>
    )
}

export default List