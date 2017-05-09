package result

import(
    "execute"
)

type EventDispatcher interface{
    StartDispatch( c <-chan execute.Event)
    StopDispatch()
}

type SimpleDispatcher struct{
    Consumers [] EventConsumer
    done chan struct{}
}

func (this *SimpleDispatcher) StartDispatch(c <-chan execute.Event){
    for _,con:= range this.Consumers{
        con.Init()
    }
    this.done = make(chan struct{})
    go func(){
        for{
        select{
            case e:= <- c:
                for _,con:= range this.Consumers{
                    con.Consume(e)
                }
            case <- this.done:
                return
        }
    }
    }()
}
func (this *SimpleDispatcher)StopDispatch(){
    close(this.done)
}