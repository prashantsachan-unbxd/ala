package main

import(
	"github.com/spf13/viper"
)

var conf_batch_interval = "execute.batch_interval"
var conf_log_level = "log_level"

//execute section
var secEx = "execute."
var conf_service_path =  secEx+"path_service_conf"
// Rule Engine section
var secRE = "ruleEngine."
var conf_re_host = secRE + "host"
var conf_re_port = secRE + "port"
var conf_re_user = secRE + "user"
var conf_re_pass = secRE + "pass"

// Kafka section
var secKafka = "kafka."
var conf_kafka_brokers = secKafka+ "brokers_hostport"

var mustConfs = []string{conf_batch_interval, 
	conf_service_path,
	conf_re_host, conf_re_port,
	conf_kafka_brokers}

func missingConfs()[]string{
	var missConfs []string  
	for _,c := range mustConfs{
		if !viper.IsSet(c){
			missConfs = append(missConfs, c)
		}
	}
	return missConfs
}