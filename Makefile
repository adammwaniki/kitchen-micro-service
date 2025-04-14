export PATH := /home/adam/go/bin:$(PATH) # Need to find an alternative to hard coding the path to my protoc

$(info PATH is $(PATH))

gen:
	@protoc \
		--proto_path=protobuf "protobuf/orders.proto" \
		--go_out=services/common/genproto/orders --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/orders --go-grpc_opt=paths=source_relative
