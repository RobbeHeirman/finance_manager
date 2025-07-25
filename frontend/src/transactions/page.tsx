import type {FormEvent} from "react";
import {toast} from "react-toastify";
import {getApi} from "./api.ts";

const api = getApi()

export function TransactionPage() {
    async function onFileInput(element: FormEvent<HTMLInputElement> | undefined) {
        const file = element?.currentTarget.files?.[0];
        if (file === undefined) {
            toast.error("Could not upload file")
            return
        }
        const response = await api.kbcTransactionsUpload(file);
        console.log(response)
        toast.success(`${file?.name} uploaded successfully`)
    }

    return (
        <input type="file" id="csvfile" name="csvfile" accept="text/csv" onInput={onFileInput}/>
    )
}