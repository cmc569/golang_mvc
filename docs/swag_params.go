package docs

type DocGetUsersInfo struct {
	LineToken string `json:"lineToken"`
}

type DocUsersSignUp struct {
	LineToken   string `json:"lineToken"`
	Phone       string `json:"phone"`
	CountryCode string `json:"countryCode"`
	VerifyCode  string `json:"verifyCode"`
	Name        string `json:"name"`
}

type DocSmsSendVerifySms struct {
	Phone       string `json:"phone"`
	CountryCode string `json:"countryCode"`
	VerifyType  int    `json:"verifyType"`
}

type DocMpgPay struct {
	LineToken  string `json:"lineToken"`
	MerchantID string `json:"merchantID"`
	Amount     int    `json:"amount"`
}

type DocGetContractList struct {
	LineToken string `json:"lineToken"`
}

type DocGetContractPeriod struct {
	LineToken string `json:"lineToken"`
	//IsCurrent 1=當前; 2=非當前
	IsCurrent int `json:"isCurrent"`
}

type DocGetProfileList struct {
	LineToken string `json:"lineToken"`
}

type DocGetProfileInfo struct {
	LineToken string `json:"lineToken"`
	ProfileId int64  `json:"profileId"`
}

type DocAddProfileInfo struct {
	LineToken string `json:"lineToken"`
	ThumbnailImage string `json:"thumbnailImage"`
	Btn1Label      string `json:"btn1Label"`
	Btn1Url        string `json:"btn1Url"`
	Btn2Label      string `json:"btn2Label"`
	Btn2Url        string `json:"btn2Url"`
}

type DocEditProfileInfo struct {
	LineToken      string `json:"lineToken"`
	ThumbnailImage string `json:"thumbnailImage"`
	Btn1Label      string `json:"btn1Label"`
	Btn1Url        string `json:"btn1Url"`
	Btn2Label      string `json:"btn2Label"`
	Btn2Url        string `json:"btn2Url"`
	// Status 1:開啟, 2:刪除
	Status         int    `json:"status"`
}
type DocDeleteProfileInfo struct {
	LineToken      string `json:"lineToken"`
	ThumbnailImage string `json:"thumbnailImage"`
	Btn1Label      string `json:"btn1Label"`
	Btn1Url        string `json:"btn1Url"`
	Btn2Label      string `json:"btn2Label"`
	Btn2Url        string `json:"btn2Url"`
	// Status 1:開啟, 2:刪除
	Status         int    `json:"status"`
}
