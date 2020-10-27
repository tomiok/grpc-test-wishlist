
# Protocol buffers
 # A little example about protobuf and Golang

 ## What is it?   
    - IDL = interface description language
    - Multilenguage
    - Optimized serialization (faster than XML, JSON)
    - Less error prone

 ## What do we need 
    - a .proto file
    - a compiler for protobuffers  => Download here: https://grpc.io/blog/installation/
    
 ## Documentation:
    - https://developers.google.com/protocol-buffers/docs/overview#simple

## Command
`protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/wishlist.proto`
