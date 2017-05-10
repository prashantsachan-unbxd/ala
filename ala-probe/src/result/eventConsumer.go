package result

import (
    )
//EventConsumer interface is contract for consumers which process an Event
type EventConsumer interface{
    //Init initializes a consumer
    Init()
    //Consume  : processes an Event
    Consume(e Event)
}