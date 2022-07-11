rm *.pem

openssl req -x509 -newkey rsa:4096 -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=UZ/ST=Tashkent/L=Chilonzor/O=No/OU=No/CN=jahongir/emailAddress=anorboev.jahongir8007@gmail.com"

openssl req -newkey rsa:4096 -keyout server-key.pem -out server-req.pem -subj "/C=UZ/ST=Tashkent/L=Chilonzor/O=No/OU=No/CN=jahongir/emailAddress=anorboev.jahongir8007@gmail.com"

openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cet.pem 

