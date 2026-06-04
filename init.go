package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/PretendoNetwork/friends/database"
	"github.com/PretendoNetwork/friends/globals"
	"github.com/PretendoNetwork/friends/types"
	"github.com/PretendoNetwork/nex-go/v2"
	nex_types "github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/plogger-go"
	"github.com/joho/godotenv"
)

func init() {
	globals.Logger = plogger.NewLogger()
	globals.ConnectedUsers = nex.NewMutexMap[uint32, *types.ConnectedUser]()

	var err error

	err = godotenv.Load()
	if err != nil {
		globals.Logger.Warningf("Error loading .env file: %s", err.Error())
	}

	globals.Config = globals.NewConfigParser(globals.Config).SetPrefix("PN_FRIENDS_CONFIG").ParseFromEnv()

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

	database.ConnectPostgres()
}
