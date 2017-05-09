package result

import(
    "fmt"
    "execute"
    )

type EventLogger struct{

}
func (this *EventLogger)Init(){

}
func (this *EventLogger) Consume(e execute.Event){
    fmt.Println("logging event: ", e)    
}
