## Event
event is model class to capture a metric's value. It represents a metric's value for a specific service at a given point of time. At each execution cycle, an event is generated for each service and each metric defined for it. This is what is pushed to kafka at each execution cycle / batch

### Syntax
following is the JSON representation syntax for an event
```
{
	"service"[object] : service for which the execution was done
	"timestamp"[int64] : Timestamp at which the execution started (milliseconds since epoch)
	"metricName"[string] : MetricName which is computed
	"value"[float64] : value of the metric
}
```

### Comments
* `service` has an object of same structure as given in [ServiceFormat|ServiceConf.md]
* `timestamp` is a 64 bit int. 

#### Example
```
{
	"service":{
		"id":"solr-4-164",
		"host":"http://54.210.4.164",
		"port":8086,
		"class":["solr"],
		"metadata":null
	},
	"timestamp":1496903769213,
	"metricName":"HTTPstatus200",
	"value":1
}
```