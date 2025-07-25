import type {CredentialResponse} from "./google-types.ts";
import {useEffect, useRef} from 'react';
import {GetUserContext, type User, type UserContextType} from "../User.tsx";
import {getApi} from "../api.ts";
import {toast} from "react-toastify";


export default function GoogleLoginButton() {
    const divRef = useRef(null);
    const userContext = GetUserContext()
    useEffect(() => {
        if (window.google && divRef.current) {
            // noinspection JSUnusedGlobalSymbols
            window.google.accounts.id.initialize({
                client_id: '63575078815-kl24b59mf9adslcut5671amaqm05een3.apps.googleusercontent.com',
                callback: (response: CredentialResponse) => handleCredentialResponse(response, userContext),
            });

            window.google.accounts.id.renderButton(divRef.current, {
                theme: 'outline',
                size: 'large',
                text: 'signin_with"',
                shape: "pill"
            });
        }
    }, [userContext]);
    return (
        <div className="flex justify-center mt-6">
            <div className="w-fit " ref={divRef}></div>
        </div>
    );
}

async function handleCredentialResponse(response: CredentialResponse, userContext: UserContextType) {
    const idToken = response.credential; // JWT ID token
    const api = getApi();
    const res = await api.googleAuth({idToken: idToken});
    const data = res.data

    if (!data.jwtToken) {
        toast.error(`Login failed: Invalid token response from backend`, {
            position: "top-center",
            autoClose: false,
            hideProgressBar: true,
            closeOnClick: true,
            pauseOnHover: true,
            draggable: false,
            progress: undefined,
        });
        return
    }

    const user: User = {
        jwtToken: data.jwtToken ?? "",
        firstName: data.firstName ?? "",
        lastName: data.lastName ?? "",
        userEmail: data?.userEmail ?? "",
        pictureUrl: data?.pictureUrl ?? ""
    }
    userContext.setUser(user)
}
