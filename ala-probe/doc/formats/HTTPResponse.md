## HttpResponse
HttpResponse is an object/model to store the response received from an HTTP server. 

### Syntax
There are following fields in HttpResponse
```
{
	"status"[int]:<http status code>,
	"headers"[map string-> list of string]: <http response headers>,
	"body"[map string-> object / string] : <response body either parsed as JSON or as String if failed>,
	"version"[string] : "http protocol version"
}
```

### Comment
* Any Metriccolletion rule which expects HTTP response to be passed to it, should assume the above structure of the data received
* Body of the actual response received from an HTTP server could / couldn't be JSON. If it is a JSON, then it would be parsed to a map of string-> object, otherwise, `body` contains response as plain string. Hence, type of field `body` would be either *string* OR *map string->object*

#### Example1
HTTPResponse object's JSON representation (where body was a JSON string)

```
{
	"status":200,
	"headers":["Content-Type":["application/json"], "Content-Length": [93]],
	"body" : {
		"success":"true",
		"data":{
			"id":"some-id",
			"aliases":["alias1", "alias2"]
		}
	},
	"version" : "HTTP/1.0"

}
```
#### Example2
HTTPResponse object's JSON representation (where body was not a valid JSON string)

```
{
	"status":200,
	"headers":["Content-Type":["text/html"], "Content-Length": [93]],
	"body" : "<html> <body>some non-json string</body></html>",
	"version" : "HTTP/1.0"

}
```