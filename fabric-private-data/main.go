package main                                                                                                                                                                     
                                                                                                                                                                                 
import (                                                                                                                                                                         
    "net/http"                                                                                                                                                                   
    "log"                                                                                                                                                                        
    "github.com/gorilla/mux"
    "os/exec"
    "fmt"                                                                                                                                                                        
)                                                                                                                                                                                
                                                                                                                                                                                 
func query(w http.ResponseWriter, r *http.Request) {                                                                                                                       
    vars := mux.Vars(r)
    id := vars["deviceid"]
    cmd := fmt.Sprintf("peer chaincode query -o orderer0.example.com:7050 -n device -c '{\"Args\":[\"readDevice\",\"%s\"]}' -C mychannel", id)                                                                                                                                                            
    out, err := exec.Command("bash", "-c",cmd).Output()
 
    if err != nil {                                                                                                                                                              
        log.Fatal(err)                                                                                                                                                           
    }                                                                                                                                                                            
    //RESPONS                                                                                                                                                                    
    w.Write([]byte(out))                                                                                                                                                         
}

func invoke(w http.ResponseWriter, r *http.Request) {                                                                                                                       
    vars := mux.Vars(r)
    id := vars["deviceid"]
    deviceType:=vars["type"]
    data:=vars["data"]

    cmd := fmt.Sprintf("peer chaincode invoke -o orderer0.example.com:7050 -n device -c '{\"Args\":[\"sendDeviceReading\",\"%s\",\"%s\",\"%s\"]}' -C mychannel", id,deviceType,data)                                                                                                                                                            
    out, err := exec.Command("bash", "-c",cmd).Output()   
    
    if err != nil {                                                                                                                                                              
        log.Fatal(err)                                                                                                                                                           
    }                                                                                                                                                                            
    //RESPONS                                                                                                                                                                    
    w.Write([]byte(out))                                                                                                                                                         
}

func delete(w http.ResponseWriter, r *http.Request) {                                                                                                                       
    vars := mux.Vars(r)
    id := vars["deviceid"]
  

    cmd := fmt.Sprintf("peer chaincode invoke -o orderer0.example.com:7050 -n device -c '{\"Args\":[\"delete\",\"%s\"]}' -C mychannel", id)                                                                                                                                                            
    out, err := exec.Command("bash", "-c",cmd).Output()   
    
    if err != nil {                                                                                                                                                              
        log.Fatal(err)                                                                                                                                                           
    }                                                                                                                                                                            
    //RESPONS                                                                                                                                                                    
    w.Write([]byte(out))                                                                                                                                                         
}

func getHistory(w http.ResponseWriter, r *http.Request) {                                                                                                                       
    vars := mux.Vars(r)
    id := vars["deviceid"]
  

    cmd := fmt.Sprintf("peer chaincode invoke -o orderer0.example.com:7050 -n device -c '{\"Args\":[\"getHistoryForDevice\",\"%s\"]}' -C mychannel", id)                                                                                                                                                            
    out, err := exec.Command("bash", "-c",cmd).Output()   
   
    if err != nil {                                                                                                                                                              
        log.Fatal(err)                                                                                                                                                           
    }                                                                                                                                                                            
    //RESPONS                                                                                                                                                                    
    w.Write([]byte(out))                                                                                                                                                         
} 
                                                                                                                                                                       
func main() {                                                                                                                                                                    
    r := mux.NewRouter()                                                                                                                                                         
    // Routes consist of a path and a handler function.                                                                                                                          
    r.HandleFunc("/query/{deviceid}", query)
    r.HandleFunc("/invoke/{deviceid}/{type}/{data}", invoke)
    r.HandleFunc("/delete/{deviceid}", delete)
    r.HandleFunc("/getHistory/{deviceid}", getHistory)
    
    //TODO for production
    //r.HandleFunc("/login/{deviceid}", login)
    //r.HandleFunc("/register/{deviceid}", register)
    //r.HandleFunc("/list/{deviceid}", list)                                                                                                                                 
                                                                                                                                                                                 
    // Bind to a port and pass our router in                                                                                                                                     
    log.Fatal(http.ListenAndServe(":8000", r))                                                                                                                                   
}                                                                                                                                                                                
