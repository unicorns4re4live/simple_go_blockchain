package network

import (
	"net"
	"encoding/json"
	"../utils"
	"../models"
	"../settings"
)

func SendPackage(pack *models.PackageTCP) {
	conn, err := net.Dial("tcp", pack.To.Addr)
	if err != nil {
		disconnectFrom(pack.To.Addr)
		return
	}
	defer conn.Close()
	pack.From.Hash = settings.User.Hash
	pack.From.Addr = settings.User.Addr.IPv4 + settings.User.Addr.Port
	data, err := json.Marshal(*pack)
	utils.CheckError(err)
	conn.Write(data)
}

func disconnectFrom(address string) {
	delete(settings.Connections, address)
}
