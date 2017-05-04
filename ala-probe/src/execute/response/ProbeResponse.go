package response

type ProbeResponse interface{
    getType()string
    getJson()string
}