github.com/micro/go-plugins/client/http/v2

go get github.com/micro/go-micro/v2

go get github.com/gin-gonic/gin

go get github.com/micro/go-plugins/registry/consul/v2

go get github.com/micro/protoc-gen-micro

在 protos目录下 protoc --micro_out=. --go_out=. Prods.proto

// go run main.go prod_main.go --server_address  127.0.0.1:8000
// go run main.go prod_main.go --server_address  127.0.0.1:8001



