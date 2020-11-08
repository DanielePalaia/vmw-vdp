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

## Design:
The project is written in GO </br>
It exposes two endpoints:  </br>
/service which is the endpoint calling the urls specified  </br>
/metrics which is the endpoint that will manage the metrics that prometheus will monitor </br> </br>

