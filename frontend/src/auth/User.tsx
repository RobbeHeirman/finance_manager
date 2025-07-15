import {createContext, type ReactNode, useContext, useState} from "react";

export type User  ={
    jwtToken: string,
    userEmail: string,
    firstName: string,
    lastName: string
    pictureUrl: string

}



export type UserContextType = {
    user: User | null;
    setUser: (_: User | null) => void;
};

const UserContext = createContext<UserContextType>({
    user: null,
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    setUser: (_) => {},
});

type UserProviderProps = {
    children: ReactNode;
};

const KEY_USER = "user";

function getUserFromLocalStorage() : User | null {
    const storedUser = localStorage.getItem(KEY_USER)
    if (!storedUser) {
        return null
    }

    try {
        return JSON.parse(storedUser)
    } catch (e) {
        console.error(`Wrong json fromat of stored user ${storedUser}`, e)
        localStorage.removeItem(KEY_USER)
    }
    return null
}

export function UserProvider({children}: UserProviderProps) {
    const [user, setUser] = useState<User | null>(getUserFromLocalStorage);
    const setUserWithLocalStorage  = (user: User | null) => {
        localStorage.setItem(KEY_USER, JSON.stringify(user))
        setUser(user)
    }
    return (
        <UserContext value={{user: user, setUser: setUserWithLocalStorage}}>
            {children}
        </UserContext>
    )
}

export function GetUserContext() {
    return useContext(UserContext)
}

export function IsUserLoggedIn(): boolean {
    return useContext(UserContext).user != null;
}