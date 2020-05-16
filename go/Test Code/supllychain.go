package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)
type SupllyChain struct{

}

// Here create a new type named Order.
type Order struct{
	// When edit this args please check argsLength in function Invoke.
	OrderID string `json:orderid`
	TruckID string `json:truckid`
	GpscordsX string `json:gpscordsx`
	GpscordsY string `json:gpscordsy`

	// New Attributes.
	MedicineType string `json:medicinetype`
	Temperature string `json:temperature`
	Humidity string `json:humidity`

	CarrierName string `json:carriername` // I think that is not better than CarrierID.
	// CarrierID string `json:carrierid`
}

func constructOrderKey(orderId string) string{
	return fmt.Sprintf("order_%s",orderId)
}
func consrtuctTruckKey(truckId string) string{
	return fmt.Sprintf("truck_%s",truckId)
}


func (c *SupllyChain) Init(stub shim.ChaincodeStubInterface) pb.Response{
	return shim.Success(nil)
}

func (c *SupllyChain) Invoke(stub shim.ChaincodeStubInterface) pb.Response{
	funcName,args:=stub.GetFunctionAndParameters()
	switch funcName{
	// Decide which function to uses here.
	case "Println":
		fmt.Println(funcName,args)
		return shim.Success(nil)
	default:
		return shim.Error(fmt.Sprintf("Unsupported function: %s.",funcName))
	}
}

func (c *SupllyChain)ordersRegister(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	argsLength:=4
	if len(args)!=argsLength {
		return shim.Error(fmt.Sprintf("It needs %d args but provides %d.",argsLength,len(args)))
	}
	// Check args whether invalid.
	// Init Orders.
	// Check whether exists this orders.
	// Put into json and check.
	// Put new order.
	return shim.Success(nil)
}

func (c *SupllyChain)gpsChange(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	if len(args)!=3{
		return shim.Error(fmt.Sprintf("It needs 3 args id,x,y but provides %d.",len(args)))
	}

	return shim.Success(nil)
}

func (c *SupllyChain)queryOrders(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	if len(args)!=1{
		return shim.Error(fmt.Sprintf("It needs 1 args id but provides %d.",len(args)))
	}
	objectID:=args[0]
	orderID:=strings.Replace(objectID," ","",-1)
	if orderID==""{
		return shim.Error("Invalid args.")
	}
	orderBytes,err:=stub.GetState(constructOrderKey(orderID))
	if err!=nil||len(orderBytes)==0{
		return shim.Error(fmt.Sprintf("Not found order %s.",orderID))
	}
	return shim.Success(orderBytes)
}

func (c *SupllyChain)queryOrdersHistory(stub shim.ChaincodeStubInterface,args []string) pb.Response{
	if len(args)!=1{
		return shim.Error(fmt.Sprintf("It needs 1 args id but provides %d.",len(args)))
	}
	objectID:=args[0]
	orderID:=strings.Replace(objectID," ","",-1)
	if orderID==""{
		return shim.Error("Invalid args.")
	}
	resultIter,err:=stub.GetHistoryForKey(constructOrderKey(orderID))
	if err!=nil{
		resultIter.Close()
		return shim.Error(fmt.Sprintf("Not fount order",orderID))
	}
	defer resultIter.Close()
	var buffer bytes.Buffer
	buffer.WriteString("[")

	isWrited:=false
	for resultIter.HasNext(){
		queryResponse,error:=resultIter.Next()
		if error!=nil{
			return shim.Error(error.Error())
		}
		if isWrited{
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(",\"Value\":")
		if queryResponse.IsDelete{
			buffer.WriteString("null")
		}else {
			buffer.WriteString(string(queryResponse.Value))
		}

		buffer.WriteString(",\"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(queryResponse.Timestamp.Seconds,int64(queryResponse.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(",\"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(queryResponse.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		isWrited=true
	}

	buffer.WriteString("]")

	fmt.Printf("getHistoryForOrders result : \n%s\n", buffer.String())
	fmt.Println("Finished getHistoryforOrders.")
	return shim.Success(buffer.Bytes())
}

func main(){
	fmt.Println("Start to run SupllyChain chaincode.")
	err:=shim.Start(new(SupllyChain))
	if err!=nil{
		fmt.Printf("Error starting SupllyChain chaincode: %s.",err)
	}
}
