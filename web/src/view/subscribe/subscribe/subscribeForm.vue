<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户UUID:" prop="user_uuid">
          <el-input v-model.number="formData.user_uuid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品ID:" prop="product_id">
          <el-input v-model.number="formData.product_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="订阅状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="通知渠道:" prop="notify_channel">
          <el-input v-model="formData.notify_channel" :clearable="true" placeholder="请输入通知渠道" />
        </el-form-item>
        <el-form-item label="最后更新:" prop="last_update">
          <el-date-picker v-model="formData.last_update" type="date" placeholder="选择日期"
            :clearable="true"></el-date-picker>
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
  createSubscribe,
  updateSubscribe,
  findSubscribe
} from '@/api/subscribe/subscribe'

defineOptions({
  name: 'SubscribeForm'
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
  user_uuid: undefined,
  product_id: undefined,
  status: undefined,
  notify_channel: '',
  last_update: new Date(),
})
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findSubscribe({ ID: route.query.id })
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
        res = await createSubscribe(formData.value)
        break
      case 'update':
        res = await updateSubscribe(formData.value)
        break
      default:
        res = await createSubscribe(formData.value)
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
.row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 1px;
  /* 行与行之间的间距 */
}

.row .el-form-item {
  flex: 1;
  /* 使每个表单项占据相等的空间 */
  margin-right: 10px;
  /* 表单项之间的间距 */
}

.el-table .cell {
  /* 文本超出容器部分隐藏 */
  overflow: hidden;
  /* 超出部分使用省略号代替 */
  text-overflow: ellipsis;
  /* 不换行 */
  white-space: nowrap;
}

.demo-form-inline .el-input {
  width: 160px;
  /* 输入框的宽度 */
}
</style>