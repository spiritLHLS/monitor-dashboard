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


        <template v-if="showAllQuery">
          <div>
            <el-form-item label="邀请码" prop="code">
              <el-input v-model="searchInfo.code" placeholder="搜索条件" />

            </el-form-item>
            <el-form-item label="使用状态" prop="status">

              <el-input v-model.number="searchInfo.status" placeholder="搜索条件" />

            </el-form-item>
            <el-form-item label="过期时间" prop="expires_at">

              <el-date-picker v-model="searchInfo.expires_at" type="datetime" placeholder="搜索条件"></el-date-picker>

            </el-form-item>
            <el-form-item label="使用者" prop="user_uuid">
              <el-input v-model="searchInfo.user_uuid" placeholder="搜索条件" />

            </el-form-item>
            <el-form-item label="备注" prop="remarks">
              <el-input v-model="searchInfo.remarks" placeholder="搜索条件" />
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
        <el-button type="success" icon="plus" style="margin-left: 10px;"
          @click="openBatchBuildDialog">批量创建邀请码</el-button>
        <el-button type="warning" icon="download" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="batchExport">批量导出邀请码</el-button>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>

        <el-table-column align="left" label="邀请码" prop="code" width="120" />
        <el-table-column align="left" label="使用状态" prop="status" width="120" />
        <el-table-column align="left" label="过期时间" prop="expires_at" width="180">
          <template #default="scope">{{ formatDate(scope.row.expires_at) }}</template>
        </el-table-column>
        <el-table-column align="left" label="使用时间" prop="used_at" width="180">
          <template #default="scope">{{ formatDate(scope.row.used_at) }}</template>
        </el-table-column>
        <el-table-column align="left" label="使用者" prop="user_uuid" width="120" />
        <el-table-column align="left" label="备注" prop="remarks" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon
                style="margin-right: 5px">
                <InfoFilled />
              </el-icon>查看详情</el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateInviteCodesFunc(scope.row)">变更</el-button>
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
        <el-form-item label="邀请码:" prop="code">
          <el-input v-model="formData.code" :clearable="false" placeholder="请输入邀请码" />
        </el-form-item>
        <el-form-item label="使用状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入使用状态" />
        </el-form-item>
        <el-form-item label="过期时间:" prop="expires_at">
          <el-date-picker v-model="formData.expires_at" type="date" style="width:100%" placeholder="选择日期"
            :clearable="true" />
        </el-form-item>
        <el-form-item label="使用时间:" prop="used_at">
          <el-date-picker v-model="formData.used_at" type="date" style="width:100%" placeholder="选择日期"
            :clearable="true" />
        </el-form-item>
        <el-form-item label="使用者:" prop="user_uuid">
          <el-input v-model="formData.user_uuid" :clearable="true" placeholder="请输入使用者" />
        </el-form-item>
        <el-form-item label="备注:" prop="remarks">
          <el-input v-model="formData.remarks" :clearable="true" placeholder="请输入备注" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="邀请码">
          {{ detailFrom.code }}
        </el-descriptions-item>
        <el-descriptions-item label="使用状态">
          {{ detailFrom.status }}
        </el-descriptions-item>
        <el-descriptions-item label="过期时间">
          {{ detailFrom.expires_at }}
        </el-descriptions-item>
        <el-descriptions-item label="使用时间">
          {{ detailFrom.used_at }}
        </el-descriptions-item>
        <el-descriptions-item label="使用者">
          {{ detailFrom.user_uuid }}
        </el-descriptions-item>
        <el-descriptions-item label="备注">
          {{ detailFrom.remarks }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

    <!-- 新增批量创建邀请码的对话框 -->
    <el-dialog v-model="batchBuildDialogVisible" title="批量创建邀请码" width="30%">
      <el-form :model="batchBuildForm" label-width="120px">
        <el-form-item label="创建数量">
          <el-input-number v-model="batchBuildForm.count" :min="1" :max="1000"></el-input-number>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="batchBuildDialogVisible = false">取消</el-button>
          <el-button type="primary" @click="confirmBatchBuild">确定</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 批量导出邀请码的结果对话框 -->
    <el-dialog v-model="exportDialogVisible" title="导出的邀请码" width="50%">
      <el-input v-model="exportedCodes" type="textarea" :rows="10" placeholder="导出的邀请码将显示在这里" readonly></el-input>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="exportDialogVisible = false">关闭</el-button>
          <el-button type="primary" @click="copyExportedCodes">复制</el-button>
        </span>
      </template>
    </el-dialog>

  </div>
</template>

<script setup>
import {
  createInviteCodes,
  deleteInviteCodes,
  deleteInviteCodesByIds,
  updateInviteCodes,
  findInviteCodes,
  getInviteCodesList,
  batchBuildCodes,
  batchExportCodes
} from '@/api/invite_codes/inviteCodes'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
  name: 'InviteCodes'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  code: '',
  status: undefined,
  expires_at: new Date(),
  used_at: new Date(),
  user_uuid: '',
  remarks: '',
})



// 验证规则
const rule = reactive({
  code: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
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
  expires_at: [{
    validator: (rule, value, callback) => {
      if (searchInfo.value.startExpires_at && !searchInfo.value.endExpires_at) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startExpires_at && searchInfo.value.endExpires_at) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startExpires_at && searchInfo.value.endExpires_at && (searchInfo.value.startExpires_at.getTime() === searchInfo.value.endExpires_at.getTime() || searchInfo.value.startExpires_at.getTime() > searchInfo.value.endExpires_at.getTime())) {
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
  const table = await getInviteCodesList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteInviteCodesFunc(row)
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
    const res = await deleteInviteCodesByIds({ IDs })
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
const updateInviteCodesFunc = async (row) => {
  const res = await findInviteCodes({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteInviteCodesFunc = async (row) => {
  const res = await deleteInviteCodes({ ID: row.ID })
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
    code: '',
    status: undefined,
    expires_at: new Date(),
    used_at: new Date(),
    user_uuid: '',
    remarks: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createInviteCodes(formData.value)
        break
      case 'update':
        res = await updateInviteCodes(formData.value)
        break
      default:
        res = await createInviteCodes(formData.value)
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
  const res = await findInviteCodes({ ID: row.ID })
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

// 批量创建邀请码相关
const batchBuildDialogVisible = ref(false)
const batchBuildForm = reactive({
  count: 1
})

const openBatchBuildDialog = () => {
  batchBuildDialogVisible.value = true
}

const confirmBatchBuild = async () => {
  try {
    const res = await batchBuildCodes({ count: batchBuildForm.count })
    if (res.code === 0) {
      const createdCodes = res.data.Codes.split('\n')
      ElMessage({
        type: 'success',
        message: `成功创建 ${createdCodes.length} 个邀请码`
      })
      batchBuildDialogVisible.value = false
      getTableData()
    }
  } catch (error) {
    ElMessage({
      type: 'error',
      message: '批量创建邀请码失败'
    })
  }
}


// 批量导出邀请码相关
const exportDialogVisible = ref(false)
const exportedCodes = ref('')

const batchExport = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要导出的邀请码'
    })
    return
  }

  try {
    const IDs = multipleSelection.value.map(item => item.ID)
    const res = await batchExportCodes({ IDs })
    if (res.code === 0) {
      exportedCodes.value = res.data.Codes
      exportDialogVisible.value = true
    }
  } catch (error) {
    ElMessage({
      type: 'error',
      message: '批量导出邀请码失败'
    })
  }
}

const copyExportedCodes = () => {
  navigator.clipboard.writeText(exportedCodes.value).then(() => {
    ElMessage({
      type: 'success',
      message: '邀请码已复制到剪贴板'
    })
  }, () => {
    ElMessage({
      type: 'error',
      message: '复制失败，请手动复制'
    })
  })
}

</script>

<style>
.row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 1px;
  /* 行与行之间的间距 */
}

.row .el-form-item {
  flex: 1;
  /* 使每个表单项占据相等的空间 */
  margin-right: 10px;
  /* 表单项之间的间距 */
}

.el-table .cell {
  /* 文本超出容器部分隐藏 */
  overflow: hidden;
  /* 超出部分使用省略号代替 */
  text-overflow: ellipsis;
  /* 不换行 */
  white-space: nowrap;
}

.demo-form-inline .el-input {
  width: 160px;
  /* 输入框的宽度 */
}
</style>
