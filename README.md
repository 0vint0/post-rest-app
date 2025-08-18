## To Generate mocked repositories
mockgen -source=persistent/repository/repositories.go -destination=persistent/repository/mocks/repositories_mock.go -package=mocks

## Build docker image ##
docker build -t post-rest-app:latest -f ./docker/Dockerfile .


## Start App in Docker ##
docker run -d -p 8080:8080 post-rest-app:latest

## Remove stoped containers
docker container prune

## Run all tests
go test ./...

## Start Docker Compose with Grafana+Prometheus
docker-compose -f docker-compose.yml up -d

