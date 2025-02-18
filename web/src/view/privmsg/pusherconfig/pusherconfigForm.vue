<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="推送类型:" prop="pushType">
          <el-select v-model="formData.pushType" placeholder="请输入推送类型">
            <el-option label="Telegram" value="telegram_bot"></el-option>
            <el-option label="Email" value="email"></el-option>
            <el-option label="微信" value="wechat"></el-option>
            <el-option label="QQ" value="qq"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="配置名称:" prop="configName">
          <el-input v-model="formData.configName" :clearable="true" placeholder="请输入配置名称" />
        </el-form-item>
        <el-form-item label="配置数值:" prop="configValue">
          <el-input v-model="formData.configValue" :clearable="true" placeholder="请输入配置数值" />
        </el-form-item>
        <el-form-item label="最大字符数:" prop="maxCharactersPerPush">
          <el-input v-model.number="formData.maxCharactersPerPush" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="最大推送次数:" prop="maxPushesPerInterval">
          <el-input v-model.number="formData.maxPushesPerInterval" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="使用间隔:" prop="intervalSeconds">
          <el-input v-model.number="formData.intervalSeconds" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="已推送次数:" prop="currentIntervalPushes">
          <el-input v-model.number="formData.currentIntervalPushes" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="最后一次推送时间:" prop="lastPushTime">
          <el-date-picker v-model="formData.lastPushTime" type="date" placeholder="选择日期"
            :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="配置激活:" prop="isActive">
          <el-switch v-model="formData.isActive" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
            inactive-text="否" clearable></el-switch>
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
  createPusherConfig,
  updatePusherConfig,
  findPusherConfig
} from '@/api/privmsg/pusherconfig'

defineOptions({
  name: 'PusherConfigForm'
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
  pushType: '',
  configName: '',
  configValue: '',
  maxCharactersPerPush: undefined,
  maxPushesPerInterval: undefined,
  intervalSeconds: undefined,
  currentIntervalPushes: undefined,
  lastPushTime: new Date(),
  isActive: false,
})
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findPusherConfig({ ID: route.query.id })
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
        res = await createPusherConfig(formData.value)
        break
      case 'update':
        res = await updatePusherConfig(formData.value)
        break
      default:
        res = await createPusherConfig(formData.value)
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
