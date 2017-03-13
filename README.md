## grpc + beego  develop rpc api
~~~run server
    go build
    ./scloud-grpc
~~~

~~~run client
    go run client/hello/main.go
~~~

## use protoc 
~~~
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
~~~

## create proto

## create rpc controllers
