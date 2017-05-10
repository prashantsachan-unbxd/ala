package result

import(
    log "github.com/Sirupsen/logrus"
    )

type EventLogger struct{

}
func (this *EventLogger)Init(){

}
func (this *EventLogger) Consume(e Event){
    log.WithFields(log.Fields{"module":"EventLogger","event":e}).Info("event received") 
}
