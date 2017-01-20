package api

import(
    "io"
    )
type Api struct{
    Method string
    Url string
    Data io.Reader
}
