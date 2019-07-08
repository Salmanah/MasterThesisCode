package main                                                                                                                                                                     
                                                                                                                                                                                 
import (                                                                                                                                                                         
    "net/http"                                                                                                                                                                   
    "log"                                                                                                                                                                        
    "github.com/gorilla/mux"
    "os/exec"
    "fmt"       
    "net" z                                                                                                                                                            
)       

  
// DeviceReading struct 
type User struct{
	IpAddr          string `json:"ipaddr"`
    Admin 			bool `json:"Admin"`
    Deviceid 			bool `json:"Deviceid"`
}
var whitelist = make([]User, 0)


func query(w http.ResponseWriter, r *http.Request){                                                                                                                       
    vars := mux.Vars(r)
    id := vars["deviceid"]

  
    
    if authorizationAdmin(r) == false{
        w.Write([]byte("Admin not detected!"))
        return
    }
    
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
    
    if authorizationIoT(r) == false{
        w.Write([]byte("Permission denied!"))
        return
    }
    
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

    if authorizationIoT(r) == false{
        w.Write([]byte("Permission denied!"))
        return
    }

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
    
    if authorizationIoT(r) == false{
        w.Write([]byte("Permission denied!"))
        return
    }

    cmd := fmt.Sprintf("peer chaincode invoke -o orderer0.example.com:7050 -n device -c '{\"Args\":[\"getHistoryForDevice\",\"%s\"]}' -C mychannel", id)                                                                                                                                                            
    out, err := exec.Command("bash", "-c",cmd).Output()   
   
    if err != nil {                                                                                                                                                              
        log.Fatal(err)                                                                                                                                                           
    }                                                                                                                                                                            
    //RESPONS                                                                                                                                                                    
    w.Write([]byte(out))                                                                                                                                                         
}




func registerAdmin(w http.ResponseWriter, r *http.Request) {                                                                                                                       
    vars := mux.Vars(r)
    incomingIP, _, err := net.SplitHostPort(r.RemoteAddr)
    admin := false

    if err != nil{
        fmt.Println("Could not get IP")
    }

    for _, checkUser := range whitelist {
        if checkUser.IpAddr == incomingIP{
            admin = true
        }
    }

    if admin != true{
        w.Write([]byte("No Admin detected"))
        return 
    }
    
    newAdmin := User{
        IpAddr:vars["ip"],
        Admin:true,
    }
    whitelist = append(whitelist,newAdmin)
    out := fmt.Sprintf("Added new admin with ip %s", vars["ip"])                                                                                                                                                               
    //RESPONS                                                                                                                                                                    
    w.Write([]byte(out))                                                                                                                                                         
}

func registerIoT(w http.ResponseWriter, r *http.Request) {                                                                                                                       
    vars := mux.Vars(r)
    incomingIP, _, err := net.SplitHostPort(r.RemoteAddr)
    admin := false


    if error != nil{
        fmt.Println("Encryption failed")
        return 
    }

    if err != nil{
        fmt.Println("Could not get IP")
    }

    for _, checkUser := range whitelist {
        if checkUser.IpAddr == incomingIP{
            admin = true
        }
    }

    if admin != true{
        w.Write([]byte("No Admin detected"))
        return 
    }
    
    newAdmin := User{
        IpAddr:vars["ip"],
        Admin:false,
        deviceId: h.sum(nil)
    }
    whitelist = append(whitelist,newAdmin)
    out := fmt.Sprintf("Added new IoT with ip %s", vars["ip"])                                                                                                                                                               
    //RESPONS                                                                                                                                                                    
    w.Write([]byte(out))                                                                                                                                                         
}

func authorizationAdmin(r *http.Request) bool {                                                                                                                       
    incomingIP, _, err := net.SplitHostPort(r.RemoteAddr)

    if err != nil{
        fmt.Println("Could not get IP")
        return false
    }

    for _, checkUser := range whitelist {
        if checkUser.IpAddr == incomingIP && checkUser.Admin == true{
            return true
        }
    }

    return false                                                                                                                           
}

func authorizationIoT(r *http.Request) bool {                                                                                                                       
    incomingIP, _, err := net.SplitHostPort(r.RemoteAddr)

    if err != nil{
        fmt.Println("Could not get IP")
        return false
    }

    for _, checkUser := range whitelist {
        if checkUser.IpAddr == incomingIP{
            return true
        }
    }

     return false
                                                                                                                                                      
} 
                                                                                                                                                                       
func main() {              
    superAdmin := User{
        IpAddr:"158.37.63.234",
        Admin:true,
    }                                                                                                                                                      
    r := mux.NewRouter()
    whitelist = append(whitelist,superAdmin)             
    // Routes consist of a path and a handler function.                                                                                                                          
    r.HandleFunc("/query/{deviceid}", query)
    r.HandleFunc("/invoke/{deviceid}/{type}/{data}", invoke)
    r.HandleFunc("/delete/{deviceid}", delete)
    r.HandleFunc("/getHistory/{deviceid}", getHistory)
    
    //TODO for production
    //r.HandleFunc("/login/{deviceid}", login)
    r.HandleFunc("/registerAdmin/{ip}", registerAdmin)
    r.HandleFunc("/registerIoT/{ip}/", registerIoT)
    //r.HandleFunc("/list/{deviceid}", list)                                                                                                                                 
                                                                                                                                                                                 
    // Bind to a port and pass our router in                                                                                                                                     
    log.Fatal(http.ListenAndServe(":8000", r))                                                                                                                                   
}                                                                                                                                                                                
