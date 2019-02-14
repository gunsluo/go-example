
build-grpc:
	@cd pb && protoc -I=. --go_out=plugins=grpc:. helloworld.proto

build-grpc-web:
	@cd web && npm install && npx webpack lib/client.js

gen-grpc-web:
	@protoc -I=pb pb/helloworld.proto --js_out=import_style=commonjs:./web/lib --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./web/lib

clean:

.PHONY: build-grpc build-grpc-web gen-grpc-web clean
