import UserState from "../userState";

export const BACKEND_SRC = "https://127.0.0.1:443/";

export function getAxiosHeaders() {
    let axiosHdrs = {
        "Content-type": "application/json"
    };
    const userState: UserState = UserState.getInstance();

    if("jwt_auth_token" in userState.state) {
        axiosHdrs["Authorization"] = `Bearer ${userState.state["jwt_auth_token"]}`;
    }

    return axiosHdrs;
}