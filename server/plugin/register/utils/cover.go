package utils

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ecsusers"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"reflect"
)

// ConvertEcsUserToSysUser 函数将 ecsusers.EcsUsers 类型转换为 system.SysUser 类型
func ConvertEcsUserToSysUser(ecsUser ecsusers.EcsUsers) system.SysUser {
	var sysUser system.SysUser                    // 创建一个 system.SysUser 类型的变量用于存储转换结果
	sourceValue := reflect.ValueOf(ecsUser)       // 获取源结构体(ecsUser)的反射值
	destValue := reflect.ValueOf(&sysUser).Elem() // 获取目标结构体(sysUser)的反射值
	// 遍历目标结构体的所有字段
	for i := 0; i < destValue.NumField(); i++ {
		destField := destValue.Type().Field(i)                 // 获取目标结构体的第i个字段
		sourceField := sourceValue.FieldByName(destField.Name) // 在源结构体中查找同名字段
		// 如果在源结构体中找到了同名字段,且该字段类型可以转换为目标字段类型
		if sourceField.IsValid() && sourceField.Type().ConvertibleTo(destField.Type) {
			destValue.Field(i).Set(sourceField.Convert(destField.Type)) // 将源字段的值转换并设置到目标字段
		}
	}
	return sysUser // 返回转换后的 system.SysUser 结构体
}
