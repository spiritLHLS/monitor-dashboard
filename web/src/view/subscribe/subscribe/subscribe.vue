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
            :disabled-date="time => searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
            :disabled-date="time => searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
        </el-form-item>
        <el-form-item label="最后更新" prop="last_update">
          <el-date-picker v-model="searchInfo.last_update" type="datetime" placeholder="搜索条件"></el-date-picker>
        </el-form-item>
        <template v-if="showAllQuery">
          <div class="row">
            <el-form-item label="用户UUID" prop="user_uuid">
              <el-input v-model.number="searchInfo.user_uuid" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="商品ID" prop="product_id">
              <el-input v-model.number="searchInfo.product_id" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="订阅状态" prop="status">
              <el-input v-model.number="searchInfo.status" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="通知渠道" prop="notify_channel">
              <el-input v-model="searchInfo.notify_channel" placeholder="搜索条件" />
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
        <el-button type="primary" icon="edit" @click="openBatchStatusDialog"
          :disabled="!multipleSelection.length">批量修改订阅状态</el-button>
        <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="onDelete">删除</el-button>
      </div>
      <el-dialog v-model="batchStatusDialogVisible" title="批量修改订阅状态" width="30%">
        <el-form :model="batchStatusForm">
          <el-form-item label="订阅状态">
            <el-select v-model="batchStatusForm.status" placeholder="请选择订阅状态">
              <el-option label="启用" :value="0"></el-option>
              <el-option label="冻结" :value="1"></el-option>
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="batchStatusDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmBatchStatusChange">确定</el-button>
          </span>
        </template>
      </el-dialog>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="用户UUID" prop="user_uuid" width="120" />
        <el-table-column align="left" label="商品ID" prop="product_id" width="120" />
        <el-table-column align="left" label="订阅状态" prop="status" width="120" />
        <el-table-column align="left" label="通知渠道" prop="notify_channel" width="120" />
        <el-table-column align="left" label="最后更新" prop="last_update" width="180">
          <template #default="scope">{{ formatDate(scope.row.last_update) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon
                style="margin-right: 5px">
                <InfoFilled />
              </el-icon>查看详情</el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateSubscribeFunc(scope.row)">变更</el-button>
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
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
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
        <el-form-item label="用户UUID:" prop="user_uuid">
          <el-input v-model.number="formData.user_uuid" :clearable="true" placeholder="请输入用户UUID" />
        </el-form-item>
        <el-form-item label="商品ID:" prop="product_id">
          <el-input v-model.number="formData.product_id" :clearable="true" placeholder="请输入商品ID" />
        </el-form-item>
        <el-form-item label="订阅状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择订阅状态">
            <el-option label="启用" :value="0"></el-option>
            <el-option label="冻结" :value="1"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="通知渠道:" prop="notify_channel">
          <el-input v-model="formData.notify_channel" :clearable="true" placeholder="请输入通知渠道" />
        </el-form-item>
        <el-form-item label="最后更新:" prop="last_update">
          <el-date-picker v-model="formData.last_update" type="date" style="width:100%" placeholder="选择日期"
            :clearable="true" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="用户UUID">
          {{ detailFrom.user_uuid }}
        </el-descriptions-item>
        <el-descriptions-item label="商品ID">
          {{ detailFrom.product_id }}
        </el-descriptions-item>
        <el-descriptions-item label="订阅状态">
          {{ detailFrom.status }}
        </el-descriptions-item>
        <el-descriptions-item label="通知渠道">
          {{ detailFrom.notify_channel }}
        </el-descriptions-item>
        <el-descriptions-item label="最后更新">
          {{ detailFrom.last_update }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

    <el-dialog v-model="batchStatusDialogVisible" title="批量修改订阅状态" width="30%">
      <el-form :model="batchStatusForm">
        <el-form-item label="订阅状态">
          <el-select v-model="batchStatusForm.status" placeholder="请选择订阅状态">
            <el-option label="启用" :value="0"></el-option>
            <el-option label="冻结" :value="1"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="batchStatusDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmBatchStatusChange">确定</el-button>
        </span>
      </template>
    </el-dialog>

  </div>
</template>

<script setup>
import {
  createSubscribe,
  deleteSubscribe,
  deleteSubscribeByIds,
  updateSubscribe,
  findSubscribe,
  getSubscribeList
} from '@/api/subscribe/subscribe'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'Subscribe'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  user_uuid: undefined,
  product_id: undefined,
  status: undefined,
  notify_channel: '',
  last_update: new Date(),
})

// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
  ],
  last_update: [{
    validator: (rule, value, callback) => {
      if (searchInfo.value.startLastUpdate && !searchInfo.value.endLastUpdate) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startLastUpdate && searchInfo.value.endLastUpdate) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startLastUpdate && searchInfo.value.endLastUpdate && (searchInfo.value.startLastUpdate.getTime() === searchInfo.value.endLastUpdate.getTime() || searchInfo.value.startLastUpdate.getTime() > searchInfo.value.endLastUpdate.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change'
  }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getSubscribeList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteSubscribeFunc(row)
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
    const res = await deleteSubscribeByIds({ IDs })
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

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateSubscribeFunc = async (row) => {
  const res = await findSubscribe({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteSubscribeFunc = async (row) => {
  const res = await deleteSubscribe({ ID: row.ID })
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

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    user_uuid: undefined,
    product_id: undefined,
    status: undefined,
    notify_channel: '',
    last_update: new Date(),
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createSubscribe(formData.value)
        break
      case 'update':
        res = await updateSubscribe(formData.value)
        break
      default:
        res = await createSubscribe(formData.value)
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

const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)

// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}

// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findSubscribe({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}

// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}

// 批量修改状态相关
const batchStatusDialogVisible = ref(false)
const batchStatusForm = ref({
  status: undefined
})

// 打开批量修改状态弹窗
const openBatchStatusDialog = () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的数据'
    })
    return
  }
  batchStatusDialogVisible.value = true
}

// 确认批量修改状态
const confirmBatchStatusChange = async () => {
  if (batchStatusForm.value.status === undefined) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的状态'
    })
    return
  }

  const updatePromises = multipleSelection.value.map(item => 
    updateSubscribe({ ...item, status: batchStatusForm.value.status })
  )

  try {
    await Promise.all(updatePromises)
    ElMessage({
      type: 'success',
      message: '批量修改成功'
    })
    batchStatusDialogVisible.value = false
    getTableData()
  } catch (error) {
    ElMessage({
      type: 'error',
      message: '批量修改失败，请重试'
    })
  }
}

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
</style>
