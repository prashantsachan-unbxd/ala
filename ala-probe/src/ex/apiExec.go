package ex

import (
)

type ApiExec interface{
    StartExec() <-chan Event
    StopExec()
}

