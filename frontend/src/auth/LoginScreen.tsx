import GoogleLoginButton from "./google/GoogleLoginButton.tsx";
import {GetUserContext} from "./User.tsx";

export default function LoginScreen() {
    const userContext = GetUserContext()
    console.log(userContext)
    return <>
        <GoogleLoginButton/>
    </>
}