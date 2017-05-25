package main

import(
    "fmt"
    "github.com/spf13/viper"
    log "github.com/Sirupsen/logrus"
    "time"
    "os"
    topo "topology"
    "execute"
    "result"
)
var confFile = "resource/appConf"
func init(){
    cliArgs := os.Args[1:]
    if( len(cliArgs) !=1){
        fmt.Println("continuing with default conf file", confFile)
    }else{
        confFile = cliArgs[0]    
    }
    
    viper.SetConfigName(confFile)     // no need to include file extension
    viper.AddConfigPath("/")
    viper.AddConfigPath(".")
    viper.AddConfigPath("$HOME")
    // viper.AddConfigPath("../../resource/")  // set the path of your config file
    err := viper.ReadInConfig()
    if err != nil {
        fmt.Println("Failed to load configs from", confFile)
        fmt.Println(err)
    }else{
        fmt.Println("Configs loaded from file:", confFile)
    }
    log.SetFormatter(&log.JSONFormatter{})
    log.SetOutput(os.Stdout)
    logLevel:= viper.GetString(conf_log_level)
    lvl,err := log.ParseLevel(logLevel)
    if !viper.IsSet(conf_log_level){
        log.SetLevel(log.DebugLevel)
        log.WithFields(log.Fields{"module":"main",
         "error":"missing conf log_level"}).Info("setting log level to default: Debug")
    }else if err !=nil{
        log.SetLevel(log.DebugLevel)
        log.WithFields(log.Fields{"module":"main", "error":err}).Info("error parsinglog level,setting to default: Debug")
    }else{
        log.SetLevel(lvl)
        log.WithFields(log.Fields{"module":"main"}).Info("setting log level to:"+logLevel)
    }
}
func main(){
    m:= missingConfs();
    if len(m)>0{
        log.WithFields(log.Fields{"module":"main","confFields":m,"error":"missing mandatory confs"}).Panic("shutting down")

    }
    // var serviceDao topo.ServiceDao = &topo.FileServiceDao{viper.GetString(conf_service_path)}
    zkConn, zkErr:= ConnectZk()
    if zkErr !=nil{
        log.WithFields(log.Fields{"module":"main","error":zkErr}).Fatal("unable to connect to ZK")
        return
    }
    var serviceDao topo.ServiceDao = &topo.ZkServiceDao{zkConn}
    serviceDao.Init()
    var REDao execute.RuleEngineDao = execute.RuleEngineDao{viper.GetString(conf_re_host),
     viper.GetInt(conf_re_port), viper.GetString(conf_re_user),viper.GetString(conf_re_pass)}
    var exec execute.Executor
    var batchInterval = 7*time.Second
    if(!viper.IsSet(conf_batch_interval)){
        log.WithFields(log.Fields{"module":"main", "error":"batch_interval conf missing"}).Error("setting default: "+batchInterval.String())
    }else{
        batchInterval,err := time.ParseDuration(viper.GetString(conf_batch_interval))
        if err !=nil{
            log.WithFields(log.Fields{"module":"main","error":err}).Error("unable to parse batch-interval from conf")
        }else{
            log.WithFields(log.Fields{"module":"main"}).Info("setting batch interval to "+batchInterval.String())    
        }
        
    }
    exec  = & execute.IntervalExec{Interval:batchInterval, ServiceStore:serviceDao, REDao: REDao}
    log.WithFields(log.Fields{"module": "main",}).Info("starting executor")    
    
    out:= exec.StartExec()
    dispatcher := result.SimpleDispatcher{Consumers:[]result.EventConsumer{&result.EventLogger{}, 
        & result.KafkaForwarder{viper.GetStringSlice(conf_kafka_brokers)}}}
    // dispatcher := result.SimpleDispatcher{Consumers:[]result.EventConsumer{&result.EventLogger{}}}
    log.WithFields(log.Fields{"module": "main",}).Info("starting dispatcher")
    dispatcher.StartDispatch(out)
    //wait forever
    select{}
    
    // exec.StopExec()
    
}


