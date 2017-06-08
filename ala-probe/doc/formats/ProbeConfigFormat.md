## ProbeConfig Format
ProbeConfig is the Configuration of a ProbeRequest. This along with service details should be sufficient to create and execute probeRequest & compute metric values from its response
### Syntax
```
{
	"probeType" [string]: <type-of-probeclient-to-use>,
	"probeData" [map]: <attributes-to-create-proberequest>,
	"metrics" [list]: <list-of-metrics-to-be-collected-from-the-proberesponse>
}
```
where details of each `metric` in the above list is of following format
```
{
	"domain":"platform.monitoring",
	"subdomain":  "metricCollect",
	"metricName"[string]: <name-of-metric>,
	"defaultMetricValue"[float64]:<default-value-to-be-returned-in-case-of-any-error>
}
```
### Comments
* Value of `probeType` indicates which type of ProbeRequest to be done. Most of the time it would correspond to a specific protocol. Application internally maintains a hard-coded binding between `probeType` and `ProbeClient` implementation.
* Elements of `probeData` are specific to the ProbeClient implementation bound to the specific `probeType` value. It is the responsibility of that specific `ProbeClient` implementation to consume these values at the time of instantiation. see [[ProbeClient Design|../design-impl/ProbeClientDesign.md]] for more details
* One ProbeConfig means exactly one ProbeRequest to be sent to the service. But any number of metrics can be collected from its response e.g. from the response of search api `q=*` we can collect : 
	* Total number of products
	* No. of facets
	* no. of products above a given threshold
  
* values of `domain` and `subdomain` are fixed and are passed to ruleEngine as it is at the time of rule-resolution as a Segment to match. 
* There should be one (& exactly one) Rule in RuleEngine corresponding to the segment defined by `domain`, `subdomain` & `metricName` which accepts the ProbeResponse returned by the service & returns a numerical value. If there are more than one rules, each returning one value, then any one of these values will be return as the metricValue (no guarantees on which one to pick)

#### Example 1
 HTTP probeConfig which pings /tomcat.gif & computes metric named `HTTPstatus200` from it
```
{
	"probeType" : "HTTP",
	"probeData" : {
		"path": "/tomcat.gif",
		"method": "HEAD",
		"connTimeout": "3",
		"readTimeout": "5"
	},
	"metrics":[
	{
		"domain":"platform.monitoring",
		"subdomain":  "metricCollect",
		"metricName": "HTTPstatus200",
		"defaultMetricValue":0
	}]
}
```
This ProbeConfig assumes that the RuleEngine when resolved with following segment: 
```
{
	"domain":"platform.monitoring",
	"subdomain":  "metricCollect",
	"metricName": "HTTPstatus200"
}
```
will resolve to a lua script which accepts an HttpResponse as cli data (`args=...` in lua) and returns back with 0/1 depending upon the HTTP status code of the response. This rule's `consequent`(lua script) would look like : 
```
"local args = ...; if (args.response.status >= 200 and args.response.status < 300) then return 1 else return 0 end"
```
#### Example 2
HTTP probeConfig which fires a `GET` reqest on path `/unbxd-search/express_com-u1456154309768/search?q=*` and retrieves two metrices from the response would be :

```
{
	"probeType" : "HTTP",
	"probeData" : {
		"path": "/unbxd-search/express_com-u1456154309768/search?q=*",
		"method": "GET"
	},
	"metrics":[
	{
		"domain":"platform.monitoring",
		"subdomain":  "metricCollect",
		"metricName": "solrNumProduct",
		"defaultMetricValue":0
	},
	{
		"domain":"platform.monitoring",
		"subdomain":  "metricCollect",
		"metricName": "solrNumFacets",
		"defaultMetricValue":0
	}]
}
```
again, it is expected that the RuleEngine will have two separate rules (which understand the response JSON structure)to compute `solrNumProduct` and `solrNumFacets` 