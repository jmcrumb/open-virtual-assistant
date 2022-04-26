import axios from "axios";
import https from "https";

export default axios.create({
  baseURL: "https://localhost:443",
  headers: {
    "Content-type": "application/json",
    "Access-Control-Allow-Origin": "http://localhost:8080",
    "Access-Control-Allow-Credentials": true
  },
  httpsAgent: new https.Agent({  
    rejectUnauthorized: false
  })
});