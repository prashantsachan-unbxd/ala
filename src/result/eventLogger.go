package result

import(
    "fmt"
    "ex"
    )

type EventLogger struct{

}

func (l *EventLogger) Consume(e ex.Event){
    fmt.Println("logging event: ", e)    
}
