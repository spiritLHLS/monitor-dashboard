
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="tag字段:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true"  placeholder="请输入tag字段" />
       </el-form-item>
        <el-form-item label="cpu字段:" prop="cpu">
          <el-input v-model="formData.cpu" :clearable="true"  placeholder="请输入cpu字段" />
       </el-form-item>
        <el-form-item label="memory字段:" prop="memory">
          <el-input v-model="formData.memory" :clearable="true"  placeholder="请输入memory字段" />
       </el-form-item>
        <el-form-item label="disk字段:" prop="disk">
          <el-input v-model="formData.disk" :clearable="true"  placeholder="请输入disk字段" />
       </el-form-item>
        <el-form-item label="traffic字段:" prop="traffic">
          <el-input v-model="formData.traffic" :clearable="true"  placeholder="请输入traffic字段" />
       </el-form-item>
        <el-form-item label="portSpeed字段:" prop="portSpeed">
          <el-input v-model="formData.portSpeed" :clearable="true"  placeholder="请输入portSpeed字段" />
       </el-form-item>
        <el-form-item label="location字段:" prop="location">
          <el-input v-model="formData.location" :clearable="true"  placeholder="请输入location字段" />
       </el-form-item>
        <el-form-item label="price字段:" prop="price">
          <el-input v-model="formData.price" :clearable="true"  placeholder="请输入price字段" />
       </el-form-item>
        <el-form-item label="additional字段:" prop="additional">
          <RichEdit v-model="formData.additional"/>
       </el-form-item>
        <el-form-item label="url字段:" prop="url">
          <el-input v-model="formData.url" :clearable="true"  placeholder="请输入url字段" />
       </el-form-item>
        <el-form-item label="billingType字段:" prop="billingType">
          <el-input v-model="formData.billingType" :clearable="true"  placeholder="请输入billingType字段" />
       </el-form-item>
        <el-form-item label="oldStock字段:" prop="oldStock">
          <el-input v-model.number="formData.oldStock" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="stock字段:" prop="stock">
          <el-input v-model.number="formData.stock" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否3次检测再推送:" prop="multiCheck">
          <el-input v-model.number="formData.multiCheck" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="间隔为-1时不进行爬虫:" prop="intervals">
          <el-input v-model.number="formData.intervals" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="自动填充-帖子ID:" prop="messageId">
          <el-input v-model="formData.messageId" :clearable="true"  placeholder="请输入自动填充-帖子ID" />
       </el-form-item>
        <el-form-item label="间隔为-1时不进行推送:" prop="pushIntervals">
          <el-input v-model.number="formData.pushIntervals" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="自动填充-最新推送的时间:" prop="pushTime">
          <el-date-picker v-model="formData.pushTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="创建者:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新者:" prop="updatedBy">
          <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="删除者:" prop="deletedBy">
          <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入" />
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
            oldStock: undefined,
            stock: undefined,
            multiCheck: undefined,
            intervals: undefined,
            messageId: '',
            pushIntervals: undefined,
            pushTime: new Date(),
            createdBy: undefined,
            updatedBy: undefined,
            deletedBy: undefined,
        })
// 验证规则
const rule = reactive({
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
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
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
</style>
