
## Goals
* api specific healthcheck
* support for in-depth response analysis
* Stats & Dashboard


![Design_Diagram](Ala_Design.png?raw=true "Design_Diagram")

## Specification
### Terminology
#### Api validator
each Api maps to a pre-determined validator, which is used to validate its response & determine, whether the API is working as expected or not. Possible examples of validators may one/more of : 
* Http status
* specific header values
* `Success` Flag  value in the JSON response
* no. of products returned in the response

#### Event
each time, an API is fired, an Event object is generated to capture its result.

### Modules 

###  Config Manger
Loads & Stores api configuration, & hands it over to Executor
### Executor
Fires http request for each api in configuration & validates the responses with the configured Validator. Generates an Event & sends it to the Event Dispatcher
### Event Dispatcher
Maintains a collection of registered Event Consumers & forwards each event to all of them.

### Event Consumer
Receives Events & performs a specific set of tasks on them such as : 
* maintaining current state of all the APIs
* storing events to a datastore / file for querying in future
* sending notification

### Dialer
Is a DAO kind of layer to make all the external API calls 

### UI

All the User Actions are handled here
