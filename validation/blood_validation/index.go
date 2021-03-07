package blood_validation

type validation struct{

}

func NewValidation() *validation{
	return &validation{}
}

func (v *validation) GetMessage() map[string]string{
	return map[string]string{
		"Pulse.required": "請輸入脈搏資料",
		"Diastolic.required": "請輸入舒張壓",
		"Systolic.required": "請輸入收縮壓",
	}
}

type GetUpdateRule struct {
	Pulse int  `json:"pulse" form:"id" binding:"required"`
	Diastolic float32 `json:"diastolic" form:"diastolic" binding:"required"`
	Systolic float32 `json:"systolic" form:"systolic" binding:"required"`
}

type GetStoreRule struct {
	Pulse int `json:"pulse" form:"pulse" binding:"required"`
	Diastolic float32 `json:"diastolic" form:"diastolic" binding:"required"`
	Systolic float32 `json:"systolic" form:"systolic" binding:"required"`
}