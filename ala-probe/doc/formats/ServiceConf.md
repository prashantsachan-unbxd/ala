## ServiceConf Format
ServiceConf contains info common across all the services (id, host, port) along with the classes (logical groups) to which it belongs to. The idea behind grouping of service into `serviceClass`es is to consider a set of services (processes) as a single entity.

### Syntax
```
{
    "id"[string]: <unique-id-of-the-service>,
    "host"[string]:<ip/DNS of the service>,
    "port"[integer]:<port at which the service listens>,
    "class"[list of string]:<serviceClasses-to-which-service-belongs-to>
    "metadata"[map string-> object]: <other meta information related to service>
}
```
### Comments
* `serviceClass` is just a logical group of some ProbeConfigs. Behind the scene, it simply maps to a set of probeConfigs applicable to the service. For all the services/ remote-processes which belong to a particular serviceClass these probeConfigs will be applicable
* A service can belong to multiple serviceClasses. A `service-to-serviceClass` relationship here is similar to a `Class-to-Interface` relationship in Object Oriented Programming domain. A serviceClass tells which probeRequests are to be done for a service. There could be other set of probeRequests also (corresponding to another serviceClass it belongs to) for the same service
* `metadata` can contain any random information related to the service. It's format is expected to be specific to the service, hence it is kept generic. For example, `metadata` for a kafka service may contain zookeeper hosts, zkRoot, userName, Password etc.
* There are three level of hierarchies here: 
	* A service can belong to multiple serviceClasses
	* A serviceClass can map to multiple ProbeConfigs
	* A ProbeConfig will fire single probeRequst, but multiple metrics can be calculated from it
Rationale behind such a structure comes from the following cases:
	* Multiple metrices may be extrated from same ProbeRequest. In such a case, it is efficient to probe the service once & compute metrices from the same response multiple times (rather than firing same request for each metric)
	* There could be multiple ProbeRequests required for a specific type / class of services. Hence, it makes sense to group these into serviceClasses. This way, whenever a new instance of such a service is deployed, it just registers itself with the servicClass & all the probeConfigs in it are applied to it
	* It is possible for services to share these groups among them (two different JAVA servies say SOLR & Elasticsearch expose their JVM stats over HTTP requiring the same set of ProbeRequests to be done for them). Hence a service should be able to subscribe any no. of these logical groups (i.e. serviceClasses)

#### Example
example of a Solr service configuration is as below
```
{
    "id":"67355A76-5669-4F6C-9F4A-A03EE41BB292",
    "host":"http://54.210.4.164",
    "port":8086,
    "class":[ "SOLRHTTPSTATUSCHECK","EXPRESS_NUMPRODUCTS"]
}
```