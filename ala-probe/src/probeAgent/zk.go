package main
import (
	"github.com/samuel/go-zookeeper/zk"
    "github.com/spf13/viper"
    log "github.com/Sirupsen/logrus"
	"time"
)
const conf_sec_zk = "zk."
const conf_zk_servers = conf_sec_zk+"servers_hostport"
const conf_zk_session_timeout = conf_sec_zk+"sessionTimeout"
const default_session_timeout = time.Second

func ConnectZk() (*zk.Conn,error) {
	zkServers := viper.GetStringSlice(conf_zk_servers)
	timeout := default_session_timeout
	if(viper.IsSet(conf_zk_session_timeout)){
		t,err:= time.ParseDuration(viper.GetString( conf_zk_session_timeout))
		if err !=nil{
			log.WithFields(log.Fields{"module":"zkDao","error":err,
			"key":conf_zk_session_timeout}).Info("error parsing conf value to time, setting default: "+default_session_timeout.String())	
			timeout = default_session_timeout
		}else{
			timeout = t
		}
	} else{
		log.WithFields(log.Fields{"module":"zkDao","error":"missing conf",
			"key":conf_zk_session_timeout}).Info("setting default: "+default_session_timeout.String())
	}
	conn, _, err := zk.Connect(zkServers, timeout)
	return conn,err
}