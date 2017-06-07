## interface EventDispatcher
EvenetDispatcher interface has similar functionality as that of a Broker in a queue system (e.g. Kafka). It has a list of `EventConsumer`s to which it forwards the incoming Events. There is only one implementations of it which is `SimpleDispatcher` which simply forwards the events to all the consumers it has without any processing/waiting etc. It has following methods 

* _StartDispatch( c <-chan Event)_ : starts dispatching events, reading from an input channel `c`
* _StopDispatch()_ : Stops dispatching, cleanup should be done here


## interface EventConsumer
EventConsumer is an interface representing an entity which receives Events as they are generated & does some processing. Follwing are the methods of this interface
    
* _Init()_ : initializes a consumer. 
* _Consume(e Event)_ processes an Event. 


### Implemantations
There are two known implementations of EventConsumer

* EventLogger : simply logs the event
* KafkaForwarder : Marshals the event to Json & pushes it to kafka with keys as ServiceId
