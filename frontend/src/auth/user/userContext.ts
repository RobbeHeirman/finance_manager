import type {User} from "./user.ts";
import {createContext, type ReactNode, useContext} from "react";

export type UserContextType = {
    user: User | null;
    setUser: (_: User | null) => void;
};
export const UserContextSupplier = createContext<UserContextType>({
    user: null,
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    setUser: (_) => {
    },
});
export type UserProviderProps = {
    children: ReactNode;
};

export function GetUserContext() {
    return useContext(UserContextSupplier)
}

export function IsUserLoggedIn(): boolean {
    return useContext(UserContextSupplier).user != null;
}