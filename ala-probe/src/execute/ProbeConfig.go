package execute

import (
    //"encoding/json"
    )

type ProbeConfig struct{
    ProbeType string `json:"probeType"`
    ProbeData map[string]interface{} `json:"probeData"`
    Metrics []map[string]string `json:"metrics"`

}