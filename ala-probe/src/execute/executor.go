//Package execute provides implementation for Metric Computation logic for services
package execute

import (
    "result"
)
//Executor interface provides contract for schedulers running the metric computation
type Executor interface{
    //StartExec signals the Executor to start scheduling the services
    // it returns a input channel of Event which will provice one Event for 
    // each of the metrics computed for each service
    // EventDispatcher should read from this channel
    StartExec() <-chan result.Event
    // StopExec terminates the scheduling of metric computation. 
    // It doesn't guarantee how long will it take to close the channel after calling this method
    StopExec()
}