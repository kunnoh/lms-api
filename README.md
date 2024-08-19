# Library Management System

## Clone repo

```sh
git clone https://github.com/kunnoh/lms-api.git
```

goto directory
```sh
cd lms-api
```
## Run app
Change `example.env` to `.env` at the root folder of project
Set the information in `.env`

NOTE:
Make sure you have postgres database running
Run using postgres container using `docker`:
```sh
docker run --name lms-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=test -p 5432:5432 -d postgres
```


**development mode**
```sh
make dev
```

**Build app**
```sh
make build
```



