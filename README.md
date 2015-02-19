# go-prismatic
Go library for accessing the [Prismatic API] (https://github.com/Prismatic/interest-graph).

go-prismatic requires Go version 1.1 or greater.

## Usage ##
```go
import "github.com/johnwesonga/go-prismatic/prismatic"
```

Construct a new Prismatic client, then use the various services on the client to 
access the API. For example to search for an interests like "Clojure":

```go
apiToken := "prismatic-api-token"
client := prismatic.NewClient(nil, apiToken)
results, _, err := client.Topics.SearchForInterest("Clojure")
```

### Authentication ###
The Prismatic API
