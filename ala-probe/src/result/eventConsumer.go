package result

import (
    )
type EventConsumer interface{
    Init()
    Consume(e Event)
}