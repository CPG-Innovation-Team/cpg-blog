package vo

/**
  @author: chenxi@cpgroup.cn
  @date:2022/4/7
  @description:
**/

type UserRolesVO struct {
	UserId    int      `json:"userId"`
	UserName  string   `json:"userName"`
	RoleNames []string `json:"roleNames"`
}
