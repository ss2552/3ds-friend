package globals

import (
	"github.com/ss2552/3ds-friend/types"
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/plogger-go"
)

var Logger *plogger.Logger
var AuthenticationServerAccount *nex.Account
var SecureServerAccount *nex.Account
var GuestAccount *nex.Account
var KerberosPassword = "password" // * Default password
var AuthenticationServer *nex.PRUDPServer
var AuthenticationEndpoint *nex.PRUDPEndPoint
var SecureServer *nex.PRUDPServer
var SecureEndpoint *nex.PRUDPEndPoint
var ConnectedUsers *nex.MutexMap[uint32, *types.ConnectedUser]
var AESKey []byte
