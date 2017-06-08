##ProbConfig APIs
This section lists out all the CRUD APIs for ProbeConfigs
### Get All ServiceClasses
**Path**
```
GET http://metric-collection.unbxdapi.com/probeconfig/class
```
**Output**
```json
[
  "solr"
]
```
### 
Get probeConfig for class `solr` 
**Path**
```
GET http://metric-collection.unbxdapi.com/probeconfig/solr
```
**Output**
```
[
  {
    "id": "solr-tomcat-status",
    "probeType": "HTTP",
    "probeData": {
      "method": "HEAD",
      "path": "/tomcat.gif"
    },
    "metrics": [
      {
        "defaultMetricValue": 0,
        "domain": "platform.monitoring",
        "metricName": "HTTPstatus200",
        "subdomain": "metricCollect"
      }
    ]
  }
]
```
### Add ProbeConfig to class `solr`
**Path**
```
PUT http://metric-collection.unbxdapi.com/probeconfig/solr
```
**Headers**
```
Content-Type : application/json
```
**Body**
```json
{
	"id":"solr-tomcat-status",
	"probeType" : "HTTP",
	"probeData" : {
		"path": "/tomcat.gif",
		"method": "HEAD"
	},
	"metrics":[
		{
			"domain":"platform.monitoring",
			"subdomain":  "metricCollect",
			"metricName": "HTTPstatus200",
			"defaultMetricValue":0
		}
	]
}
```
**Output**
```
successfully added new service
```
### Delete a ProbeConfig named `solr-tomcat-status` from serviceClass `solr`
**Path**
```
DELETE http://metric-collection.unbxdapi.com/probeconfig/solr/solr-tomcat-status
```
**Output**
```
deleted the probeConfig with Id: solr-tomcat-status
```