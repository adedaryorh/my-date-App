package constants

type (
	ContentViewer string
)

const (
	ContentViewerEveryone         ContentViewer = "everyone"
	ContentViewerFriends          ContentViewer = "friends"
	ContentViewerFriendsOfFriends ContentViewer = "friendsOfFriends"
	ContentViewerSelectedFriends  ContentViewer = "selected-friends"
)
