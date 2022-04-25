rm *.pem

openssl genrsa -out client.key.pem 4096

# Generate CA's private key and self-signed certificate
openssl req -new -key client.key.pem -out client.csr -subj "/C=US/ST=California/L=San Diego/O=USD/OU=NOVA/CN=localhost:8080/emailAddress=jcrumb@sandiego.edu"

openssl x509 -req -in client.csr -passin pass:usdnova -CA /root/tls/intermediate/certs/ca-chain-bundle.cert.pem -CAkey /root/tls/intermediate/private/intermediate.cakey.pem -out client.cert.pem -CAcreateserial -days 365 -sha256 -extfile client_cert_ext.cnf
