# GoBunnyAPI
Rabbit API built using Go.
This API will be used to store and retrieve data about pet rabbits. 
Sample data is provided in the models.go file

To use:

Start server with 
`` go run *.go" ``

Test endpoints:
Endpoint to test server running: 
`` curl localhost:8000 ``

Endpoint to list all test data: 
`` curl localhost:8000/api/v1/test/bunnies `` 
