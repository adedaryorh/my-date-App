package constants

type (
	FriendRequestStatus string
)

const (
	FriendRequestStatusNew      FriendRequestStatus = "new"
	FriendRequestStatusAccepted FriendRequestStatus = "accepted"
	FriendRequestStatusDeclined FriendRequestStatus = "declined"
)
