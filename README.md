# Mandatory 4
To run our program, cd into the mandatory3 directory
```console
cd mandatory4
```
Then to start a client run this command
```console
go run mutex_client.go <client-id> <client-ip> <known-peer-ip>
```
Where:
- client-id is the name of the client
- client-ip is the listening address of the client
- know-peer-ip is an address of a known peer in the network

> **_NOTE:_** that the first client has to be started without a known-peer-ip, as there aren't any other clients on the network yet

### These were the commands we used in the logs
```console
go run mutex_client.go client1 localhost:5001
```
```console
go run mutex_client.go client2 localhost:5002 localhost:5001
```
```console
go run mutex_client.go client3 localhost:5003 localhost:5002
```
```console
go run mutex_client.go client4 localhost:5004 localhost:5001
```

# Mandatory 3
To run our program, cd into the mandatory3 directory
```console
cd mandatory3
```
Then run the server first

```console
go run server.go
```
and the client
```console
go run client.go
```
To disconnect a client, simply stop the client program (ctrl+c)
