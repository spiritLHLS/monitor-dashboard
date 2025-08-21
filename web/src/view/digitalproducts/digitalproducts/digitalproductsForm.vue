
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="TAG:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true"  placeholder="请输入TAG" />
       </el-form-item>
        <el-form-item label="核心:" prop="cpu">
          <el-input 
            v-model="formData.cpu" 
            type="number" 
            step="0.001"
            :clearable="true" 
            placeholder="请输入"
            @input="handleNumberInput('cpu', $event)"
          />
        </el-form-item>
        <el-form-item label="内存:" prop="memory">
          <el-input 
            v-model="formData.memory" 
            type="number" 
            step="0.001"
            :clearable="true" 
            placeholder="请输入"
            @input="handleNumberInput('memory', $event)"
          />
        </el-form-item>
        <el-form-item label="硬盘:" prop="disk">
          <el-input 
            v-model="formData.disk" 
            type="number" 
            step="0.001"
            :clearable="true" 
            placeholder="请输入"
            @input="handleNumberInput('disk', $event)"
          />
        </el-form-item>
        <el-form-item label="流量:" prop="traffic">
          <el-input 
            v-model="formData.traffic" 
            type="number" 
            step="0.001"
            :clearable="true" 
            placeholder="请输入"
            @input="handleNumberInput('traffic', $event)"
          />
        </el-form-item>
        <el-form-item label="带宽:" prop="portSpeed">
          <el-input 
            v-model="formData.portSpeed" 
            type="number" 
            step="0.001"
            :clearable="true" 
            placeholder="请输入"
            @input="handleNumberInput('portSpeed', $event)"
          />
        </el-form-item>
        <el-form-item label="位置:" prop="location">
          <el-input v-model="formData.location" :clearable="true"  placeholder="请输入位置" />
       </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input 
            v-model="formData.price" 
            type="number" 
            step="0.001"
            :clearable="true" 
            placeholder="请输入"
            @input="handleNumberInput('price', $event)"
          />
        </el-form-item>
        <el-form-item label="价格单位:" prop="priceUnit">
          <el-input v-model="formData.priceUnit" :clearable="true"  placeholder="请输入价格单位" />
       </el-form-item>
        <el-form-item label="其他:" prop="additional">
          <el-input v-model="formData.additional" :clearable="true"  placeholder="请输入其他" />
       </el-form-item>
        <el-form-item label="原表ID:" prop="originId">
          <el-input v-model.number="formData.originId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createDigitalProducts,
  updateDigitalProducts,
  findDigitalProducts
} from '@/api/digitalproducts/digitalproducts'

defineOptions({
    name: 'DigitalProductsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            tag: '',
            cpu: undefined,
            memory: undefined,
            disk: undefined,
            traffic: undefined,
            portSpeed: undefined,
            location: '',
            price: undefined,
            priceUnit: '',
            additional: '',
            originId: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDigitalProducts({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

// 添加数字输入处理方法
const handleNumberInput = (prop, value) => {
  if (value === '' || value === null) {
    formData.value[prop] = undefined
  } else {
    const numValue = parseFloat(value)
    formData.value[prop] = isNaN(numValue) ? undefined : numValue
  }
}

// 批量编辑的数字输入处理方法
const handleBatchNumberInput = (prop, value) => {
  if (value === '' || value === null) {
    batchEditData.value[prop] = undefined
  } else {
    const numValue = parseFloat(value)
    batchEditData.value[prop] = isNaN(numValue) ? undefined : numValue
  }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createDigitalProducts(formData.value)
               break
             case 'update':
               res = await updateDigitalProducts(formData.value)
               break
             default:
               res = await createDigitalProducts(formData.value)
               break
           }
           btnLoading.value = false
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
