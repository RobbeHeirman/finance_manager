import {Configuration, TransactionsApi} from "../api/generated";
import apiClient from "../api/global_axios.ts";

const baseUrl = import.meta.env.VITE_API_BASE_URL;

export function getApi(): TransactionsApi {
    const config = new Configuration({
            "basePath": baseUrl
        }
    )
    return new TransactionsApi(config, baseUrl, apiClient)
}