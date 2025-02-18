<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="rules.search"
        @keyup.enter="onSubmit">
        <template v-if="showAllQuery">
          <div v-for="(fields, index) in searchFields" :key="index" class="search-row">
            <el-form-item v-for="field in fields" :key="field.prop" :label="field.label" :prop="field.prop">
              <el-input v-model="searchInfo[field.prop]" :placeholder="field.placeholder || '搜索条件'"
                :type="field.type || 'text'" />
            </el-form-item>
          </div>
          <el-form-item label="推送日期" prop="pushTime">
            <template #label>
              <span>
                推送日期
                <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                  <el-icon>
                    <QuestionFilled />
                  </el-icon>
                </el-tooltip>
              </span>
            </template>
            <el-date-picker v-model="searchInfo.startPushTime" type="datetime" placeholder="开始日期"
              :disabled-date="time => searchInfo.endPushTime ? time.getTime() > searchInfo.endPushTime.getTime() : false" />
            —
            <el-date-picker v-model="searchInfo.endPushTime" type="datetime" placeholder="结束日期"
              :disabled-date="time => searchInfo.startPushTime ? time.getTime() < searchInfo.startPushTime.getTime() : false" />
          </el-form-item>
        </template>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" :icon="showAllQuery ? 'arrow-up' : 'arrow-down'"
            @click="showAllQuery = !showAllQuery">
            {{ showAllQuery ? '收起' : '展开' }}
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <!-- 按钮列表区域修改 -->
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-button icon="delete" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        <el-button type="primary" icon="edit" :disabled="!multipleSelection.length"
          @click="openBatchEdit">批量修改</el-button>
        <el-button type="warning" icon="delete" :disabled="!multipleSelection.length"
          @click="openBatchClear">批量清空</el-button>
      </div>
      <!-- 添加批量修改抽屉 -->
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
          <el-form-item label="TAG" prop="tag">
            <el-input v-model="batchEditForm.tag" placeholder="留空则不修改" clearable />
          </el-form-item>
          <el-form-item label="推送库存" prop="pushStock">
            <el-input v-model.number="batchEditForm.pushStock" type="number" placeholder="留空则不修改" clearable />
          </el-form-item>
          <el-form-item label="历史库存" prop="oldStock">
            <el-input v-model.number="batchEditForm.oldStock" type="number" placeholder="留空则不修改" clearable />
          </el-form-item>
          <el-form-item label="现有库存" prop="stock">
            <el-input v-model.number="batchEditForm.stock" type="number" placeholder="留空则不修改" clearable />
          </el-form-item>
          <el-form-item label="爬虫类型" prop="billingType">
            <el-select v-model="batchEditForm.billingType" clearable placeholder="留空则不修改">
              <el-option label="single" value="single" />
              <el-option label="multi" value="multi" />
            </el-select>
          </el-form-item>
          <el-form-item label="爬取间隔" prop="intervals">
            <el-input v-model.number="batchEditForm.intervals" type="number" placeholder="留空则不修改" clearable />
          </el-form-item>
          <el-form-item label="推送间隔" prop="pushIntervals">
            <el-input v-model.number="batchEditForm.pushIntervals" type="number" placeholder="留空则不修改" clearable />
          </el-form-item>
          <el-form-item label="消息ID" prop="messageId">
            <el-input v-model="batchEditForm.messageId" placeholder="留空则不修改" clearable />
          </el-form-item>
          <el-form-item label="重复检测" prop="multiCheck">
            <el-input v-model.number="batchEditForm.multiCheck" type="number" placeholder="留空则不修改" clearable />
          </el-form-item>
          <el-form-item label="防御类型" prop="flag">
            <el-input v-model="batchEditForm.flag" placeholder="留空则不修改" clearable />
          </el-form-item>
          <el-form-item label="其他信息" prop="additional">
            <el-input v-model="batchEditForm.additional" placeholder="留空则不修改" clearable />
          </el-form-item>
        </el-form>
      </el-drawer>
      <el-table ref="multipleTable" style="width: 100%" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" fixed="left" />
        <el-table-column v-for="col in tableColumns" :key="col.prop" :prop="col.prop" :label="col.label"
          :width="col.width" :align="col.align || 'left'">
          <template #default="scope" v-if="col.formatter">
            {{ col.formatter(scope.row[col.prop]) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" width="240" fixed="right">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px">
                <InfoFilled />
              </el-icon>查看详情
            </el-button>
            <el-button type="primary" link icon="edit" @click="updateProductsFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
          layout="total, sizes, prev, pager, next, jumper" @current-change="handleCurrentChange"
          @size-change="handleSizeChange" />
      </div>
    </div>
    <!-- 表单抽屉 -->
    <el-drawer v-model="dialogs.form.visible" :size="800" destroy-on-close :title="type === 'create' ? '新增' : '编辑'"
      :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{ type === 'create' ? '新增' : '编辑' }}</span>
          <div>
            <el-button type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>
      <el-form :model="formData" ref="elFormRef" :rules="rules.form" label-position="top" label-width="80px">
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
            :placeholder="field.placeholder || `请输入${field.label.replace(':', '')}`"
            :clearable="field.clearable !== false" />
        </el-form-item>
      </el-form>
    </el-drawer>
    <!-- 详情抽屉 -->
    <el-drawer v-model="dialogs.detail.visible" :size="800" destroy-on-close title="详情">
      <el-descriptions :column="1" border>
        <el-descriptions-item v-for="field in detailFields" :key="field.prop" :label="field.label">
          <template v-if="field.type === 'rich-view'">
            <RichView v-model="detailFrom[field.prop]" />
          </template>
          <template v-else>
            {{ detailFrom[field.prop] }}
          </template>
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
    <!-- 添加批量清空抽屉 -->
    <el-drawer v-model="dialogs.batchClear.visible" :size="800" destroy-on-close title="批量清空字段"
      :before-close="closeBatchClear">
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
            <el-checkbox label="cpu">CPU</el-checkbox>
            <el-checkbox label="memory">内存</el-checkbox>
            <el-checkbox label="disk">磁盘</el-checkbox>
            <el-checkbox label="traffic">流量</el-checkbox>
            <el-checkbox label="portSpeed">端口</el-checkbox>
            <el-checkbox label="location">地点</el-checkbox>
            <el-checkbox label="price">价格</el-checkbox>
            <el-checkbox label="additional">其他信息</el-checkbox>
            <el-checkbox label="flag">防御类型</el-checkbox>
            <el-checkbox label="pushStock">推送库存</el-checkbox>
            <el-checkbox label="oldStock">历史库存</el-checkbox>
            <el-checkbox label="stock">现有库存</el-checkbox>
            <el-checkbox label="intervals">爬虫间隔</el-checkbox>
            <el-checkbox label="pushIntervals">推送间隔</el-checkbox>
            <el-checkbox label="multiCheck">重复检测</el-checkbox>
            <el-checkbox label="messageId">消息编号</el-checkbox>
          </el-checkbox-group>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  createProducts,
  deleteProducts,
  deleteProductsByIds,
  updateProducts,
  findProducts,
  getProductsList
} from '@/api/products/products'
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichView from '@/components/richtext/rich-view.vue'
import { formatDate } from '@/utils/format'

defineOptions({ name: 'Products' })
// 状态定义
const showAllQuery = ref(false)
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({
  startPushTime: undefined,
  endPushTime: undefined,
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
  pushStock: undefined,
  oldStock: undefined,
  stock: undefined,
  intervals: undefined,
  messageId: '',
  pushIntervals: undefined,
  multiCheck: undefined,
  flag: '',
})
const multipleSelection = ref([])
const type = ref('')
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
// 配置
const searchFields = [
  [
    { label: 'TAG', prop: 'tag' },
    { label: 'CPU', prop: 'cpu' },
    { label: '内存', prop: 'memory' },
    { label: '磁盘', prop: 'disk' },
    { label: '流量', prop: 'traffic' },
    { label: '端口', prop: 'portSpeed' },
  ],
  [
    { label: '地点', prop: 'location' },
    { label: '价格', prop: 'price' },
    { label: '链接', prop: 'url' },
    { label: '其他', prop: 'additional' },
    { label: 'flag', prop: 'flag' },
    { label: '爬虫类型', prop: 'billingType' },
  ],
  [
    { label: '推送库存', prop: 'pushStock', type: 'number' },
    { label: '历史库存', prop: 'oldStock', type: 'number' },
    { label: '现有库存', prop: 'stock', type: 'number' },
    { label: '爬虫间隔', prop: 'intervals', type: 'number' },
    { label: '消息编号', prop: 'messageId' }
  ]
]
const tableColumns = [
  { prop: 'tag', label: 'TAG', width: 90 },
  { prop: 'cpu', label: 'CPU', width: 80 },
  { prop: 'memory', label: '内存', width: 80 },
  { prop: 'disk', label: '磁盘', width: 80 },
  { prop: 'traffic', label: '流量', width: 80 },
  { prop: 'portSpeed', label: '端口', width: 80 },
  { prop: 'location', label: '地点', width: 80 },
  { prop: 'price', label: '价格', width: 80 },
  { prop: 'url', label: '链接', width: 90, className: 'url-column' },
  { prop: 'billingType', label: '爬虫类型', width: 90 },
  { prop: 'pushStock', label: '推送库存', width: 80 },
  { prop: 'oldStock', label: '历史库存', width: 80 },
  { prop: 'stock', label: '现有库存', width: 80 },
  { prop: 'intervals', label: '爬虫间隔', width: 80 },
  { prop: 'pushIntervals', label: '推送间隔', width: 80 },
  { prop: 'multiCheck', label: '重复检测', width: 80 },
  { prop: 'messageId', label: '消息编号', width: 80 },
  { prop: 'pushTime', label: '推送时间', width: 180, formatter: formatDate },
  { prop: 'flag', label: 'flag', width: 80 },
  { prop: 'additional', label: '其他', width: 120 }
]
const formFields = [
  { label: 'TAG:', prop: 'tag', required: true },
  { label: '链接:', prop: 'url', required: true },
  {
    label: '爬虫类型:', prop: 'billingType', type: 'select', required: true, options: [
      { label: 'single', value: 'single' },
      { label: 'multi', value: 'multi' }
    ]
  },
  { label: '爬虫间隔:', prop: 'intervals', inputType: 'number' },
  { label: '推送间隔:', prop: 'pushIntervals', inputType: 'number' },
  { label: '其他:', prop: 'additional', type: 'rich-edit' },
  { label: '消息编号:', prop: 'messageId' },
  { label: '现有库存:', prop: 'stock', inputType: 'number' },
  { label: '历史库存:', prop: 'oldStock', inputType: 'number' },
  { label: '推送库存:', prop: 'pushStock', inputType: 'number' },
  { label: '重复检测:', prop: 'multiCheck', inputType: 'number' },
  { label: 'flag', prop: 'flag' },
  { label: 'CPU:', prop: 'cpu' },
  { label: '内存:', prop: 'memory' },
  { label: '磁盘:', prop: 'disk' },
  { label: '流量:', prop: 'traffic' },
  { label: '端口:', prop: 'portSpeed' },
  { label: '地点:', prop: 'location' },
  { label: '价格:', prop: 'price' },
  { label: '推送时间:', prop: 'pushTime', type: 'datetime' },
]
const detailFields = tableColumns.map(col => ({
  label: col.label,
  prop: col.prop,
  type: col.prop === 'additional' ? 'rich-view' : undefined
}))
const rules = {
  form: {
    tag: [{
      required: true,
      message: '',
      trigger: ['input', 'blur'],
    },
    {
      whitespace: true,
      message: '不能只输入空格',
      trigger: ['input', 'blur'],
    }],
    url: [{
      required: true,
      message: '',
      trigger: ['input', 'blur'],
    }],
    billingType: [{
      required: true,
      message: '',
      trigger: ['input', 'blur'],
    }]
  },
  search: {
    pushTime: [{
      validator: (rule, value, callback) => {
        if (searchInfo.value.startPushTime && !searchInfo.value.endPushTime) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startPushTime && searchInfo.value.endPushTime) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startPushTime && searchInfo.value.endPushTime &&
          (searchInfo.value.startPushTime.getTime() === searchInfo.value.endPushTime.getTime() ||
            searchInfo.value.startPushTime.getTime() > searchInfo.value.endPushTime.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      },
      trigger: 'change'
    }]
  }
}
const paginationConfig = reactive({
  layout: 'total, sizes, prev, pager, next, jumper',
  total,
  'current-page': page,
  'page-size': pageSize,
  'page-sizes': [10, 30, 50, 100]
})
// 方法定义
const getTableData = async () => {
  const table = await getProductsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
  flag: '',
  intervals: 0,
  messageId: '',
  pushIntervals: 0,
  pushTime: new Date(),
  createdBy: null,
  updatedBy: null,
  deletedBy: null
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
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => deleteProductsFunc(row))
}
const deleteProductsFunc = async (row) => {
  const res = await deleteProducts({ ID: row.ID })
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
const onDelete = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const IDs = multipleSelection.value.map(item => item.ID)
    const res = await deleteProductsByIds({ IDs })
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
const updateProductsFunc = async (row) => {
  const res = await findProducts({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    // 确保数字类型字段的正确类型
    const data = {
      ...res.data,
      pushIntervals: Number(res.data.pushIntervals),
      pushStock: Number(res.data.pushStock),
      oldStock: Number(res.data.oldStock),
      stock: Number(res.data.stock),
      intervals: Number(res.data.intervals),
      multiCheck: Number(res.data.multiCheck)
    }
    formData.value = data
    dialogs.form.visible = true
  }
}
const openDialog = () => {
  type.value = 'create'
  formData.value = resetFormData()
  dialogs.form.visible = true
}
const closeDialog = () => {
  dialogs.form.visible = false
  formData.value = resetFormData()
}
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    // 确保数字字段的类型
    const submitData = {
      ...formData.value,
      pushIntervals: Number(formData.value.pushIntervals),
      pushStock: Number(formData.value.pushStock),
      oldStock: Number(formData.value.oldStock),
      stock: Number(formData.value.stock),
      intervals: Number(formData.value.intervals),
      multiCheck: Number(formData.value.multiCheck)
    }
    switch (type.value) {
      case 'create':
        res = await createProducts(submitData)
        break
      case 'update':
        res = await updateProducts(submitData)
        break
      default:
        res = await createProducts(submitData)
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
const getDetails = async (row) => {
  const res = await findProducts({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    dialogs.detail.visible = true
  }
}
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
      // 只更新表单中有值的字段
      Object.keys(batchEditForm.value).forEach(key => {
        if (batchEditForm.value[key] !== undefined &&
          batchEditForm.value[key] !== null &&
          batchEditForm.value[key] !== '') {
          updateData[key] = batchEditForm.value[key]
        }
      })
      return updateProducts(updateData)
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
    ElMessage({
      type: 'warning',
      message: '请选择要操作的数据'
    })
    return
  }
  try {
    await ElMessageBox.confirm(
      `确定要对 ${multipleSelection.value.length} 条记录进行批量清空操作吗？`,
      '确认清空',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    dialogs.batchClear.visible = true
    batchClearFields.value = []
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage({
        type: 'error',
        message: '操作取消'
      })
    }
  }
}
const closeBatchClear = () => {
  dialogs.batchClear.visible = false
  batchClearFields.value = []
}
const enterBatchClear = async () => {
  if (batchClearFields.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要清空的字段'
    })
    return
  }
  try {
    const updatePromises = multipleSelection.value.map(item => {
      const updateData = { ...item }
      batchClearFields.value.forEach(field => {
        updateData[field] = ''
      })
      return updateProducts(updateData)
    })
    const results = await Promise.all(updatePromises)
    const hasError = results.some(res => res.code !== 0)
    if (!hasError) {
      ElMessage({
        type: 'success',
        message: '批量清空成功'
      })
      closeBatchClear()
      getTableData()
    }
  } catch (error) {
    ElMessage({
      type: 'error',
      message: '批量清空失败'
    })
  }
}
// 初始化
getTableData()
</script>

<style>
.search-row {
  display: flex;
  justify-content: flex-start;
  flex-wrap: wrap;
  gap: 1rem;
  margin-bottom: 1rem;
}

.demo-form-inline .el-input {
  width: 160px;
}

.el-table .cell {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.el-table .url-column .cell {
  direction: rtl;
  text-align: left;
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
  .search-row {
    flex-direction: column;
  }

  .search-row .el-form-item {
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