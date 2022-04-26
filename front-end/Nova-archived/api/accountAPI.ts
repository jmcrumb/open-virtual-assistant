import axios from "axios";
import { backEndSource } from "./helper";

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

    public static async createAccount(fname: string, lname: string, email: string, password: string): Promise<any> {
        try {
            const response = await axios.post(`${backEndSource}account/`, {
                first_name: fname,
                last_name: lname,
                password: password,
                email: email
            }, 
            // {
            //     httpsAgent: httpsAgent
            // }
            );
            return new Account(response.data);
        } catch (error) {
            return error;
        }
    }

    public static async getAccountByID(id: string): Promise<any> {
        try {
            const response = await axios.get(`${backEndSource}account/${id}`);
            return new Account(response.data);
        } catch (error) {
            return error;
        }
    }

    public static async updateAccount(account: Account): Promise<any> {
        try {
            const response = await axios.put(`${backEndSource}account/`, JSON.stringify(account));
        } catch (error) {
            return error;
        }
    }

    public static async getProfileByID(id: string): Promise<any> {
        try {
            const response = await axios.get(`${backEndSource}account/profile/${id}`);
            return new Profile(response.data);
        } catch (error) {
            return error;
        }
    }

    public static async updateProfile(profile: Profile) {
        try {
            const response = await axios.put(`${backEndSource}account/profile/`, JSON.stringify(profile));
        } catch (error) {
            return error;
        }
    }
}