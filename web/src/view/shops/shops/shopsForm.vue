<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="频道TAG:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true" placeholder="请输入频道TAG" />
        </el-form-item>
        <el-form-item label="消息编号:" prop="mid">
          <el-input v-model="formData.mid" :clearable="true" placeholder="请输入消息编号" />
        </el-form-item>
        <el-form-item label="商家类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="请输入商家类型" />
        </el-form-item>
        <el-form-item label="推广链接:" prop="affLink">
          <el-input v-model="formData.affLink" :clearable="true" placeholder="请输入推广链接" />
        </el-form-item>
        <el-form-item label="商家类型:" prop="type">
          <el-select v-model="formData.type" placeholder="请选择商家类型" clearable>
            <el-option label="whmcs" value="whmcs"></el-option>
            <el-option label="hostbill" value="hostbill"></el-option>
            <el-option label="blesta" value="blesta"></el-option>
            <el-option label="other" value="other"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="额外标签:" prop="additionalTags">
          <el-input v-model="formData.additionalTags" :clearable="true" placeholder="请输入额外标签" />
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
  createShops,
  updateShops,
  findShops
} from '@/api/shops/shops'

defineOptions({
  name: 'ShopsForm'
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
  mid: '',
  type: '',
  affLink: '',
  shopLink: '',
  additionalTags: '',
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
  type: [{
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
    const res = await findShops({ ID: route.query.id })
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
        res = await createShops(formData.value)
        break
      case 'update':
        res = await updateShops(formData.value)
        break
      default:
        res = await createShops(formData.value)
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