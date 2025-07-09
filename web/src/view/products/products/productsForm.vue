<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="TAG:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true" placeholder="请输入TAG" />
        </el-form-item>
        <el-form-item label="链接:" prop="url">
          <el-input v-model="formData.url" :clearable="true" placeholder="请输入链接" />
        </el-form-item>
        <el-form-item label="爬虫类型:" prop="billingType">
          <el-select v-model="formData.billingType" placeholder="请选择爬虫类型" clearable>
            <el-option label="single" value="single"></el-option>
            <el-option label="multi" value="multi"></el-option>
            <el-option label="single_cf5s" value="single_cf5s"></el-option>
            <el-option label="multi_cf5s" value="multi_cf5s"></el-option>
            <el-option label="single_dynamic" value="single_dynamic"></el-option>
            <el-option label="multi_dynamic" value="multi_dynamic"></el-option>
            <el-option label="single_only_price" value="single_only_price"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="爬虫间隔:" prop="intervals">
          <el-input v-model.number="formData.intervals" type="number" :clearable="true"
            placeholder="请输入爬虫间隔，以秒为单位，写-1则不爬虫" />
        </el-form-item>
        <el-form-item label="推送间隔:" prop="pushIntervals">
          <el-input v-model.number="formData.pushIntervals" type="number" :clearable="true"
            placeholder="请输入推送间隔，以天为单位，写-1则不推送" />
        </el-form-item>
        <el-form-item label="其他:" prop="additional">
          <RichEdit v-model="formData.additional" />
        </el-form-item>
        <el-form-item label="消息编号:" prop="messageId">
          <el-input v-model="formData.messageId" :clearable="true" placeholder="请输入消息编号" />
        </el-form-item>
        <el-form-item label="现有库存:" prop="stock">
          <el-input v-model.number="formData.stock" type="number" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="历史库存:" prop="oldStock">
          <el-input v-model.number="formData.oldStock" type="number" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="推送库存:" prop="pushStock">
          <el-input v-model.number="formData.pushStock" type="number" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="重复检测次数:" prop="multiCheck">
          <el-input v-model.number="formData.multiCheck" type="number" :clearable="true" placeholder="请输入重复检测的次数" />
        </el-form-item>
        <el-form-item label="推送时间:" prop="pushTime">
          <el-date-picker v-model="formData.pushTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="CPU:" prop="cpu">
          <el-input v-model="formData.cpu" :clearable="true" placeholder="请输入CPU" />
        </el-form-item>
        <el-form-item label="内存:" prop="memory">
          <el-input v-model="formData.memory" :clearable="true" placeholder="请输入内存" />
        </el-form-item>
        <el-form-item label="磁盘:" prop="disk">
          <el-input v-model="formData.disk" :clearable="true" placeholder="请输入磁盘" />
        </el-form-item>
        <el-form-item label="流量:" prop="traffic">
          <el-input v-model="formData.traffic" :clearable="true" placeholder="请输入流量" />
        </el-form-item>
        <el-form-item label="端口:" prop="portSpeed">
          <el-input v-model="formData.portSpeed" :clearable="true" placeholder="请输入端口" />
        </el-form-item>
        <el-form-item label="地点:" prop="location">
          <el-input v-model="formData.location" :clearable="true" placeholder="请输入地点" />
        </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input v-model="formData.price" :clearable="true" placeholder="请输入价格" />
        </el-form-item>
        <!-- <el-form-item label="创建者:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新者:" prop="updatedBy">
          <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="删除者:" prop="deletedBy">
          <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入" />
       </el-form-item> -->
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
  createProducts,
  updateProducts,
  findProducts
} from '@/api/products/products'

defineOptions({
  name: 'ProductsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  tag: '',
  cpu: '',
  memory: '',
  disk: '',
  traffic: '',
  portSpeed: '',
  location: '',
  price: '',
  additional: '',
  url: '',
  billingType: '',
  pushStock: 0,
  oldStock: 0,
  stock: 0,
  multiCheck: 0,
  intervals: 0,
  messageId: '',
  pushIntervals: 0,
  pushTime: new Date(),
  createdBy: undefined,
  updatedBy: undefined,
  deletedBy: undefined,
})
// 验证规则
const rule = reactive({
  tag: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  url: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  stock: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  billingType: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  multiCheck: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  intervals: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  pushIntervals: [{
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
    const res = await findProducts({ ID: route.query.id })
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
        res = await createProducts(formData.value)
        break
      case 'update':
        res = await updateProducts(formData.value)
        break
      default:
        res = await createProducts(formData.value)
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