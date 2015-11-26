package main 

import (
    "fmt"
    "github.com/pmylund/go-cache"
    "time"
    "log"
    "net/http"
    "strings"
    "encoding/json"
    "bytes"
    "encoding/binary"
)

type Response [7]struct{
    Key string
    Value string
}

var r Response
var b []byte
var str []string
var i int=0
var c *cache.Cache

func prettyprint(b []byte) ([]byte, error) {
    var out bytes.Buffer
    err := json.Indent(&out, b, "", "  ")
    return out.Bytes(), err
}
func getting(rw http.ResponseWriter, req *http.Request){
    if(req.Method=="GET"){//------------------------> GET---------------------------------->
    s1:=req.URL.Path[1:]
    st1:=string(s1[4:])
    //str=strings.Split(st1,"/")
    if(st1==""){
        b,_=json.Marshal(r) 
        bx, _ := prettyprint(b)
        n:=binary.Size(bx)
        s := string(bx[:n])
        fmt.Fprintf(rw,s)
        }
    }
}
func posting(rw http.ResponseWriter, req *http.Request) {
    
    if(req.Method=="PUT"){//------------------------> PUT---------------------------------->


    s1:=req.URL.Path[1:]
    st1:=string(s1[5:])
    str=strings.Split(st1,"/")
    
    c.Set(str[0], str[1], cache.DefaultExpiration)
    foo, found := c.Get(str[0])
    if found {
        r[i].Key=str[0]
        r[i].Value=str[1]
        fmt.Println(foo)
        i++
    }

    
}else {
        s1:=req.URL.Path[1:]
        st1:=string(s1[5:])
    
        foo, found := c.Get(st1)
        if found{
            fmt.Println(foo,found)
            for i=0;i<7;i++{
                if(r[i].Value==foo){
                    b,_=json.Marshal(r[i]) 
                    bx, _ := prettyprint(b)
                    n:=binary.Size(bx)
                    s := string(bx[:n])
                      
                    fmt.Fprintf(rw,s)
                    
                }
            }
        
        }
    }
}
func main() {
    c = cache.New(5*time.Minute, 30*time.Second)
    http.HandleFunc("/keys",getting)
    http.HandleFunc("/keys/",posting)
    log.Fatal(http.ListenAndServe(":3001", nil))
    
}