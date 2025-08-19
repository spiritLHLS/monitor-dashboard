
// 自动生成模板AgentIpRecord
package claweragent
import (
	"server/global"
)

// IP记录 结构体  AgentIpRecord
type AgentIpRecord struct {
    global.GVA_MODEL
    AgentIp  *string `json:"agentIp" form:"agentIp" gorm:"column:agent_ip;comment:代理IP地址;size:25;" binding:"required"`  //代理IP
    Remark  *string `json:"remark" form:"remark" gorm:"column:remark;comment:备注信息;size:350;"`  //备注
    Frozen  *bool `json:"frozen" form:"frozen" gorm:"column:frozen;comment:是否冻结使用;"`  //冻结
}


// TableName IP记录 AgentIpRecord自定义表名 agent_ip_record
func (AgentIpRecord) TableName() string {
    return "agent_ip_record"
}





