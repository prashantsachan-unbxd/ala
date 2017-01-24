package ex

import(
    "api"
    "time"
    )

type Event struct{
    Api api.Api
    Timestamp time.Time
    Status api.ApiStatus
}
