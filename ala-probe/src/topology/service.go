package topology

import (

    )
type Service struct{
    Id string `json:"id"`
    Host string `json:"host"`
    Port int `json:"port"`
    Class []string `json:"class"`
    Metadata map[string]interface{} `json:"metadata"`
}

