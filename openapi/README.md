###  Install Tools
 
```
brew install openapi-generator
```

### Generate SDK
```
openapi-generator config-help -g go

openapi-generator generate -i petstore.yaml -g go -o sdk \
    --git-user-id ory \
    --git-repo-id "github.com/gunsluo/go-example/openapi/sdk" \
    --git-host github.com \
    -c go.yml
```


### Generate Gin Stub

```
openapi-generator config-help -g go-gin-server

openapi-generator generate -i petstore.yaml -g go-gin-server -o stub --additional-properties=packageName=petstore --additional-properties=packageVersion=1.0.0 --additional-properties=serverPort=8080
```


