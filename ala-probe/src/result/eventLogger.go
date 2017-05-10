package result

import(
    log "github.com/Sirupsen/logrus"
    )
//EventLogger is an EventConsumer which simply logs each of the event received
type EventLogger struct{

}
func (this *EventLogger)Init(){

}
func (this *EventLogger) Consume(e Event){
    log.WithFields(log.Fields{"module":"EventLogger","event":e}).Info("event received") 
}
