<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true" placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item label="昵称:" prop="nickname">
          <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入昵称" />
        </el-form-item>
        <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true" placeholder="请输入密码" />
        </el-form-item>
        <el-form-item label="头像:" prop="avatar">
          <SelectImage v-model="formData.avatar" file-type="image" />
        </el-form-item>
        <el-form-item label="是否冻结:" prop="isFrozen">
          <el-switch v-model="formData.isFrozen" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
            inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item label="推送渠道1:" prop="pushChannel1">
          <el-input v-model="formData.pushChannel1" :clearable="true" placeholder="请输入推送渠道1" />
        </el-form-item>
        <el-form-item label="推送渠道2:" prop="pushChannel2">
          <el-input v-model="formData.pushChannel2" :clearable="true" placeholder="请输入推送渠道2" />
        </el-form-item>
        <el-form-item label="推送渠道3:" prop="pushChannel3">
          <el-input v-model="formData.pushChannel3" :clearable="true" placeholder="请输入推送渠道3" />
        </el-form-item>
        <el-form-item label="TGID:" prop="tgID">
          <el-input v-model="formData.tgID" :clearable="true" placeholder="请输入TGID" />
        </el-form-item>
        <el-form-item label="QQ号:" prop="qqNumber">
          <el-input v-model="formData.qqNumber" :clearable="true" placeholder="请输入QQ号" />
        </el-form-item>
        <el-form-item label="微信号:" prop="weChatNumber">
          <el-input v-model="formData.weChatNumber" :clearable="true" placeholder="请输入微信号" />
        </el-form-item>
        <el-form-item label="邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="备注:" prop="additional">
          <el-input v-model="formData.additional" :clearable="true" placeholder="请输入备注" />
        </el-form-item>
        <el-form-item label="用户等级:" prop="level">
          <el-input v-model.number="formData.level" :clearable="false" placeholder="请输入" />
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
  createEcsUsers,
  updateEcsUsers,
  findEcsUsers
} from '@/api/ecsusers/ecsusers'

defineOptions({
  name: 'EcsUsersForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  username: '',
  nickname: '',
  avatar: "",
  password: '',
  isFrozen: false,
  pushChannel1: '',
  pushChannel2: '',
  pushChannel3: '',
  tgID: '',
  qqNumber: '',
  weChatNumber: '',
  email: '',
  additional: '',
  level: undefined,
})
// 验证规则
const rule = reactive({
  username: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  nickname: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  password: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  tgID: [{
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
    const res = await findEcsUsers({ ID: route.query.id })
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
        res = await createEcsUsers(formData.value)
        break
      case 'update':
        res = await updateEcsUsers(formData.value)
        break
      default:
        res = await createEcsUsers(formData.value)
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

<style></style>
