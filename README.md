# tureloGo - RESTful API

[![Go Report Card](https://goreportcard.com/badge/github.com/kebabmane/tureloGo)](https://goreportcard.com/report/github.com/kebabmane/tureloGo)

[![CircleCI](https://circleci.com/gh/kebabmane/tureloGo.svg?style=svg)](https://circleci.com/gh/kebabmane/tureloGo)


## Pull and Use Local DynamoDB

I used https://github.com/dwmkerr/docker-dynamodb

To setup an ephermeral instance (which I generally use for testing)


```
docker pull dwmkerr/dynamodb
docker run -p 8000:8000 dwmkerr/dynamodb
```

This instance will be available on port 8000 localhost


## Seed Local DB Instance

```
cd dbSeed
go run create_sessions.go
```

This will create the requried sessions table that is used in the auth middleware later, if required this could be swapped away from DynamoDB to REDIS or simliar relativly easyily


## Environment Config

Open up dev.sh, here we are exporting some session vars for use by the microservice - you may note in server.go we are loading these into a struct and validaing that some mandatory items are populated or else we exit immeditly.

To ensure there is no dependency in the app its'self for hard coded files, we specificly use OS vars only.

If you wanted to create a differnet set of vars, clone dev.sh into something like test.sh and change out the exported vars

## Logging

By default we log into a JSON format, below is an example of the server starting up and a consumer hitting the health endpoint

```
INFO[0000] {Port:8080 IsProduction:false HealthName:tureloGo AwsRegion:us-west-2 NegroniLoggerName:web VerboseLogging:true}
 
INFO[0006] started handling request                      method=GET remote=[::1]:59862 request=/health
{"level":"info","method":"GET","msg":"started handling request","remote":"[::1]:59862","request":"/health","time":"2018-02-20T21:28:52+11:00"}
{"level":"info","measure#web.latency":539389,"method":"GET","msg":"completed handling request","remote":"[::1]:59862","request":"/health","status":200,"text_status":"OK","time":"2018-02-20T21:28:52+11:00","took":539389}
INFO[0006] completed handling request                    measure#web.latency=755551 method=GET remote=[::1]:59862 request=/health status=200 text_status=OK took=755.551µs
[negroni] 2018-02-20T21:28:52+11:00 | 200 | 	 941.811µs | localhost:8080 | GET /health 
```

A flag is in our env config to enable more verbose logging (boolean VerboseLogging) to enable more detailed logging

## Health Check

In this skeleton I have used github.com/InVisionApp/go-health for health checks, it provides an enapoint we can expose that gives us multiple health checks with an overall status to make it easy to view whats going on, a basic good and bad health check have been implemented to prove a point, however the overall endpoint will return a status 200 because the good known health check is just getting to google.com

```
{
    "details": {
        "bad-check": {
            "name": "bad-check",
            "status": "failed",
            "error": "Ran into error while performing 'GET' request: Get google.com: unsupported protocol scheme \"\"",
            "check_time": "2018-02-21T10:08:42.944697967+11:00"
        },
        "good-check": {
            "name": "good-check",
            "status": "ok",
            "check_time": "2018-02-21T10:08:43.158383591+11:00"
        }
    },
    "status": "ok"
}
``` 

## Auth Middleware

In this skeleton i'am using negroni middle in server.go that sits in the middle of every single API call to the app and verifies the JWT - at this stage it creates some delay as we are not caching the session so it's going through the check every single time.

```
myToken := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
```       

Make sure every API call has the Authorization header, if not the request is automatially set to unauthorized

## Session Management

DynamoDB is used for session management, if we have seen the session before and it doesnt appear to be tampered then no need to very the JWT, simply let the user through.

```
SAMPLE
```

## Docker

To build the slimiest executate possible in a docker contrainer run

```
docker build -t turelogo/development .
```

This is a multistage dockerfile that,
* spins up a container to build the binary in - this container will be larger then required to operate in production (~1GB)
* create another container based off alpine that is as tiny as possible to run the binary

## Run the server

./dev.sh