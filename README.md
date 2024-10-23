# Library Management System
## Run `lms-api`
Clone repo.
```sh
git clone https://github.com/kunnoh/lms-api.git
```

Go to directory.
```sh
cd lms-api
```



**Configuration**  
Change `example.env` to `.env` at the root folder of project.
Set the information in `.env`.

Genarate ECDSA keys used to generate `jwt` token. Save them in `./keys` folder.
Create private key.
```sh
openssl ecparam -genkey -name prime256v1 -noout -out ecdsa_private_key.pem
```

Generate public key
```sh
openssl ec -in ecdsa_private_key.pem -pubout -out ecdsa_public_key.pem
```
  

NOTE:  
Make sure you have postgres database running.  
Run using postgres container using `docker`:
```sh
docker run --name lms-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=test -p 5432:5432 -d postgres
```
  
  
## Build

Build app.  
```sh
make build
```
  
or

```sh
go build -o ./bin/
```

  
Create `docker` image.  
```sh
make build-image
```

or


```sh
docker build -t lms-api:latest .
```


**Start app**
```sh
make dev
```

or

```sh
docker run --name lms-app -p 7755:7755 lms-api:latest
```


## Testing
Register.  
```sh
curl -X POST http://192.168.39.2:30001/auth/register -H "Content-Type: application/json" -d '{
  "Email": "example@example.com", 
  "Name": "John Doe",
  "Password": "securepassword123",
  "Phone": "+1234567890",
  "IdNumber": "ID12345678"
}'
```

Login.  
```sh
curl -X POST http://192.168.39.2:30001/auth/login -H "Content-Type: application/json" -d '{
  "Email": "example@example.com",
  "Password": "securepassword123"
}'
```