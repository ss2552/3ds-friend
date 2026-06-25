package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/ss2552/3ds-friend/globals"
	"github.com/ss2552/3ds-friend/types"
	"github.com/PretendoNetwork/nex-go/v2"
	nex_types "github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/plogger-go"
)

func init() {
	globals.Logger = plogger.NewLogger()
	globals.ConnectedUsers = nex.NewMutexMap[uint32, *types.ConnectedUser]()

	var err error

	kerberosPassword := make([]byte, 0x10)
	_, err = rand.Read(kerberosPassword)
	if err != nil {
		globals.Logger.Error("Error generating Kerberos password")
		os.Exit(0)
	}

	globals.KerberosPassword = string(kerberosPassword)

	globals.AuthenticationServerAccount = nex.NewAccount(nex_types.NewPID(1), "Quazal Authentication", globals.KerberosPassword, false)
	globals.SecureServerAccount = nex.NewAccount(nex_types.NewPID(2), "Quazal Rendez-Vous", globals.KerberosPassword, false)
	globals.GuestAccount = nex.NewAccount(nex_types.NewPID(100), "guest", "MMQea3n!fsik", false)
	globals.AESKey, err = hex.DecodeString(globals.Config.AESKey)
	if err != nil {
		globals.Logger.Criticalf("Failed to decode AES key: %v", err)
		os.Exit(0)
	}
}
