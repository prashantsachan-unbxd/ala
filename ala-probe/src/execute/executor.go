package execute

import (
    "result"
)

type Executor interface{
    StartExec() <-chan result.Event
    StopExec()
}