https://developers.google.com/protocol-buffers/docs/reference/go-generated

`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

`protoc --go-grpc_out=. --go_out=. ecommerce/product.proto`
