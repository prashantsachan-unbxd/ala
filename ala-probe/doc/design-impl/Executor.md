## interface Executor
executor interface takes SerivceDao & executes the probe process for all the services in it as per implementation. There is only one implemenation of it currently `IntervalExecutor`

### Comments

* Executor wakes up every `interval` time interval (defined in application config) & does a batch execution
* While running a batch, it first fetches all the services & collects unique serviceClasses
* Then it fetches all the ProbeConfigs (to avoid fetching probeConfigs for a serviceClass more than once)
* Once ProbeConfigs have been fetched, It runs a separate goRoutine for each service, each probeconf. This goroutine 
	* Instantiate the client of appropriate `probeType`
	* Calls execute on the ProbeClient & parse resonse to corresponding ProbeResponse implementation
	* For each metric to collect, call RuleEngine with this ProbeResponse
	* Create one event for each service, each metric & push this to an output channel
* Event Dispatcher then takes this event from that channel and through its eventConsumer, pushes it to kafka
* If executor encounters any error after fetching the ProbeConfig such as : 
	* error in instantiating the `ProbeClient`
	* network/API error in sending `ProbeRequest` to the service
	* error in parsing its response to `ProbeReponse`
	* error in sending `ProbeResponse` to RuleEngine to compute metric values
	* no matching rule with given `domain`, `subdomain` & `metricName` as defined in ProbeConfig
	* Matching & running rule encounters a runtime error
Then, defulat value of the metric(s) is sent
* Therefore, there is a guaranee that as far as zookeeper is reachable (i.e. probAgent is able to retrieve serviceConfs & ProbeConfigs for them), then for each service, each metric an event will be generated at each batch-execution
