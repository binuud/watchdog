# API

* HTTP server enabled, [OpenApiSpec](../gen/web/v1/watchdog/openapi.json)
* GRPC enabled, reflection enabled by default
* Typescript interface at /gen/web/v1/watchdog/*.ts
* Docker image by default starts in server mode
* For looking at OpenApiSpec, clone the project, and run below command, from root of project.

```
docker run  -p 10030:8080 -v ./gen/web/v1/watchdog/openapi.json:/tmp/swagger.json -e SWAGGER_FILE=/tmp/swagger.json docker.swagger.io/swaggerapi/swagger-editor
```
*  After running above command, launch browser http://localhost:10030/


## Curl Commands

Once server is running, REST api can be used for getting details of domains, refer to swagger interface 
for details of the REST api's
```
curl -X GET 'http://localhost:9080/v1/watchdog/getAll?page=1&perPage=10' | jq
```

```
curl -X GET 'http://localhost:9080/v1/watchdog/getProjects?page=1&perPage=10' | jq
```

Get Details of a particular domain
```
curl -X GET 'http://localhost:9080/v1/watchdog/get?name=www.google.com' | jq
```

ReFetch all the information in the background (non blocking)
```
curl -X POST 'http://localhost:9080/v1/watchdog/reload' -d '{}'
```