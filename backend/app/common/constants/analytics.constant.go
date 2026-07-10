package constants

type (
	AppEvent string
)

const (
	AppEventError                  AppEvent = "error"
	AppEventSignUpStarted          AppEvent = "sign-up-started"
	AppEventSignUpCompleted        AppEvent = "sign-up-completed"
	AppEventUserAccountCreated     AppEvent = "user-account-created"
	AppEventLoginStarted           AppEvent = "login-started"
	AppEventLoginSuccessful        AppEvent = "login-successful"
	AppEventLoginFailed            AppEvent = "login-failed"
	AppEventLogout                 AppEvent = "logout"
	AppEventProfileViewed          AppEvent = "profile-viewed"
	AppEventUserAccountDeactivated AppEvent = "user-account-deactivated"
	AppEventProfileUpdated         AppEvent = "profile-updated"
	AppEventCelebrationCreated     AppEvent = "celebration-created"
	AppEventCelebrationEdited      AppEvent = "celebration-edited"
	AppEventCelebrationDeleted     AppEvent = "celebration-deleted"
)
