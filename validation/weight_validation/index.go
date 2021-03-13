package weight_validation

type validation struct {

}

func NewValidation() *validation {
	return &validation{}
}

func (v *validation) GetMessage() map[string]string {
	return map[string]string {
		"Number.required": "請輸入體重資料",
	}
}

type GetUpdateRule struct {
	Number int  `json:"number" form:"id" binding:"required"`
}
