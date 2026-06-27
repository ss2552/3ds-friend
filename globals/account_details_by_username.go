package globals

import (
	"fmt"
	"strconv"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

func AccountDetailsByUsername(username string) (*nex.Account, *nex.Error) {
	if username == AuthenticationEndpoint.ServerAccount.Username {
		return AuthenticationEndpoint.ServerAccount, nil
	}

	if username == SecureEndpoint.ServerAccount.Username {
		return SecureEndpoint.ServerAccount, nil
	}

	if username == GuestAccount.Username {
		return GuestAccount, nil
	}

	// TODO - This is fine for our needs, but not for servers which use non-PID usernames?
	pid, err := strconv.Atoi(username)
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
		return nil, nex.NewError(nex.ResultCodes.RendezVous.InvalidUsername, "Invalid username")
	}

	// * Trying to use AccountDetailsByPID here led to weird nil checks?
	// * Would always return an error even when it shouldn't.
	// TODO - Look into this more

	password := "nupHf1bMOjIs4FoX"

	account := nex.NewAccount(types.NewPID(uint64(pid)), username, password, false)

	return account, nil
}
