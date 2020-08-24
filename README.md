# u2json
![Go](https://github.com/saka1/u2json/workflows/Go/badge.svg)

A Utility to convert URL to JSON containing each parts.

```shell
$ u2json 'https://example.com:80/foo/bar'
{"host":"example.com","path":"/foo/bar","port":80,"scheme":"https"}
```

## Usage

### Full example

Pretty print with [jq](https://stedolan.github.io/jq/).

```shell
$ u2json 'https://usr:pass@example.com:8080/foo/bar?q=v#fr' | jq '.'
{
  "fragment": "fr",
  "host": "example.com",
  "password": "pass",
  "path": "/foo/bar",
  "port": 8080,
  "query": {
    "q": "v"
  },
  "rawQuery": "q=v",
  "scheme": "https",
  "user": "usr"
}
```

### multiple values in query parameters

`--query-array` allows multiple values with a single field(see the difference).

```shell
$ u2json 'http://example.com?q=v1&q=v2' | jq '.'
{
  "host": "example.com",
  "query": {
    "q": "v2"
  },
  "rawQuery": "q=v1&q=v2",
  "scheme": "http"
}
```

```shell
$ u2json --query-array 'http://example.com?q=v1&q=v2' | jq '.'
{
  "host": "example.com",
  "query": {
    "q": [
      "v1",
      "v2"
    ]
  },
  "rawQuery": "q=v1&q=v2",
  "scheme": "http"
}
```
