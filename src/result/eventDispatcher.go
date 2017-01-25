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
    Done *chan struct{}
}

func (d SimpleDispatcher) StartDispatch(c <-chan ex.Event){
    c1 := make(chan struct{})
    * (d.Done) = c1
    go func(){
    for{
        select{
            case e:= <- c:
                for _,c:= range d.Consumers{
                    c.Consume(e)
                }
            case <- *(d.Done):
                return
        }
    }
    }()
}
func (d SimpleDispatcher)StopDispatch(){
    close(*(d.Done))
}
