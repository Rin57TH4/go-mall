
#Khoi tao project

mkdir go-mall
cd  go-mall
go mod init
goctl api new order
goctl api new user
go work init
go work use order
go work use user

#Add Grpc 

Tao file user.proto

syntax ="proto3";


package user;

option go_package="./user";

message IdRequest{
    string id=1;
}


message UserResponse{
    string id =1;
    string name =2;
    string gender =3;

}

service User{
 rpc getUser(IdRequest) returns (UserResponse);
}

Run command:
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

gen xong go mod tidy lai neu can

Gen order api.
goctl api go -api order.api -dir ./gen

