openssl genrsa -out client.key 2048
openssl req -new -x509 -key client.key -out client.pem -days 3650

openssl genrsa -out server.key 2048
openssl req -new -x509 -key server.key -out server.pem -days 3650

Make CA:
$ openssl genrsa -out rootCA.key 2048
$ openssl req -x509 -new -nodes -key rootCA.key -out rootCA.pem -days 1024
... install that to Firefox
Make cert:
$ openssl genrsa -out server.key 2048
$ openssl req -new -key server.key -out server.csr
$ openssl x509 -req -in server.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out server.crt -days 500