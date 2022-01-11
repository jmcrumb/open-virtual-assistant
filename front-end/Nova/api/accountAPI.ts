import { Profiler } from "react";

export class User {
    id: string;
    username: string;
    firstName: string;
    lastName: string;
    email: string;
    profile: Profile;

    constructor(json: {[key: string]: any}) {
        this.id = json.id;
        this.username = json.username;
        this.firstName = json.firstName;
        this.lastName = json.lastName;
        this.email = json.email;
    }
}

export class Profile {
    userId: string;
    bio: string;
    photo: any;

    constructor(json: {[key: string]: any}) {
        this.userId = json.userId;
        this.bio = json.bio;
        this.photo = json.photo;
    }
}

export class AccountAPI {

}