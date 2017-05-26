package zk

import (	
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"unicode/utf8"
)
var separator = "/"

var flags = int32(0)
var acl = zk.WorldACL(zk.PermAll)
//createDirPath creates all the znodes for the given path
// equivalent to bash command `mkdir -p`
func CreatePath(Conn *zk.Conn, path string, tailData []byte)error{
	first,size:= utf8.DecodeRuneInString(path)
	if string(first) == separator{
		path = path[size:]
	}
	parts:= strings.Split(path, separator)
	for i:= range parts{
		subPath:= "/"+strings.Join(parts[:i], separator)
		exists,_,err := Conn.Exists(subPath)
		if err!=nil{
			return err
		}else if !exists{
			if i == len(parts)-1{
				_,cErr:= Conn.Create(subPath, tailData, flags, acl)	
				return cErr
			}else{
				_,cErr:= Conn.Create(subPath, make([]byte,0), flags, acl)	
				return cErr
			}
			
			
		}
	}
	return nil
}