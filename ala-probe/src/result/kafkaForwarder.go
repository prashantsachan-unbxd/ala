package result

import(
    sarama "gopkg.in/Shopify/sarama.v1"
    "ex"
    "encoding/json"
    "fmt"
    "time"
    )
var brokerlist = []string{"localhost:9092"}
type KafkaForwarder struct{
    producer sarama.AsyncProducer

}

func (l *KafkaForwarder) Consume(e ex.Event){
    fmt.Println("forwarding event to kafka :", e)
    if(l.producer == nil ){
        l.producer = newAsyncProducer(brokerlist);    
    }
    msg, err := json.Marshal(e)
    if err !=  nil{
        fmt.Println(" erro in marsalling event to Json")
        return
    }
    l.producer.Input() <- &sarama.ProducerMessage{
            Topic: "monitoring",
            Key:   sarama.StringEncoder(e.Api.Url),
            Value: sarama.StringEncoder(msg),
        }

}

func newAsyncProducer(brokerList []string) sarama.AsyncProducer {

    // For the access log, we are looking for AP semantics, with high throughput.
    // By creating batches of compressed messages, we reduce network I/O at a cost of more latency.
    config := sarama.NewConfig()
 
    config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
    config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
    config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

    producer, err := sarama.NewAsyncProducer(brokerList, config)
    if err != nil {
        fmt.Println("Failed to start Sarama producer:", err)
    }

    // We will just log to STDOUT if we're not able to produce messages.
    // Note: messages will only be returned here after all retry attempts are exhausted.
    go func() {
        for err := range producer.Errors() {
            fmt.Println("Failed to write access log entry:", err)
        }
    }()

    return producer
}