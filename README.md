# API-Gateway

Example API-Gateway with proxy microservice and authentication microservice.

## Usage

To run the application on [docker](https://www.docker.com/), you will need:
- installed docker
- run ```make start-docker-compose```
- test app using postman with collection ```./Postman/API-Gateway.postman_collection.json```


## Swagger
If yor wont to see swagger documentation, yor will need:
- installed [redoc-cli](https://redocly.com/docs/redoc/deployment/cli/)
- run ```make swagger-user-serve``` for user service documentation
- run ```make swagger-authenticator-serve``` for authenticator service documentation
- open http://127.0.0.1:8080 on your browser
