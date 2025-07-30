import {useEffect, useState} from "react";
import {setApiLogoutHandler} from "../../api/global_axios.ts";
import {getUserFromLocalStorage, setUserToLocalStorage, type User} from "./user.ts";
import {UserContextSupplier, type UserProviderProps} from "./userContext.ts";

export function UserContextProvider({children}: UserProviderProps) {
    const [user, setUser] = useState<User | null>(getUserFromLocalStorage);
    const setUserWithLocalStorage = (user: User | null) => {
        setUserToLocalStorage(user)
        setUser(user)
    }
    useEffect(() => {
        setApiLogoutHandler(() => setUserWithLocalStorage(null))
    }, [])
    return (
        <UserContextSupplier value={{user: user, setUser: setUserWithLocalStorage}}>
            {children}
        </UserContextSupplier>
    )
}


