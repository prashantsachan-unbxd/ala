package result

import (
    "execute"
    )
type EventConsumer interface{
    Init()
    Consume(e execute.Event)
}