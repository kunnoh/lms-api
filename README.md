# Library Management System


Clone `lms-api` repository.
```sh
git clone https://github.com/kunnoh/lms-api.git
```

Go to directory. And run `./run.sh help` to see available options.
```sh
cd lms-api
```
```sh
./run.sh
```

## Configuration  
Change `example.env` to `.env` at the root folder of project.
Set the information in `.env`.

Export variables to environment from `.env` file to the current terminal session.

```sh
./env.sh
```


## JWT key generation
Genarate ECDSA keys used to generate `jwt` token. Save them in `./keys` folder.

1. Create private key.
```sh
openssl ecparam -genkey -name prime256v1 -noout -out ./keys/ecdsa_private_key.pem
```

2. Generate public key
```sh
openssl ec -in ./keys/ecdsa_private_key.pem -pubout -out ./keys/ecdsa_public_key.pem
```

NOTE:  
Make sure you have postgres database running.  
Run `postgres` container using `docker`:

```sh
./run.sh start-db
```

**Start app**
Run the app for development.
```sh
./run.sh dev
```

## Run tests
```sh
./run.sh test
```

## Build

**Build executable**
Build app.
```sh
./run.sh build-app
```

**Build docker image**

```sh
./run.sh build-image
```

## API testing

Register.

```sh
curl -X POST http://localhost:7755/auth/register -H "Content-Type: application/json" -d '{
  "Email": "example@example.com",
  "Name": "John Doe",
  "Password": "securepassword123",
  "Phone": "+1234567890",
  "IdNumber": "912345678"
}'
```

Login.

```sh
curl -X POST http://localhost:7755/auth/login -H "Content-Type: application/json" -d '{
  "Email": "example@example.com",
  "Password": "securepassword123"
}'
```
