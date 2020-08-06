package middleware

import "github.com/terryli1643/apidemo/domain/model"

type Profile struct {
	//ID
	ID int64
	//名称
	Username string
	//帐号
	Account string
	//玩家类型
	UserType model.UserType
	//帐号状态,  正常/锁定/冻结/停用
	State model.AccountState
	//电话号码
	Tel string
	//是否锁定
	Locked bool
	//是否登录
	Logined bool
	//角色
	Role string
	//当前站点
	CurrentSite string
	//ip地址
	IP string
	//具体地址
	IPAddr string
	//是否已经设置谷歌验证码
	SetGoogleToken bool
	//当期语言
	Lang string
	//当前域名
	Host string
	//用户层级
	UserLevel int64
	//是否代理
	IsAgent model.IsAgent
	//代理等级
	AgentLevel int
	//用户层级名称
	UserLevelName string
	// app标识码
	AppCode string
}
