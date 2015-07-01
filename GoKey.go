package main

import (

	"net/http"
    
    "strings"
     
    "encode"
    
    "logger"
    
    "encoding/json"
)

type Secret struct {
    Unix int64 `json:"Unix"`
    Sign string `json:"Signature"`
}

func main() {

    logger.Log("Started")
    
    // Check to test web server is running
    http.HandleFunc("/", test)
    
    // Deeper test stuff
    http.HandleFunc("/secure/", func(w http.ResponseWriter, r *http.Request) {

        key := strings.SplitN(r.URL.Path, "/", 3)[2]
        
        signature, unix := encode.HmacSha256Encode(key, "bq1Yl5thyhy6uQUvghh67j7ijG7HbcvrrKUMc3fDdNYn/+d=") 
        
        signature = strings.Replace(signature,"-","%2b",-1)
        
        signature = strings.Replace(signature,"_","%2f",-1)
        
        signature = strings.Replace(signature,"=","%3d",-1)
        
        secretKey := &Secret{ Unix: unix, Sign: signature }    
        
        address := strings.Split(r.RemoteAddr, ":")
        
        logger.Log("Request from IP " + address[0] + " on port " + address[1] + ".  Response is: " + key)
        
        jsonResult, _ := json.Marshal(secretKey)
        
        result := string(jsonResult)

        w.Write([]byte(result))
    })

    http.ListenAndServe(":8080", nil)
}

func test(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, I'm working ok :)"))
}

