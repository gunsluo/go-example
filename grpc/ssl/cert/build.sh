
----------------------------

ca: 

openssl req -newkey rsa:2048 -nodes -sha256 -keyout ca.key -x509 -days 365 -out ca.crt

openssl req -newkey rsa:2048 -nodes -sha256 -keyout server.key -new -out server.csr

openssl x509 -CA ca.crt -CAkey ca.key -in server.csr -req -days 365 -out server.crt -CAcreateserial -sha256


check: 

openssl x509 -text -noout -in ca.crt

openssl x509 -text -noout -in server.crt

