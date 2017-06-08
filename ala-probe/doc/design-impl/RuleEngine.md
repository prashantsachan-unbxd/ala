## RuleEngine Integration

Rule Engine is used to compute metrics from a `ProbeResponse`. Each rule stores a lua script which takes a Map (Table) as CLI arg & outputs a numerical value (float64). 

* The Cli arg is nothing but the ProbeResponse as a map & hence, the lua script is tighly coupled with the specific ProbeResponse implementation for which it is written.
* There should be only one rule with the given name. In case there are many rules with the same name, first rule-value will be returned as the metric value
* For the sake of exclusivity from other micro-services (so that the rules don't get mixed due to accidentally same attribute values) there are two fields `domain` & `subdomain` passed along with the `metricName`. Value of `domain` should remain same across the service, wheras value of `subdomain` may change across modules. As of now, these two values will remain fixed to `platform.monitoring` & `metricCollect` respectively.

### Example
Rule to extract availability status from an HTTPResponse (http implementation of ProbeResponse ) would look like: 
```
{
	"metadata": {
      "name": "httpStatusCheck",
      "description": "accepts a http response in json format : {body:\"resp body\", headers:[header1:\"value1\"], status:200, version:\"HTTP/1.1 200 OK\"}, & checks whether it has status code 200 (1) or not(0)",
      "class": "platform.monitoring",
      "tags": [ "availability"]
    },
	"metricName": "HTTPstatus200",
    "domain": "platform.monitoring",
    "subdomain": "metricCollect",
    "consequent": "local args = ...; if (args.response.status >= 200 and args.response.status < 300) then return 1 else return 0 end"
}
```

and the rule resolution Request (the body) to compute metric `HTTPstatus200` from an HttpResponse would look like: 

```
{
   "domain":"platform.monitoring",
   "subdomain":"metricCollect",
   "metricName":"HTTPstatus200",
   "response":{
      "body":{
		"searchMetaData":{
    		"status":0,
    		"queryTime":59,
    		"queryParams":{
    		"q":"*"}
    	},
		"response":{"numberOfProducts":4692,"start":0,"products":[]
		}
	},
      "headers":{"Accept-Ranges":["bytes"],"Content-Length":["2066"],"Content-Type":["image/gif"]},
      "status":200,
      "version":"HTTP/1.1"
   }
}	
```