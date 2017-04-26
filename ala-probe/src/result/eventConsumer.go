package result

import(
    "ex"
    )

type EventConsumer interface{
    Consume(e ex.Event)
}
