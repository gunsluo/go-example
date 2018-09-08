#!/bin/sh

mkdir -p {certs,crl,newcerts}
touch index.txt
echo 1000 > serial


# CA private key (unencrypted)
openssl genrsa -out ca.key 4096
# Certificate Authority (self-signed certificate)
openssl req -config openssl.cnf -new -x509 -days 3650 -sha256 -key ca.key -extensions v3_ca -out ca.crt -subj "/CN=fake-ca"

# Server private key (unencrypted)
openssl genrsa -out server.key 2048
# Server certificate signing request (CSR)
openssl req -config openssl.cnf -new -sha256 -key server.key -out server.csr -subj "/CN=ses.grpc"
# Certificate Authority signs CSR to grant a certificate
openssl ca -batch -config openssl.cnf -extensions server_cert -days 365 -notext -md sha256 -in server.csr -out server.crt -cert ca.crt -keyfile ca.key

# Client private key (unencrypted)
openssl genrsa -out client.key 2048
# Signed client certificate signing request (CSR)
openssl req -config openssl.cnf -new -sha256 -key client.key -out client.csr -subj "/CN=ses.client"
# Certificate Authority signs CSR to grant a certificate
openssl ca -batch -config openssl.cnf -extensions usr_cert -days 365 -notext -md sha256 -in client.csr -out client.crt -cert ca.crt -keyfile ca.key

#openssl x509 -text -noout -in ca.crt

rm *.csr
