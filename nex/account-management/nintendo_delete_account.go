package nex_account_management

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"

	"github.com/PretendoNetwork/friends/globals"
	"github.com/PretendoNetwork/friends/utility"
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	account_management "github.com/PretendoNetwork/nex-protocols-go/v2/account-management"
)

func NintendoDeleteAccount(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureEndpoint.LibraryVersions(), globals.SecureEndpoint.ByteStreamSettings())

	// 送信するデータの書き込み *.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseBody)
	rmcResponse.ProtocolID = account_management.ProtocolID
	rmcResponse.MethodID = // account_management.MethodNintendoDeleteAccount
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
