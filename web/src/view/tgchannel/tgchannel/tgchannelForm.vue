<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="频道地址:" prop="channel">
          <el-input v-model="formData.channel" :clearable="true" placeholder="请输入频道地址" />
        </el-form-item>
        <el-form-item label="频道ID:" prop="tgid">
          <el-input v-model="formData.tgid" :clearable="true" placeholder="请输入频道ID" />
        </el-form-item>
        <el-form-item label="推送类型:" prop="flag">
          <el-select v-model="formData.flag" placeholder="请选择推送类型标志">
            <el-option label="ONLY" value="ONLY"></el-option>
            <el-option label="ALL WITHOUT" value="ALL WITHOUT"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="Tokens:" prop="tokens">
          <el-input v-model="formData.tokens" :clearable="true" placeholder="请输入Tokens" />
        </el-form-item>
        <el-form-item label="推送商家:" prop="tags">
          <RichEdit v-model="formData.tags" />
        </el-form-item>
        <el-form-item label="其他信息关键词:" prop="additionalKey">
          <el-input v-model="formData.additionalKey" :clearable="true" placeholder="请输入其他信息关键词" />
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
  createTgchannel,
  updateTgchannel,
  findTgchannel
} from '@/api/tgchannel/tgchannel'

defineOptions({
  name: 'TgchannelForm'
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
  channel: '',
  tgid: '',
  flag: '',
  tokens: '',
  tags: '',
  additionalKey: '',
  createdBy: undefined,
  updatedBy: undefined,
  deletedBy: undefined,
})
// 验证规则
const rule = reactive({
  channel: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  tgid: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  flag: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  tokens: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  tags: [{
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
    const res = await findTgchannel({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()

// 处理富文本内容的函数
const processTags = (tags) => {
  return tags
    .replace(/<[^>]+>/g, '') // 移除所有HTML标签
    .trim() // 去掉前后空格
    .replace(/\s+/g, ' '); // 将多个空格替换为单个空格
}

// 保存按钮
const save = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    
    // 处理推送商家字段
    formData.value.tags = processTags(formData.value.tags);
    
    let res
    switch (type.value) {
      case 'create':
        res = await createTgchannel(formData.value)
        break
      case 'update':
        res = await updateTgchannel(formData.value)
        break
      default:
        res = await createTgchannel(formData.value)
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