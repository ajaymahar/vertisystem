# vertisystem
Verti System assignment.

### Insctuction to build and use the service.
#### 1. Clone the git project.
`git clone https://github.com/ajaymahar/vertisystem.git && cd vertisystem`

#### 2. Build binary and run it.
###  **Linux**

#### 64-bit
```bash
$ GOOS=linux GOARCH=amd64 go build -o vertisystem cmd/main.go&& ./vertisystem
```
#### 32-bit
```bash
$ GOOS=linux GOARCH=386 go build -o vertisystem cmd/main.go&& ./vertisystem 
```


------------



### **MacOS**
#### 64-bit
```bash
$ GOOS=darwin GOARCH=amd64 go build -o  vertisystem cmd/main.go && ./vertisystem
```

#### 32-bit
```bash
$ GOOS=darwin GOARCH=386 go build -o vertisystem cmd/main.go && ./vertisystem
```

------------


### **Windows**
#### 64-bit
```bash
$ GOOS=windows GOARCH=amd6
Verti System assignment 

4 go build -o vertisystem.exe cmd/main.go
```

#### 32-bit
```bash
$ GOOS=windows GOARCH=386 go build -o vertisystem.exe cmd/main.go
```

## API Documentation
- HOST: localhost
- PORT: 8080
- METHOD: POST
- PATH: /api/text
- Payload: {"text": "<strings>"}

Example: 
```bash
$ curl -X POST localhost:8080/api/text -d '{"text":"The Go Playground is a web service that runs on go.devs servers. The service receives a Go program, vets, compiles, links, and runs the program inside a sandbox, then returns the output."}' -v | jq
```
Output:
```bash
Note: Unnecessary use of -X or --request, POST is already inferred.
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0*   Trying ::1:8080...
* Connected to localhost (::1) port 8080 (#0)
> POST /api/text HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.77.0
> Accept: */*
> Content-Length: 197
> Content-Type: application/x-www-form-urlencoded
>
} [197 bytes data]
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Sun, 09 Jan 2022 05:00:37 GMT
< Content-Length: 54
< Content-Type: text/plain; charset=utf-8
<
{ [54 bytes data]
100   251  100    54  100   197   5716  20853 --:--:-- --:--:-- --:--:--  122k
* Connection #0 to host localhost left intact
{
  "job": {
    "id": "b6e02336-0cdb-441d-8a9b-03cd8eda0e07"
  }
}
```

- HOST: localhost
- PORT: 8080
- METHOD: GET
- PATH: /api/text/{id}

Example:
```bash
$ curl localhost:8080/api/text/b6e02336-0cdb-441d-8a9b-03cd8eda0e07 | jq
```

Output:
```bash
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   326  100   326    0     0  33019      0 --:--:-- --:--:-- --:--:--  318k
{
  "job": {
    "ID": "b6e02336-0cdb-441d-8a9b-03cd8eda0e07",
    "Frequency": [
      {
        "key": "the",
        "value": 3
      },
      {
        "key": "and",
        "value": 2
      },
      {
        "key": "runs",
        "value": 2
      },
      {
        "key": "service",
        "value": 2
      },
      {
        "key": "The",
        "value": 2
      },
      {
        "key": "Go",
        "value": 2
      },
      {
        "key": "program",
        "value": 2
      },
      {
        "key": "web",
        "value": 1
      },
      {
        "key": "go.devs",
        "value": 1
      },
      {
        "key": "compiles,",
        "value": 1
      }
    ]
  }
}

