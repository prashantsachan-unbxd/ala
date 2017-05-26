package probe
import(
	"fmt"
	"errors"
	"github.com/samuel/go-zookeeper/zk"
    log "github.com/Sirupsen/logrus"
    "encoding/json"
    zkUtil "util/zk"
)
const rootNode = "/metricCollect/services"
var flags = int32(0)
var acl = zk.WorldACL(zk.PermAll)
type ZkPCDao struct{
	Conn *zk.Conn
}
//Init initializes the object
func (this *ZkPCDao)Init(){
	// create if the root path doesn't exist
	zkErr:= zkUtil.CreatePath(this.Conn, rootNode, []byte("root for all ProbeConfigs") )
	log.WithFields(log.Fields{"module":"ZkPCDao","action":"create",
			"path":rootNode, "error":zkErr}).Info("unable to create rootNode")
}

//AddService Adds a probeConfig to the existing list
	func (this *ZkPCDao) AddProbeConf(p ProbeConfig, serviceClass string)error{
	// store this service as a child of root
	path := rootNode+"/"+serviceClass+"/"+p.Id
	data,jErr := json.Marshal(p)
	if jErr !=nil{
		return jErr
	}else{
		return zkUtil.CreatePath(this.Conn, path, data)
	}
}
//DeleteProbeConf deletes a probeConfig with given ID
func (this *ZkPCDao) DeleteProbeConf(serviceClass string, id string) error{
	// delete this service
	path := rootNode+"/"+serviceClass+"/"+id
	zkErr := this.Conn.Delete(path, -1)
	return zkErr
	
}
// should return non-nil error if id passed doesn't match with any service
func (this *ZkPCDao)GetProbeConf(serviceClass string,id string)(ProbeConfig,error){
	path := rootNode+"/"+serviceClass+"/"+ id
	exists,_,zkErr:= this.Conn.Exists(path)
	if zkErr!=nil{
		return ProbeConfig{},zkErr
	}else if  !exists{
		return ProbeConfig{},errors.New("No service with id: "+id)
	}else{
		data,_,err:= this.Conn.Get(path)
		if err!=nil{
			return ProbeConfig{}, err
		}
		var p ProbeConfig
		jErr:= json.Unmarshal(data, &p)
		if jErr !=nil{
			return ProbeConfig{}, jErr
		}else{
			return p,nil
		}
	}
}
func (this *ZkPCDao) GetAllProbeConfs(serviceClass string)([]ProbeConfig,error){
	cPath := rootNode+"/"+serviceClass
	// get ProbeConfig ids
	ids,_,err := this.Conn.Children(cPath)
	if err!=nil{
		return make([]ProbeConfig,0),err
	}

	var confs []ProbeConfig
	var failedIds []string
	for _,id := range ids{
		path := cPath+"/"+id
		data,_,zkErr :=this.Conn.Get(path)
		if zkErr !=nil{
			log.WithFields(log.Fields{"module":"zkPCDao","action":"getAll","serviceClass":serviceClass,
				"id":id,"error":zkErr}).Error("error fetching service from zk")
			failedIds = append(failedIds, id)
		}else{
			var p ProbeConfig
			jErr:= json.Unmarshal(data, &p)
			if jErr !=nil{
				log.WithFields(log.Fields{"module":"zkPCDao","action":"getAll","serviceClass":serviceClass,
					"data":string(data),"error":jErr}).Error("error parsing to Service")
				failedIds = append(failedIds, id)	
			}else{
				confs = append(confs, p)
			}
		}
	}
	if len(failedIds)==0 {
		return confs, nil
	}else{
		msg:= fmt.Sprintf("failed to retrieve services for Ids: %v",failedIds)
		return confs, errors.New(msg)
	}

}
func (this *ZkPCDao) GetAllClasses()([]string, error){
	// get all children of the RootNode
	classes,_,err := this.Conn.Children(rootNode)
	return classes,err
}