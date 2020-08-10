package service

import (
	"github.com/terryli1643/apidemo/domain/dao"
	"github.com/terryli1643/apidemo/domain/model"
)

// import (
// 	"errors"
// 	"fmt"
// 	"strconv"
// 	"strings"

// 	"github.com/sirupsen/logrus"
// 	log "github.com/sirupsen/logrus"
// 	"gitlab.99safe.org/magpie/proto/auth/authclient"
// 	"gitlab.99safe.org/magpie/proto/auth/authproto"
// 	"gitlab.99safe.org/rrp/rrp-backend/data/dao"
// 	"gitlab.99safe.org/rrp/rrp-backend/data/model"
// 	"gitlab.99safe.org/rrp/rrp-backend/server"
// 	"gitlab.99safe.org/rrp/rrp-backend/service/dto"
// )

type AdminService struct {
	encryptKey []byte
}

var adminServiceObj *AdminService

func NewAdminService() *AdminService {
	if adminServiceObj == nil {
		l.Lock()
		if adminServiceObj == nil {
			adminServiceObj = &AdminService{
				encryptKey: []byte("TvggrXNWpvjRZ5GwVWLLtPMQxuXe28ya"),
			}
		}
		l.Unlock()
	}
	return adminServiceObj
}

func (service *AdminService) LoadUserByUsername(username string) model.Admin {
	db := getDB()
	adminDao := dao.NewAdminDao(db)
	admin := model.Admin{
		Account: username,
	}
	adminDao.FindOne(&admin)

	return admin
}

// func (service *UserService) Authenticate(authentication shadowsecurity.IAuthentication) shadowsecurity.IAuthentication {
// 	if requestAuthenticationToken, ok := authentication.(*shadowsecurity.TRequestAuthenticationToken); ok {
// 		details := authentication.GetDetails()
// 		if policy, ok := details.(shadowsecurity.TCasbinPolicyDetails); ok {
// 			//调用rpc请求权限服务器
// 			client, err := authclient.NewAuthClient(service.ServerConfig.AuthServer.AppID, service.ServerConfig.AuthServer.Addr)
// 			if err != nil {
// 				log.Errorf("NewAuthClient failed. err=[%v]", err)
// 				return nil
// 			}
// 			defer client.Close()

// 			privilegeService := NewPrivilegeService()
// 			privilege := privilegeService.GetResourceByURL(policy.Obj, policy.Act)

// 			resp, err := client.CheckPrivilegeRule(policy.Sub, privilege.ResourceCode, privilege.PrivilegeCode)
// 			if err != nil {
// 				log.Errorf("AssignPrivilegeRule failed. err=[%v]", err)
// 				return nil
// 			}

// 			if resp.Response.Code != "0" {
// 				log.Errorf("AssignPrivilegeRule failed. ErrMsg=[%v]", resp.Response.Message)
// 				return nil
// 			}
// 			log.WithFields(logrus.Fields{
// 				"roleCode": policy.Sub,
// 				"privCode": privilege.ResourceCode,
// 				"action":   privilege.PrivilegeCode,
// 				"pass":     resp.Pass,
// 			}).Debug("Authenticate")
// 			if resp.Pass {
// 				requestAuthenticationToken.SetAuthenticated(true)
// 			}

// 		}
// 		return authentication
// 	}
// 	return nil
// }

// func (adminManager *UserService) GetAdminByLoginName(login string) (admin model.Admin, err error) {
// 	if login == "" {
// 		return model.Admin{}, errors.New("no login name")
// 	}
// 	db := getDB()
// 	adminDao := dao.NewAdminDao(db)
// 	result := model.Admin{
// 		Account: login,
// 	}
// 	log.Infoln("GetAdminByLoginName", login)
// 	err = adminDao.FindOne(&result)
// 	if err != nil {
// 		log.Error(err)
// 		return model.Admin{}, err
// 	}
// 	return result, nil
// }

// func (service UserService) GetAdminByID(adminID int64) (admin model.Admin, err error) {
// 	db := getDB()
// 	adminDao := dao.NewAdminDao(db)
// 	admin = model.Admin{
// 		ID: adminID,
// 	}
// 	err = adminDao.Get(&admin)
// 	return admin, err
// }

// func (service UserService) FindAdminListPaging(condition model.SearchAdminCondition, pageNum int, pageSize int) (result []model.Admin, count int, err error) {
// 	rowbound := model.NewRowBound(pageNum, pageSize)

// 	db := getDB()
// 	adminDao := dao.NewAdminDao(db)
// 	result, count, err = adminDao.FindAdmin(condition, rowbound)
// 	return
// }

// func (service UserService) VerifySecurePwd(adminID int64, securePwd string) (err error) {
// 	db := getDB()
// 	adminDao := dao.NewAdminDao(db)
// 	admin := model.Admin{
// 		ID: adminID,
// 	}
// 	err = adminDao.Get(&admin)
// 	if err != nil {
// 		return err
// 	}
// 	passwordEncoder := shadowsecurity.PasswordEncoderInstance(shadowsecurity.PASSWORD_ENCODER)
// 	//默认最多可重试6次
// 	defaultTryTimes := 6
// 	wrongTimes := 0
// 	if !passwordEncoder.Matches(securePwd, admin.SecurePassword) {
// 		//记录输入错误次数
// 		val, ok := wrongSecurePwdCountHolder.Load(adminID)
// 		if ok {
// 			wrongTimes = val.(int) + 1
// 		} else {
// 			wrongTimes = 1
// 		}
// 		//如果错误次数达到最大，则冻结帐号
// 		if wrongTimes == defaultTryTimes {
// 			db := getDB()
// 			adminDao := dao.NewAdminDao(db)
// 			adminDao.Update(adminID, "state", model.AccountDisable)
// 			wrongSecurePwdCountHolder.Delete(adminID)
// 			return shadowsecurity.AccountLockedError{}
// 		}
// 		wrongSecurePwdCountHolder.Store(adminID, wrongTimes)
// 		err = InvalidSecurePwd{
// 			errors.New("invalid secure pwd"),
// 			defaultTryTimes - wrongTimes,
// 		}
// 		return err
// 	}
// 	wrongSecurePwdCountHolder.Delete(adminID)
// 	return nil
// }

// //创建管理员
// func (service UserService) CreateAdmin(ctx model.GlobalHandler, admin model.Admin) (err error) {
// 	db := getDB()
// 	adminDao := dao.NewAdminDao(db)

// 	// 验证用户账户已经存在
// 	if ok := adminDao.Existed(&model.Admin{
// 		// 帐号不区分大小写，所有帐号转成大写验证是否存在
// 		Account: strings.ToUpper(admin.Account),
// 	}); ok {
// 		err = AccountExistError{
// 			error: errors.New("account is exist"),
// 		}
// 		log.Error(err)
// 		return err
// 	}

// 	//加密密码
// 	passwordEncoder := shadowsecurity.PasswordEncoderInstance(shadowsecurity.PASSWORD_ENCODER)
// 	admin.LoginPassword = passwordEncoder.Encode(admin.LoginPassword)
// 	admin.SecurePassword = passwordEncoder.Encode(admin.SecurePassword)

// 	admin.UserType = model.UserTypeAdmin
// 	admin.State = model.AccountEnable

// 	err = adminDao.Create(&admin)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}

// 	go NewSysLogService().CreateSysLog(ctx.SysLog, nil, admin)
// 	return nil
// }

// func (service UserService) UpdateAdminPassword(ctx model.GlobalHandler, UserID int64, password string) (err error) {
// 	db := getDB()
// 	adminDao := dao.NewAdminDao(db)
// 	passwordEncoder := shadowsecurity.PasswordEncoderInstance(shadowsecurity.PASSWORD_ENCODER)
// 	obj := model.Admin{
// 		ID: UserID,
// 	}
// 	err = adminDao.FindOne(&obj)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}
// 	password = passwordEncoder.Encode(password)
// 	err = adminDao.Updates(UserID, map[string]interface{}{"login_password": password})
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}
// 	go NewSysLogService().CreateSysLog(ctx.SysLog, obj, map[string]interface{}{"login_password": password})
// 	//清除session
// 	//todo:
// 	return nil
// }

// //UpdateAdminState 修改管理员状态
// func (service UserService) UpdateAdminState(ctx model.GlobalHandler, adminID int64, state model.AccountState) (err error) {
// 	db := getDB()
// 	adminDao := dao.NewAdminDao(db)

// 	admin, err := service.GetAdminByID(adminID)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	if admin.State == state {
// 		return
// 	}

// 	err = adminDao.Update(adminID, "state", state)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}

// 	new := admin
// 	new.State = state
// 	go NewSysLogService().CreateSysLog(ctx.SysLog, admin, new)
// 	return
// }

// //DeleteAdmin 删除admin
// func (service UserService) DeleteAdmin(ctx model.GlobalHandler, adminID int64) (err error) {
// 	db := getDB()
// 	adminDao := dao.NewAdminDao(db)

// 	admin, err := service.GetAdminByID(adminID)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	err = adminDao.Delete(&admin)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}

// 	roleRelDao := dao.NewRoleUserRelDao(db)
// 	err = roleRelDao.DeleteBYuserId(admin.ID)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}

// 	go NewSysLogService().CreateSysLog(ctx.SysLog, admin, nil)
// 	return
// }

// //UpdateAdmin 更新admin
// func (service UserService) UpdateAdmin(ctx model.GlobalHandler, new model.Admin) (err error) {
// 	db := getDB()
// 	adminDao := dao.NewAdminDao(db)

// 	admin, err := service.GetAdminByID(new.ID)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	if new.Account != admin.Account {
// 		new.Account = admin.Account
// 	}

// 	if new.LoginPassword != "" {
// 		//加密密码
// 		passwordEncoder := shadowsecurity.PasswordEncoderInstance(shadowsecurity.PASSWORD_ENCODER)
// 		new.LoginPassword = passwordEncoder.Encode(new.LoginPassword)
// 	}

// 	if new.SecurePassword != "" {
// 		passwordEncoder := shadowsecurity.PasswordEncoderInstance(shadowsecurity.PASSWORD_ENCODER)
// 		new.SecurePassword = passwordEncoder.Encode(new.SecurePassword)
// 	}

// 	new.CreatedAt = admin.CreatedAt
// 	err = adminDao.Save(&new)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}

// 	go NewSysLogService().CreateSysLog(ctx.SysLog, admin, new)
// 	return
// }

// //UpdateAdminRole 分配角色
// func (service UserService) UpdateAdminRole(ctx model.GlobalHandler, adminID int64, roleID []int64) (err error) {
// 	db := getDB()
// 	tx := db.Begin()
// 	defer closeTx(tx, &err)

// 	roleUserRelDao := dao.NewRoleUserRelDao(tx)
// 	roleDao := dao.NewRoleDao(tx)

// 	oldRoles, err := roleUserRelDao.Find(&model.RoleUserRel{
// 		UserID: adminID,
// 	})
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	roles, err := roleDao.BatchGet(roleID)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	//删除旧关系
// 	result, err := roleUserRelDao.Find(&model.RoleUserRel{
// 		UserID: adminID,
// 	})
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}
// 	ids := make([]int64, len(result))
// 	for i, v := range result {
// 		ids[i] = v.ID
// 	}

// 	if len(ids) > 0 {
// 		err = roleUserRelDao.BatchDelete(ids)
// 		if err != nil {
// 			log.Error(err)
// 			return err
// 		}
// 	}

// 	//添加新关系
// 	for _, r := range roles {
// 		err = roleUserRelDao.Create(&model.RoleUserRel{
// 			UserID: adminID,
// 			RoleID: r.ID,
// 		})
// 		if err != nil {
// 			log.Error(err)
// 			return err
// 		}
// 	}

// 	//调用rpc请求权限服务器
// 	client, err := authclient.NewAuthClient(service.ServerConfig.AuthServer.AppID, service.ServerConfig.AuthServer.Addr)
// 	if err != nil {
// 		log.Errorf("NewAuthClient failed. err=[%v]", err)
// 		return
// 	}
// 	defer client.Close()

// 	var gParam []*authproto.PolicyGParam
// 	for _, v := range roles {
// 		gParam = append(gParam, &authproto.PolicyGParam{
// 			InheritRoleCode: v.RoleCode,
// 		})
// 	}
// 	resp, err := client.UpdatePrivilegeRule(fmt.Sprintf("%d", adminID), "g", nil, gParam)
// 	if err != nil {
// 		log.Errorf("GetPermissionsForRole failed. err=[%v]", err)
// 		return
// 	}
// 	if resp.Response.Code != "0" {
// 		log.Errorf("GetPermissionsForRole failed. ErrMsg=[%v]", resp.Response.Message)
// 		return
// 	}

// 	go NewSysLogService().CreateSysLog(ctx.SysLog, oldRoles, roles)
// 	return
// }

// //FindRoles 查询角色
// func (service UserService) FindRoles(role model.Role) (result []model.Role, err error) {
// 	db := getDB()
// 	roleDao := dao.NewRoleDao(db)

// 	result, err = roleDao.Find(&role)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}
// 	return
// }

// //CreateRole 创建角色
// func (service UserService) CreateRole(ctx model.GlobalHandler, pRole model.Role) (err error) {
// 	db := getDB()
// 	roleDao := dao.NewRoleDao(db)
// 	role := model.Role{
// 		RoleName: pRole.RoleName,
// 	}
// 	result, err := roleDao.Find(&role)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}
// 	if len(result) > 0 {
// 		err = errors.New("role is exist")
// 		log.Error(err)
// 		return
// 	}

// 	pRole.ID = GenRoleID()
// 	pRole.RoleCode = "ROLE_" + strconv.FormatInt(pRole.ID, 10)
// 	pRole.State = model.AccountEnable
// 	err = roleDao.Create(&pRole)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	go NewSysLogService().CreateSysLog(ctx.SysLog, nil, pRole)
// 	return nil
// }

// //UpdateRoleState 修改角色状态
// func (service UserService) UpdateRoleState(ctx model.GlobalHandler, roleID int64, state model.AccountState) (err error) {
// 	db := getDB()
// 	roleDao := dao.NewRoleDao(db)

// 	role, err := service.GetRoleByID(roleID)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}
// 	//默认角色不允许修改
// 	if role.RoleCode == "DEFAULT_ROLE" {
// 		err = DefaultRoleUpdateError{
// 			errors.New("default role does not allow to update"),
// 		}
// 		log.Error(err)
// 		return
// 	}

// 	if role.State == state {
// 		return
// 	}

// 	err = roleDao.Update(roleID, "state", state)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}

// 	new := role
// 	new.State = state

// 	//调用rpc请求权限服务器
// 	client, err := authclient.NewAuthClient(service.ServerConfig.AuthServer.AppID, service.ServerConfig.AuthServer.Addr)
// 	if err != nil {
// 		log.Errorf("NewAuthClient failed. err=[%v]", err)
// 		return
// 	}
// 	defer client.Close()

// 	if state == model.AccountEnable {
// 		resp, err := client.CreateRole(new.RoleName, new.RoleCode, new.Description)
// 		if err != nil {
// 			log.Errorf("GetPermissionsForRole failed. err=[%v]", err)
// 			return err
// 		}
// 		if resp.Response.Code != "0" {
// 			log.Errorf("GetPermissionsForRole failed. ErrMsg=[%v]", resp.Response.Message)
// 			return err
// 		}
// 	} else {
// 		resp, err := client.DeleteRole(role.RoleCode)
// 		if err != nil {
// 			log.Errorf("GetPermissionsForRole failed. err=[%v]", err)
// 			return err
// 		}
// 		if resp.Response.Code != "0" {
// 			log.Errorf("GetPermissionsForRole failed. ErrMsg=[%v]", resp.Response.Message)
// 			return err
// 		}
// 	}
// 	go NewSysLogService().CreateSysLog(ctx.SysLog, role, new)
// 	return
// }

// //CreateRole 查询角色
// func (service UserService) GetRoleByID(roleID int64) (role model.Role, err error) {
// 	db := getDB()
// 	roleDao := dao.NewRoleDao(db)
// 	role.ID = roleID
// 	err = roleDao.Get(&role)

// 	if err != nil {
// 		log.Error(err)
// 	}
// 	return
// }

// //UpdateRole 修改角色
// func (service UserService) UpdateRole(ctx model.GlobalHandler, roleID int64, roleName string, desc string) (err error) {
// 	db := getDB()
// 	roleDao := dao.NewRoleDao(db)
// 	role := model.Role{
// 		ID: roleID,
// 	}
// 	err = roleDao.Get(&role)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	//默认角色不允许修改
// 	if role.RoleCode == "DEFAULT_ROLE" {
// 		err = DefaultRoleUpdateError{
// 			errors.New("default role does not allow to update"),
// 		}
// 		log.Error(err)
// 		return
// 	}

// 	new := &model.Role{
// 		ID:          role.ID,
// 		RoleCode:    role.RoleCode,
// 		RoleName:    roleName,
// 		Description: desc,
// 		State:       role.State,
// 	}
// 	err = roleDao.Save(new)

// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	go NewSysLogService().CreateSysLog(ctx.SysLog, role, new)
// 	return nil
// }

// //DeleteRole 删除角色
// func (service UserService) DeleteRole(ctx model.GlobalHandler, roleID int64) (err error) {
// 	db := getDB()
// 	roleDao := dao.NewRoleDao(db)
// 	role := model.Role{
// 		ID: roleID,
// 	}
// 	err = roleDao.Get(&role)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	//默认角色不允许删除
// 	if role.RoleCode == "DEFAULT_ROLE" {
// 		err = DefaultRoleDeleteError{
// 			errors.New("default role does not allow to delete"),
// 		}
// 		log.Error(err)
// 		return
// 	}

// 	//如果角色有分配给用户，则不可以删除
// 	roleUserRelDao := dao.NewRoleUserRelDao(db)
// 	roleUserRels, err := roleUserRelDao.Find(&model.RoleUserRel{
// 		RoleID: role.ID,
// 	})
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}
// 	if len(roleUserRels) > 0 {
// 		err = NotEmptyROleDeleteError{
// 			errors.New("not empty role does not allow to delete"),
// 		}
// 		log.Error(err)
// 		return
// 	}

// 	roleRelDao := dao.NewRoleUserRelDao(db)
// 	result, err := roleRelDao.Find(&model.RoleUserRel{
// 		RoleID: role.ID,
// 	})
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}
// 	if len(result) > 0 {
// 		err = errors.New("can not delete role which is not empty")
// 		log.Error(err)
// 		return
// 	}

// 	err = roleDao.Delete(&role)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}

// 	//调用rpc请求权限服务器
// 	client, err := authclient.NewAuthClient(service.ServerConfig.AuthServer.AppID, service.ServerConfig.AuthServer.Addr)
// 	if err != nil {
// 		log.Errorf("NewAuthClient failed. err=[%v]", err)
// 		return
// 	}
// 	defer client.Close()
// 	resp, err := client.DeleteRole(role.RoleCode)
// 	if err != nil {
// 		log.Errorf("GetPermissionsForRole failed. err=[%v]", err)
// 		return err
// 	}
// 	if resp.Response.Code != "0" {
// 		log.Errorf("GetPermissionsForRole failed. ErrMsg=[%v]", resp.Response.Message)
// 		return err
// 	}

// 	go NewSysLogService().CreateSysLog(ctx.SysLog, role, nil)
// 	return nil
// }

// //GetRolePrivilege 查看角色已有权限
// func (service UserService) GetRolePrivileges(roleID int64) (role model.Role, allPrivs []model.Privilege, rolePrivs []model.Privilege, err error) {
// 	db := getDB()
// 	roleDao := dao.NewRoleDao(db)
// 	role.ID = roleID
// 	err = roleDao.Get(&role)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}
// 	privilegeDao := dao.NewPrivilegeDao(db)
// 	allPrivs, err = privilegeDao.FindAll(&model.Privilege{})
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}
// 	//默认角色默认全部权限
// 	if role.RoleCode == "DEFAULT_ROLE" {
// 		return role, allPrivs, allPrivs, nil
// 	}

// 	roleResPrivRelDao := dao.NewRolePrivsRelDao(db)
// 	allPrivsRels, err := roleResPrivRelDao.Find(&model.RolePrivsRel{
// 		RoleID: roleID,
// 	})

// 	m := make(map[int64]model.Privilege)
// 	for _, v := range allPrivs {
// 		m[v.ID] = v
// 	}

// 	for _, rel := range allPrivsRels {
// 		if p, ok := m[rel.PrivilegeID]; ok {
// 			rolePrivs = append(rolePrivs, p)
// 		}
// 	}

// 	return
// }

// //SetPrivilegeToRole 分配给角色权限
// func (service UserService) SetPrivilegeToRole(ctx model.GlobalHandler, roleID int64, privilegeID []int64) (err error) {
// 	if len(privilegeID) == 0 {
// 		return
// 	}
// 	db := getDB()
// 	roleDao := dao.NewRoleDao(db)
// 	role := model.Role{
// 		ID: roleID,
// 	}
// 	err = roleDao.Get(&role)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	//默认角色不允许删除
// 	if role.RoleCode == "DEFAULT_ROLE" {
// 		err = DefaultRoleUpdateError{
// 			errors.New("default role does not allow to update"),
// 		}
// 		log.Error(err)
// 		return
// 	}

// 	tx := db.Begin()
// 	defer closeTx(tx, &err)

// 	//删除旧关系
// 	rolePrivRelDao := dao.NewRolePrivsRelDao(tx)
// 	err = rolePrivRelDao.DeleteByRoleID(roleID)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	privDao := dao.NewPrivilegeDao(tx)
// 	privs, err := privDao.BatchGet(privilegeID)
// 	if err != nil {
// 		log.Error(err)
// 		return err
// 	}

// 	for _, v := range privs {
// 		rolePrivRelDao.Create(&model.RolePrivsRel{
// 			RoleID:      roleID,
// 			PrivilegeID: v.ID,
// 		})
// 	}

// 	//调用rpc请求权限服务器
// 	client, err := authclient.NewAuthClient(service.ServerConfig.AuthServer.AppID, service.ServerConfig.AuthServer.Addr)
// 	if err != nil {
// 		log.Errorf("NewAuthClient failed. err=[%v]", err)
// 		return
// 	}
// 	defer client.Close()

// 	privilegeDao := dao.NewPrivilegeDao(db)
// 	privileges, err := privilegeDao.BatchGet(privilegeID)
// 	var pParam []*authproto.PolicyPParam
// 	for _, pri := range privileges {
// 		pParam = append(pParam, &authproto.PolicyPParam{
// 			PrivCode: pri.ResourceCode,
// 			Action:   pri.PrivilegeCode,
// 		})
// 	}

// 	pParam = append(pParam, &authproto.PolicyPParam{
// 		PrivCode: "Basic",
// 		Action:   "Basic_*",
// 	})
// 	resp, err := client.UpdatePrivilegeRule(role.RoleCode, "p", pParam, nil)
// 	if err != nil {
// 		log.Errorf("AssignPrivilegeRule failed. err=[%v]", err)
// 		return err
// 	}

// 	if resp.Response.Code != "0" {
// 		log.Errorf("AssignPrivilegeRule failed. ErrMsg=[%v]", resp.Response.Message)
// 		return err
// 	}
// 	return nil
// }

// func (service UserService) GetPrivilegeToRoleIDs(adminID int64) (privilege []model.Privilege, err error) {
// 	db := getDB()
// 	//获取角色
// 	roleUserRelDao := dao.NewRoleUserRelDao(db)
// 	roles, err := roleUserRelDao.Find(&model.RoleUserRel{
// 		UserID: adminID,
// 	})
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}
// 	var roleIDs []int64

// 	for _, v := range roles {
// 		roleIDs = append(roleIDs, v.RoleID)
// 	}

// 	roleResPrivRelDao := dao.NewRolePrivsRelDao(db)
// 	allPrivsRels, err := roleResPrivRelDao.BatchGetRoleID(roleIDs)
// 	if err != nil {
// 		log.Error(err)
// 		return nil, err
// 	}
// 	var privsIDs []int64
// 	privsMap := make(map[int64]interface{})

// 	for _, v := range allPrivsRels {
// 		privsMap[v.PrivilegeID] = v
// 	}

// 	for k := range privsMap {
// 		privsIDs = append(privsIDs, k)
// 	}

// 	privilegeDao := dao.NewPrivilegeDao(db)
// 	allPrivs, err := privilegeDao.BatchGet(privsIDs)
// 	if err != nil {
// 		log.Error(err)
// 		return
// 	}
// 	return allPrivs, nil
// }

// func (adminService UserService) UpdateAdminPwd(userid int64, password dto.UpdatePwd) (err error) {
// 	log.WithFields(logrus.Fields{
// 		"userid":   userid,
// 		"password": password,
// 	}).Debug("UpdateAdminPwd")
// 	db := getDB()
// 	if password.NewPwd == password.OldPwd {
// 		err = NewPasswordSameAsOldPasswordError{errors.New("newPassword is samewith oldPassword!")}
// 		log.Error(err)
// 		return
// 	}

// 	if password.NewPwd != password.EnsurePwd {
// 		err = NewPasswordNotSameAsEnsurePasswordError{errors.New("newPassword did not match ensurePassword!")}
// 		log.Error(err)
// 		return
// 	}
// 	adminDao := dao.NewAdminDao(db)
// 	admin := model.Admin{
// 		ID: userid,
// 	}
// 	err = adminDao.Get(&admin)
// 	if err != nil {
// 		return
// 	}
// 	passwordEncoder := shadowsecurity.PasswordEncoderInstance(shadowsecurity.PASSWORD_ENCODER)
// 	var oldPwd string
// 	oldPwd = admin.LoginPassword
// 	if oldPwd == "" {
// 		err = NoLoginPasswordError{errors.New("LoginPassword is nil!")}
// 		log.Error(err)
// 		return
// 	}

// 	if !passwordEncoder.Matches(password.OldPwd, oldPwd) {
// 		err = ChangePasswordError{errors.New("password is wrong!")}
// 		log.Error(err)
// 		return
// 	}

// 	err = adminDao.Update(userid, "login_password", passwordEncoder.Encode(password.NewPwd))
// 	if err != nil {
// 		return
// 	}
// 	return
// }
