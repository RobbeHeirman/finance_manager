import {StrictMode} from 'react'
import {createRoot} from 'react-dom/client'
import './index.css'
import {App} from "./app.tsx";
import {UserProvider} from "./auth/User.tsx";


createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <UserProvider>
            <App/>
        </UserProvider>
    </StrictMode>,
)
