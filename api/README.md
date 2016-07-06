# Elastic API [![GoDoc](https://godoc.org/github.com/Rakanixu/elastic/api?status.svg)](https://godoc.org/github.com/Rakanixu/elastic/api)

This is the Elastic API for consuming elascticsearch service through HTTP.

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


## Usage 

### Create
```
http[domain:micro API port]/elastic/create
{
    "index":"flags", 
    "type": "flag", 
    "id": "flag-id", 
    "data":  {
        "att1": "value1", 
        "bool": false, 
        "innerobj": {
            "attr1": 46,
            "bool": true
        }
    }
}

{}
```

### Read
```
http[domain:micro API port]/elastic/read
{
    "index":"flags", 
    "type": "flag", 
    "id": "flag-id"
}

{
  "att1": "value1",
  "bool": false,
  "innerobj": {
    "attr1": 46,
    "bool": true
  }
}
```

### Update
```
http[domain:micro API port]/elastic/update
{
    "index":"flags", 
    "type": "flag", 
    "id": "flag-id",
    "data":  {
        "update": true,
        "att1": "value1", 
        "bool": false, 
        "innerobj": {
            "attr1": 46,
            "bool": true
        }
    }
}

{
    "update": true,
    "att1": "value1", 
    "bool": false, 
    "innerobj": {
        "attr1": 46,
        "bool": true
    }
}
```

### Delete
```
http[domain:micro API port]/elastic/delete
{
    "index":"flags", 
    "type": "flag", 
    "id": "flag-id"
}

{}
```

### Search
```
http[domain:micro API port]/elastic/search
{
    "index":"flags", 
    "type": "flag", 
    "query":"47", 
    "limit": 20, 
    "offset": 0
}

{
  "took": 1,
  "timed_out": false,
  "_shards": {
    "total": 5,
    "successful": 5,
    "failed": 0
  },
  "hits": {
    "total": 1,
    "max_score": 0.19178301,
    "hits": [
      {
        "_index": "flags",
        "_type": "flag",
        "_id": "flag-id3",
        "_score": 0.19178301,
        "_source": {
          "att1": "value2",
          "bool": false,
          "innerobj": {
            "attr1": 47,
            "bool": true
          }
        }
      }
    ]
  }
}
```

### Query
```
http[domain:micro API port]/elastic/query
{
    "index":"flags", 
    "type": "flag", 
    "query": {
        "query": { 
            "match" : {
                "att1" : "value2"
            }
        }
    }
}

{
  "took": 1,
  "timed_out": false,
  "_shards": {
    "total": 5,
    "successful": 5,
    "failed": 0
  },
  "hits": {
    "total": 2,
    "max_score": 0.30685282,
    "hits": [
      {
        "_index": "flags",
        "_type": "flag",
        "_id": "flag-id2",
        "_score": 0.30685282,
        "_source": {
          "att1": "value2",
          "bool": false,
          "innerobj": {
            "attr1": 48,
            "bool": true
          },
          "update": true
        }
      },
      {
        "_index": "flags",
        "_type": "flag",
        "_id": "flag-id3",
        "_score": 0.30685282,
        "_source": {
          "att1": "value2",
          "bool": false,
          "innerobj": {
            "attr1": 47,
            "bool": true
          }
        }
      }
    ]
  }
}
```
