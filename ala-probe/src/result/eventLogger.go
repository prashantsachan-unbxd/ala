package result

import(
    "fmt"
    )

type EventLogger struct{

}
func (this *EventLogger)Init(){

}
func (this *EventLogger) Consume(e Event){
    fmt.Println("logging event: ", e)    
}
