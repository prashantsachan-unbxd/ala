package main

import(
	"github.com/spf13/viper"
)

const conf_batch_interval = "execute.batch_interval"
const conf_log_level = "log_level"

//execute section
const secEx = "execute."
const conf_service_path =  secEx+"path_service_conf"
// Rule Engine section
const secRE = "ruleEngine."
const conf_re_host = secRE + "host"
const conf_re_port = secRE + "port"
const conf_re_user = secRE + "user"
const conf_re_pass = secRE + "pass"

// Kafka section
const secKafka = "kafka."
const conf_kafka_brokers = secKafka+ "brokers_hostport"
const conf_kafka_topic = secKafka+ "topic"

var mustConfs = []string{conf_batch_interval, 
	conf_service_path,
	conf_re_host, conf_re_port,
	conf_kafka_brokers, conf_kafka_topic}

func missingConfs()[]string{
	var missConfs []string  
	for _,c := range mustConfs{
		if !viper.IsSet(c){
			missConfs = append(missConfs, c)
		}
	}
	return missConfs
}