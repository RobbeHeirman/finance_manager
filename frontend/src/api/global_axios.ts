import axios, {type AxiosInstance, type InternalAxiosRequestConfig} from "axios";

import {getUserFromLocalStorage} from "../auth//user/user.ts";

export let apiLogoutHandler: (() => void) | null = null;
const baseUrl = import.meta.env.VITE_API_BASE_URL;

export function setApiLogoutHandler(f: () => void) {
    apiLogoutHandler = f
}

function getBearerToken() {
    return `Bearer ${getUserFromLocalStorage()?.jwtToken}`
}

const apiClient: AxiosInstance = axios.create({
    baseURL: baseUrl,
    headers: {
        "Authorization": getBearerToken()
    }
})

apiClient.interceptors.request.use(
    (config: InternalAxiosRequestConfig) => {
        config.headers.Authorization = getBearerToken()
        return config
    },
    (error) => error
)


apiClient.interceptors.response.use(
    (response) => response, // For successful responses, just pass them through
    (error) => {
        if (!axios.isAxiosError) {
            return Promise.reject(error)
        }
        if (error.response.status === 401) {
            if (apiLogoutHandler !== null) {
                apiLogoutHandler()
                // delete apiClient.defaults.headers.common["Authorization"]
            } else {
                console.warn("logout handler not set. Ignoring")
            }
        }
        return Promise.reject(error);
    }
);

export default apiClient;