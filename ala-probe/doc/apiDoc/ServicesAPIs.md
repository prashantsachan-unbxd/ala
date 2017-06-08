## Service APIs
This section lists out the CRUD api's for metricCollection platform (Ala). These are a part of ala-probe project. 


### Add new Service
**Path**
```
PUT http://metric-collection.unbxdapi.com/service
```
**Headers**
```
Content-Type : application/json
```
**Body**
```json
{
	"id":"solr-4-164",
	"host":"http://54.210.4.164",
	"port":8086,
	"class":["solr"]
}
```
**Output**
```
successfully added new service
```

### Get All Services
**Path**
```
GET http://metric-collection.unbxdapi.com/service
```
**Output**
```
[
  {
    "id": "solr-4-164",
    "host": "http://54.210.4.164",
    "port": 8086,
    "class": [
      "solr"
    ],
    "metadata": null
  }
]
```

### Get Service
**Path**
```
GET http://metric-collection.unbxdapi.com/service/solr-4-164
```
**Output**
```
{
  "id": "solr-4-164",
  "host": "http://54.210.4.164",
  "port": 8086,
  "class": [
    "solr"
  ],
  "metadata": null
}
```

### Delete Service
**Path**
```
DELETE http://metric-collection.unbxdapi.com/service/solr-4-164
```
**Output**
```
deleted the service with Id: solr-4-164
```