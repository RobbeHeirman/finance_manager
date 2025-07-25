import {Configuration, TransactionsApi} from "../api/generated";

const baseUrl = import.meta.env.VITE_API_BASE_URL;

export function getApi(): TransactionsApi {
    const config = new Configuration({
            "basePath": baseUrl
        }
    )
    return new TransactionsApi(config)
}