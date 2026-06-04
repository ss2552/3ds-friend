package types

import (
	"github.com/PretendoNetwork/nex-go/v2"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/v2/friends-3ds/types"
)

type ConnectedUser struct {
	PID        uint32
	Platform   Platform
	Connection *nex.PRUDPConnection
	Presence   friends_3ds_types.NintendoPresence
}

func NewConnectedUser() *ConnectedUser {
	return &ConnectedUser{
		Presence:   friends_3ds_types.NewNintendoPresence(),
	}
}
