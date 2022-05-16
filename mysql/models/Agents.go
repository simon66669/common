package models

type Agents struct {
	Id            string  `json:"id"`              // 代理商表
	UserId        string  `json:"user_id"`         // 关联用户id
	Username      string  `json:"username"`        // 登录代理商后台的帐号
	Password      string  `json:"password"`        // 登录代理商后台的密码
	ParentAgentId string  `json:"parent_agent_id"` // 父级代理商ID，0表示该用户是超级代理商
	Level         string  `json:"level"`           // 代理商等级,0:超级管理员； 1：一级代理商；2:二级代理商；3:三级代理商；4:四级代理商
	AgentPath     string  `json:"agent_path"`      // 代理商关系，用,拼接成字符串
	IsAdmin       int     `json:"is_admin"`        // 是否为超级管理员
	IsLock        int     `json:"is_lock"`         // 该代理商是否锁定
	IsAddson      int     `json:"is_addson"`       // 是否拥有开设下级代理商的权限
	ProLoss       float32 `json:"pro_loss"`        // 头寸比例
	ProSer        float32 `json:"pro_ser"`         // 手续费比例
	RegTime       int     `json:"reg_time"`        // 代理商注册时间
	LockTime      int     `json:"lock_time"`       // 代理商锁定时间
}
