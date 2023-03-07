
package main

import (
    "fmt"
    "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
    "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
    "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
    "github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)


func main() {
    
    // Create SDK setup for the integration tests
    sdk, err := fabsdk.New(config.FromFile("./connection.yaml"))
    if err != nil {
        fmt.Println("Failed to create new SDK: ", err)
        return
    }
    
    // Prepare client context
    orgName := "org1"
    orgAdmin := "Admin"
    orgMsp := "Org1MSP"

    clientChannelContext := sdk.ChannelContext("mychannel", channel.WithUser(orgAdmin), channel.WithOrg(orgName))
    if clientChannelContext == nil {
    fmt.Println("Failed to create client channel context")
    return
    }
    
    client, err := channel.New(clientChannelContext)
     if err != nil {
        fmt.Println("Failed to create new client: ", err)
        return
    }
    
}
