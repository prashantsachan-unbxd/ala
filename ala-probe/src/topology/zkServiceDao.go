package topology
import(
	"fmt"
	"errors"
	"github.com/samuel/go-zookeeper/zk"
    log "github.com/Sirupsen/logrus"
    "encoding/json"
    zkUtil "util/zk"
)
const RootNode = "/metricCollect/services"
var flags = int32(0)
var acl = zk.WorldACL(zk.PermAll)
type ZkServiceDao struct{
	Conn *zk.Conn
}

func (this *ZkServiceDao) Init(){
	// create if the root path doesn't exist
	zkErr:= zkUtil.CreatePath(this.Conn, RootNode,[]byte("Root Node for services") )
	if(zkErr!=nil){
		log.WithFields(log.Fields{"module":"zkServiceDao","action":"create",
			"path":RootNode, "error":zkErr}).Info("unable to create RootNode")
	}
		
}
// Fetches to get all The services, returns the first 
func (this *ZkServiceDao) GetAllServices()([]Service,error){
	// get all children of the RootNode
	ids,_,err := this.Conn.Children(RootNode)
	if err!=nil{
		return nil,err
	}else{
		var services []Service
		var failedIds []string
		for _,id := range ids{
			path := RootNode+"/"+id
			data,_,zkErr :=this.Conn.Get(path)
			if zkErr !=nil{
				log.WithFields(log.Fields{"module":"zkServiceDao","action":"getAll",
					"id":id,"error":zkErr}).Error("error fetching service from zk")
				failedIds = append(failedIds, id)
			}else{
				var s Service;
				jErr:= json.Unmarshal(data, &s)
				if jErr !=nil{
					log.WithFields(log.Fields{"module":"zkServiceDao","action":"getAll",
						"data":string(data),"error":jErr}).Error("error parsing to Service")
					failedIds = append(failedIds, id)	
				}else{
					services = append(services, s)
				}
			}
		}
		if len(failedIds)==0 {
			return services, nil
		}else{
			msg:= fmt.Sprintf("failed to retrieve services for Ids: %v",failedIds)
			return services, errors.New(msg)
		}
	}
}

func (this *ZkServiceDao) AddService(s Service)error{
	// store this service as a child of root
	path := RootNode+"/"+s.Id
	data,jErr := json.Marshal(s)
	if jErr !=nil{
		return jErr
	}else{
		_, zkErr := this.Conn.Create(path, data, flags, acl)	
		return zkErr
	}
}
func (this *ZkServiceDao) DeleteService(id string) error{
	// delete this service
	path := RootNode+"/"+id
	zkErr := this.Conn.Delete(path, -1)
	return zkErr
	
}
func (this *ZkServiceDao)GetService(id string)(Service,error){
	path := RootNode+"/"+ id
	exists,_,zkErr:= this.Conn.Exists(path)
	if zkErr!=nil{
		return Service{},zkErr
	}else if  !exists{
		return Service{},errors.New("No service with id: "+id)
	}else{
		data,_,err:= this.Conn.Get(path)
		if err!=nil{
			return Service{}, err
		}
		var s Service
		jErr:= json.Unmarshal(data, &s)
		if jErr !=nil{
			return Service{}, jErr
		}else{
			return s,nil
		}
	}
}
