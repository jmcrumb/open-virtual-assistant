export class Account {
  id: string;
  email: string;
  first_name: string;
  last_name: string;
  password: string;
  last_login: Date;
  date_joined: Date;
  is_admin: boolean;

  constructor(json: { [key: string]: any }) {
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

  constructor(json: { [key: string]: any }) {
    this.account_id = json.account_id;
    this.bio = json.bio;
    this.photo = json.photo;
  }
}

export class PublicProfile extends Profile {
  first_name: string;
  last_name: string;

  constructor(json: { [key: string]: any }) {
    super(json);
    this.first_name = json.first_name;
    this.last_name = json.last_name;
  }
}

export function useQueryAccountByID() {
  const accountID = "29b744e6-0a2f-48a9-aeb0-1a162f546763";
  return useQuery(["account", accountID], async () => {
    const { data } = await axios.get(
      `https://127.0.0.1:443/account/${accountID}`
    );
    return new Account(data);
  });
}

