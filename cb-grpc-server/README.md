# cloudbees-grpc-server

### Build server ###
```
cd .\cb-grpc-server\
go mod download
go build -o server.exe .\cmd\server.go
```

### Starting server  ###
```
.\cb-grpc-server\server --conf=configs
```
