package errors

type Error interface {
	Code() int
	String() string
}

type ServiceError int

func (e ServiceError) String() string {
	return errorCodeStrings[e]
}
func (e ServiceError) Code() int {
	return int(e)
}

const (
	NOOP ServiceError = iota
	ServerInternalErr
	DBNotFound
	StateConflict
	InvalidParam
	UUIDGenErr
	MissParam
	NoThisPerson
	FailToCreateOrder
	NoThisForm
	FailToPushFormID
	FailToCreateComment
	FailToCreateSolution
	FailToCreateZanSolution
	FailToCreateZanComment
	FailToDeleteZanSolution
	FailToDeleteZanComment
	FailToCreateCaina
	FailToSubscribe
	FailToGetInviteList
	FailToCreateInvite
	FailToGetFormIDNum
	FailToCreateMessage
	FailToCreateFeedback
	WrongPassword
	FailToCreateToken
	InvalidToken
	FailToVerifyJWT
	FailToReadStory
	AccountExist
	FailToUpdateOrder
	AtoiErr
	FailToDeleteOrder
	Openidnull
	NoWechat
	NoPhone
	NoRealName
	NoGraduateYear
	AlreadySetRealname
	ActionTypeNotHandle
	NoThisSchool
	NoSetSchool
	FailToCreateFindOrder
	FailToUpdateFindOrder
	FailToDeleteFindOrder
	FailToUpshelfFindOrder
	FailToUpshelfOrder
	FailToCreateMenu
	FailToUpdateMenu
	FailToDeleteMenu
	FailToDeleteComment
	FailToAddZan
	FailToDeleteZan
	FailToCheckAuthority
	FailToAddRocketToOrder
	FailToRocket
	FailToCreateQrSpread
	NoThisMenu
	FailToCreateMenuComment
	FailToFinishOrder
	FailToObtainCoupon
	PhoneWrong
	FailToCreateWeekOrder
	HaveObtained
	FailToInsertDailing
	PaytypeConflict
	NoFruit
	NoEgg
	FailToGetOpenid
	FailToReadAll
	FailToUnmarshal
	FailToCreateUser
	FailToFindUser
	FailToUpdateUser
	FailToFindMenu
	FailToCreateCoupon
	FailToUpdateCoupon
	FailToPayResHandle
	ThisDayHaveMenu
	ParseTimeError
	EatTimeError
	NoCoupon
	CouponStateError
	FailToDoneCoupon
	WeekdaycountError
	FailToDelayOrder
	FailToParseTime
	OneCouponOneTime
	UHaveCommentToThisMenu
	FailToPostUnifiedorder
	CantDelayToToday
	NoThisOrder
	OrderHaveDone
	FailToAddLocation
	FailToGetMyLocation
	FailToGetYuanqu
	FailToNewChineseDay
)

var errorCodeStrings = []string{

	"",
	"server_internal_err",
	"db_not_found",
	"state_conflict_err",
	"Invalid_param_err",
	"uuid_generate_error",
	"miss or misstype some req",
	"no this person",
	"fail to create order",
	"no this form",
	"FailToPushFormID",
	"FailToCreateComment",
	"FailToCreateSolution",
	"FailToCreateZanSolution",
	"FailToCreateZanComment",
	"FailToDeleteZanSolution",
	"FailToDeleteZanComment",
	"FailToCreateCaina",
	"FailToSubscribe",
	"FailToGetInviteList",
	"FailToCreateInvite",
	"FailToGetFormIDNum",
	"FailToCreateMessage",
	"FailToCreateFeedback",
	"WrongPassword",
	"FailToCreateToken",
	"InvalidToken",
	"FailToVerifyJWT",
	"FailToReadStory",
	"AccountExist",
	"FailToUpdateOrder",
	"AtoiErr",
	"FailToDeleteOrder",
	"Openidnull",
	"NoWechat",
	"NoPhone",
	"NoRealName",
	"NoGraduateYear",
	"AlreadySetRealname",
	"ActionTypeNotHandle",
	"NoThisSchool",
	"NoSetSchool",
	"FailToCreateFindOrder",
	"FailToUpdateFindOrder",
	"FailToDeleteFindOrder",
	"FailToUpshelfFindOrder",
	"FailToUpshelfOrder",
	"FailToCreateMenu",
	"FailToUpdateMenu",
	"FailToDeleteMenu",
	"FailToDeleteComment",
	"FailToAddZan",
	"FailToDeleteZan",
	"FailToCheckAuthority",
	"FailToAddRocketToOrder",
	"FailToRocket",
	"FailToCreateQrSpread",
	"NoThisMenu",
	"FailToCreateMenuComment",
	"FailToFinishOrder",
	"FailToObtainCoupon",
	"PhoneWrong",
	"FailToCreateWeekOrder",
	"HaveObtained",
	"FailToInsertDailing",
	"PaytypeConflict",
	"NoFruit",
	"NoEgg",
	"FailToGetOpenid",
	"FailToReadAll",
	"FailToUnmarshal",
	"FailToCreateUser",
	"FailToFindUser",
	"FailToUpdateUser",
	"FailToFindMenu",
	"FailToCreateCoupon",
	"FailToUpdateCoupon",
	"FailToPayResHandle",
	"ThisDayHaveMenu",
	"ParseTimeError",
	"EatTimeError",
	"NoCoupon",
	"CouponStateError",
	"FailToDoneCoupon",
	"WeekdaycountError",
	"FailToDelayOrder",
	"FailToParseTime",
	"OneCouponOneTime",
	"UHaveCommentToThisMenu",
	"FailToPostUnifiedorder",
	"不能改到当天",
	"NoThisOrder",
	"OrderHaveDone",
	"FailToAddLocation",
	"FailToGetMyLocation",
	"FailToGetYuanqu",
	"FailToNewChineseDay",
}
