import { StringMappingType } from "typescript"

export interface reg_user{
    Username: string
    Email: string
    Password: string
    Balance: number
    isRegistered: boolean
    isLoggedin: boolean
}

export interface log_user{
	Username: string
	Email: string
	Password: string
    Balance: number
    isLoggedin: boolean
}
