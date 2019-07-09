package main

import (
	"encoding/json"
    "fmt"
  //  "strconv"
  //  "strings"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)
type FoodChainCode struct{	
}

//Data Structure
type FoodInfo struct{
    FoodID string `json:FoodID`                             //foodId
    FoodProInfo ProInfo `json:FoodProInfo`                  //Product Information
    FoodIngInfo []IngInfo `json:FoodIngInfo`                //Ingredient Information
    FoodLogInfo LogInfo `json:FoodLogInfo`                  //Logistics Information
}

type FoodAllInfo struct{
    FoodID string `json:FoodId`
    FoodProInfo ProInfo `json:FoodProInfo`
    FoodIngInfo []IngInfo `json:FoodIngInfo`
    FoodLogInfo []LogInfo `json:FoodLogInfo`
}

//Product Information
type ProInfo struct{
    FoodName string `json:FoodName`                         //Food name
    FoodSpec string `json:FoodSpec`                         //Food specification
    FoodMFGDate string `json:FoodMFGDate`                   //Food Production date
    FoodEXPDate string `json:FoodEXPDate`                   //Food Shelf life
    FoodLOT string `json:FoodLOT`                           //Food batch number
    FoodQSID string `json:FoodQSID`                         //Food Production license number
    FoodMFRSName string `json:FoodMFRSName`                 //Food Manufacturer name
    FoodProPrice string `json:FoodProPrice`                 //Food Production Price
    FoodProPlace string `json:FoodProPlace`                 //Food Production Location
}
type IngInfo struct{
    IngID string `json:IngID`                               //Ingredients Id
    IngName string `json:IngName`                           //Ingredients name
}

type LogInfo struct{
    LogDepartureTm string `json:LogDepartureTm`             //Departure time
    LogArrivalTm string `json:LogArrivalTm`                 //Time of Arrival
    LogMission string `json:LogMission`                     //Processing Business(storage or transportation)
    LogDeparturePl string `json:LogDeparturePl`             //Departure
    LogDest string `json:LogDest`                           //Destination
    LogToSeller string `json:LogToSeller`                   //Sellers
    LogStorageTm string `json:LogStorageTm`                 //Storage time
    LogMOT string `json:LogMOT`                             //Shipping Method
    LogCopName string `json:LogCopName`                     //Logistics Company name
    LogCost string `json:LogCost`                           //Cost
}

func (a *FoodChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}

func (a *FoodChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    fn,args := stub.GetFunctionAndParameters()
    if fn == "addProInfo"{
        return a.addProInfo(stub,args)
    } else if fn == "addIngInfo"{
        return a.addIngInfo(stub,args)
    } else if fn == "getFoodInfo"{
        return a.getFoodInfo(stub,args)
    }else if fn == "addLogInfo"{
        return a.addLogInfo(stub,args)
    }else if fn == "getProInfo"{
        return a.getProInfo(stub,args)
    }else if fn == "getLogInfo"{
        return a.getLogInfo(stub,args)
    }else if fn == "getIngInfo"{
        return a.getIngInfo(stub,args)
    }else if fn == "getLogInfo_l"{
        return a.getLogInfo_l(stub,args)
    }

    return shim.Error("Recevied unkown function invocation")
}

func (a *FoodChainCode) addProInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error 
    var FoodInfos FoodInfo

    if len(args)!=10{
        return shim.Error("Incorrect number of arguments.")
    }
    FoodInfos.FoodID = args[0]
    if FoodInfos.FoodID == ""{
        return shim.Error("FoodID can not be empty.")
    }
    
    
    FoodInfos.FoodProInfo.FoodName = args[1]
    FoodInfos.FoodProInfo.FoodSpec = args[2]
    FoodInfos.FoodProInfo.FoodMFGDate = args[3]
    FoodInfos.FoodProInfo.FoodEXPDate = args[4]
    FoodInfos.FoodProInfo.FoodLOT = args[5]
    FoodInfos.FoodProInfo.FoodQSID = args[6]
    FoodInfos.FoodProInfo.FoodMFRSName = args[7]
    FoodInfos.FoodProInfo.FoodProPrice = args[8]
    FoodInfos.FoodProInfo.FoodProPlace = args[9]
    ProInfosJSONasBytes,err := json.Marshal(FoodInfos)
    if err != nil{
        return shim.Error(err.Error())
    }

    err = stub.PutState(FoodInfos.FoodID,ProInfosJSONasBytes)
    if err != nil{
        return shim.Error(err.Error())
    }

    return shim.Success(nil)
}

func(a *FoodChainCode) addIngInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
        
    var FoodInfos FoodInfo
    var IngInfoitem IngInfo

    if  (len(args)-1)%2 != 0 || len(args) == 1{
        return shim.Error("Incorrect number of arguments")
    }

    FoodID := args[0]
    for i :=1;i < len(args);{   
        IngInfoitem.IngID = args[i]
        IngInfoitem.IngName = args[i+1]
        FoodInfos.FoodIngInfo = append(FoodInfos.FoodIngInfo,IngInfoitem)
        i = i+2
    }
    
    
    FoodInfos.FoodID = FoodID
  /*  FoodInfos.FoodIngInfo = foodIngInfo*/
    IngInfoJsonAsBytes,err := json.Marshal(FoodInfos)
    if err != nil {
    return shim.Error(err.Error())
    }

    err = stub.PutState(FoodInfos.FoodID,IngInfoJsonAsBytes)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
        
}

func(a *FoodChainCode) addLogInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
 
    var err error
    var FoodInfos FoodInfo

    if len(args)!=11{
        return shim.Error("Incorrect number of arguments.")
    }
    FoodInfos.FoodID = args[0]
    if FoodInfos.FoodID == ""{
        return shim.Error("FoodID can not be empty.")
    }
    FoodInfos.FoodLogInfo.LogDepartureTm = args[1]
    FoodInfos.FoodLogInfo.LogArrivalTm = args[2]
    FoodInfos.FoodLogInfo.LogMission = args[3]
    FoodInfos.FoodLogInfo.LogDeparturePl = args[4]
    FoodInfos.FoodLogInfo.LogDest = args[5]
    FoodInfos.FoodLogInfo.LogToSeller = args[6]
    FoodInfos.FoodLogInfo.LogStorageTm = args[7]
    FoodInfos.FoodLogInfo.LogMOT = args[8]
    FoodInfos.FoodLogInfo.LogCopName = args[9]
    FoodInfos.FoodLogInfo.LogCost = args[10]
    
    LogInfosJSONasBytes,err := json.Marshal(FoodInfos)
    if err != nil{
        return shim.Error(err.Error())
    } 
    err = stub.PutState(FoodInfos.FoodID,LogInfosJSONasBytes)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
}



func(a *FoodChainCode) getFoodInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }
    FoodID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(FoodID)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    var foodAllinfo FoodAllInfo

    for resultsIterator.HasNext(){
        var FoodInfos FoodInfo
        response,err :=resultsIterator.Next()
        if err != nil {
             return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&FoodInfos)
        if FoodInfos.FoodProInfo.FoodName !=""{
            foodAllinfo.FoodProInfo = FoodInfos.FoodProInfo
        }else if FoodInfos.FoodIngInfo != nil{
            foodAllinfo.FoodIngInfo = FoodInfos.FoodIngInfo
        }else if FoodInfos.FoodLogInfo.LogMission !=""{
            foodAllinfo.FoodLogInfo = append(foodAllinfo.FoodLogInfo,FoodInfos.FoodLogInfo)
        }

    }
    
    jsonsAsBytes,err := json.Marshal(foodAllinfo)
    if err != nil{
        return shim.Error(err.Error())
    }

    return shim.Success(jsonsAsBytes)
}
 

func(a *FoodChainCode) getProInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
    
    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }
    FoodID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(FoodID)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    var foodProInfo ProInfo

    for resultsIterator.HasNext(){
        var FoodInfos FoodInfo
        response,err :=resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&FoodInfos)
        if FoodInfos.FoodProInfo.FoodName != ""{
            foodProInfo = FoodInfos.FoodProInfo
            continue
        }
    }
    jsonsAsBytes,err := json.Marshal(foodProInfo)   
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}

func(a *FoodChainCode) getIngInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{
 
    if len(args) !=1{
        return shim.Error("Incorrect number of arguments.")
    }
    FoodID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(FoodID)

    if err != nil{
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()
    
    var foodIngInfo []IngInfo
    for resultsIterator.HasNext(){
        var FoodInfos FoodInfo
        response,err := resultsIterator.Next()
        if err != nil{
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&FoodInfos)
        if FoodInfos.FoodIngInfo != nil{
            foodIngInfo = FoodInfos.FoodIngInfo
            continue
        }
    }
    jsonsAsBytes,err := json.Marshal(foodIngInfo)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}

func(a *FoodChainCode) getLogInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{

    var LogInfos []LogInfo

    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }

    FoodID := args[0]
    resultsIterator,err :=stub.GetHistoryForKey(FoodID)
    if err != nil{
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

   
    for resultsIterator.HasNext(){
        var FoodInfos FoodInfo
        response,err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&FoodInfos)
        if FoodInfos.FoodLogInfo.LogMission != ""{
            LogInfos = append(LogInfos,FoodInfos.FoodLogInfo)
        }
    }
    jsonsAsBytes,err := json.Marshal(LogInfos)
    if err != nil{
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}

func(a *FoodChainCode) getLogInfo_l(stub shim.ChaincodeStubInterface,args []string) pb.Response{
    var Loginfo LogInfo

    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }

    FoodID := args[0]
    resultsIterator,err :=stub.GetHistoryForKey(FoodID)
    if err != nil{
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

   
    for resultsIterator.HasNext(){
        var FoodInfos FoodInfo
        response,err := resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&FoodInfos)
        if FoodInfos.FoodLogInfo.LogMission != ""{
           Loginfo = FoodInfos.FoodLogInfo
           continue 
       }
    }
    jsonsAsBytes,err := json.Marshal(Loginfo)
    if err != nil{
        return shim.Error(err.Error ())
    }
    return shim.Success(jsonsAsBytes)
}


func main(){
     err := shim.Start(new(FoodChainCode))
     if err != nil {
         fmt.Printf("Error starting Food chaincode: %s ",err)
     }
}
