import LoginScreen from "./auth/LoginScreen.tsx";
import {IsUserLoggedIn} from "./auth/User.tsx";
import {ToastContainer} from "react-toastify";

export function App() {
    return <>
        {IsUserLoggedIn() ? <p>Logged in!</p> : <LoginScreen/>}
        <ToastContainer position="top-right" autoClose={5000} hideProgressBar={false} newestOnTop={false} closeOnClick
                        rtl={false} pauseOnFocusLoss draggable pauseOnHover aria-label={undefined}/>
    </>
}