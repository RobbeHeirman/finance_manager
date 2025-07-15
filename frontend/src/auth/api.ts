import {AuthApi, Configuration} from "../api/generated";

const baseUrl = import.meta.env.VITE_API_BASE_URL;

export function getApi(): AuthApi {
    const config = new Configuration({
            "basePath": baseUrl
        }
    )
    return new AuthApi(config)
}