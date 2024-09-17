<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="Tag:" prop="tgTag">
          <el-input v-model="formData.tgTag" :clearable="true" placeholder="请输入Tag" />
        </el-form-item>
        <el-form-item label="分区名字:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入分区名字" />
        </el-form-item>
        <el-form-item label="分区链接:" prop="link">
          <el-input v-model="formData.link" :clearable="true" placeholder="请输入分区链接" />
        </el-form-item>
        <el-form-item label="分区类型:" prop="type">
          <el-select v-model="formData.type" placeholder="请选择分区类型" clearable>
            <el-option label="whmcs1" value="whmcs1"></el-option>
            <el-option label="whmcs2" value="whmcs2"></el-option>
            <el-option label="whmcs3" value="whmcs3"></el-option>
            <el-option label="whmcs4" value="whmcs4"></el-option>
            <el-option label="hostbill1" value="hostbill1"></el-option>
            <el-option label="blesta1" value="blesta1"></el-option>
            <el-option label="blesta2" value="blesta2"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="识别数量:" prop="num">
          <el-input v-model.number="formData.num" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="爬虫间隔:" prop="intervals">
          <el-input v-model.number="formData.intervals" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="其他信息:" prop="additional">
          <RichEdit v-model="formData.additional" />
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
  createPartitionspage,
  updatePartitionspage,
  findPartitionspage
} from '@/api/partitions/partitionspage'

defineOptions({
  name: 'PartitionspageForm'
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
  tgTag: '',
  name: '',
  link: '',
  type: '',
  num: undefined,
  additional: '',
  intervals: undefined,
  createdBy: undefined,
  updatedBy: undefined,
  deletedBy: undefined,
})
// 验证规则
const rule = reactive({
  tg_tag: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  link: [{
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
    const res = await findPartitionspage({ ID: route.query.id })
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
        res = await createPartitionspage(formData.value)
        break
      case 'update':
        res = await updatePartitionspage(formData.value)
        break
      default:
        res = await createPartitionspage(formData.value)
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