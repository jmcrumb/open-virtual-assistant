import React from "react";
import { GlobalStateContext } from "../globalState";


export const BACKEND_SRC = "https://127.0.0.1:443/";

export function getAxiosHeaders() {
    let axiosHdrs = {
        "Content-type": "application/json"
    };
    const context = React.useContext(GlobalStateContext);

    if(context.token != "") {
        axiosHdrs["Authorization"] = `Bearer ${context.token}`;
    }

    return axiosHdrs;
}