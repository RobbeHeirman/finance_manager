import {StrictMode} from 'react'
import {createRoot} from 'react-dom/client'
import './index.css'
import {App} from "./app.tsx";
import {UserContextProvider} from "./auth/user/UserContexProvider.tsx";


createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <UserContextProvider>
            <App/>
        </UserContextProvider>
    </StrictMode>,
)
