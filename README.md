#Testing Go HTTP Endpoints

Running the API server

```
#Make sure that you set your GOPATH correctly
go run src/main.go
```

```
#Creating user
curl -v -X POST localhost:3000/users --data '{"username": "dennis", "balance": 100}' -H 'Content-Type: application/json'
```

```
#Listing users
curl localhost:3000/users
```