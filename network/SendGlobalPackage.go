package network

import (
	"../models"
	"../settings"
)

func SendGlobalPackage(pack *models.PackageTCP) {
	for addr := range settings.Connections {
		pack.To.Addr = addr
		SendPackage(pack)
	}
}
