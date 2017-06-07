## interface ProbeResponse 
Probe Response is an interface to capture response of a probeRequest. A `ProbeClient` is expected to return a ProbeResponse Implementation on calling `execute` method. ProbeResponse interface has following methods : 

* _GetType()string_ : returns the ProbeType.Ideally this value should be the same as the ProbeClient's ProbeType value]
* _AsMap()map[string]interface{}_ : converts the response into a Map which marshalled & sent (to RuleEngine) for computing metrics


### Comments
* Normally, each type of ProbeClient will have its specific ProbeResponse.
* it is the individual implementation's responsibility to handle multiple calls of this method. Users should be able to call this method multiple times getting the same value each time


### Implementations: 
Following implementations of ProbeResponse are available
* [HttpResponse|../formats/HTTPResponse.md]