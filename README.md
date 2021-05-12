# GoGin Framework
## Boilerplate Code
This is complete framework for quick building of backend apis using GoLang and GoGin framework.
It includes
- logging
- command line tools
- Swagger Specs for apis
- all required modules in go.mod file



## How to run application
```bash
> swag init; rm ./app; go build -o app; ./app
```
Once started go to [http://localhost:8088/swagger/index.html](http://localhost:8088/swagger/index.html) for swagger specs.

You can set address and port to listen on by passing "address:port" (0.0.0.0:8088) as command line args as shown below
```bash
> ./app --addr 0.0.0.0:8088
```
OR
set GOGIN_ADDRESS environment variable with 0.0.0.0:8088

```bash
> export GOGIN_ADDRESS=0.0.0.0:8088
```

check ./app -h for more command like args

## Deploy as docker container (Recommended)
Run below commands
```bash
> doclker build -t gogin_framework .
> docker run -d --name gogin_framework -p 8088:8088 -e GOGIN_ADDRESS=0.0.0.0:8088 gogin_framework
```

To stop the container run
```basah
> docker stop gogine_framework
> docker rm gogine_framework
```