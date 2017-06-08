# catchinspect

Simple go program for inspecting http requests.

[![Go Report Card](https://goreportcard.com/badge/github.com/delabroj/catchinspect)](https://goreportcard.com/report/github.com/delabroj/catchinspect)

## Usage

Start service:

`go run catchinspect.go`

### Catch Service

Send request to http://localhost:8080

### Inspect Service

The most recent request made to the catch service can be seen at http://localhost:8081

### Example

#### Catch

```shell
curl -X POST --url localhost:8080/?url1=hello -d "body1=world"
```

#### Inspect (http://localhost:8081)

```
Form Values: url.Values{"url1":[]string{"hello"}, "body1":[]string{"world"}}

Remote Address: 127.0.0.1:53814

POST /?url1=hello HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 11
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/7.54.0

body1=world
```
