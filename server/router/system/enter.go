package system

type RouterGroup struct {
	JwtRouter
	SysRouter
	BaseRouter
	MenuRouter
	UserRouter
	CasbinRouter
	AuthorityRouter
	OperationRecordRouter
	AuthorityBtnRouter
}
