
.PHONY: build
build:
	@cd src/go && GOOS=js GOARCH=wasm go build -o add.wasm add.go
	@mv src/go/*.wasm public
	@npm run build

