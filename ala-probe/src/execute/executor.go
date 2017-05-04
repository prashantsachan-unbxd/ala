package execute

import (
)

type Executor interface{
    StartExec() <-chan Event
    StopExec()
}