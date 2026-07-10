package constants

type (
	CelebrationType      string
	CelebrationFrequency string
	CelebrationAudience  string
	CelebrationStatus    string
	CelebrationKind      string
	CelebrationOwner     string
)

const (
	CelebrationTypeGeneral       CelebrationType = "general"
	CelebrationTypePrivate       CelebrationType = "private"
	CelebrationTypeSpecial       CelebrationType = "special"
	CelebrationTypeSurpriseEvent CelebrationType = "surprise-event"
	CelebrationTypePremiumEvent  CelebrationType = "premium-event"

	CelebrationFrequencyOnetime CelebrationFrequency = "one-time"
	CelebrationFrequencyMonthly CelebrationFrequency = "monthly"
	CelebrationFrequencyAnnual  CelebrationFrequency = "annual"

	CelebrationKindBirthday           CelebrationKind = "birthday"
	CelebrationKindBabyShower         CelebrationKind = "baby-shower"
	CelebrationKindArrivalOfNewBaby   CelebrationKind = "arrival-of-new-baby"
	CelebrationKindEngaged            CelebrationKind = "engaged"
	CelebrationKindCeremony           CelebrationKind = "wedding-ceremony"
	CelebrationKindWeddingAnniversary CelebrationKind = "wedding-anniversary"
	CelebrationKindGraduation         CelebrationKind = "graduation"
	CelebrationKindFathersDay         CelebrationKind = "fathers-day"
	CelebrationKindMothersDay         CelebrationKind = "mothers-day"
	CelebrationKindRetirement         CelebrationKind = "retirement"
	CelebrationKindGetWellSoon        CelebrationKind = "get-well-soon"
	CelebrationKindTributes           CelebrationKind = "tributes"
	CelebrationKindCondolences        CelebrationKind = "condolences"

	CelebrationAudienceEveryOne        CelebrationAudience = "everyone"
	CelebrationAudienceFriendsOnly     CelebrationAudience = "friends-only"
	CelebrationAudienceSelectedFriends CelebrationAudience = "selected-friends"

	CelebrationStatusCreated CelebrationStatus = "created"
	CelebrationStatusPending CelebrationStatus = "pending"
	CelebrationStatusActive  CelebrationStatus = "active"
	CelebrationStatusExpired CelebrationStatus = "expired"

	CelebrationOwnerSelf   CelebrationOwner = "self"
	CelebrationOwnerFriend CelebrationOwner = "friend"

	MaxCelebrationFileSize int64 = 1024 * 1024 * 10 // 10 MB
)

func (c CelebrationOwner) IsValid() bool {
	switch c {
	case CelebrationOwnerSelf,
		CelebrationOwnerFriend:
		return true
	}
	return false
}

func (c CelebrationStatus) IsValid() bool {
	switch c {
	case CelebrationStatusCreated,
		CelebrationStatusActive,
		CelebrationStatusExpired:
		return true
	}
	return false
}

func (c CelebrationAudience) IsValid() bool {
	switch c {
	case CelebrationAudienceEveryOne,
		CelebrationAudienceSelectedFriends,
		CelebrationAudienceFriendsOnly:
		return true
	}
	return false
}

func (c CelebrationKind) IsValid() bool {
	switch c {
	case CelebrationKindBirthday,
		CelebrationKindArrivalOfNewBaby,
		CelebrationKindCeremony,
		CelebrationKindEngaged,
		CelebrationKindMothersDay,
		CelebrationKindFathersDay,
		CelebrationKindRetirement,
		CelebrationKindGetWellSoon,
		CelebrationKindTributes,
		CelebrationKindCondolences,
		CelebrationKindGraduation,
		CelebrationKindWeddingAnniversary,
		CelebrationKindBabyShower:
		return true
	}
	return false
}

func (c CelebrationFrequency) IsValid() bool {
	switch c {
	case CelebrationFrequencyOnetime,
		CelebrationFrequencyMonthly,
		CelebrationFrequencyAnnual:
		return true
	}
	return false
}

func (c CelebrationType) IsValid() bool {
	switch c {
	case CelebrationTypeGeneral,
		CelebrationTypePrivate,
		CelebrationTypeSpecial,
		CelebrationTypeSurpriseEvent,
		CelebrationTypePremiumEvent:
		return true
	}
	return false
}
