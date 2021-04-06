# grpc-go-distributed-tracing-example
Small example of distributed tracing in grpc-go applications. Using [OpenCensus](https://opencensus.io/) for collecting traces data and [Jaeger](https://www.jaegertracing.io/) for export.

## Overview
Example has 3 parts. 

Client:
* reads "radius" from console
* calls "Circle.area" grpc method to calculate an area of a circle with given radius

Circle:
* Circle.area calls Math.sqr grpc method to calculate the radius squared
* Circle.area returns value of area

Math:
* Math.sqr returns the radius squared

## Prerequisites
1) Start Jaeger Exporter locally in Docker
```dockerfile
docker run -d --name jaeger \
  -p 16686:16686 \
  -p 14268:14268 \
  jaegertracing/all-in-one:1.22
```
* 16686 - Jaeger UI port
* 14268 - Collector port

2) Run main func of each part
3) See your traces in http://localhost:16686/

## Example
![example](https://user-images.githubusercontent.com/12141268/113720557-74775c00-96f7-11eb-9011-af52bb381dcc.png)
