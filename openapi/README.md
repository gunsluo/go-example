###

```
brew install openapi-generator
```

```
openapi-generator config-help -g go

openapi-generator generate -i petstore.yaml -g go -o sdk --additional-properties=packageName=petstore --additional-properties=packageVersion=1.0.0 --additional-properties=isGoSubmodule=true
```
