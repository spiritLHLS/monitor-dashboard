<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="TAG:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true" placeholder="请输入TAG" />
        </el-form-item>
        <el-form-item label="标志:" prop="flag">
          <el-input v-model="formData.flag" :clearable="true" placeholder="请输入标志" />
        </el-form-item>
        <el-form-item label="计费类型:" prop="billingType">
          <el-input v-model="formData.billingType" :clearable="true" placeholder="请输入计费类型" />
        </el-form-item>
        <el-form-item label="起始编号:" prop="pdId">
          <el-input v-model.number="formData.pdId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="结束编号:" prop="endId">
          <el-input v-model.number="formData.endId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="CPU:" prop="cpu">
          <el-input v-model="formData.cpu" :clearable="true" placeholder="请输入CPU" />
        </el-form-item>
        <el-form-item label="内存:" prop="memory">
          <el-input v-model="formData.memory" :clearable="true" placeholder="请输入内存" />
        </el-form-item>
        <el-form-item label="硬盘:" prop="disk">
          <el-input v-model="formData.disk" :clearable="true" placeholder="请输入硬盘" />
        </el-form-item>
        <el-form-item label="流量:" prop="traffic">
          <el-input v-model="formData.traffic" :clearable="true" placeholder="请输入流量" />
        </el-form-item>
        <el-form-item label="带宽:" prop="portSpeed">
          <el-input v-model="formData.portSpeed" :clearable="true" placeholder="请输入带宽" />
        </el-form-item>
        <el-form-item label="地址:" prop="location">
          <el-input v-model="formData.location" :clearable="true" placeholder="请输入地址" />
        </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input v-model="formData.price" :clearable="true" placeholder="请输入价格" />
        </el-form-item>
        <el-form-item label="其他信息:" prop="additional">
          <el-input v-model="formData.additional" :clearable="true" placeholder="请输入其他信息" />
        </el-form-item>
        <el-form-item label="链接:" prop="url">
          <el-input v-model="formData.url" :clearable="true" placeholder="请输入链接" />
        </el-form-item>
        <el-form-item label="历史库存:" prop="oldStock">
          <el-input v-model.number="formData.oldStock" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="现有库存:" prop="stock">
          <el-input v-model.number="formData.stock" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="消息ID:" prop="messageId">
          <el-input v-model="formData.messageId" :clearable="true" placeholder="请输入消息ID" />
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
  createFindallpd,
  updateFindallpd,
  findFindallpd
} from '@/api/findallpd/findallpd'

defineOptions({
  name: 'FindallpdForm'
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
  tag: '',
  flag: '',
  billingType: '',
  pdId: undefined,
  endId: undefined,
  cpu: '',
  memory: '',
  disk: '',
  traffic: '',
  portSpeed: '',
  location: '',
  price: '',
  additional: '',
  url: '',
  oldStock: undefined,
  stock: undefined,
  messageId: '',
})
// 验证规则
const rule = reactive({
  tag: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  flag: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  billingType: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  endId: [{
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
    const res = await findFindallpd({ ID: route.query.id })
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
        res = await createFindallpd(formData.value)
        break
      case 'update':
        res = await updateFindallpd(formData.value)
        break
      default:
        res = await createFindallpd(formData.value)
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