## Zookeeper Integration

All the data for this project stored on zookeeper resides at root-path /metricCollect. Zookeeper is used for two purposes: 

* Storing Service Configurations (root-path : `/metricCollect/services`)
	* Root directory contains multiple znodes each for a service with name = serviceId
	* `root-path/serviceId` contains the `ServiceConf` object in JSON format as its data
* Storing ServiceClass to ProbeConfig mapping ( root-path : `/metricCollect/probeConfigs`)
	* Root directory contains multiple znodes each for a serviceClass with name = ServiceClass
	* `root-path/serviceClass` contains one or more znodes one for each probeConfig. node-name = probeConfigId
	* `root-path/serviceClass/probeConfigId` contains the `probeConfig` object in JSON format as its data

 zookeeper Hierarchy looks like the follwing : 
```
/
|__metricCollect
  |__services
  |  |__serviceId1 (data ={"id":"id1","host":"host1","port":8086,"class":["class1"],"metadata":null}) 
  |  |__serviceId2 (data ={"id":"id2","host":"host2","port":8086,"class":["class2"],"metadata":null}) 
  |__probeConfigs
     |__serviceClass1
     |  |__probeConfId1 (data = {.....})
     |  |__probeConfId2 (data = {......})
     |__serviceClass2
```

#### Examples
running commands on zookeeper cli 
* Get Ids of all services
```
$ ls /metricCollect/services
[solr-4-164, solr-4-222]
```
* Get serviceConf for serviceId `solr-4-164`
```
$ get  /metricCollect/services/solr-4-164
{"id":"solr-4-164","host":"http://54.210.4.164","port":8086,"class":["solr"],"metadata":null}
Zxid = 0x1a97
ctime = Mon May 29 18:28:15 IST 2017
mZxid = 0x1a97
mtime = Mon May 29 18:28:15 IST 2017
.......
.......
```
* Get List of all the serviceClasses
```
$ ls /metricCollect/probeConfigs
[solr,feed]
```
* Get Ids of all the probeConfigs for serviceClass `solr`
```
$ ls /metricCollect/probeConfigs/solr
[solr-tomcat-status, solr-express-qAll]
```
* Get ProbeConfig with id=solr-tomcat-status & serviceClass `solr`
```
$ get /metricCollect/probeConfigs/solr/solr-tomcat-status
{"id":"solr-tomcat-status","probeType":"HTTP","probeData":{"method":"HEAD","path":"/tomcat.gif"},"metrics":[{"defaultMetricValue":0,"domain":"platform.monitoring","metricName":"HTTPstatus200","subdomain":"metricCollect"}]}
Zxid = 0x1a97
ctime = Mon May 29 18:28:15 IST 2017
mZxid = 0x1a97
mtime = Mon May 29 18:28:15 IST 2017
.....
.....
```