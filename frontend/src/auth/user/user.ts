export const KEY_USER = "user";

export type User = {
    jwtToken: string,
    userEmail: string,
    firstName: string,
    lastName: string
    pictureUrl: string

}

export function setUserToLocalStorage(user: User | null) {
    if (user === null) {
        localStorage.removeItem(KEY_USER)
    } else {
        const userJson = JSON.stringify(user)
        localStorage.setItem(KEY_USER, userJson)
    }
}

export function getUserFromLocalStorage(): User | null {
    const userJson = localStorage.getItem(KEY_USER);
    console.log(userJson)
    if (userJson === null) {
        return null
    }
    console.log(userJson)
    return JSON.parse(userJson)
}
