package response

type ProbeResponse interface{
    GetType()string
    AsMap()map[string]interface{}
}