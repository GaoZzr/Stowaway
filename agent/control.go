package agent

import (
	"Stowaway/node"
	"fmt"
	"strconv"
	"time"
)

//not in use,add todo
func TryReconnect(gap string) {
	for {
		lag, _ := strconv.Atoi(gap)
		time.Sleep(time.Duration(lag) * time.Second)

		controlConnToAdmin, dataConnToAdmin, _, err := node.StartNodeConn(Monitor, ListenPort, NODEID, AESKey)
		if err != nil {
			fmt.Println("[*]Admin seems still down")
		} else {
			fmt.Println("[*]Admin up! Reconnect successful!")
			ControlConnToAdmin = controlConnToAdmin
			DataConnToAdmin = dataConnToAdmin
			return
		}
	}
}
