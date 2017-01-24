package result

import(
    "ex"
)

type EventDispatcher interface{
    Dispatch(e ex.Event)
}

type SimpleDispatcher struct{
    Consumers [] EventConsumer
}

func (d SimpleDispatcher) Dispatch(e ex.Event){
    for _,c:= range d.Consumers{
        c.Consume(e)
    }
}
