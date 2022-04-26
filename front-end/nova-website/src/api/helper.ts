import { AxiosProxyConfig } from 'axios';
// import https from 'https';
// import fs from 'fs';

// export const backEndSource: string = 'http://127.0.0.1:443/';
export const backEndSource: string = '/api/';

// export const httpsProxy: AxiosProxyConfig = {
//     protocol: 'https',
//     host: '127.0.0.1',
//     port: 9000
//   };

// let caCrt: Buffer | undefined = undefined;
// try {
//     caCrt = fs.readFileSync('./server.crt')
// } catch(err) {
//     console.log('Make sure that the CA cert file is named ca.crt', err);
// }

// export const httpsAgent = new https.Agent({
//   ca: caCrt, 
//   keepAlive: false
// });