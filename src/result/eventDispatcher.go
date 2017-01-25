package result

import(
    "ex"
)

type EventDispatcher interface{
    StartDispatch( c <-chan ex.Event)
    StopDispatch()
}

type SimpleDispatcher struct{
    Consumers [] EventConsumer
    done chan struct{}
}

func (d *SimpleDispatcher) StartDispatch(c <-chan ex.Event){
    d.done = make(chan struct{})
    go func(){
    for{
        select{
            case e:= <- c:
                for _,con:= range d.Consumers{
                    con.Consume(e)
                }
            case <- (d.done):
                return
        }
    }
    }()
}
func (d *SimpleDispatcher)StopDispatch(){
    close((d.done))
}
