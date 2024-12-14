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

Export variables to environment from `.env` file.

```sh
. env.sh
```

Genarate ECDSA keys used to generate `jwt` token. Save them in `./keys` folder.
Create private key.

```sh
openssl ecparam -genkey -name prime256v1 -noout -out ./keys/ecdsa_private_key.pem
```

Generate public key

```sh
openssl ec -in ecdsa_private_key.pem -pubout -out ./keys/ecdsa_public_key.pem
```

NOTE:  
Make sure you have postgres database running.  
Run `postgres` container using `docker`:

```sh
make start-db
```

**Start app**

```sh
make dev
```

## Run tests

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

## Testing

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
