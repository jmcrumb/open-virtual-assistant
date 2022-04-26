import axios from "axios";
import { useQuery } from "react-query";
import { backEndSource } from "./helper";
import https from "./http-common";

// use react query, but note, designed to be used with react 17 so some issues may occur

export class Account {
    id: string;
    email: string;
    first_name: string;
    last_name: string;
    password: string;
    last_login: Date;
    date_joined: Date;
    is_admin: boolean;

    constructor(json: {[key: string]: any}) {
        this.id = json.id;
        this.first_name = json.first_name;
        this.last_name = json.last_name;
        this.email = json.email;
        this.last_login = new Date(json.last_login);
        this.password = json.password;
        this.date_joined = new Date(json.date_joined);
        this.is_admin = json.is_admin;
    }
}

export class Profile {
    account_id: string;
    bio: string;
    photo: any;

    constructor(json: {[key: string]: any}) {
        this.account_id = json.account_id;
        this.bio = json.bio;
        this.photo = json.photo;
    }
}

export class AccountAPI {
    
    public static useQueryAccountByID() {
        const accountID = "425bb481-dd67-41e5-820a-5ec0c6cc2bbd"
        return useQuery(["account", accountID], async () => {
          const { data } = await axios.get(
            `https://127.0.0.1:443/account/${accountID}`
          );
          return new Account(data);
        });
      }
}