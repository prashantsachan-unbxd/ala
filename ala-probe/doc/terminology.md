# Terminology

## Service
a service a process on a remote machine exposed over network. Each service has a/an
  * `id` : to uniquely identify the service
  * `host` : public DNS/ip at which it is running
  * `port` : port on which it listens
  * `serviceClass`(s) : tells what kind of a service it is ! (SOLR, AGGREGATOR, RECOMMENDER, CONSOLE, FEED-MONGO etc.) these values determine which probe requests will be made to it ( & hence, what all metrics will be collected from it). A service could belong to multiple `serviceClass`es.


## Probe
probing is an action of connecting to a service over network & fetch a response.
### ProbeRequest
Represents a request of a specific type / protocol. 
### ProbeResponse
Represents a response received for a specific type of `ProbeRequest`

an `HTTPClient` will send and `HTTPRequest` to a service & parse the response into an `HttpResponse` object
### ProbeClient
ProbeClient is an entity which sends a `ProbeRequest` to a service and collects the response. There could be multiple types of ProbeClient, one for each particular protocol e.g. HTTPClient, to probe a service over HTTP protocol. To connect to a request in some other protocol (Mongo, solr-binary etc.) a specific HTTP client should be implemented first, which understands that protocol




### ProbeConfig
a ProbeConfig contains following information: 
* `probeType` : used to select correct type of ProbeClient
* `probeData` : used to instantiate the ProbeClient (specific to the `probeType`)
* `metrics`: list of metrics to collect
```
{
	"probeType" : "HTTP",
	"probeData" : {
		"path": "/tomcat.gif",
		"method": "GET",
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

### ProbeAgent
Main Runner process of Ala. It 
* Collects all the `services` (from file as of now)
* For each of the `serviceClass`, finds all the `ProbeConfig`s 
* Send `ProbeRequest`s according to the `ProbeConfig`s & collect `ProbeResponse`s
* For each of the ProbeResponse, collect all the `metric`es listed out in the `ProbeConfig`
* Create an `Event` of from each of the metrices & send that to kafka

# Metric
is a key(string)-value(float64) pair stored as a time series for each of the service
This is what we see on grafana
# Event 
is an object emitted for each metric collected for each of the service

