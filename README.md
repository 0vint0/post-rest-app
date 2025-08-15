## Build docker image ##
docker build -t post-rest-app:latest -f ./docker/Dockerfile .


## Start App in Docker ##
docker run -d -p 8080:8080 post-rest-app:latest

## Remove stoped containers
docker container prune