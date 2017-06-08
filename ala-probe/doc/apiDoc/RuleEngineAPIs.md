## Rule Engine APIs
Following section lists out the Rule Engine API's which a user may need to execute in order to add metric collection logic to RuleEngine

### Add a Metric collection logic 
Add a Rule with metricName `HTTPstatus200` from an HttpResponse. it assumes that the object passed to it (the HttpResponse) will have a key `response` with value as the `HttpResponse` object's structure (HttpResponse has 4 keys `status`[int], `body`[map], `version`[string] & `headers`[map of string-> list of strings]). 
**Path**
```
PUT ec2-54-250-167-237.ap-northeast-1.compute.amazonaws.com:8080/rule-engine/rules/
```
**Headers**
```
Content-Type : application/json
```
**Body**
```
{
    "domain": "platform.monitoring",  
	"subdomain":"metricCollect",
	"metricName":"HTTPstatus200",
    "consequent": "local args = ...; if (args.response.status >= 200 and args.response.status < 300) then return 1 else return 0 end", 
    "metadata": {
        "class": "platform.monitoring", 
        "tags": ["platform","monitoring", "http", "response", "status", "availability"], 
        "description": "accepts a http response in json format : {body:\"resp body\", headers:[header1:\"value1\"], status:200, version:\"HTTP/1.1 200 OK\"}, & checks whether it has status code 200 (1) or not(0)", 
        "name": "httpStatus200"
    }
}
```
**Response**
```
{
  "id": "AVyHN-HycGGQyzVBRIer"
}
```
### Get a Rule with id `AVyHN-HycGGQyzVBRIer`
**Path**
```
GET http://ec2-54-173-96-124.compute-1.amazonaws.com:8081/rule-engine/rules/AVyHN-HycGGQyzVBRIer
```
**Response**
```
{
  "metadata": {
    "name": "httpStatusCheck",
    "description": "accepts a http response in json format : {body:\"resp body\", headers:[header1:\"value1\"], status:200, version:\"HTTP/1.1 200 OK\"}, & checks whether it has status code 200 or not",
    "class": "platform.monitoring",
    "tags": [
      "platform",
      "monitoring",
      "http",
      "response",
      "status",
      "availability"
    ]
  },
  "domain": "platform.monitoring",  
  "subdomain":"metricCollect",
  "metricName":"HTTPstatus200",
  "consequent": "local args = ...; if (args.response.status >= 200 and args.response.status < 300) then return 1 else return 0 end",
}
```
### Get All the Rules / metric Collection logics 
GET all rules with class = `platform.monitoring`
**Path**
```
POST http://ec2-54-173-96-124.compute-1.amazonaws.com:8081/rule-engine/rules-by-metadata/
```
**Headers**
```
Content-Type : application/json
```
**Body**
```
{"class": "platform.monitoring"}
```
**Response**
```
{
  "AVvsYU-DGwmeuud0OBd-": {
    "metadata": {
      "name": "httpStatusCheck",
      "description": "accepts a http response in json format : {body:\"resp body\", headers:[header1:\"value1\"], status:200, version:\"HTTP/1.1 200 OK\"}, & checks whether it has status code 200 (1) or not(0)",
      "class": "platform.monitoring",
      "tags": [
        "platform",
        "monitoring",
        "http",
        "response",
        "status",
        "availability"
      ]
    },
    "metricName": "HTTPstatus200",
    "domain": "platform.monitoring",
    "subdomain": "metricCollect",
    "consequent": "local args = ...; if (args.response.status >= 200 and args.response.status < 300) then return 1 else return 0 end"
  },
  ......
  ......
}
```
### Delete a Rule with id `xyz`
**Path**
```
DELETE ec2-54-250-167-237.ap-northeast-1.compute.amazonaws.com:8080/rule-engine/rules/xyz
```
**Response**
```

```

### Compute Metric for HTTP Response
This API is internally executed by the ProbeAgent. However, can be useful for debugging purpose. Follwing is the request that can be used to evaluate the rule with metricName `HTTPstatus200`. Notice that the body contains a key `response` with value as the `HttpResponse` object (serialized).
**Path**
```
POST http://ec2-54-173-96-124.compute-1.amazonaws.com:8081/rule-engine/rule-results/
```
**Headers**
```
Content-Type : application/json
```
**Body**
```
{
  "domain": "platform.monitoring",  
  "subdomain":"metricCollect",
  "metricName":"HTTPstatus200",
  "response": {
    "version":"HTTP\/1.1 200 OK",
    "status":300,
    "headers":{"Content-Type":"text\/html","connection":"close"},
    "body":"Response Body"
  }
}
```
**Response**
```
{
  "AVvsYU-DGwmeuud0OBd-": {
    "metadata": {
      "name": "httpStatusCheck",
      "description": "accepts a http response in json format : {body:\"resp body\", headers:[header1:\"value1\"], status:200, version:\"HTTP/1.1 200 OK\"}, & checks whether it has status code 200 (1) or not(0)",
      "class": "platform.monitoring",
      "tags": [
        "platform",
        "monitoring",
        "http",
        "response",
        "status",
        "availability"
      ]
    },
    "value": 0
  }
}
```
Here, the field `value` contains the numerical value of the metric `HTTPstatus200`




