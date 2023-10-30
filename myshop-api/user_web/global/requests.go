package global

type PassWordLoginReq struct {
	Mobile       string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord     string `form:"password" json:"password" binding:"required,min=3,max=20"`
	VerifyCode   string `form:"verify_code" json:"verify_code" binding:"required,gte=6"`
	VerifyCodeId string `form:"verify_code_id" json:"verify_code_id" binding:"required"`
}

type SendSmsReq struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	Type uint `form:"type" json:"type" binding:"required,oneof=1 2"` //1. 1:注册发送短信验证码;2:动态验证码登录发送验证码
}

type RegisterReq struct{
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Code string `form:"code" json:"code" binding:"required,gte=6"`
}