import axios from "axios";
import https from "https";

export default axios.create({
  baseURL: "https://localhost:443",
  headers: {
    "Content-type": "application/json",
    "Authorization": `Bearer ${}`
  },
  httpsAgent: new https.Agent({
    rejectUnauthorized: false,
  }),
});