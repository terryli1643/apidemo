package model

import (
	"github.com/gin-gonic/gin"
	"github.com/terryli1643/apidemo/libs/datasource"
	log "github.com/terryli1643/apidemo/libs/logger"
)

// swagger:model
type (
	Device                      int
	UserType                    int
	AccountState                int
	MerchantState               int
	SettlementRecordState       int
	RecordType                  int
	BankAccountState            int
	AmountFormat                int
	TradingMode                 int
	FundchangeType              int
	FundSourceType              int
	RemarkType                  int
	PayerState                  int
	PayerType                   int
	TerminalType                int
	Location                    int
	OrderType                   int
	Switch                      int
	PayerProfitsState           int
	ChannelState                int
	OperateAccountType          int
	RateState                   int
	ChannelType                 int
	Channel                     int
	RateChannel                 int
	OrderState                  int
	PaymentState                int
	AccountType                 int
	OperateAccountOrderType     int
	FinancialSubjectType        int
	OperateAccountOrderState    int
	FillType                    int
	FillState                   int
	TransferType                int
	Online                      int
	ParamType                   int
	ParamState                  int
	PushState                   int
	BuildInType                 int
	MatchState                  int
	OperationType               int
	SourceType                  int
	BuildInMerchant             int
	Domain                      int
	FillchangeType              int
	FillTypeForMerchant         int
	IsAgent                     int
	CompleteState               int
	TypeJpush                   int
	IsFill                      int
	ErrType                     int
	QueryTimeType               int
	QRType                      int
	IsPointType                 int
	PointUporDown               int
	FixedCode                   int
	AgentRecalculateProfitState int
	OrderSettled                int
	YesOrNo                     int
	PayerComplainState          int
)

const (
	_                                     AgentRecalculateProfitState = iota
	AgentRecalculateProfitStateSettle                                 //未发放
	AgentRecalculateProfitStateAllSettled                             //全部已发放
)
const (
	_ OrderSettled = iota
	Unsettle
	Settled
)

const (
	RoleAnonymous = "ROLE_ANONYMOUS" // 陌生人
	RoleAdmin     = "ROLE_ADMIN"     // 管理员
)

type IEnum interface {
	Val() int
}

const (
	IDSpaceUser          = "user"
	IDSpaceOrder         = "order"
	IDSpaceRecommendCode = "recommendcode"
	IDSpaceReq           = "req"
	IDSPaceRole          = "role"
)

const (
	USD = "USD" //美元
	EUR = "EUR" //欧元
	JPY = "JPY" //日元
	GBP = "GBP" //英镑
	CHF = "CHF" //瑞士法郎
	AUD = "AUD" //澳大利亚元
	NZD = "NZD" //新西兰元
	CAD = "CAD" //加拿大元
	CNY = "CNY" //人民币
	RUB = "RUB" //卢比
	HKD = "HKD" //港币
	IDR = "IDR" //印尼盾
	KRW = "KRW" //韩国元
	SAR = "SAR" //亚尔
	THB = "THB" //泰铢
)

const (
	_       Device = iota
	PC             //来源终端: 电脑       1
	IOS            //来源终端: 苹果手机 	2
	Android        //来源终端: 安卓手机 	3
)

const (
	_ YesOrNo = iota
	Yes
	No
)

const (
	_                    OperationType = iota
	OperationTypeQuery                 //查询
	OperationTypeCreate                //添加
	OperationTypeUpdate                //编辑
	OperationTypeDelete                //删除
	OperationTypeUnknown               //未知
)

const (
	_           SourceType = iota
	CLIENT                 //来源终端: 后台系统      1
	APP                    //来源终端: APP          2
	MerchantAPI            //来源终端: 商户API       3
	ExeAPI                 //来源终端: ExeAPI       4

)

const (
	_              TypeJpush = iota
	GoodsNotice              //来源终端: 后台系统      1
	ApprovalNotice           //来源终端: APP          2
	OfflineNotice            //来源终端: 商户API       3
)

const (
	_         IsFill = iota
	IsFillNo         //不是填报     1
	IsFillYes        //是填报       2
)

const (
	_            FixedCode = iota
	FixedCodeNo            //不是固码     1
	FixedCodeYes           //固定二维码      2
)

const (
	_    PointUporDown = iota
	Up                 //
	Down               //
)

const (
	_              IsPointType = iota
	IsPointTypeNo              //不是小数点形式     1
	IsPointTypeYes             //是小数点形式       2
)

const (
	_                  BuildInMerchant = iota
	BuildInMerchantYes                 // 内置商户
	BuildInMerchantNo                  // 真实商户
)

const (
	_              AccountState = iota
	AccountEnable               // 正常		1
	AccountDisable              // 帐号冻结	2
)

const (
	_                   CompleteState = iota
	Complete                          //绑定	1
	UnComplete                        //未完成	2
	CompleteStateDetete               //删除	3
)

const (
	_         Domain = iota
	DomainHt         // 后台		1
	DomainAPI        // api	2
	DomainAPP        // APP	3

)

const (
	_                                OperateAccountOrderType = iota
	OperateAccountOrderTypeConverted                         // 内转		1
	OperateAccountOrderTypeOutSide                           //填报		2
)

const (
	_               TransferType = iota
	TransferTypeIn               // 转入		1
	TransferTypeOut              //转出		2
)

const (
	_         Online = iota
	POnline          // 在线		1
	POffline         // 离线		2
	PDropline        // 掉线		3
)

const (
	_        ParamType = iota
	BizParam           //业务参数
	SysParam           //系统参数
)

const (
	_ ParamState = iota
	ParamEnable
	ParamDisable
)

const (
	_         Location = iota
	CHINA              //中国	1
	TAIWAN             //台湾	2
	XIANGGANG          //香港    3
	US                 //美国	4
	VIETNAM            //越南	5
	THAILAND           //泰国    6
	KOREA              //韩国   7
)

const (
	_                       BankAccountState = iota
	BankAccountStateCreate                   //待审核 1
	BankAccountStateEnable                   // 审核通过		2
	BankAccountStateDisable                  // 审核不通过	3
)

const (
	_                                      FinancialSubjectType = iota
	FinancialSubjectTypeForeign                                 // 外汇支出		1
	FinancialSubjectTypeAmount                                  // 费用支出		2
	FinancialSubjectTypeFinancialInjection                      // 财务注入		3
	FinancialSubjectTypeAdditionalIncome                        // 额外收入		4
)
const (
	_             QueryTimeType = iota
	Today                       // 今天		1
	Yesterday                   // 昨天		2
	LastSevenDays               // 最近七天		3
	LastOneMonth                // 最近一月		4
)

const (
	_                    OperateAccountType = iota
	OperateAccountTypeSh                    //商户卡	1
	OperateAccountTypeZj                    //资金卡	2
	OperateAccountTypeSf                    //收付卡	3

)

const (
	_                  UserType = iota
	UserTypeMerchent            //商户	1
	UserTypePayer               //付客	2
	UserTypeOperateAcc          //财务账户	3
	UserTypeAdmin               //管理员	4
	UserTypeAgent               //付客代理 5
)

const (
	_        IsAgent = iota
	IsAgentY         //代理1
	IsAgentN         //非代理2
)

const (
	_                RateState = iota
	RateStateEnable            //启用	1
	RateStateDisable           //禁用	2
)
const (
	_                 ChannelType = iota
	ChannelTypeWeChat             //微信
	ChannelTypeAlipay             //支付宝
	ChannelTypeBank               //银行
	ChannelTypeJy                 //金燕E商
	ChannelTypeYSF                //云闪付
)

const (
	AlipayH5   = 1    //支付宝H5
	Alipay     = 2    //支付宝
	AlipayKs   = 3    //支付宝快速
	AlipayCard = 4    //支付宝卡收
	UNIONPAY   = 5    //云闪付
	JY         = 6    //金燕e商
	WeiXin     = 7    //微信
	STUNIONPAY = 8    //st云闪付
	CMBC       = 1001 //中国民生银行
	ICBC       = 1002 //中国工商银行
	BOC        = 1003 //中国银行
	BOCOM      = 1004 //交通银行
	PINGAN     = 1005 //中国平安银行
	CMB        = 1006 //招商银行
	ABC        = 1007 //中国农业银行
	CCB        = 1008 //中国建设银行
	PSBC       = 1009 //中国邮政储蓄银行
	CEBB       = 1010 //中国光大银行
	CIB        = 1011 //兴业银行
	SPDB       = 1012 //浦发银行
	CGB        = 1013 //广发银行
	CITIC      = 1014 //中信银行
	HXB        = 1015 //华夏银行
	BCCB       = 1016 //北京银行
	BOSC       = 1017 //上海银行
	GZCB       = 1018 //广州银行
	CZB        = 1019 //网商银行

	BANK = 2000 //银行
)

var channelNameMap = map[int]string{
	AlipayH5:   "支付宝H5",
	Alipay:     "支付宝",
	AlipayKs:   "支付宝快速",
	AlipayCard: "支付宝转卡",
	UNIONPAY:   "云闪付",
	JY:         "金燕e商",
	WeiXin:     "微信",
	STUNIONPAY: "ST云闪付",
	CMBC:       "民生银行",
	ICBC:       "工商银行",
	BOC:        "中国银行",
	BOCOM:      "交通银行",
	PINGAN:     "平安银行",
	CMB:        "招商银行",
	ABC:        "农业银行",
	CCB:        "建设银行",
	PSBC:       "邮政储蓄银行",
	CEBB:       "光大银行",
	CIB:        "兴业银行",
	SPDB:       "浦发银行",
	CGB:        "广发银行",
	CITIC:      "中信银行",
	HXB:        "华夏银行",
	BCCB:       "北京银行",
	BOSC:       "上海银行",
	GZCB:       "广州银行",
	CZB:        "网商银行",
	BANK:       "银行",
}

var channelValueMap = map[string]int{
	"AlipayH5":   AlipayH5,
	"Alipay":     Alipay,
	"AlipayKs":   AlipayKs,
	"AlipayCard": AlipayCard,
	"CMBC":       CMBC,
	"ICBC":       ICBC,
	"BOC":        BOC,
	"BOCOM":      BOCOM,
	"PINGAN":     PINGAN,
	"CMB":        CMB,
	"ABC":        ABC,
	"CCB":        CCB,
	"PSBC":       PSBC,
	"CEBB":       CEBB,
	"CIB":        CIB,
	"SPDB":       SPDB,
	"CGB":        CGB,
	"CITIC":      CITIC,
	"HXB":        HXB,
	"BCCB":       BCCB,
	"BOSC":       BOSC,
	"GZCB":       GZCB,
	"CZB":        CZB,
	"UNIONPAY":   UNIONPAY,
	"STUNIONPAY": STUNIONPAY,
	"JY":         JY,
	"WECHAT":     WeiXin,
	"BANK":       BANK,
}

var channelCodeMap = map[int]string{
	AlipayH5:   "AlipayH5",
	Alipay:     "Alipay",
	AlipayKs:   "AlipayKs",
	AlipayCard: "AlipayCard",
	CMBC:       "CMBC",
	ICBC:       "ICBC",
	BOC:        "BOC",
	BOCOM:      "BOCOM",
	PINGAN:     "PINGAN",
	CMB:        "CMB",
	ABC:        "ABC",
	CCB:        "CCB",
	PSBC:       "PSBC",
	CEBB:       "CEBB",
	CIB:        "CIB",
	SPDB:       "SPDB",
	CGB:        "CGB",
	CITIC:      "CITIC",
	HXB:        "HXB",
	BCCB:       "BCCB",
	BOSC:       "BOSC",
	GZCB:       "GZCB",
	CZB:        "CZB",
	UNIONPAY:   "UNIONPAY",
	STUNIONPAY: "STUNIONPAY",
	JY:         "JY",
	WeiXin:     "WECHAT",
	BANK:       "BANK",
}

func (channel Channel) ToChannelName() string {
	ch := channelNameMap[int(channel)]
	return ch
}

func (channel Channel) ToChannelCode() string {
	ch := channelCodeMap[int(channel)]
	return ch
}

func ToChannel(channelCode string) Channel {
	ch := channelValueMap[channelCode]
	return Channel(ch)
}

func (rateChannel RateChannel) ToChannelName() string {
	ch := channelNameMap[int(rateChannel)]
	return ch
}

func PkgToBankCode(bankname string) string {
	channelMap := map[string]string{}
	channelMap["com.eg.android.AlipayGphone"] = "Alipay"
	channelMap["com.tencent.mm"] = "WeiXin"
	channelMap["com.chinamworld.bocmbci"] = "BOC"
	channelMap["cmb.pb"] = "CMB"
	channelMap["com.mybank.android.phone"] = "CZB"
	channelMap["com.chinamworld.main"] = "CCB"
	channelMap["com.icbc"] = "ICBC"
	channelMap["com.android.bankabc"] = "ABC"
	channelMap["com.bankcomm.Bankcomm"] = "BOCOM"
	channelMap["com.unionpay"] = "UNIONPAY"
	channelMap["com.hnnx.sh.mbank"] = "JY"
	channelMap["com.yitong.mbank.psbc"] = "PSBC"
	channelMap["cn.com.spdb.mobilebank.per"] = "SPDB"

	ch := channelMap[bankname]
	return ch
}

func (channel Channel) ToChannelType() ChannelType {
	if channel == AlipayH5 {
		return ChannelTypeAlipay
	}
	if channel == Alipay {
		return ChannelTypeAlipay
	}
	if channel == AlipayKs {
		return ChannelTypeAlipay
	}
	if channel == JY {
		return ChannelTypeJy
	}
	if channel == AlipayCard {
		return ChannelTypeBank
	}
	if channel == UNIONPAY {
		return ChannelTypeYSF
	}

	if channel == STUNIONPAY {
		return ChannelTypeYSF
	}

	if channel == WeiXin {
		return ChannelTypeWeChat
	}
	return ChannelTypeBank
}

func (channel Channel) ToRateChannel() RateChannel {
	if channel > 1000 {
		return BANK
	}
	return RateChannel(channel)
}

func (channel RateChannel) ToChannelType() ChannelType {
	if channel == AlipayH5 {
		return ChannelTypeAlipay
	}
	if channel == WeiXin {
		return ChannelTypeWeChat
	}
	if channel == Alipay {
		return ChannelTypeAlipay
	}
	if channel == AlipayKs {
		return ChannelTypeAlipay
	}
	if channel == AlipayCard {
		return ChannelTypeBank
	}
	if channel == JY {
		return ChannelTypeJy
	}
	if channel == UNIONPAY {
		return ChannelTypeYSF
	}
	if channel == STUNIONPAY {
		return ChannelTypeYSF
	}
	if channel == WeiXin {
		return ChannelTypeWeChat
	}
	return ChannelTypeBank
}

const (
	_                 PayerState = iota
	PayerStateEnable             //启用	1
	PayerStateDisable            //禁用	2
)

const (
	_                   ChannelState = iota
	ChannelStateEnable               //启用	1
	ChannelStateDisable              //禁用	2
)
const (
	_                        PayerProfitsState = iota
	PayerProfitsStateEnable                    //启用	1
	PayerProfitsStateDisable                   //禁用	2
)

const (
	_               PayerType = iota
	PayerTypeMember           //会员	1
	PayerTypeVIP              //VIP	2
)

const (
	_               TerminalType = iota
	TerminalTypeRrf              //rrfApp	1
	TerminalTypeDg               //代购app	2
)

const (
	_                               OperateAccountOrderState = iota
	OperateAccountOrderStateCreate                           //新建
	OperateAccountOrderStateSuccess                          //审核通过	1
	OperateAccountOrderStateFail                             //审核不通过	2
)

const (
	_                                 SettlementRecordState = iota
	SettlementRecordStatePending                            //未处理 	    1
	SettlementRecordStateProcessing                         //处理中 		2
	SettlementRecordStatePreCompleted                       //预完成 		3
	SettlementRecordStatePreClose                           //预关闭 		4
	SettlementRecordStateComplete                           //已完成	    5
	SettlementRecordStateClose                              //已关闭      6
)

const (
	_                    RecordType = iota
	RecordTypeInjection             //注资  	1
	RecordTypeSettlement            //结算		2
)

const (
	_                 TradingMode = iota
	TradingModeByHand             //手动  	1
	TradingModeByAPI              //API		2
)

const (
	_                 FillchangeType = iota
	FillIncrease                     //增加余额
	FillDecrease                     //减少余额
	FillFrezeIncrease                //增加冻结
	FillFrezeDecrease                //减少冻结

)

const (
	_               FundchangeType = iota
	BalanceIncrease                //账户余额增加 1
	BalanceDecrease                //账户余额减少 2
	FrozenIncrease                 //账户冻结余额增加 3
	FrozenDecrease                 //账户冻结余额减少 4
	Frozen                         //冻结 5
	UnFrozen                       //解冻 6
)

const (
	_                        FundSourceType = iota
	FundSourceTypeDeposit                   //订单充值 1 (已废弃)
	FundSourceTypeWithdraw                  //订单提现 2 (已废弃)
	FundSourceTypeExeRecord                 //异常记录 3 (已废弃)
	FundSourceTypeSettlement                //结算/注资 4 (已废弃)
	FundSourceTypeFill                      //付客填报 5 (已废弃)
	FundSourceTypeManual                    //人工填报 6 (已废弃)
	FundSourceTypeAgent                     //代理收益 7 (已废弃)

	FundSourceTypeDepositMBI                    //商户充值 8
	FundSourceTypeDepositFeeMBD                 //商户充值手续费 9
	FundSourceTypeMerchantAgentDepositProfitMBI //商户代理充值收益 10
	FundSourceTypePayerAssignedOrderPFR         //付客接单冻结 11
	FundSourceTypePayerAssignedOrderPUF         //付客接单解冻 12
	FundSourceTypePayerReceivedMoneySuccessPBD  //付客收款成功 13
	FundSourceTypeDepositPayerProfitPBI         //付客接单收益 14
	FundSourceTypeDepositPayerAgentProfitPBI    //付客代理收益 15
	FundSourceTypeMatchExeptionFailedPFR        //流水未匹配 16
	FundSourceTypeManualExeptionHandlePUF       //人工异常处理解冻 17
	FundSourceTypeOperateAccountOBD             //扣减内部财务账户金额 18
	FundSourceTypeOperateAccountOBI             //增加内部财务账户金额 19

	FundSourceTypeCreateMerchantWithdrawMFR          //商户提现 20
	FundSourceTypeCreateMerchantWithdrawMUF          //商户提现冻结 21
	FundSourceTypeCreateMerchantWithdrawSuccessMBD   //商户提现扣款 22
	FundSourceTypeCreateMerchantWithdrawFeeMBD       //商户提现手续费 23
	FundSourceTypeMerchantAgentWithdrawProfitMBI     //商户提现代理收益 24
	FundSourceTypePayerSeckillingOrderPBI            //付客秒宝贝 25
	FundSourceTypePayerSeckillingOrderProfitPBI      //付客秒宝贝收益 26
	FundSourceTypePayerSeckillingOrderAgentProfitPBI //付客代理秒宝贝收益 27

	FundSourceTypeCreateMerchantSettlementMFR        //商户结算冻结 28
	FundSourceTypeCreateMerchantSettlementMUF        //商户结算解冻 29
	FundSourceTypeCreateMerchantSettlementSuccessMBD //商户结算扣款 30
	FundSourceTypeCreateMerchantSettlementFeeMFR     //商户结算手续费冻结 31
	FundSourceTypeCreateMerchantSettlementFeeMUF     //商户结算手续费解冻 32
	FundSourceTypeCreateMerchantSettlementFeeMBD     //商户结算手续费 33
	FundSourceTypeCreateMerchantInjectionMBI         //商户注资加款 34

	FundSourceTypeManualAdjustMerchantBanlanceMBI //人工商户余额账户增加 35
	FundSourceTypeManualAdjustMerchantBanlanceMBD //人工商户余额账户减少 36
	FundSourceTypeManualAdjustMerchantFreezeMFI   //人工商户冻结账户增加 37
	FundSourceTypeManualAdjustMerchantFreezeMFD   //人工商户冻结账户减少 38

	FundSourceTypeManualAdjustPayerBanlancePBI //付客填报调整增加 39
	FundSourceTypeManualAdjustPayerBanlancePBD //付客填报调整减少 40
	FundSourceTypeManualAdjustPayerFreezePFI   //付客填报调整冻结增加 41
	FundSourceTypeManualAdjustPayerFreezePFD   //付客填报调整冻结减少 42

	FundSourceTypeManualActivityPayerBanlancePBI //付客填报活动增加 43
	FundSourceTypeManualActivityPayerBanlancePBD //付客填报活动减少 44
	FundSourceTypeManualActivityPayerBanlancePFI //付客填报活动冻结增加 45
	FundSourceTypeManualActivityPayerBanlancePFD //付客填报活动冻结减少 46
	FundSourceTypeTransferOutPayerBanlancePBD    //付客转出 47
	FundSourceTypeTransferInPayerBanlancePBI     //付客转入 48

	FundSourceTypePayerDepositPBI         //付客充值 49
	FundSourceTypeCreatePayerWithdrawPFR  //付客提现冻结 50
	FundSourceTypePayerWithdrawPUF        //付客提现解冻 51
	FundSourceTypePayerWithdrawSuccessPBD //付客提现扣款 52
	FundSourceTypePayerWithdrawFeePBD     //付客提现手续费 53

	FundSourceTypeOperateAccountInOBI  //财务账号转入 54
	FundSourceTypeOperateAccountOutOFR //财务账号转出冻结 55
	FundSourceTypeOperateAccountOutOUF //财务账号转出解冻 56
	FundSourceTypeOperateAccountOutOBD //财务账号转出 57

)

const (
	_                      RemarkType = iota
	RemarkTypePayer                   //付客 1
	RemarkTypeOrder                   //订单 2
	RemarkTypeOperateAcc              //运营账户 3
	RemarkTypeFill                    //填报备注 4
	RemarkTypeExceptionRec            //异常记录 5
	RemarkTypeFillImg                 //填报备注 6

)

const (
	_                 OrderType = iota
	OrderTypeDeposit            //充值	收入 1
	OrderTypeWithdraw           //提现	支出 2
)

const (
	_                           OrderState = iota
	OrderStatesFreeze                      //冻结 1
	OrderStatesNew                         //未处理 2
	OrderStatesProcessing                  //处理中3
	OrderStatesPreComplete                 //预完成4
	OrderStatesPreClose                    //预关闭5
	OrderStatesCompleted                   //完成6
	OrderStatesPartialCompleted            //部分完成7
	OrderStatesClose                       //关闭8
	OrderStatesCloseByPayer                //付客关闭9
)

const (
	_                         PaymentState = iota
	PaymentStatesNew                       //初始状态 , 冻结三方帐号资金
	PaymentStatesSettleFailed              //结算失败
	PaymentStatesSettled                   //已结算, 三方扣款
	PaymentStatesAbandoned                 //丢弃的, 解冻三方帐号资金
)

const (
	_              AccountType = iota
	TypeThirdparty             //三方
	TypeBank                   //银行卡
	TypeMerchant               //商户
	TypeQrCode                 //二维码
)

const (
	_   Switch = iota
	On         //开
	Off        //关
)

const (
	_        FillType = iota
	Activity          //活动
	Adjust            //调整
	Transfer          //转账
)

const (
	_                   FillState = iota
	FillStatePending              //未处理 	    1
	FillStateProcessing           //处理中 		2
	FillStateComplete             //已完成	    3
	FillStateClose                //已关闭      4
)

const (
	_             FillTypeForMerchant = iota
	FillPayment                       //付款
	FillGathering                     //收款
)

const (
	_           PushState = iota
	Pushing               //推送中 1
	PushSuccess           //推送成功2
	PushFailed            //推送失败3
)

const (
	_       BuildInType = iota
	BuildIn             //内置类型 1
	Real                //真实类型 2
)

const (
	_              MatchState = iota
	MatchedNew                //新建 1
	AutoMatched               //自动匹配 2
	NotMatch                  //未匹配 3
	MatchedClosed             //已关闭 4
	ManualMatched             //人工匹配 5
	ExcepitonMatch            //异常 6
	SmsMatch                  //需要短信匹配 7

)

const (
	_              ErrType = iota
	MerchantHuabei         //商户版花呗
	WrongProofNo           //错误版流水号
	OutofMemory            //内存不足
)

const (
	Unlimited   AmountFormat = iota // 无限制
	Integer                         // 整数
	TwoDicemal                      // 两位小数
	FixedAmount                     // 固定金额
)

const (
	_                         PayerComplainState = iota
	PayerComplainStateInit                       //未处理
	PayerComplainStateHandled                    //已处理
	PayerComplainStateClosed                     //已关闭
)

const (
	QRRawType      = "raw"      //原生二维码
	QRPlatformType = "platform" //新二维码
)

const (
	RequestResolveFailed        gin.ErrorType = 401   //请求解析失败
	MerchantResolveFailed       gin.ErrorType = 402   //商户解析失败
	SginVerifyFailed            gin.ErrorType = 403   //请求校验失败
	IPVerifyFailed              gin.ErrorType = 404   //IP校验失败
	QueryChannelListFailed      gin.ErrorType = 9000  //查询通道失败
	QueryOrderFailed            gin.ErrorType = 9001  //查询订单失败
	WithdrawVerifyFailed        gin.ErrorType = 9002  //提现请求失败
	ThirdPartyFailed            gin.ErrorType = 10001 //三方支付请求失败
	ChannelIsNotSupported       gin.ErrorType = 10002 //请求的通道不被支持
	OrderVerifyFailed           gin.ErrorType = 10003 //订单验证失败
	MerchantInsufficientBalance gin.ErrorType = 10003 //商户余额不足
	OrderFailedApprove          gin.ErrorType = 10004 //提现订单被拒
	OrderCreateFailed           gin.ErrorType = 10005 //创建订单失败
	OrderRefuse                 gin.ErrorType = 10006 //拒绝支付
)

func InitialModels() {
	log.Info("Register models")
	t := []interface{}{
		new(Admin),
	}
	datasource.RegisterModels(t...)
}
