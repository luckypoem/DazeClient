package obscure

import (
	"net"
	"net/http"
	"math/rand"
	"github.com/crabkun/DazeClient/util"
	"bytes"
	"bufio"
)

type HttpPost struct {
}

func (this *HttpPost) Action(conn net.Conn , param string) (error){
	var err error
	body:=make([]byte,0)
	bodystr:=util.GetRandomString(rand.Intn(10))+"="+
	util.GetRandomString(rand.Intn(512))+"&"+
	util.GetRandomString(rand.Intn(10))+"="+
	util.GetRandomString(rand.Intn(512))
	body=[]byte(bodystr)
	req,err:=http.NewRequest("POST","http://"+param+"/"+util.GetRandomString(rand.Intn(10))+".php",bytes.NewReader(body))
	if err!=nil{
		return err
	}
	req.Header=make(http.Header)
	req.Header.Add("Connection","Keep-Alive")
	req.Header.Add("Accept","*/*")
	req.Write(conn)
	reader:=bufio.NewReader(conn)
	_,err=http.ReadResponse(reader,nil)
	if err!=nil{
		return err
	}
	req.Write(conn)
	return nil
}
func init(){
	RegisterObscure("http_post",new(HttpPost))
}