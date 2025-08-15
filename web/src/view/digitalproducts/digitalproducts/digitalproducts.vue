<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <!-- 第一行：创建日期 -->
        <div class="search-row">
          <el-form-item label="创建日期" prop="createdAt">
            <template #label>
              <span>
                创建日期
                <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                  <el-icon><QuestionFilled /></el-icon>
                </el-tooltip>
              </span>
            </template>
            <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
          </el-form-item>
          
          <el-form-item>
            <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
            <el-button icon="refresh" @click="onReset">重置</el-button>
            <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
            <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
          </el-form-item>
        </div>
        
        <!-- 展开的搜索条件 - 分多行显示 -->
        <template v-if="showAllQuery">
          <!-- 第二行：TAG、核心、内存、硬盘 -->
          <div class="search-row">
            <el-form-item label="TAG" prop="tag">
              <el-input v-model="searchInfo.tag" placeholder="搜索TAG" />
            </el-form-item>
            <el-form-item label="核心" prop="cpu">
              <el-input v-model.number="searchInfo.cpu" placeholder="搜索核心数" />
            </el-form-item>
            <el-form-item label="内存" prop="memory">
              <el-input v-model.number="searchInfo.memory" placeholder="搜索内存" />
            </el-form-item>
            <el-form-item label="硬盘" prop="disk">
              <el-input v-model.number="searchInfo.disk" placeholder="搜索硬盘" />
            </el-form-item>
          </div>
          
          <!-- 第三行：流量、带宽、位置、价格 -->
          <div class="search-row">
            <el-form-item label="流量" prop="traffic">
              <el-input v-model.number="searchInfo.traffic" placeholder="搜索流量" />
            </el-form-item>
            <el-form-item label="带宽" prop="portSpeed">
              <el-input v-model.number="searchInfo.portSpeed" placeholder="搜索带宽" />
            </el-form-item>
            <el-form-item label="位置" prop="location">
              <el-input v-model="searchInfo.location" placeholder="搜索位置" />
            </el-form-item>
            <el-form-item label="价格" prop="price">
              <el-input v-model.number="searchInfo.price" placeholder="搜索价格" />
            </el-form-item>
          </div>
          
          <!-- 第四行：价格单位 -->
          <div class="search-row">
            <el-form-item label="价格单位" prop="priceUnit">
              <el-input v-model="searchInfo.priceUnit" placeholder="搜索条件" />
            </el-form-item>
          </div>
        </template>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog()">新增</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        <el-button type="warning" icon="edit" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="openBatchEditDialog">批量修改</el-button>
        <el-button type="success" icon="refresh" style="margin-left: 10px;" @click="onConvert" :loading="convertLoading">一键转换</el-button>
        <el-button type="info" icon="upload" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onBatchConvert" :loading="batchConvertLoading">批量转换</el-button>
        <el-button type="primary" icon="view" style="margin-left: 10px;" @click="checkConversionStatus" :loading="statusLoading" plain>查询状态</el-button>
      </div>
      
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="TAG" prop="tag" width="120" />
        <el-table-column align="left" label="核心" prop="cpu" width="120" />
        <el-table-column align="left" label="内存" prop="memory" width="120" />
        <el-table-column align="left" label="硬盘" prop="disk" width="120" />
        <el-table-column align="left" label="流量" prop="traffic" width="120" />
        <el-table-column align="left" label="带宽" prop="portSpeed" width="120" />
        <el-table-column align="left" label="位置" prop="location" width="120" />
        <el-table-column align="left" label="价格" prop="price" width="120" />
        <el-table-column align="left" label="价格单位" prop="priceUnit" width="120" />
        <el-table-column align="left" label="其他" prop="additional" width="120" />
        <el-table-column align="left" label="原表ID" prop="originId" width="120" />
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="updateDigitalProductsFunc(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
          <div>
            <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
            <el-button @click="closeDialog">取 消</el-button>
          </div>
        </div>
      </template>
      <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item v-for="field in formFields" :key="field.prop" :label="field.label + ':'" :prop="field.prop">
          <el-input 
            v-if="field.type === 'number' || field.type === 'float'"
            v-model.number="formData[field.prop]" 
            :clearable="true" 
            :placeholder="'请输入' + field.label"
          />
          <el-input 
            v-else
            v-model="formData[field.prop]" 
            :clearable="true" 
            :placeholder="'请输入' + field.label"
          />
        </el-form-item>
      </el-form>
    </el-drawer>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="batchEditVisible" :show-close="false" :before-close="closeBatchEditDialog">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">批量修改 (已选择 {{multipleSelection.length}} 条记录)</span>
          <div>
            <el-button :loading="batchEditLoading" type="primary" @click="enterBatchEditDialog">确 定</el-button>
            <el-button @click="closeBatchEditDialog">取 消</el-button>
          </div>
        </div>
      </template>
      <el-alert
        title="批量修改说明"
        description="只有填写了值的字段才会被更新，空白字段将保持原值不变"
        type="info"
        :closable="false"
        style="margin-bottom: 20px;"
      />
      <el-form :model="batchEditData" label-position="top" ref="elBatchEditFormRef" label-width="80px">
        <el-form-item v-for="field in batchEditFields" :key="field.prop" :label="field.label + ':'">
          <el-input 
            v-if="field.type === 'number' || field.type === 'float'"
            v-model.number="batchEditData[field.prop]" 
            :clearable="true" 
            :placeholder="'留空则不修改' + field.label"
          />
          <el-input 
            v-else
            v-model="batchEditData[field.prop]" 
            :clearable="true" 
            :placeholder="'留空则不修改' + field.label"
          />
        </el-form-item>
      </el-form>
    </el-drawer>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
      <el-descriptions :column="1" border>
        <el-descriptions-item v-for="field in formFields" :key="field.prop" :label="field.label">
          {{ detailFrom[field.prop] }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createDigitalProducts,
  deleteDigitalProducts,
  deleteDigitalProductsByIds,
  updateDigitalProducts,
  findDigitalProducts,
  getDigitalProductsList,
  convertProductsToDigital,
  batchConvertProductsToDigital,
  getConversionStatus
} from '@/api/digitalproducts/digitalproducts'

import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive, computed } from 'vue'
import { useBtnAuth } from '@/utils/btnAuth'
import { useAppStore } from "@/pinia"

defineOptions({
  name: 'DigitalProducts'
})

// 字段配置
const formFields = [
  { prop: 'tag', label: 'TAG', type: 'string' },
  { prop: 'cpu', label: '核心', type: 'float' },
  { prop: 'memory', label: '内存', type: 'float' },
  { prop: 'disk', label: '硬盘', type: 'float' },
  { prop: 'traffic', label: '流量', type: 'float' },
  { prop: 'portSpeed', label: '带宽', type: 'float' },
  { prop: 'location', label: '位置', type: 'string' },
  { prop: 'price', label: '价格', type: 'float' },
  { prop: 'priceUnit', label: '价格单位', type: 'string' },
  { prop: 'additional', label: '其他', type: 'string' },
  { prop: 'originId', label: '原表ID', type: 'number' }
]

// 初始化表单数据
const getInitialFormData = () => {
  const data = {}
  formFields.forEach(field => {
    if (field.type === 'number' || field.type === 'float') {
      data[field.prop] = undefined
    } else {
      data[field.prop] = ''
    }
  })
  return data
}

// 批量编辑字段（排除不需要批量修改的字段）
const batchEditFields = formFields.filter(field => !['originId'].includes(field.prop))

// 基础配置
const btnAuth = useBtnAuth()
const appStore = useAppStore()
const btnLoading = ref(false)
const convertLoading = ref(false)
const batchEditLoading = ref(false)
const showAllQuery = ref(false)

// 表单数据
const formData = ref(getInitialFormData())
const batchEditData = ref(getInitialFormData())

// 验证规则
const rule = reactive({})
const searchRule = reactive({
  createdAt: [
    { 
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && 
                  (searchInfo.value.startCreatedAt.getTime() >= searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, 
      trigger: 'change' 
    }
  ],
})

const elFormRef = ref()
const elBatchEditFormRef = ref()
const elSearchFormRef = ref()

// 表格控制
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const multipleSelection = ref([])

// 弹窗控制
const type = ref('')
const dialogFormVisible = ref(false)
const batchEditVisible = ref(false)
const detailShow = ref(false)
const detailFrom = ref({})
const batchConvertLoading = ref(false)
const statusLoading = ref(false)

// 表格操作方法
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const getTableData = async() => {
  const table = await getDigitalProductsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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

// 删除操作
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteDigitalProductsFunc(row)
  })
}

const onDelete = async() => {
  if (multipleSelection.value.length === 0) {
    ElMessage({ type: 'warning', message: '请选择要删除的数据' })
    return
  }
  
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    const IDs = multipleSelection.value.map(item => item.ID)
    const res = await deleteDigitalProductsByIds({ IDs })
    if (res.code === 0) {
      ElMessage({ type: 'success', message: '删除成功' })
      if (tableData.value.length === IDs.length && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}

const deleteDigitalProductsFunc = async (row) => {
  const res = await deleteDigitalProducts({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '删除成功' })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 单个编辑操作
const updateDigitalProductsFunc = async(row) => {
  const res = await findDigitalProducts({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}

const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = getInitialFormData()
}

const enterDialog = async () => {
  btnLoading.value = true
  elFormRef.value?.validate(async (valid) => {
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
      ElMessage({ type: 'success', message: '创建/更改成功' })
      closeDialog()
      getTableData()
    }
  })
}

// 批量编辑操作
const openBatchEditDialog = () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({ type: 'warning', message: '请选择要修改的数据' })
    return
  }
  batchEditData.value = getInitialFormData()
  batchEditVisible.value = true
}

const closeBatchEditDialog = () => {
  batchEditVisible.value = false
  batchEditData.value = getInitialFormData()
}

const enterBatchEditDialog = async () => {
  batchEditLoading.value = true
  try {
    const updateData = {}
    Object.keys(batchEditData.value).forEach(key => {
      const value = batchEditData.value[key]
      if (value !== '' && value !== undefined && value !== null) {
        updateData[key] = value
      }
    })
    if (Object.keys(updateData).length === 0) {
      ElMessage({ type: 'warning', message: '请至少填写一个要修改的字段' })
      batchEditLoading.value = false
      return
    }
    const promises = multipleSelection.value.map(item => {
      const data = { ...item, ...updateData }
      return updateDigitalProducts(data)
    })
    const results = await Promise.all(promises)
    const successCount = results.filter(res => res.code === 0).length
    if (successCount === multipleSelection.value.length) {
      ElMessage({ type: 'success', message: `批量修改成功，共更新 ${successCount} 条记录` })
    } else {
      ElMessage({ type: 'warning', message: `部分更新成功，成功 ${successCount} 条，失败 ${multipleSelection.value.length - successCount} 条` })
    }  
    closeBatchEditDialog()
    getTableData()
  } catch (error) {
    ElMessage({ type: 'error', message: '批量修改失败，请重试' })
  } finally {
    batchEditLoading.value = false
  }
}

// 详情查看
const getDetails = async (row) => {
  const res = await findDigitalProducts({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    detailShow.value = true
  }
}

const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}

// 转换操作 - 一键转换（全量转换）
const onConvert = async() => {
  ElMessageBox.confirm('确定要将products表数据转换为数字商品表吗？此操作会同步所有数据。', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'info'
  }).then(async() => {
    convertLoading.value = true
    try {
      const res = await convertProductsToDigital()
      if (res.code === 0) {
        ElMessage({ 
          type: 'success', 
          message: '转换任务已启动，请点击"查询状态"按钮查看进度',
          duration: 5000 
        })
      }
    } catch (error) {
      ElMessage({ type: 'error', message: '转换失败，请重试' })
    } finally {
      convertLoading.value = false
    }
  })
}

// 批量转换操作
const onBatchConvert = async() => {
  if (multipleSelection.value.length === 0) {
    ElMessage({ type: 'warning', message: '请选择要重新转换的数据' })
    return
  }
  ElMessageBox.confirm(`确定要重新转换选中的 ${multipleSelection.value.length} 条记录吗？此操作会从原表重新解析数据并覆盖当前记录。`, '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    batchConvertLoading.value = true
    try {
      const digitalProductIds = multipleSelection.value.map(item => item.ID)
      
      const res = await batchConvertProductsToDigital(digitalProductIds)
      if (res.code === 0) {
        ElMessage({ 
          type: 'success', 
          message: res.msg + '，请点击"查询状态"按钮查看进度',
          duration: 5000 
        })
      }
    } catch (error) {
      ElMessage({ type: 'error', message: '批量转换失败，请重试' })
    } finally {
      batchConvertLoading.value = false
    }
  })
}

// 手动查询转换状态
const checkConversionStatus = async () => {
  statusLoading.value = true
  try {
    const res = await getConversionStatus()
    if (res.code === 0) {
      if (!res.data) {
        ElMessage({ 
          type: 'info', 
          message: '当前没有运行中的转换任务' 
        })
        return
      }
      const status = res.data
      if (status.is_running) {
        const progressPercent = status.total_products > 0 
          ? Math.round((status.progress / status.total_products) * 100) 
          : 0
        ElMessage({ 
          type: 'info', 
          message: `转换任务进行中 - ${status.current_step}
进度：${status.progress}/${status.total_products} (${progressPercent}%)
成功：${status.success_count} 条，失败：${status.fail_count} 条`,
          duration: 6000,
          dangerouslyUseHTMLString: false
        })
      } else {
        const completionTime = new Date(status.start_time).toLocaleString()
        const messageType = status.fail_count > 0 ? 'warning' : 'success'
        ElMessage({ 
          type: messageType, 
          message: `最近转换任务已完成，
开始时间：${completionTime}
结果：成功 ${status.success_count} 条，失败 ${status.fail_count} 条
${status.last_error ? '错误信息：' + status.last_error : ''}`,
          duration: 8000
        })
        getTableData()
      }
    } else {
      ElMessage({ type: 'error', message: '查询状态失败：' + (res.msg || '未知错误') })
    }
  } catch (error) {
    console.error('查询转换状态失败:', error)
    ElMessage({ type: 'error', message: '查询状态失败，请重试' })
  } finally {
    statusLoading.value = false
  }
}

getTableData()
</script>

<style scoped>
.search-row {
  display: flex;
  flex-wrap: wrap;
  align-items: flex-end;
  margin-bottom: 10px;
}

.search-row .el-form-item {
  margin-right: 20px;
  margin-bottom: 10px;
}

.search-row:last-of-type {
  margin-bottom: 0;
}

/* 响应式布局 */
@media (max-width: 1400px) {
  .search-row .el-form-item {
    margin-right: 15px;
  }
}

@media (max-width: 1200px) {
  .search-row {
    flex-direction: column;
    align-items: flex-start;
  }
  
  .search-row .el-form-item {
    margin-right: 0;
    width: 100%;
  }
}
</style>