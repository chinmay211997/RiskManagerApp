# RiskManagerApp

The Project Implements the below endpoints for the Risk Manager Service

- GET `/v1/risks` (List all risks)
- POST `/v1/risks` (Create a risk)
- GET `/v1/risks/{Id}` (Get a risk by Id)
- GET `/_health` (Get health of service)

The service is hosted on `127.0.0.1:8080` by default. to change, and edit configurations from `./config/config.go`

Run the project with the below command from the app directory
```go run server.go```

Build the project from the app directory
```go build -o ./bin/```

To test the API use Postman collection or the below curl command
```
curl --location 'http://127.0.0.1:8080/v1/risks' \
--header 'Content-Type: application/json' \
--data '{
    "title": "test risk",
    "description": "risk description",
    "state":   "open" 
}'
```

```
curl --location 'http://localhost:8080/v1/risks'
```

```
curl --location --globoff 'http://localhost:8080/v1/risks/{Id}'
```