package globals

import (
	"context"
	"strconv"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

func AccountDetailsByPID(pid types.PID) (*nex.Account, *nex.Error) {
	if pid.Equals(AuthenticationServerAccount.PID) {
		return AuthenticationServerAccount, nil
	}

	if pid.Equals(SecureServerAccount.PID) {
		return SecureServerAccount, nil
	}

	if pid.Equals(GuestAccount.PID) {
		return GuestAccount, nil
	}

	// pidから
	Password

	username := strconv.Itoa(int(pid))
	account := nex.NewAccount(pid, username, Password, false)

	return account, nil
}
