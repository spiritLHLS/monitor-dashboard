<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="邀请码:" prop="code">
          <el-input v-model="formData.code" :clearable="false" placeholder="请输入邀请码" />
        </el-form-item>
        <el-form-item label="使用状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="过期时间:" prop="expires_at">
          <el-date-picker v-model="formData.expires_at" type="date" placeholder="选择日期"
            :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="使用时间:" prop="used_at">
          <el-date-picker v-model="formData.used_at" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="使用者:" prop="user_uuid">
          <el-input v-model="formData.user_uuid" :clearable="true" placeholder="请输入使用者" />
        </el-form-item>
        <el-form-item label="备注:" prop="remarks">
          <el-input v-model="formData.remarks" :clearable="true" placeholder="请输入备注" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createInviteCodes,
  updateInviteCodes,
  findInviteCodes
} from '@/api/invite_codes/inviteCodes'

defineOptions({
  name: 'InviteCodesForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  code: '',
  status: undefined,
  expires_at: new Date(),
  used_at: new Date(),
  user_uuid: '',
  remarks: '',
})
// 验证规则
const rule = reactive({
  code: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findInviteCodes({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createInviteCodes(formData.value)
        break
      case 'update':
        res = await updateInviteCodes(formData.value)
        break
      default:
        res = await createInviteCodes(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>


<style>
.demo-form-inline .el-input {
  width: 130px;
  /* 输入框的宽度 */
}
</style>