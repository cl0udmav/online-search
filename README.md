# online-search
Programatically search online for stuff.

## Prereqs
* go
* [SerpApi](https://serpapi.com) API key

## Usage
1. Create an `.env` file in your root directory
2. Update the `.env` file with parameters required by the map in [search](search.go)
3. `go run main.go`

## Results
You should get a `results.json` output in your root directory with your query results.

## Test coverage
* Empty search arguments