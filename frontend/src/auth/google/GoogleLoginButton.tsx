import type {CredentialResponse} from "./google-types.ts";
import {useEffect, useRef} from 'react';
import {GetUserContext, type User} from "../User.tsx";
import {Configuration, DefaultApi} from "../../api";

const baseUrl = import.meta.env.VITE_API_BASE_URL;


export default function GoogleLoginButton() {
    const divRef = useRef(null);
    const userContext = GetUserContext()
    useEffect(() => {
        if (window.google && divRef.current) {
            window.google.accounts.id.initialize({
                client_id: '63575078815-kl24b59mf9adslcut5671amaqm05een3.apps.googleusercontent.com',
                callback: handleCredentialResponse,
            });

            window.google.accounts.id.renderButton(divRef.current, {
                theme: 'outline',
                size: 'large',
                text: 'signin_with"',
                shape: "pill"
            });
        }
    }, []);

    async function handleCredentialResponse(response: CredentialResponse) {
        console.log("Google ID token:", response.credential);
        const config = new Configuration({
                "basePath": baseUrl
            }
        )
        const api = new DefaultApi(config)
        const idToken = response.credential; // JWT ID token
        const res  = await api.authGoogleAuthPost({idToken: idToken})



        console.log('Backend response:', user);
    }

    return (
        <div className="flex justify-center mt-6">
            <div className="w-fit " ref={divRef}></div>
        </div>
    );
}
