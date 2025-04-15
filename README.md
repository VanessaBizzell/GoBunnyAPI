# GoBunnyAPI
Rabbit API built using Go.
This API will be used to store and retrieve data about pet rabbits. 
Sample data is provided in the models.go file

To use:

Start server with:
`` go run . ``

Run tests with:
``go test``

Test endpoints:
Endpoint to test server running: 
`` curl localhost:8000 ``

Endpoint to list all test data: 
`` curl localhost:8000/api/v1/test/bunnies `` 

Endpoint to find a bunny by ID (replace last number with id you want to test):
``curl "http://localhost:8000/api/v1/test/bunny?id=1"``

Endpoint to convert roman numerals at end of URL into an ID integer and then find a bunny by that ID (replace last characters with roman numerals you want to test):
``curl "localhost:8000/api/v1/test/IV"``
