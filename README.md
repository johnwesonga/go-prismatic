# go-prismatic
go-prismatic is a Go library for accessing the Alpha version of the [Prismatic API] (https://github.com/Prismatic/interest-graph).

_Disclaimer_ :Because the API is in alpha state many things might break and thus affect the library..

*Documentation:* [![GoDoc](https://godoc.org/github.com/johnwesonga/go-prismatic/prismatic?status.svg)](https://godoc.org/github.com/johnwesonga/go-prismatic/prismatic)

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
The Prismatic API uses a token based authentication mechanism. The API token can be sent with the header
or as a querystring. The library uses the header method.
To request a token go to ........


For complete usage of go-prismatic, see the full [package docs] (http://godoc.org/github.com/johnwesonga/go-prismatic/prismatic).


### License ###
This library is distributed under the BSD-style license
