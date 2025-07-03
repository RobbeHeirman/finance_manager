import type {CredentialResponse} from "./google-types.ts";
const baseUrl = import.meta.env.VITE_API_BASE_URL;


import { useEffect, useRef } from 'react';
import {SetUser, type User} from "../User.tsx";

export default function GoogleLoginButton() {
    const divRef = useRef(null);

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
        const idToken = response.credential; // JWT ID token
        const res = await fetch(`${baseUrl}/auth/google_auth`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ idToken }),
        });
        const user = await res.json() as User;
        SetUser(user)
        console.log('Backend response:', user);
    }

    return (
        <div className="flex justify-center mt-6">
            <div className="w-fit " ref={divRef}></div>
        </div>
    );
}
