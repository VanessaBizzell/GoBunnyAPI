# GoBunnyAPI

GoBunnyAPI is a Rabbit API built using Go, designed to store and retrieve data about pet rabbits. This API provides sample data for testing purposes, which can be found in the `models.go` file.

## Features

- Retrieve data about pet rabbits.
- Convert Roman numerals into ID integers and retrieve rabbit data.

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/) (1.x or higher) installed on your machine.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/VanessaBizzell/GoBunnyAPI.git
   cd GoBunnyAPI

2. Run the Server:
  ```bash
  go run .
```

## API Endpoints
 Health Check

 Verify if the server is running:
```bash
curl localhost:8000
```

 List All Test Data
(Retrieve all sample bunny data):

```bash
curl localhost:8000/api/v1/test/bunnies
```
Find a Bunny by ID
(Replace the last integer with the bunny ID you want to retrieve):

```bash
curl "http://localhost:8000/api/v1/test/bunny?id=2"
```

 Find a Bunny by Roman Numeral ID
(Replace Roman numerals with the Roman numeral representing the bunny ID you want to retreive):

```bash
curl "http://localhost:8000/api/v1/test/IV"
```


   
