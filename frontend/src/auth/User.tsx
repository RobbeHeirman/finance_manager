import {createContext, type ReactNode, useContext, useState} from "react";

export type User  ={
    jwtToken: string,
    userEmail: string,
    firstName: string,
    lastName: string
    pictureUrl: string

}



type UserContextType = {
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
        <UserContext value={{user, setUser: setUserWithLocalStorage}}>
            {children}
        </UserContext>
    )
}

export function GetUser(): User | null {
    return useContext(UserContext).user;
}

export function SetUser(user: User | null) {
    useContext(UserContext).setUser(user);
}

export function IsUserLoggedIn(): boolean {
    return useContext(UserContext).user != null;
}