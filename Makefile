greeter: hello.pb.go *.go
	go build -o $@ .

hello.pb.go: hello.proto
	protoc --go_out=plugins=grpc:. $<

clean:
	rm hello.pb.go greeter
