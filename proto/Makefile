.PHONY: all
all: protogen

.PHONY: protogen
protogen:
	echo "## generating Go stubs ##"
	protoc -I . --go_out=plugins=grpc:. myftp.proto

clean:
	rm *.pb.go