# cb-services
Cloudbees grpc services

### Build server ###
```
cd .\cb-grpc-server\
go mod download
go build -o server.exe .\cmd\server.go
```

### Build client ###
```
cd .\cb-grpc-client\
go mod download
go build -o client.exe .\cmd\client.go
```


### Starting server  ###
```
.\cb-grpc-server\server --conf=configs
```

### Test gRPC connection ###
```
.\cb-grpc-client\client.exe test --conf=.\configs\
```

### Booking Ticket ###
```
.\cb-grpc-client\client.exe book --conf=.\configs\ --lname="colt" --fname="jay" --email="jaeee@colt.in"
```

### View Seat Arrangement ###
```
.\cb-grpc-client\client.exe view --conf=.\configs\ --section="A"
```

### View booking receipt ###
```
.\cb-grpc-client\client.exe receipt --conf=.\configs\ --userID="6361438830677966656"
```

### Remove user from train ###
```
.\cb-grpc-client\client.exe remove --conf=.\configs\ --userID="6361438830677966656"
```

### Modify seat ###
```
.\cb-grpc-client\client.exe modify --conf=.\configs\ --seat="2" --userID="8436004956245676728"
```