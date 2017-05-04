package execute

import (
    "time"
    topo "topology"
    )


type Event struct{
    Srvc topo.Service
    Timestamp time.Time
    Metric float64
}
