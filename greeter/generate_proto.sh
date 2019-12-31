# !/bin/bash
# Command used to generate .pb.go file from a .proto file
protoc greetpb/greet.proto --go_out=plugins=grpc:.