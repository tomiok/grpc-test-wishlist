
*** Protocol buffers

    IDL = interface description language
    Multilenguaje
    Serializacion optimizada (mas rapido que XML y json, menos propenso a errores)

    *** Que necesitamos?
    - un archivo .proto
    - compilador (protoc) => descargar aca: https://grpc.io/blog/installation/
    
documentacion: https://developers.google.com/protocol-buffers/docs/overview#simple

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/wishlist.proto