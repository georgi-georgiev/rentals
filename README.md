## Run the application
- go mod tidy
- go mod vendor
- docker compose -f "docker-compose.yml" up -d --build

## Execute the tests
- go test tests/rentals_test.go

## Possible Improvements
- add more tests
- db indexing for faster search
- caching
- return information about the page size and total results in the response
- add relation urls in the response for previous nad next page accordingly