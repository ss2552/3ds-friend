package nex

import (
	"github.com/PretendoNetwork/friends/globals"
	nex_account_management "github.com/PretendoNetwork/friends/nex/account-management"
	nex_friends_3ds "github.com/PretendoNetwork/friends/nex/friends-3ds"
	account_management "github.com/PretendoNetwork/nex-protocols-go/v2/account-management"
	friends_3ds "github.com/PretendoNetwork/nex-protocols-go/v2/friends-3ds"
)

func registerSecureServerProtocols() {
	accountManagementProtocol := account_management.NewProtocol()
	friends3DSProtocol := friends_3ds.NewProtocol()

	// * Account Management protocol handles
	accountManagementProtocol.NintendoCreateAccount = nex_account_management.NintendoCreateAccount

	// * Friends (3DS) protocol handles
	friends3DSProtocol.UpdateProfile = nex_friends_3ds.UpdateProfile
	friends3DSProtocol.UpdateMii = nex_friends_3ds.UpdateMii
	friends3DSProtocol.UpdatePreference = nex_friends_3ds.UpdatePreference
	friends3DSProtocol.SyncFriend = nex_friends_3ds.SyncFriend
	friends3DSProtocol.UpdatePresence = nex_friends_3ds.UpdatePresence
	friends3DSProtocol.UpdateFavoriteGameKey = nex_friends_3ds.UpdateFavoriteGameKey
	friends3DSProtocol.UpdateComment = nex_friends_3ds.UpdateComment
	friends3DSProtocol.AddFriendByPrincipalID = nex_friends_3ds.AddFriendByPrincipalID
	friends3DSProtocol.GetFriendPersistentInfo = nex_friends_3ds.GetFriendPersistentInfo
	friends3DSProtocol.GetFriendMii = nex_friends_3ds.GetFriendMii
	friends3DSProtocol.GetFriendPresence = nex_friends_3ds.GetFriendPresence
	friends3DSProtocol.RemoveFriendByPrincipalID = nex_friends_3ds.RemoveFriendByPrincipalID
	friends3DSProtocol.RemoveFriendByLocalFriendCode = nex_friends_3ds.RemoveFriendByLocalFriendCode
	friends3DSProtocol.GetPrincipalIDByLocalFriendCode = nex_friends_3ds.GetPrincipalIDByLocalFriendCode
	friends3DSProtocol.GetAllFriends = nex_friends_3ds.GetAllFriends

	globals.SecureEndpoint.RegisterServiceProtocol(accountManagementProtocol)
	globals.SecureEndpoint.RegisterServiceProtocol(friends3DSProtocol)
}
