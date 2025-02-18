<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon>
                  <QuestionFilled />
                </el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期"
            :disabled-date="time => searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false">
          </el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
            :disabled-date="time => searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false">
          </el-date-picker>
        </el-form-item>
        <template v-if="showAllQuery">
          <div class="row">
            <el-form-item v-for="field in searchFields" :key="field.prop" :label="field.label" :prop="field.prop">
              <el-input v-model="searchInfo[field.prop]" :placeholder="field.placeholder || '搜索条件'"
                :type="field.type || 'text'" />
            </el-form-item>
          </div>
        </template>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery = true"
            v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery = false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="onDelete">删除</el-button>
        <el-button type="primary" icon="edit" :disabled="!multipleSelection.length"
          @click="openBatchEdit">批量修改</el-button>
        <el-button type="warning" icon="delete" :disabled="!multipleSelection.length"
          @click="openBatchClear">批量清空</el-button>
      </div>
      <!-- 批量修改抽屉 -->
      <el-drawer v-model="dialogs.batchEdit.visible" :size="800" destroy-on-close title="批量修改">
        <template #header>
          <div class="flex justify-between items-center">
            <span class="text-lg">批量修改</span>
            <div>
              <el-button type="primary" @click="enterBatchEdit">确 定</el-button>
              <el-button @click="closeBatchEdit">取 消</el-button>
            </div>
          </div>
        </template>
        <el-form :model="batchEditForm" label-position="top" label-width="80px">
          <el-form-item v-for="field in formFields" :key="field.prop" :label="field.label" :prop="field.prop">
            <el-input v-model="batchEditForm[field.prop]" :placeholder="`留空则不修改${field.label}`" clearable />
          </el-form-item>
        </el-form>
      </el-drawer>
      <!-- 批量清空抽屉 -->
      <el-drawer v-model="dialogs.batchClear.visible" :size="800" destroy-on-close title="批量清空字段">
        <template #header>
          <div class="flex justify-between items-center">
            <span class="text-lg">批量清空字段</span>
            <div>
              <el-button type="primary" @click="enterBatchClear">确 定</el-button>
              <el-button @click="closeBatchClear">取 消</el-button>
            </div>
          </div>
        </template>
        <el-form>
          <el-form-item label="选择要清空的字段">
            <el-checkbox-group v-model="batchClearFields">
              <el-checkbox v-for="field in formFields" :key="field.prop" :label="field.prop">
                {{ field.label }}
              </el-checkbox>
            </el-checkbox-group>
          </el-form-item>
        </el-form>
      </el-drawer>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column v-for="col in tableColumns" :key="col.prop" :prop="col.prop" :label="col.label"
          :width="col.width || 90" :align="col.align || 'left'">
          <template #default="scope" v-if="col.formatter">
            {{ col.formatter(scope.row[col.prop]) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>查看详情
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updatePartitionspageFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="page" :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]" :total="total" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <!-- 表单抽屉 -->
    <el-drawer destroy-on-close size="800" v-model="dialogs.form.visible" :show-close="false"
      :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '添加' : '修改' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>
      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item v-for="field in formFields" :key="field.prop" :label="field.label" :prop="field.prop">
          <template v-if="field.type === 'rich-edit'">
            <RichEdit v-model="formData[field.prop]" />
          </template>
          <template v-else-if="field.type === 'select'">
            <el-select v-model="formData[field.prop]" :placeholder="field.placeholder"
              :clearable="field.clearable !== false">
              <el-option v-for="opt in field.options" :key="opt.value" :label="opt.label" :value="opt.value" />
            </el-select>
          </template>
          <el-input v-else v-model="formData[field.prop]" :type="field.inputType || 'text'"
            :placeholder="field.placeholder || `请输入${field.label}`" :clearable="field.clearable !== false" />
        </el-form-item>
      </el-form>
    </el-drawer>
    <!-- 详情抽屉 -->
    <el-drawer destroy-on-close size="800" v-model="dialogs.detail.visible" :show-close="true"
      :before-close="closeDetailShow">
      <el-descriptions :column="1" border>
        <el-descriptions-item v-for="field in formFields" :key="field.prop" :label="field.label">
          <template v-if="field.type === 'rich-view'">
            <RichView v-model="detailFrom[field.prop]" />
          </template>
          <template v-else>
            {{ detailFrom[field.prop] }}
          </template>
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createPartitionspage,
  deletePartitionspage,
  deletePartitionspageByIds,
  updatePartitionspage,
  findPartitionspage,
  getPartitionspageList
} from '@/api/partitions/partitionspage'
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichView from '@/components/richtext/rich-view.vue'
import { getDictFunc, formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'Partitionspage'
})

// 配置定义
const searchFields = [
  { label: 'Tag', prop: 'tg_tag' },
  { label: '分区名字', prop: 'name' },
  { label: '分区链接', prop: 'link' },
  { label: '分区类型', prop: 'type' },
  { label: '识别数量', prop: 'num', type: 'number' },
  { label: '爬虫间隔', prop: 'intervals', type: 'number' },
  { label: '其他信息', prop: 'additional' }
]

const tableColumns = [
  { prop: 'tgTag', label: 'Tag' },
  { prop: 'name', label: '分区名字' },
  { prop: 'link', label: '分区链接' },
  { prop: 'type', label: '分区类型' },
  { prop: 'num', label: '识别数量' },
  { prop: 'intervals', label: '爬虫间隔' },
  { prop: 'additional', label: '其他信息', width: 120 },
  { prop: 'CreatedAt', label: '日期', width: 180, formatter: formatDate }
]

const formFields = [
  { label: 'Tag:', prop: 'tgTag', required: true },
  { label: '分区名字:', prop: 'name', required: true },
  {
    label: '分区类型:', prop: 'type', type: 'select', required: true,
    options: [
      { label: 'whmcs1', value: 'whmcs1' },
      { label: 'whmcs2', value: 'whmcs2' },
      { label: 'whmcs3', value: 'whmcs3' },
      { label: 'whmcs4', value: 'whmcs4' },
      { label: 'hostbill1', value: 'hostbill1' },
      { label: 'blesta1', value: 'blesta1' },
      { label: 'blesta2', value: 'blesta2' }
    ]
  },
  { label: '分区链接:', prop: 'link', required: true },
  { label: '识别数量:', prop: 'num', inputType: 'number' },
  { label: '爬虫间隔:', prop: 'intervals', inputType: 'number' },
  { label: '其他信息:', prop: 'additional', type: 'rich-edit' }
]

// 状态定义
const showAllQuery = ref(false)
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const type = ref('')
const multipleSelection = ref([])
const dialogs = reactive({
  form: { visible: false },
  detail: { visible: false },
  batchEdit: { visible: false },
  batchClear: { visible: false }
})
const batchEditForm = ref({})
const batchClearFields = ref([])
const formData = ref({})
const detailFrom = ref({})

// 引用
const elFormRef = ref()
const elSearchFormRef = ref()

// 验证规则
const rule = reactive({
  tgTag: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }],
  name: [{
    required: true,
    message: '请输入分区名字',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }],
  link: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }],
  type: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }]
})

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deletePartitionspageFunc(row)
  })
}

// 多选删除
const onDelete = async () => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const IDs = []
    if (multipleSelection.value.length === 0) {
      ElMessage({
        type: 'warning',
        message: '请选择要删除的数据'
      })
      return
    }
    multipleSelection.value &&
      multipleSelection.value.map(item => {
        IDs.push(item.ID)
      })
    const res = await deletePartitionspageByIds({ IDs })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '删除成功'
      })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

// 更新行
const updatePartitionspageFunc = async (row) => {
  const res = await findPartitionspage({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogs.form.visible = true
  }
}

// 删除行
const deletePartitionspageFunc = async (row) => {
  const res = await deletePartitionspage({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogs.form.visible = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogs.form.visible = false
  formData.value = resetFormData()
}

// 弹窗确定
const enterDialog = async () => {
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
      closeDialog()
      getTableData()
    }
  })
}

// 查看详情
const getDetails = async (row) => {
  const res = await findPartitionspage({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    dialogs.detail.visible = true
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  dialogs.detail.visible = false
  detailFrom.value = {}
}

const searchRule = reactive({
  createdAt: [{
    validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt &&
        (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() ||
          searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change'
  }]
})

// 方法定义
const getTableData = async () => {
  const table = await getPartitionspageList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const resetFormData = () => ({
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

const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 批量操作相关方法
const openBatchEdit = () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要修改的数据')
    return
  }
  dialogs.batchEdit.visible = true
}

const closeBatchEdit = () => {
  dialogs.batchEdit.visible = false
  batchEditForm.value = {}
}

const enterBatchEdit = async () => {
  try {
    const updatePromises = multipleSelection.value.map(item => {
      const updateData = { ...item }
      Object.keys(batchEditForm.value).forEach(key => {
        if (batchEditForm.value[key] !== undefined &&
          batchEditForm.value[key] !== null &&
          batchEditForm.value[key] !== '') {
          updateData[key] = batchEditForm.value[key]
        }
      })
      return updatePartitionspage(updateData)
    })
    const results = await Promise.all(updatePromises)
    if (results.every(res => res.code === 0)) {
      ElMessage.success('批量修改成功')
      closeBatchEdit()
      getTableData()
    }
  } catch (error) {
    ElMessage.error('批量修改失败')
  }
}

const openBatchClear = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage.warning('请选择要操作的数据')
    return
  }
  dialogs.batchClear.visible = true
  batchClearFields.value = []
}

const closeBatchClear = () => {
  dialogs.batchClear.visible = false
  batchClearFields.value = []
}

const enterBatchClear = async () => {
  if (batchClearFields.value.length === 0) {
    ElMessage.warning('请选择要清空的字段')
    return
  }
  try {
    const updatePromises = multipleSelection.value.map(item => {
      const updateData = { ...item }
      batchClearFields.value.forEach(field => {
        updateData[field] = ''
      })
      return updatePartitionspage(updateData)
    })
    const results = await Promise.all(updatePromises)
    if (results.every(res => res.code === 0)) {
      ElMessage.success('批量清空成功')
      closeBatchClear()
      getTableData()
    }
  } catch (error) {
    ElMessage.error('批量清空失败')
  }
}

// 初始化
getTableData()
</script>

<style>
.row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 1px;
}

.row .el-form-item {
  flex: 1;
  margin-right: 10px;
}

.el-table .cell {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.demo-form-inline .el-input {
  width: 160px;
}

.gva-btn-list {
  display: flex;
  gap: 8px;
  margin-bottom: 1rem;
  flex-wrap: wrap;
}

.gva-pagination {
  margin-top: 1rem;
  display: flex;
  justify-content: flex-end;
}

@media (max-width: 768px) {
  .row {
    flex-direction: column;
  }

  .row .el-form-item {
    width: 100%;
  }

  .demo-form-inline .el-input {
    width: 100%;
  }

  .el-table .el-button {
    margin: 2px 0;
    display: block;
    width: 100%;
  }
}
</style>