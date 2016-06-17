# Elasticsearch service

This is the elasticsearch service with fqdn go.micro.srv.elasticsearch

## Getting Started

### Prerequisites

Install Consul
[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

Run Consul
```
$ consul agent -dev -advertise=127.0.0.1
```

### Run Service

```
$ go run main.go
```

Create
{"record" : {"index":"flags", "type": "flag", "id": "flag-2", "data":  "{\"fieldX\": 33, \"bb\": true}"}}


Search
{
    "index":"flags", 
    "type": "flag",
    "query": "*",
    "limit": 20,
    "offset": 0
}

Query
{
    "index":"flags", 
    "type": "flag",
    "query": "{\"query\":{\"match\" : {\"fieldY\" : \"yy\"}}}"
}
