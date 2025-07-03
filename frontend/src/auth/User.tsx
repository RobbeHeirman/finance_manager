import  {createContext, type ReactNode, useState} from "react";

export type User  ={
    jwtToken: string,
    userEmail: string,
    firstName: string,
    lastName: string
    pictureUrl: string

}



type UserContextType = {
    user: User | null;
    setUser: (user: User) => void
};

const UserContext = cr
eateContext<UserContextType>({
    user: null,
    setUser: () => {
    },
});

type UserProviderProps = {
    children: ReactNode;
};

function getUserFromLocalStorage() : User | null {
    const storedUser = localStorage.getItem("user")
    if (!storedUser) {
        return null
    }

    try {
        return JSON.parse(storedUser)
    } catch (e) {
        console.error(`Wrong json fromat of stored user ${storedUser}`, e)
        localStorage.removeItem("user")
    }
    return null
}

export function UserProvider({children}: UserProviderProps) {
    const [user, setUser] = useState<User | null>(getUserFromLocalStorage);

    const setUserWithLocalStorage  = (user: User ) => {
        setUser(user)
    }

    return (
        <UserContext value={{user, setUser: setUserWithLocalStorage}}>
            {children}
        </UserContext>
    )
}