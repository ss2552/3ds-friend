package nex

import (
	"github.com/ss2552/3ds-friend/globals"
	nex "github.com/PretendoNetwork/nex-go/v2"
	_ "github.com/PretendoNetwork/nex-protocols-go/v2"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewPRUDPServer()
	globals.SecureEndpoint = nex.NewPRUDPEndPoint(1)

	globals.SecureEndpoint.ServerAccount = globals.SecureServerAccount
	globals.SecureEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.SecureEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername

	globals.SecureEndpoint.OnConnectionEnded(func(connection *nex.PRUDPConnection) {
		pid := uint32(connection.PID())
		user, ok := globals.ConnectedUsers.Get(pid)

		if !ok || user == nil {
			return
		}

		globals.ConnectedUsers.Delete(pid)
	})

	registerCommonSecureServerProtocols()
	registerSecureServerProtocols()

	globals.SecureEndpoint.IsSecureEndPoint = true
	globals.SecureServer.SetFragmentSize(962)
	globals.SecureServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(1, 1, 0))
	globals.SecureServer.SessionKeyLength = 16
	globals.SecureServer.AccessKey = "ridfebb9"
	globals.SecureServer.BindPRUDPEndPoint(globals.SecureEndpoint)
	globals.SecureServer.Listen(int(globals.Config.SecureServerPort))
}
