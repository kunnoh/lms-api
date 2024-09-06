docker network create lms-net
docker run --network lms-net --name lms-postgres -v ./initdb.sql:/docker-entrypoint-initdb.d/init.sql -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=test -p 5432:5432 -d postgres
docker run --network lms-net -p 8080:8080 adminer