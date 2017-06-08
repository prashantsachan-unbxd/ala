## interface ProbeClient
ProbeClient is an interface for a client which can send a ProbeRequest for a specific protocol. 
It has following methods : 

* _isEmpty()bool_ : checks whether the instance is empty or not (equivalent to null check)
* _New(config map[string]interface{}, service topo.Service) (ProbeClient,error)_ returns a new Instance of this type, initialized with the supplied config 
* _Execute()(resp.ProbeResponse, error)_ : runs the probe request as per config & return a ProbeResponse implementation

### Comments
* For each implementation of ProbeClient, there should be a ProbeResponse implementation to capture the response & convert it to JSON
* Ideally, creation of ProbeClient should be lightweight & should be of little overhead. All the processing, I/O should be done in execute phase
* ProbeClient is passed `probeData` (declared in `ProbeConfig`), which should be used to instantiate & configure the client

### Implementations

There are following implementations of ProbeClient 
#### HttpClient
HttpClient is an HTTP implementation of ProbeClient which communicates with a service over HTTP and returns the response as HttpResponse object (implementation of ProbeResponse). Apart from `host` & `port` of a service, it supports following configurations (fields of `probeData`):

* `path`[string] : http url to send request for
* `method`[string] : http method to send request with {GET, PUT, HEAD, POST, DELETE}
* `data`[string] : body of the request (in case of put/post requests)
