# vmw-vdp
This is an exercice project using the following tech stack: Golang, Docker, K8s and Prometheus. </br>

## Requirements:
* A service written in golang that queries 2 urls (https://httpstat.us/503 & https://httpstat.us/200) </br>
* The service will check the external urls (https://httpstat.us/503 & https://httpstat.us/200 ) are up (based on http status code 200) and response time in milliseconds </br>
* The service will run a simple http service that produces  metrics (on /metrics) and output a prometheus format when hitting the service /metrics url </br></br>
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
The project is written in GO </br></br>
The project is divided in several packages to make code modular and reusable. It is composed by the following modules: </br>
* **controllers:** This package contains the logic of the endpoints functions</br>
* **environments:** This package contains the environment settings of the project</br>
* **prometheus:** This packate embeds the prometheus functionalities in order for the callback to set the metrics in the prometheus gauges </br>
* **utilities:** This package contains utilities functions which ba be used in other projects if necessary</br>
* **main:** This is the main package, it is composed by the service.go file which contains the main function and the route.go file which is used for manage and routing callback using gorilla. </br>

Given that the project needs to be run locally, on Docker and on Kubernetes, all initialization parameters of the project are taken in input from O.S. environment variables. I will explaines in a next paragraph how to set them for Docker and K8s. For the moment just to mention that the variables needed are the following:
* **host**: This one set the address where teh service will listen/will be binded
* **port**: We need the port as well
* **numUrls**: This represents the number or Urls the service will call. I wanted to extend the requirements of the project and allow it to be able to manage as many urls we want
* **url1 - url2 - url3...**: The url we want to manage (specify as many as you put in numUrls </br>

The service exposes two endpoints:  </br>
* **/service** which is the endpoint calling the urls specified  </br>
* **/metrics** which is the endpoint that will manage the metrics that prometheus will monitor </br> </br>

The framework **Gorilla** has been chosen to make the web-server extensible in order to be able to add new endpoints for different. To add a new endpoint is enough to add the logic inside routes.go and define a new function inside the controllers package </br>

The main endpoint is the /service endpoint when this one is invoked for example with a:</br></br>
**curl localhost:8080/service**</br></br>
The endpoint defined will submit internally an http request to the url specified and still store the result: 0 or 1 depending on the http code returned and store the response time inside a different metric, the endpoint will finally reply with a 200OK to the client.</br>
The endpoint /metrics is by default implemented by the prometheus Handler promhttp.Handler()

## Design:




