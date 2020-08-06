package service

// func TestCreateAdmin(t *testing.T) {
// 	userService := NewUserService()
// 	err := userService.CreateAdmin(model.GlobalHandler{}, model.Admin{
// 		//用户类型
// 		UserType: model.UserTypeAdmin,
// 		//帐号
// 		Account: "testadmin",
// 		//工号
// 		JobNumber: "001",
// 		//登录密码
// 		LoginPassword: "111111",
// 		//安全密码
// 		SecurePassword: "111111",
// 		//部门
// 		Department: "test department",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

// func TestFindAdmin(t *testing.T) {
// 	userService := NewUserService()
// 	result, _, _ := userService.FindAdminListPaging(model.SearchAdminCondition{
// 		Account:    "alan",
// 		RoleID:     1,
// 		Department: "test",
// 		State:      model.AccountEnable,
// 	}, 1, 20)
// 	t.Log(result)
// }
