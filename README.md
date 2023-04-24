# synonym-service-go
The complete system should act as a thesaurus - it should enable users to store and fetch sets of synonyms. Please also consider testability, readability, algorithmic complexity, and maintainability!
Your implementation should fulfil the following requirements:

* Endpoint to add new sets of synonyms. For example, adding a pair of synonyms such as “begin” and “start”.
* Endpoint to search for synonyms. In the above example, searching for either “begin” or “start” should return the respective synonym (symmetrical relationship).
* A word may have multiple synonyms, and all should be returned at a user request.
* The solution needs to support concurrent requests in a thread-safe way.
* Make the solution with simple data structures in memory - no persistence/database needed.
* Transitive rule. For example, if “A” is added as a synonym for “B”, and “B” is added as a synonym for “C”, then searching the word “C” should return both “B” and “A”.

## Start
```
cd app
go run .
```

## Endpoints
Query all synonyms
```
curl localhost:8080/synonyms
```
Add a synonym set
```
curl http://localhost:8080/synonyms \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"words": ["pretty", "attractive", "lovely"]}'
```

## Testing
```
cd synonym
go test -v
```


## Goal
The complete system should act as a thesaurus - it should enable users to store and fetch sets of synonyms. Please also consider testability, readability, algorithmic complexity, and maintainability!
Your implementation should fulfil the following requirements:

* Endpoint to add new sets of synonyms. For example, adding a pair of synonyms such as “begin” and “start”.
* Endpoint to search for synonyms. In the above example, searching for either “begin” or “start” should return the respective synonym (symmetrical relationship).
* A word may have multiple synonyms, and all should be returned at a user request.
* The solution needs to support concurrent requests in a thread-safe way.
* Make the solution with simple data structures in memory - no persistence/database needed.

**Bonus***: Transitive rule. For example, if “A” is added as a synonym for “B”, and “B” is added as a synonym for “C”, then searching the word “C” should return both “B” and “A”.