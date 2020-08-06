package service

import (
	"errors"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/terryli1643/apidemo/libs/datasource"
	"gorm.io/gorm"
)

var (
	l sync.Mutex
)

const (
	TIME_FORMAT_WITH_MS         = "2006-01-02 15:04:05.000"
	TIME_FORMAT                 = "2006-01-02 15:04:05"
	TIME_FORMAT_WO_SEC_COMPACT  = "200601021504"
	TIME_FORMAT_COMPACT         = "20060102150405"
	TIME_FORMAT_WITH_MS_COMPACT = "20060102150405000"
	DATE_FORMAT                 = "2006-01-02"
	DATE_FORMAT_COMPACT         = "20060102"
	MONTH_FORMAT                = "2006-01"
)

const (
	HSET         = "HSET"
	HINCRBYFLOAT = "HINCRBYFLOAT"
	HINCRBY      = "HINCRBY"
	SADD         = "SADD"
	KEYS         = "KEYS"
	HGETALL      = "HGETALL"
	SISMEMBER    = "SISMEMBER"
	HMSET        = "HMSET"
	HMGET        = "HMGET"
)

const (
	DATASOURCE_SUFFIX_MOCK = "mock"
)

//安全密码错误次数记录
var wrongSecurePwdCountHolder sync.Map

func getDB() *gorm.DB {
	db := datasource.GetDB()
	return db
}

func closeTx(tx *gorm.DB, err *error) {
	r := recover()
	if r != nil {
		tx.Rollback()
		log.Error(r)
		*err = errors.New("panic")
		return
	}

	if *err != nil {
		tx.Rollback()
		log.Errorf("%+v", *err)
		return
	}
	tx.Commit()
}

type TanslateModel struct {
	Number int
	Value  string
}

// func TranslateTypeToName(T i18n.TranslateFunc) (result map[string]string) {
// 	tanslateRecordTypeMap := make(map[string]string)
// 	tanslateRecordTypeMap["UserType"] = T("key_t_usertype")
// 	tanslateRecordTypeMap["AccountState"] = T("key_t_accountstate")
// 	tanslateRecordTypeMap["MerchantState"] = T("key_t_merchantstate")
// 	tanslateRecordTypeMap["SettlementRecordState"] = T("key_t_settlementrecordstate")
// 	tanslateRecordTypeMap["RecordType"] = T("key_t_recordtype")
// 	tanslateRecordTypeMap["BankAccountState"] = T("key_t_bankaccountstate")
// 	tanslateRecordTypeMap["TradingMode"] = T("key_t_tradingmode")
// 	tanslateRecordTypeMap["FundchangeType"] = T("key_t_fundchangetype")
// 	tanslateRecordTypeMap["FundSourceType"] = T("key_t_fundsourcetype")
// 	tanslateRecordTypeMap["RemarkType"] = T("key_t_remarktype")
// 	tanslateRecordTypeMap["PayerState"] = T("key_t_payerstate")
// 	tanslateRecordTypeMap["PayerType"] = T("key_t_payertype")
// 	tanslateRecordTypeMap["TerminalType"] = T("key_t_terminaltype")
// 	tanslateRecordTypeMap["Location"] = T("key_t_location")
// 	tanslateRecordTypeMap["OrderType"] = T("key_t_ordertype")
// 	tanslateRecordTypeMap["Switch"] = T("key_t_switch")
// 	tanslateRecordTypeMap["PayerProfitsState"] = T("key_t_payerprofitsstate")
// 	tanslateRecordTypeMap["OperateAccountType"] = T("key_t_operateaccounttype")
// 	tanslateRecordTypeMap["RateState"] = T("key_t_ratestate")
// 	tanslateRecordTypeMap["ChannelType"] = T("key_t_channeltype")
// 	tanslateRecordTypeMap["OrderState"] = T("key_t_orderstate")
// 	tanslateRecordTypeMap["PaymentState"] = T("key_t_paymentstate")
// 	tanslateRecordTypeMap["AccountType"] = T("key_t_accounttype")
// 	tanslateRecordTypeMap["OperateAccountOrderType"] = T("key_t_operateaccountordertype")
// 	tanslateRecordTypeMap["FinancialSubjectType"] = T("key_t_financialsubjecttype")
// 	tanslateRecordTypeMap["OperateAccountOrderState"] = T("key_t_operateaccountorderstate")
// 	tanslateRecordTypeMap["FillType"] = T("key_t_filltype")
// 	tanslateRecordTypeMap["FillState"] = T("key_t_fillstate")
// 	tanslateRecordTypeMap["FillUserType"] = T("key_t_fillusertype")
// 	tanslateRecordTypeMap["FillTypeForMerchant"] = T("key_t_filltypeformerchant")
// 	tanslateRecordTypeMap["TransferType"] = T("key_t_transfertype")
// 	tanslateRecordTypeMap["NoticeState"] = T("key_t_noticestate")
// 	tanslateRecordTypeMap["Online"] = T("key_t_online")
// 	tanslateRecordTypeMap["ParamType"] = T("key_t_paramtype")
// 	tanslateRecordTypeMap["ParamState"] = T("key_t_paramstate")
// 	tanslateRecordTypeMap["PushState"] = T("key_t_pushstate")
// 	tanslateRecordTypeMap["Channel"] = T("key_t_channel")
// 	tanslateRecordTypeMap["MatchState"] = T("key_t_matchstate")
// 	tanslateRecordTypeMap["OperationType"] = T("key_t_operationtype")

// 	tanslateRecordTypeMap["SourceType"] = T("key_t_sourcetype")
// 	tanslateRecordTypeMap["Domain"] = T("key_t_domain")
// 	tanslateRecordTypeMap["BuildInMerchant"] = T("key_t_buildinmerchant")
// 	tanslateRecordTypeMap["FillchangeType"] = T("key_t_fillchangetype")
// 	tanslateRecordTypeMap["CompleteState"] = T("key_t_completestate")
// 	tanslateRecordTypeMap["BuildInType"] = T("key_t_buildintype")
// 	tanslateRecordTypeMap["IsAgent"] = T("key_t_isagent")

// 	tanslateRecordTypeMap["DepositChannel"] = T("key_t_channel")
// 	tanslateRecordTypeMap["WithDrawChannel"] = T("key_t_channel")
// 	tanslateRecordTypeMap["IsFill"] = T("key_t_isfill")
// 	tanslateRecordTypeMap["ManualHandle"] = T("key_t_manualhandle")
// 	tanslateRecordTypeMap["PointUporDown"] = T("key_t_pointupordown")

// 	tanslateRecordTypeMap["AgentRecalculateProfitState"] = T("key_t_agentrecalculateprofitstate")

// 	return tanslateRecordTypeMap

// }

// func TranslateModelToName(T i18n.TranslateFunc) (result map[string]map[string]string) {
// 	modelMap := make(map[string]map[string]string)

// 	modelMap["BuildInMerchant"] = map[string]string{
// 		fmt.Sprint(model.BuildIn): T("key_t_buildinmerchantyes"),
// 		fmt.Sprint(model.Real):    T("key_t_buildinmerchantno"),
// 	}

// 	modelMap["AgentRecalculateProfitState"] = map[string]string{
// 		fmt.Sprint(model.AgentRecalculateProfitStateSettle):     T("key_t_agentrecalculateprofitstatesettle"),
// 		fmt.Sprint(model.AgentRecalculateProfitStateAllSettled): T("key_t_agentrecalculateprofitstateallsettled"),
// 	}

// 	modelMap["BuildInMerchant"] = map[string]string{
// 		fmt.Sprint(model.BuildIn): T("key_t_buildinmerchantyes"),
// 		fmt.Sprint(model.Real):    T("key_t_buildinmerchantno"),
// 	}

// 	modelMap["PointUporDown"] = map[string]string{
// 		fmt.Sprint(model.Up):   T("key_t_up"),
// 		fmt.Sprint(model.Down): T("key_t_down"),
// 	}

// 	modelMap["ManualHandle"] = map[string]string{
// 		fmt.Sprint(model.IsFillNo):  T("key_t_isagentn"),
// 		fmt.Sprint(model.IsFillYes): T("key_t_isagenty"),
// 	}

// 	modelMap["IsAgent"] = map[string]string{
// 		fmt.Sprint(model.IsAgentY): T("key_t_isagenty"),
// 		fmt.Sprint(model.IsAgentN): T("key_t_isagentn"),
// 	}

// 	modelMap["FillchangeType"] = map[string]string{
// 		fmt.Sprint(model.FillIncrease):      T("key_t_fillincrease"),
// 		fmt.Sprint(model.FillDecrease):      T("key_t_filldecrease"),
// 		fmt.Sprint(model.FillFrezeIncrease): T("key_t_fill_freze_increase"),
// 		fmt.Sprint(model.FillFrezeDecrease): T("key_t_filld_freze_ecrease"),
// 	}

// 	modelMap["Domain"] = map[string]string{
// 		fmt.Sprint(model.DomainHt):  T("key_t_domainht"),
// 		fmt.Sprint(model.DomainAPI): T("key_t_domainapi"),
// 		fmt.Sprint(model.DomainAPP): T("key_t_domainapp"),
// 	}

// 	modelMap["OperationType"] = map[string]string{
// 		fmt.Sprint(model.OperationTypeQuery):  T("key_t_operationtypequery"),
// 		fmt.Sprint(model.OperationTypeCreate): T("key_t_operationtypecreate"),
// 		fmt.Sprint(model.OperationTypeUpdate): T("key_t_operationtypeupdate"),
// 		fmt.Sprint(model.OperationTypeDelete): T("key_t_operationtypedelete"),
// 	}

// 	modelMap["SourceType"] = map[string]string{
// 		fmt.Sprint(model.CLIENT): T("key_t_client"),
// 		fmt.Sprint(model.APP):    T("key_t_app"),
// 		fmt.Sprint(model.ExeAPI): T("key_t_exe"),
// 	}

// 	modelMap["UserType"] = map[string]string{
// 		fmt.Sprint(model.UserTypeMerchent):   T("key_t_usertypemerchent"),
// 		fmt.Sprint(model.UserTypePayer):      T("key_t_usertypepayer"),
// 		fmt.Sprint(model.UserTypeOperateAcc): T("key_t_usertypeoperateacc"),
// 		fmt.Sprint(model.UserTypeAdmin):      T("key_t_usertypeadmin"),
// 		fmt.Sprint(model.UserTypeAgent):      T("key_t_usertypeagent"),
// 	}

// 	modelMap["AccountState"] = map[string]string{
// 		fmt.Sprint(model.AccountEnable):  T("key_t_accountenable"),
// 		fmt.Sprint(model.AccountDisable): T("key_t_accountdisable"),
// 	}

// 	modelMap["MerchantState"] = map[string]string{
// 		fmt.Sprint(model.AccountEnable):  T("key_t_merchantstateenable"),
// 		fmt.Sprint(model.AccountDisable): T("key_t_merchantstatedisable"),
// 	}

// 	modelMap["SettlementRecordState"] = map[string]string{
// 		fmt.Sprint(model.SettlementRecordStatePending):      T("key_t_settlementrecordstatepending"),
// 		fmt.Sprint(model.SettlementRecordStateProcessing):   T("key_t_settlementrecordstateprocessing"),
// 		fmt.Sprint(model.SettlementRecordStatePreCompleted): T("key_t_settlementrecordstatePrecompleted"),
// 		fmt.Sprint(model.SettlementRecordStatePreClose):     T("key_t_settlementrecordstatepreclose"),
// 		fmt.Sprint(model.SettlementRecordStateComplete):     T("key_t_settlementrecordstatecomplete"),
// 		fmt.Sprint(model.SettlementRecordStateClose):        T("key_t_settlementrecordstateclose"),
// 	}

// 	modelMap["RecordType"] = map[string]string{
// 		fmt.Sprint(model.RecordTypeInjection):  T("key_t_recordtypeinjection"),
// 		fmt.Sprint(model.RecordTypeSettlement): T("key_t_recordtypesettlement"),
// 	}

// 	modelMap["BankAccountState"] = map[string]string{
// 		fmt.Sprint(model.BankAccountStateCreate):  T("key_t_bankaccountstatecreate"),
// 		fmt.Sprint(model.BankAccountStateEnable):  T("key_t_bankaccountstateenable"),
// 		fmt.Sprint(model.BankAccountStateDisable): T("key_t_bankaccountstatedisable"),
// 	}

// 	modelMap["TradingMode"] = map[string]string{
// 		fmt.Sprint(model.TradingModeByHand): T("key_t_tradingmodebyhand"),
// 		fmt.Sprint(model.TradingModeByAPI):  T("key_t_tradingmodebyapi"),
// 	}

// 	modelMap["FundchangeType"] = map[string]string{
// 		fmt.Sprint(model.BalanceIncrease): T("key_t_balanceincrease"),
// 		fmt.Sprint(model.BalanceDecrease): T("key_t_balancedecrease"),
// 		fmt.Sprint(model.FrozenIncrease):  T("key_t_frozenincrease"),
// 		fmt.Sprint(model.FrozenDecrease):  T("key_t_frozendecrease"),
// 		fmt.Sprint(model.Frozen):          T("key_t_frozen"),
// 		fmt.Sprint(model.UnFrozen):        T("key_t_unfrozen"),
// 	}

// 	modelMap["FundSourceType"] = map[string]string{
// 		fmt.Sprint(model.FundSourceTypeDeposit):    T("key_t_fundsourcetypedeposit"),
// 		fmt.Sprint(model.FundSourceTypeWithdraw):   T("key_t_fundsourcetypewithdraw"),
// 		fmt.Sprint(model.FundSourceTypeExeRecord):  T("key_t_FundSourceTypeExeRecord"),
// 		fmt.Sprint(model.FundSourceTypeSettlement): T("key_t_FundSourceTypeSettlement"),
// 		fmt.Sprint(model.FundSourceTypeFill):       T("key_t_FundSourceTypeFill"),
// 		fmt.Sprint(model.FundSourceTypeManual):     T("key_t_FundSourceTypeManual"),
// 		fmt.Sprint(model.FundSourceTypeAgent):      T("key_t_fundsourcetypeagent"),

// 		fmt.Sprint(model.FundSourceTypeDepositMBI):                         T("key_t_FundSourceTypeDepositMBI"),
// 		fmt.Sprint(model.FundSourceTypeDepositFeeMBD):                      T("key_t_FundSourceTypeDepositFeeMBD"),
// 		fmt.Sprint(model.FundSourceTypeMerchantAgentDepositProfitMBI):      T("key_t_FundSourceTypeMerchantAgentDepositProfitMBI"),
// 		fmt.Sprint(model.FundSourceTypePayerAssignedOrderPFR):              T("key_t_FundSourceTypePayerAssignedOrderPFR"),
// 		fmt.Sprint(model.FundSourceTypePayerAssignedOrderPUF):              T("key_t_FundSourceTypePayerAssignedOrderPUF"),
// 		fmt.Sprint(model.FundSourceTypePayerReceivedMoneySuccessPBD):       T("key_t_FundSourceTypePayerReceivedMoneySuccessPBD"),
// 		fmt.Sprint(model.FundSourceTypeDepositPayerProfitPBI):              T("key_t_FundSourceTypeDepositPayerProfitPBI"),
// 		fmt.Sprint(model.FundSourceTypeDepositPayerAgentProfitPBI):         T("key_t_FundSourceTypeDepositPayerAgentProfitPBI"),
// 		fmt.Sprint(model.FundSourceTypeMatchExeptionFailedPFR):             T("key_t_FundSourceTypeMatchExeptionFailedPFR"),
// 		fmt.Sprint(model.FundSourceTypeManualExeptionHandlePUF):            T("key_t_FundSourceTypeManualExeptionHandlePUF"),
// 		fmt.Sprint(model.FundSourceTypeOperateAccountOBD):                  T("key_t_FundSourceTypeOperateAccountOBD"),
// 		fmt.Sprint(model.FundSourceTypeOperateAccountOBI):                  T("key_t_FundSourceTypeOperateAccountOBI"),
// 		fmt.Sprint(model.FundSourceTypeCreateMerchantWithdrawMFR):          T("key_t_FundSourceTypeCreateMerchantWithdrawMFR"),
// 		fmt.Sprint(model.FundSourceTypeCreateMerchantWithdrawMUF):          T("key_t_FundSourceTypeCreateMerchantWithdrawMUF"),
// 		fmt.Sprint(model.FundSourceTypeCreateMerchantWithdrawSuccessMBD):   T("key_t_FundSourceTypeCreateMerchantWithdrawSuccessMBD"),
// 		fmt.Sprint(model.FundSourceTypeCreateMerchantWithdrawFeeMBD):       T("key_t_FundSourceTypeCreateMerchantWithdrawFeeMBD"),
// 		fmt.Sprint(model.FundSourceTypeMerchantAgentWithdrawProfitMBI):     T("key_t_FundSourceTypeMerchantAgentWithdrawProfitMBI"),
// 		fmt.Sprint(model.FundSourceTypePayerSeckillingOrderPBI):            T("key_t_FundSourceTypePayerSeckillingOrderPBI"),
// 		fmt.Sprint(model.FundSourceTypePayerSeckillingOrderProfitPBI):      T("key_t_FundSourceTypePayerSeckillingOrderProfitPBI"),
// 		fmt.Sprint(model.FundSourceTypePayerSeckillingOrderAgentProfitPBI): T("key_t_FundSourceTypePayerSeckillingOrderAgentProfitPBI"),
// 		fmt.Sprint(model.FundSourceTypeCreateMerchantSettlementMFR):        T("key_t_FundSourceTypeCreateMerchantSettlementMFR"),
// 		fmt.Sprint(model.FundSourceTypeCreateMerchantSettlementMUF):        T("key_t_FundSourceTypeCreateMerchantSettlementMUF"),
// 		fmt.Sprint(model.FundSourceTypeCreateMerchantSettlementSuccessMBD): T("key_t_FundSourceTypeCreateMerchantSettlementSuccessMBD"),
// 		fmt.Sprint(model.FundSourceTypeCreateMerchantSettlementFeeMFR):     T("key_t_FundSourceTypeCreateMerchantSettlementFeeMFR"),
// 		fmt.Sprint(model.FundSourceTypeCreateMerchantSettlementFeeMUF):     T("key_t_FundSourceTypeCreateMerchantSettlementFeeMUF"),
// 		fmt.Sprint(model.FundSourceTypeCreateMerchantSettlementFeeMBD):     T("key_t_FundSourceTypeCreateMerchantSettlementFeeMBD"),
// 		fmt.Sprint(model.FundSourceTypeCreateMerchantInjectionMBI):         T("key_t_FundSourceTypeCreateMerchantInjectionMBI"),
// 		fmt.Sprint(model.FundSourceTypeManualAdjustMerchantBanlanceMBI):    T("key_t_FundSourceTypeManualAdjustMerchantBanlanceMBI"),
// 		fmt.Sprint(model.FundSourceTypeManualAdjustMerchantBanlanceMBD):    T("key_t_FundSourceTypeManualAdjustMerchantBanlanceMBD"),
// 		fmt.Sprint(model.FundSourceTypeManualAdjustMerchantFreezeMFI):      T("key_t_FundSourceTypeManualAdjustMerchantFreezeMFI"),
// 		fmt.Sprint(model.FundSourceTypeManualAdjustMerchantFreezeMFD):      T("key_t_FundSourceTypeManualAdjustMerchantFreezeMFD"),
// 		fmt.Sprint(model.FundSourceTypeManualAdjustPayerBanlancePBI):       T("key_t_FundSourceTypeManualAdjustPayerBanlancePBI"),
// 		fmt.Sprint(model.FundSourceTypeManualAdjustPayerBanlancePBD):       T("key_t_FundSourceTypeManualAdjustPayerBanlancePBD"),
// 		fmt.Sprint(model.FundSourceTypeManualAdjustPayerFreezePFI):         T("key_t_FundSourceTypeManualAdjustPayerFreezePFI"),
// 		fmt.Sprint(model.FundSourceTypeManualAdjustPayerFreezePFD):         T("key_t_FundSourceTypeManualAdjustPayerFreezePFD"),
// 		fmt.Sprint(model.FundSourceTypeManualActivityPayerBanlancePBI):     T("key_t_FundSourceTypeManualActivityPayerBanlancePBI"),
// 		fmt.Sprint(model.FundSourceTypeManualActivityPayerBanlancePBD):     T("key_t_FundSourceTypeManualActivityPayerBanlancePBD"),
// 		fmt.Sprint(model.FundSourceTypeManualActivityPayerBanlancePFI):     T("key_t_FundSourceTypeManualActivityPayerBanlancePFI"),
// 		fmt.Sprint(model.FundSourceTypeManualActivityPayerBanlancePFD):     T("key_t_FundSourceTypeManualActivityPayerBanlancePFD"),
// 		fmt.Sprint(model.FundSourceTypeTransferOutPayerBanlancePBD):        T("key_t_FundSourceTypeTransferOutPayerBanlancePBD"),
// 		fmt.Sprint(model.FundSourceTypeTransferInPayerBanlancePBI):         T("key_t_FundSourceTypeTransferInPayerBanlancePBI"),
// 		fmt.Sprint(model.FundSourceTypePayerDepositPBI):                    T("key_t_FundSourceTypePayerDepositPBI"),
// 		fmt.Sprint(model.FundSourceTypeCreatePayerWithdrawPFR):             T("key_t_FundSourceTypeCreatePayerWithdrawPFR"),
// 		fmt.Sprint(model.FundSourceTypePayerWithdrawPUF):                   T("key_t_FundSourceTypePayerWithdrawPUF"),
// 		fmt.Sprint(model.FundSourceTypePayerWithdrawSuccessPBD):            T("key_t_FundSourceTypePayerWithdrawSuccessPBD"),
// 		fmt.Sprint(model.FundSourceTypePayerWithdrawFeePBD):                T("key_t_FundSourceTypePayerWithdrawFeePBD"),
// 		fmt.Sprint(model.FundSourceTypeOperateAccountInOBI):                T("key_t_FundSourceTypeOperateAccountInOBI"),
// 		fmt.Sprint(model.FundSourceTypeOperateAccountOutOFR):               T("key_t_FundSourceTypeOperateAccountOutOFR"),
// 		fmt.Sprint(model.FundSourceTypeOperateAccountOutOUF):               T("key_t_FundSourceTypeOperateAccountOutOUF"),
// 		fmt.Sprint(model.FundSourceTypeOperateAccountOutOBD):               T("key_t_FundSourceTypeOperateAccountOutOBD"),
// 	}

// 	modelMap["RemarkType"] = map[string]string{
// 		fmt.Sprint(model.RemarkTypePayer): T("key_t_remarktypepayer"),
// 		fmt.Sprint(model.RemarkTypeOrder): T("key_t_remarktypeorder"),
// 	}

// 	modelMap["PayerState"] = map[string]string{
// 		fmt.Sprint(model.PayerStateEnable):  T("key_t_payerstateenable"),
// 		fmt.Sprint(model.PayerStateDisable): T("key_t_payerstatedisable"),
// 	}

// 	modelMap["PayerType"] = map[string]string{
// 		fmt.Sprint(model.PayerTypeMember): T("key_t_payertypemember"),
// 		fmt.Sprint(model.PayerTypeVIP):    T("key_t_payertypevip"),
// 	}

// 	modelMap["TerminalType"] = map[string]string{
// 		fmt.Sprint(model.TerminalTypeRrf): T("key_t_terminaltyperrf"),
// 		fmt.Sprint(model.TerminalTypeDg):  T("key_t_terminaltypedg"),
// 	}

// 	modelMap["Location"] = map[string]string{
// 		fmt.Sprint(model.CHINA):     T("key_t_china"),
// 		fmt.Sprint(model.TAIWAN):    T("key_t_taiwan"),
// 		fmt.Sprint(model.XIANGGANG): T("key_t_xianggang"),
// 		fmt.Sprint(model.US):        T("key_t_us"),
// 		fmt.Sprint(model.VIETNAM):   T("key_t_vietnam"),
// 		fmt.Sprint(model.THAILAND):  T("key_t_thailand"),
// 		fmt.Sprint(model.KOREA):     T("key_t_korea"),
// 	}

// 	modelMap["OrderType"] = map[string]string{
// 		fmt.Sprint(model.OrderTypeDeposit):  T("key_t_ordertypedeposit"),
// 		fmt.Sprint(model.OrderTypeWithdraw): T("key_t_ordertypewithdraw"),
// 	}

// 	modelMap["Switch"] = map[string]string{
// 		fmt.Sprint(model.Off): T("key_t_off"),
// 		fmt.Sprint(model.On):  T("key_t_on"),
// 	}

// 	modelMap["PayerProfitsState"] = map[string]string{
// 		fmt.Sprint(model.PayerProfitsStateEnable):  T("key_t_payerprofitsstateenable"),
// 		fmt.Sprint(model.PayerProfitsStateDisable): T("key_t_payerprofitsstatedisable"),
// 	}

// 	modelMap["OperateAccountType"] = map[string]string{
// 		fmt.Sprint(model.OperateAccountTypeSh): T("key_t_operateaccounttypesh"),
// 		fmt.Sprint(model.OperateAccountTypeZj): T("key_t_operateaccounttypezj"),
// 		fmt.Sprint(model.OperateAccountTypeSf): T("key_t_operateaccounttypesf"),
// 	}

// 	modelMap["RateState"] = map[string]string{
// 		fmt.Sprint(model.RateStateEnable):  T("key_t_ratestateenable"),
// 		fmt.Sprint(model.RateStateDisable): T("key_t_ratestatedisable"),
// 	}

// 	modelMap["OrderState"] = map[string]string{
// 		fmt.Sprint(model.OrderStatesFreeze):           T("key_t_orderstatesfreeze"),
// 		fmt.Sprint(model.OrderStatesNew):              T("key_t_orderstatesnew"),
// 		fmt.Sprint(model.OrderStatesProcessing):       T("key_t_orderstatesprocessing"),
// 		fmt.Sprint(model.OrderStatesPreComplete):      T("key_t_orderstatesprecomplete"),
// 		fmt.Sprint(model.OrderStatesPreClose):         T("key_t_orderstatespreclose"),
// 		fmt.Sprint(model.OrderStatesCompleted):        T("key_t_orderstatescompleted"),
// 		fmt.Sprint(model.OrderStatesPartialCompleted): T("key_t_orderstatespartialcompleted"),
// 		fmt.Sprint(model.OrderStatesClose):            T("key_t_orderstatesclose"),
// 		fmt.Sprint(model.OrderStatesCloseByPayer):     T("key_t_orderstatesclosebypayer"),
// 	}

// 	modelMap["PaymentState"] = map[string]string{
// 		fmt.Sprint(model.PaymentStatesNew):          T("key_t_paymentstatesnew"),
// 		fmt.Sprint(model.PaymentStatesSettleFailed): T("key_t_paymentstatessettlefailed"),
// 		fmt.Sprint(model.PaymentStatesSettled):      T("key_t_paymentstatessettled"),
// 		fmt.Sprint(model.PaymentStatesAbandoned):    T("key_t_paymentstatesabandoned"),
// 	}

// 	modelMap["AccountType"] = map[string]string{
// 		fmt.Sprint(model.TypeThirdparty): T("key_t_typethirdparty"),
// 		fmt.Sprint(model.TypeBank):       T("key_t_typebank"),
// 		fmt.Sprint(model.TypeMerchant):   T("key_t_typemerchant"),
// 		fmt.Sprint(model.TypeQrCode):     T("key_t_typeqrcode"),
// 	}

// 	modelMap["OperateAccountOrderType"] = map[string]string{
// 		fmt.Sprint(model.OperateAccountOrderTypeConverted): T("key_t_operateaccountordertypeconverted"),
// 		fmt.Sprint(model.OperateAccountOrderTypeOutSide):   T("key_t_operateaccountordertypeoutside"),
// 	}

// 	modelMap["FinancialSubjectType"] = map[string]string{
// 		fmt.Sprint(model.FinancialSubjectTypeForeign):            T("key_t_financialsubjecttypeforeign"),
// 		fmt.Sprint(model.FinancialSubjectTypeAmount):             T("key_t_financialsubjecttypeamount"),
// 		fmt.Sprint(model.FinancialSubjectTypeFinancialInjection): T("key_t_financialsubjecttypefinancialinjection"),
// 		fmt.Sprint(model.FinancialSubjectTypeAdditionalIncome):   T("key_t_financialsubjecttypeadditionalincome"),
// 	}

// 	modelMap["OperateAccountOrderState"] = map[string]string{
// 		fmt.Sprint(model.OperateAccountOrderStateCreate):  T("key_t_operateaccountorderstatecreate"),
// 		fmt.Sprint(model.OperateAccountOrderStateSuccess): T("key_t_operateaccountorderstatesuccess"),
// 		fmt.Sprint(model.OperateAccountOrderStateFail):    T("key_t_operateaccountorderstatefail"),
// 	}

// 	modelMap["FillType"] = map[string]string{
// 		fmt.Sprint(model.Activity): T("key_t_activity"),
// 		fmt.Sprint(model.Adjust):   T("key_t_adjust"),
// 		fmt.Sprint(model.Transfer): T("key_t_filld_fillfrezetransfersbetween"),
// 	}

// 	modelMap["FillState"] = map[string]string{
// 		fmt.Sprint(model.FillStatePending):    T("key_t_fillstatepending"),
// 		fmt.Sprint(model.FillStateProcessing): T("key_t_fillstateprocessing"),
// 		fmt.Sprint(model.FillStateComplete):   T("key_t_fillstatecomplete"),
// 		fmt.Sprint(model.FillStateClose):      T("key_t_fillstateclose"),
// 	}

// 	modelMap["FillTypeForMerchant"] = map[string]string{
// 		fmt.Sprint(model.FillPayment):   T("key_t_fillpayment"),
// 		fmt.Sprint(model.FillGathering): T("key_t_fillgathering"),
// 	}

// 	modelMap["TransferType"] = map[string]string{
// 		fmt.Sprint(model.TransferTypeIn):  T("key_t_transfertypein"),
// 		fmt.Sprint(model.TransferTypeOut): T("key_t_transfertypeout"),
// 	}

// 	modelMap["Online"] = map[string]string{
// 		fmt.Sprint(model.POnline):   T("key_t_ponline"),
// 		fmt.Sprint(model.POffline):  T("key_t_poffline"),
// 		fmt.Sprint(model.PDropline): T("key_t_pdropline"),
// 	}
// 	modelMap["BuildInType"] = map[string]string{
// 		fmt.Sprint(model.BuildIn): T("key_t_buildin"),
// 		fmt.Sprint(model.Real):    T("key_t_real"),
// 	}

// 	modelMap["ParamType"] = map[string]string{
// 		fmt.Sprint(model.BizParam): T("key_t_bizparam"),
// 		fmt.Sprint(model.SysParam): T("key_t_sysparam"),
// 	}

// 	modelMap["ParamState"] = map[string]string{
// 		fmt.Sprint(model.ParamEnable):  T("key_t_paramenable"),
// 		fmt.Sprint(model.ParamDisable): T("key_t_paramdisable"),
// 	}

// 	modelMap["PushState"] = map[string]string{
// 		fmt.Sprint(model.Pushing):     T("key_t_pushing"),
// 		fmt.Sprint(model.PushSuccess): T("key_t_pushsuccess"),
// 		fmt.Sprint(model.PushFailed):  T("key_t_pushfailed"),
// 	}

// 	modelMap["CompleteState"] = map[string]string{
// 		fmt.Sprint(model.Complete):            T("key_t_complete"),
// 		fmt.Sprint(model.UnComplete):          T("key_t_uncomplete"),
// 		fmt.Sprint(model.CompleteStateDetete): T("key_t_completestatedetete"),
// 	}

// 	modelMap["Channel"] = map[string]string{
// 		fmt.Sprint(model.AlipayH5):   T("key_t_alipayh5"),
// 		fmt.Sprint(model.Alipay):     T("key_t_alipay"),
// 		fmt.Sprint(model.AlipayKs):   T("ket_t_alipayks"),
// 		fmt.Sprint(model.AlipayCard): T("ket_t_alipaycard"),
// 		fmt.Sprint(model.WeiXin):     T("key_t_wechat"),
// 		fmt.Sprint(model.CMBC):       T("ket_t_cmbc"),
// 		fmt.Sprint(model.ICBC):       T("ket_t_icbc"),
// 		fmt.Sprint(model.BOC):        T("ket_t_boc"),
// 		fmt.Sprint(model.BOCOM):      T("ket_t_bocom"),
// 		fmt.Sprint(model.PINGAN):     T("ket_t_pingan"),
// 		fmt.Sprint(model.CMB):        T("ket_t_cmb"),
// 		fmt.Sprint(model.ABC):        T("ket_t_abc"),
// 		fmt.Sprint(model.CCB):        T("ket_t_ccb"),
// 		fmt.Sprint(model.PSBC):       T("ket_t_psbc"),
// 		fmt.Sprint(model.CEBB):       T("ket_t_cebb"),
// 		fmt.Sprint(model.CIB):        T("ket_t_cib"),
// 		fmt.Sprint(model.SPDB):       T("ket_t_spdb"),
// 		fmt.Sprint(model.CGB):        T("ket_t_cgb"),
// 		fmt.Sprint(model.CITIC):      T("ket_t_citic"),
// 		fmt.Sprint(model.HXB):        T("ket_t_hxb"),
// 		fmt.Sprint(model.BCCB):       T("ket_t_bccb"),
// 		fmt.Sprint(model.BOSC):       T("ket_t_bosc"),
// 		fmt.Sprint(model.GZCB):       T("ket_t_gzcb"),
// 		fmt.Sprint(model.BANK):       T("ket_t_bank"),
// 		fmt.Sprint(model.JY):         T("ket_t_jy"),
// 		fmt.Sprint(model.UNIONPAY):   T("ket_t_unionpay"),
// 		fmt.Sprint(model.STUNIONPAY): T("ket_t_st_unionpay"),
// 	}

// 	modelMap["ChannelType"] = map[string]string{
// 		fmt.Sprint(model.ChannelTypeWeChat): T("key_t_wechat"),
// 		fmt.Sprint(model.ChannelTypeAlipay): T("key_t_alipay"),
// 		fmt.Sprint(model.ChannelTypeBank):   T("key_t_bank"),
// 		fmt.Sprint(model.ChannelTypeJy):     T("ket_t_jy"),
// 		fmt.Sprint(model.ChannelTypeYSF):    T("ket_t_unionpay"),
// 	}

// 	modelMap["DepositChannel"] = map[string]string{
// 		fmt.Sprint(model.AlipayH5):   T("key_t_alipayh5"),
// 		fmt.Sprint(model.Alipay):     T("key_t_alipay"),
// 		fmt.Sprint(model.BANK):       T("ket_t_bank"),
// 		fmt.Sprint(model.JY):         T("ket_t_jy"),
// 		fmt.Sprint(model.UNIONPAY):   T("ket_t_unionpay"),
// 		fmt.Sprint(model.STUNIONPAY): T("ket_t_st_unionpay"),
// 		fmt.Sprint(model.WeiXin):     T("key_t_wechat"),
// 	}

// 	modelMap["WithDrawChannel"] = map[string]string{
// 		fmt.Sprint(model.Alipay): T("key_t_alipay"),
// 		fmt.Sprint(model.BANK):   T("ket_t_bank"),
// 	}

// 	modelMap["MatchState"] = map[string]string{
// 		fmt.Sprint(model.AutoMatched):    T("key_t_automatched"),
// 		fmt.Sprint(model.NotMatch):       T("key_t_notmatch"),
// 		fmt.Sprint(model.MatchedClosed):  T("key_t_matchedclosed"),
// 		fmt.Sprint(model.MatchedNew):     T("key_t_matchednew"),
// 		fmt.Sprint(model.ManualMatched):  T("key_t_manual_matched"),
// 		fmt.Sprint(model.ExcepitonMatch): T("key_t_matchedexcepiton"),
// 		fmt.Sprint(model.SmsMatch):       T("key_t_smsmatch"),
// 	}
// 	return modelMap
// }

const (
	OperationTime           = "OperationTime"
	Maintenance             = "Maintenance"
	RecalculateAgentProfit  = "RecalculateAgentProfit"
	PayerInnerTransfer      = "PayerInnerTransfer"
	IsPointType             = "IsPointType"
	DestructionCountDown    = "DestructionCountDown"
	QServerURL              = "QServerURL"
	PayerWithdrawRate       = "PayerWithdrawRate"
	DailyDepositCountLimit  = "DailyDepositCountLimit"
	DailyDepositTotalLimit  = "DailyDepositTotalLimit"
	SelectByIP              = "SelectByIP"
	notice_parse            = "notice_parse"
	sms_parse               = "sms_parse"
	DailyWithdrawCountLimit = "DailyWithdrawCountLimit"
	DailyWithdrawTotalLimit = "DailyWithdrawTotalLimit"
	ip_check                = "ip_check"
	ProxyURL                = "ProxyURL"
	MerchanrSettlementRate  = "MerchanrSettlementRate"
	MaxBalance              = "MaxBalance"
)
