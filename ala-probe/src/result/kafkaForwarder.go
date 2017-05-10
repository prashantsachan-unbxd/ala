package result

import(
    sarama "gopkg.in/Shopify/sarama.v1"
    "encoding/json"
    log "github.com/Sirupsen/logrus"
    "time"
    )
// var brokerlist = []string{"localhost:9092"}

var TOPIC_NAME = "monitoring"
var producer sarama.AsyncProducer
type KafkaForwarder struct{
    BrokerList []string

}
func (this *KafkaForwarder) Init(){
    producer = newAsyncProducer(this.BrokerList)
}
func (this *KafkaForwarder) Consume(e Event){
    log.WithFields(log.Fields{"module":"KafkaForwarder","event":e}).Debug(
        "sending to kafka") 
    if(producer == nil ){
        log.WithFields(log.Fields{"module":"KafkaForwarder",
            "brokers":this.BrokerList}).Fatal(
            "nil Kafka Producer, unable to send event")
    }
    msg, err := json.Marshal(e)
    if err !=  nil{
        log.WithFields(log.Fields{"module":"KafkaForwarder","event":e,
            "error":err}).Error("unable to unmarshal event")
        return
    }
    producer.Input() <- &sarama.ProducerMessage{
            Topic: TOPIC_NAME,
            Key:   sarama.StringEncoder(e.Srvc.Id),
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
        log.WithFields(log.Fields{"module":"KafkaForwarder","brokers":brokerList,
            "error":err}).Fatal("Failed to start Sarama producer") 
    }
    log.WithFields(log.Fields{"module":"KafkaForwarder","brokers":brokerList,
        "config":config}).Info("Initialized Kafka producer") 
    // We will just log to STDOUT if we're not able to produce messages.
    // Note: messages will only be returned here after all retry attempts are exhausted.
    go func() {
        for err := range producer.Errors() {
            log.WithFields(log.Fields{"module":"KafkaForwarder","error":err}).Error(
                "Failed to write access log entry") 
        }
    }()

    return producer
}