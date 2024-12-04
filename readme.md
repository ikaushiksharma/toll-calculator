# toll-calculator-go-microservice

![diagram](https://github.com/user-attachments/assets/75d5a41f-b8d8-433b-bbb3-eed7f2e5100a)

## Project Overview

In the main.go file, We will simulate an OBU - Onboard unit that sits in the truck and that's going to send out GPS conrdinates after each interval and we are going to send that/ replicate that using some kind of web sockets connection that will basically send these messages over webs sockets and we are gonna receive that in ou 1st microservice and put them on kafka. Now the another microservice i.e Distance calculator will use these corrdinates and calculate the distance Travelled by the vehichle and will send it to the invoicer (the another microservice) and now the Invoicer will send data to Invoice calculator to calculate and generate an Innvoive and then it will send it back to invoicer. Also this invoicer is connected to a database so it will store the generated invoice in the db and also query it back and send it to the client via the API Gateway.

Note : We have kept invoice calculator service isolated and treating it as a standalone microservice because what if someone wants to calculate the the amount of money they need to pay for travelling from place A to place B, In that case it will do the calculation and send it to user directly.

## Project Dependencies

### websoket

```
go get github.com/gorilla/websocket
```

#### [Kafka](https://github.com/confluentinc/confluent-kafka-go)

#### Kafka Go-client

```
go get github.com/confluentinc/confluent-kafka-go/v2/kafka
```

#### kafka docker installation

```
docker-compose up -d
```

### Logger

```bash
go get github.com/sirupsen/logrus
```

## How to run?
- run this command : `docker-compose up -d` to run the kafka and zookeeper on your local machine
- run these commands seperately
```
make receiver
make obu
make calculator
make agg
```
- Now you can open your thunderclient and check both the REST API's
  1. `/invoice` : This calculates the invoice of the  given OBU ID in the query
  - endpoint : `http://localhost:3000/invoice?obu=6428921451518044973`
  - METHOD : `GET`
  - Response body : 
    ```
    {
        "obuID": 6428921451518044973,
        "totalDistance": 35.86001233283358,
        "totalAmount": 112.95903884842576
    }
    ```
 2. `/aggregate` : This is used to aggregate all the data coming from distance calculator automatically, but you can also send some data as JSON if you want to do it manually here as:
    - endpoint : `http://localhost:3000/aggregate`
    - METHOD : `POST`
    - Request Body : 
        ```
        {
            "value": 20.12,
            "obuID": 1838,
            "unix": 73378
        }
        ```
    - Response (Server log):
        ```
       
        HTTP Transport running at port :3000...
        INFO[0003] aggregating distance         distance=20.12 obuid=1838 unix=73378
        INFO[0003] AggregateDistance            err="<nil>" took="47.03µs"
        
        ```
