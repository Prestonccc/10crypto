
export interface reg_user{
    Username: string
    Email: string
    Password: string
    Balance: number
    isRegistered: boolean
    isLoggedin: boolean
}

export interface log_user{
    Id: string
	Username: string
	Email: string
	Password: string
    Balance: number
    isLoggedin: boolean
}

export interface IState {
    crypto: {
      id: number
      code: string
      name: string
      price: number
    }[]
  }
