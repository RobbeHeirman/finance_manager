import LoginScreen from "./auth/LoginScreen.tsx";
import {IsUserLoggedIn, UserProvider} from "./auth/User.tsx";

export function App() {
    return <UserProvider>
        {IsUserLoggedIn() ? <LoginScreen/> : <p>Logged in!</p>}
    </UserProvider>
}