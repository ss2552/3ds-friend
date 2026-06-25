package nex

import (
	"github.com/PretendoNetwork/friends/globals"
	nex_account_management "github.com/ss2552/3ds-friend/nex/account-management"
	account_management "github.com/PretendoNetwork/nex-protocols-go/v2/account-management"
)

func registerSecureServerProtocols() {
	accountManagementProtocol := account_management.NewProtocol()

	// * Account Management protocol handles
	accountManagementProtocol.NintendoCreateAccount = nex_account_management.NintendoCreateAccount

	globals.SecureEndpoint.RegisterServiceProtocol(accountManagementProtocol)
}
