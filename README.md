# TaskSync
Developed a distributed, fault-tolerant job scheduler that uses Redis to assign jobs.
## Getting Started

* Install go-lang
* Install dependencies
```
go get -u github.com/go-redis/redis
go get -u github.com/streadway/amqp
```
* Build the project
```
go build
```
## To create a job : 
```
POST / HTTP/1.1
Host: localhost:3000
Content-Type: application/json
{
  "jobName": "FirstTask",
  "runTime": 1899142237,
  "repeatAfterSec": null,
  "destIP": null,
  "destPort": null,
  "destExchange": null
}
```