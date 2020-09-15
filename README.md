# 買い物　（Kaimono）
- This personal project is meant to be an easily extendable WebShop, Store and showcase the use of GRPC and WebGRPC in modern applications using Go language for server code and Angular for the WebUI as a Single Page Application.

My main goals with this project are:
- Zero configuration and easy deployment using Docker
- Integration with Stripe
- Article builder UI

#### Environment variables:
- SHOP_BIND_ADDR
- GRPC_ADDR
- GRPC_PORT

##### Update envs.txt then:
- Run: `. envs.txt`

### Example protoc generate:
- cd to protobuf folder
command explained:`grpc:kaimono` - refers to the package definition, without specifing go_out=plugins=grpc you do not get a gRPC server generated!
- run: `protoc --proto_path= --go_out=plugins=grpc:kaimono kaimono.proto`