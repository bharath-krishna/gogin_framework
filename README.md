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

## Run behind a reverse proxy
### Oauth2-Proxy

download oauth2-proxy as 
```bash
go get github.com/oauth2-proxy/oauth2-proxy/v7
```

Assuming we are using keycloak as auth provider running at `http://localhost:8099/auth/`

Use below command line arguments

```bash
oauth2-proxy \
--provider=keycloak \
--client-id=<client_id> \
--client-<client_secret> \
--login-url="auth_url" \
--redeem-url="token_url" \
--profile-url="userinfo_url" \
--validate-url="userinfo_url" \
--cookie-secret="cookie_secret_text" \
--email-domain="*" \
--upstream="http://localhost:8088" \ # Upstream api to be protected
--http-address="http://localhost:8888" \ # address for proxy to tun
--cookie-secure=false \
--ssl-upstream-insecure-skip-verify=true \
--redirect-url="http://localhost:8888/oauth2/callback" \ # Proxy callback URL
--pass-access-token=true \
--whitelist-domain="localhost:8099" \ # keycloak URL
--skip-provider-button=true \
--skip-auth-route="/swagger/*,/info" \ #endpoints to skip proxy auth
--pass-access-token=true \ # To call proxy with bearer token for UIs
--skip-jwt-bearer-tokens \ # To call proxy with bearer token for UIs
--extra-jwt-issuers="http://localhost:8099/auth/realms/demo=demo-client" # to validate bearer token
```

Now you can directly call `http://localhost:8888/` with all the endpoints your api provides and proxy endpoints (make sure not overlap)