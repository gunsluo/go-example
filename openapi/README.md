
# Server

### Go Swagger

```
https://github.com/go-swagger/go-swagger
```


### Write swagger source
```
// swagger:meta

// swagger:route

// swagger:parameters

// swagger:model
```


### Generate a spec from source

```
swagger generate spec -m -o spec/swagger.json \
    -c github.com/gunsluo/go-example/openapi/h \
    -c github.com/gunsluo/go-example/openapi/v
```
 

# Client

###  Install Tools

```
brew install openapi-generator
```

### Generate SDK
```
ory dev swagger sanitize ./spec/swagger.json

swagger validate ./spec/swagger.json

openapi-generator config-help -g go

ory dev openapi migrate \
    --health-path-tags metadata \
    spec/swagger.json spec/api.json

openapi-generator generate -i spec/api.json -g go -o client/api \
    --git-user-id gunsluo \
    --git-repo-id "go-example/openapi/client/api" \
    --git-host github.com \
    -t client/templates/go \
    -c client/go.yml
```


