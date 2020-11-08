# vmw-vdp
This is an exercice project using the following tech stack: Golang, Docker, K8s and Prometheus. </br>

## Requirements:
A service written in golang that queries 2 urls (https://httpstat.us/503 & https://httpstat.us/200) </br>
The service will check the external urls (https://httpstat.us/503 & https://httpstat.us/200 ) are up (based on http status code 200) and response time in milliseconds </br>
The service will run a simple http service that produces  metrics (on /metrics) and output a prometheus format when hitting the service /metrics url </br></br>
Expected response format:

```
sample_external_url_up{url="https://httpstat.us/503 "}  = 0
sample_external_url_response_ms{url="https://httpstat.us/503 "}  = [value]
sample_external_url_up{url="https://httpstat.us/200 "}  = 1
sample_external_url_response_ms{url="https://httpstat.us/200 "}  = [value]
```
Others requirements:
* Dockerfile to build image
* Kubernetes Deployment Specification to deploy Image to Kubernetes Cluster
* Unit Tests

## Design:
The project is written in GO </br>
The project is divided in several packages to make code modular and reusable. It is composed by the following modules: </br>
**controllers**: This package contains the logic of the endpoints functions</br>
**environments**: This package contains the environment settings of the project</br>
**prometheus**: This packate embed the prometheus functionalities in order for the callback to set the metrics </br>
**utilities**: This package contains utilities functions which ba be used in other projects if necessary</br>
**main**: This is the main package, it is composed by the service.go file which contains the main function and the route.go file which is used for manage and routing callback using gorilla. </br>

The service exposes two endpoints:  </br>
**/service** which is the endpoint calling the urls specified  </br>
**/metrics** which is the endpoint that will manage the metrics that prometheus will monitor </br> </br>
The library Gorilla has been chosen to make the web-server extensible in order to be able to add new endpoints for different

