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
          <div class="row">
          <el-form-item label="Tag" prop="tg_tag">
            <el-input v-model="searchInfo.tg_tag" placeholder="搜索条件" />
          </el-form-item>
          <el-form-item label="分区名字" prop="name">
            <el-input v-model="searchInfo.name" placeholder="搜索条件" />
          </el-form-item>
          <el-form-item label="分区链接" prop="link">
            <el-input v-model="searchInfo.link" placeholder="搜索条件" />
          </el-form-item>
          <el-form-item label="分区类型" prop="type">
            <el-input v-model="searchInfo.type" placeholder="搜索条件" />
          </el-form-item>
          <el-form-item label="识别数量" prop="num">
            <el-input v-model.number="searchInfo.num" placeholder="搜索条件" />
          </el-form-item>
          <el-form-item label="爬虫间隔" prop="intervals">
            <el-input v-model.number="searchInfo.intervals" placeholder="搜索条件" />
          </el-form-item>
          </div>
          <!-- <div class="row"> -->
          <el-form-item label="其他信息" prop="additional">
            <el-input v-model="searchInfo.additional" placeholder="搜索条件" />
          </el-form-item>
          <!-- </div> -->
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
        <!-- 修改爬取间隔对话框 -->
        <el-dialog v-model="modifyIntervalsDialogVisible" title="修改爬取间隔" width="30%">
          <el-form label-width="120px">
            <el-form-item label="新的爬取间隔">
              <el-input v-model.number="modifyIntervals" placeholder="请输入新的爬取间隔"></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="modifyIntervalsDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmModifyIntervals">确定</el-button>
          </div>
        </el-dialog>
        <!-- 修改识别数量对话框 -->
        <el-dialog v-model="modifyNumDialogVisible" title="修改识别数量" width="30%">
          <el-form label-width="120px">
            <el-form-item label="新的识别数量">
              <el-input v-model.number="modifyNum" placeholder="请输入新的识别数量"></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="modifyNumDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmModifyNum">确定</el-button>
          </div>
        </el-dialog>
        <!-- 修改其他信息对话框 -->
        <el-dialog v-model="modifyAdditionalDialogVisible" title="修改其他信息" width="30%">
          <el-form label-width="120px">
            <el-form-item label="新的其他信息">
              <el-input v-model.trim="modifyAdditional" placeholder="请输入新的其他信息"></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="modifyAdditionalDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmmodifyAdditional">确定</el-button>
          </div>
        </el-dialog>
        <!-- 修改频道TAG对话框 -->
        <el-dialog v-model="modifyTGTAGDialogVisible" title="修改TAG" width="30%">
          <el-form label-width="120px">
            <el-form-item label="新的TAG">
              <el-input v-model.trim="modifyTGTAG" placeholder="请输入新的TAG"></el-input>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="modifyTGTAGDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="confirmmodifyTGTAG">确定</el-button>
          </div>
        </el-dialog>
        <el-button icon="edit" class="table-button" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="updateIntervals">修改爬虫间隔</el-button>
        <el-button icon="edit" class="table-button" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="updateNum">修改识别数量</el-button>
        <el-button icon="edit" class="table-button" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="updateAdditional">修改其他信息</el-button>
        <el-button icon="edit" class="table-button" style="margin-left: 10px;" :disabled="!multipleSelection.length"
          @click="updateTGTAG">修改TAG</el-button>
      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="Tag" prop="tgTag" width="90" />
        <el-table-column align="left" label="分区名字" prop="name" width="90" />
        <el-table-column align="left" label="分区链接" prop="link" width="90" />
        <el-table-column align="left" label="分区类型" prop="type" width="90" />
        <el-table-column align="left" label="识别数量" prop="num" width="90" />
        <el-table-column label="其他信息" prop="additional" width="120">
          <template #default="scope">
            [富文本内容]
          </template>
        </el-table-column>
        <el-table-column align="left" label="爬虫间隔" prop="intervals" width="90" />
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <!-- <el-table-column align="left" label="创建者" prop="createdBy" width="90" />
        <el-table-column align="left" label="更新者" prop="updatedBy" width="90" />
        <el-table-column align="left" label="删除者" prop="deletedBy" width="90" /> -->
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon
                style="margin-right: 5px">
                <InfoFilled />
              </el-icon>查看详情</el-button>
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
        <el-form-item label="Tag:" prop="tgTag">
          <el-input v-model="formData.tgTag" :clearable="true" placeholder="请输入Tag" />
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
        <el-form-item label="分区链接:" prop="link">
          <el-input v-model="formData.link" :clearable="true" placeholder="请输入分区链接" />
        </el-form-item>
        <el-form-item label="分区类型:" prop="type">
          <el-input v-model="formData.type" :clearable="true" placeholder="请输入分区类型" />
        </el-form-item>
        <el-form-item label="识别数量:" prop="num">
          <el-input v-model.number="formData.num" :clearable="true" placeholder="请输入识别数量" />
        </el-form-item>
        <el-form-item label="其他信息:" prop="additional">
          <RichEdit v-model="formData.additional" />
        </el-form-item>
        <el-form-item label="爬虫间隔:" prop="intervals">
          <el-input v-model.number="formData.intervals" :clearable="true" placeholder="请输入爬虫间隔" />
        </el-form-item>
        <!-- <el-form-item label="创建者:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入创建者" />
        </el-form-item>
        <el-form-item label="更新者:" prop="updatedBy">
          <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入更新者" />
        </el-form-item>
        <el-form-item label="删除者:" prop="deletedBy">
          <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入删除者" />
        </el-form-item> -->
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="Tag">
          {{ detailFrom.tgTag }}
        </el-descriptions-item>
        <el-descriptions-item label="分区名字">
          {{ detailFrom.name }}
        </el-descriptions-item>
        <el-descriptions-item label="分区链接">
          {{ detailFrom.link }}
        </el-descriptions-item>
        <el-descriptions-item label="分区类型">
          {{ detailFrom.type }}
        </el-descriptions-item>
        <el-descriptions-item label="识别数量">
          {{ detailFrom.num }}
        </el-descriptions-item>
        <el-descriptions-item label="其他信息">
          {{ detailFrom.additional }}
        </el-descriptions-item>
        <el-descriptions-item label="爬虫间隔">
          {{ detailFrom.intervals }}
        </el-descriptions-item>
        <!-- <el-descriptions-item label="创建者">
          {{ detailFrom.createdBy }}
        </el-descriptions-item>
        <el-descriptions-item label="更新者">
          {{ detailFrom.updatedBy }}
        </el-descriptions-item>
        <el-descriptions-item label="删除者">
          {{ detailFrom.deletedBy }}
        </el-descriptions-item> -->
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
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
  name: 'Partitionspage'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  link: [{
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
  type: [{
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
  const table = await getPartitionspageList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updatePartitionspageFunc = async (row) => {
  const res = await findPartitionspage({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
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
  }
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
  const res = await findPartitionspage({ ID: row.ID })
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

// 批量更新爬取间隔、识别数量
const modifyIntervalsDialogVisible = ref(false)
const modifyNumDialogVisible = ref(false)
const modifyAdditionalDialogVisible = ref(false)
const modifyTGTAGDialogVisible = ref(false)

// 更新所选记录的爬取间隔的方法
const updateIntervals = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的记录'
    })
    return
  }
  modifyIntervalsDialogVisible.value = true
}

// 更新所选记录的识别数量的方法
const updateNum = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的记录'
    })
    return
  }
  modifyNumDialogVisible.value = true
}

// 更新所选记录的其他信息的方法
const updateAdditional = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的记录'
    })
    return
  }
  modifyAdditionalDialogVisible.value = true
}

// 更新所选记录的TGTAG的方法
const updateTGTAG = async () => {
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要修改的记录'
    })
    return
  }
  modifyTGTAGDialogVisible.value = true
}

// 修改的新值
const modifyIntervals = ref(0)
const modifyNum = ref(0)
const modifyAdditional = ref("")
const modifyTGTAG = ref("")

// 确认修改爬取间隔
const confirmModifyIntervals = async () => {
  // 获取用户输入的新值
  const newIntervals = modifyIntervals.value
  // 调用更新函数
  await updateSelectedRecords(newIntervals, 'intervals')
  // 关闭对话框
  modifyIntervalsDialogVisible.value = false
}

// 确认修改识别数量
const confirmModifyNum = async () => {
  // 获取用户输入的新值
  const newNum = modifyNum.value
  // 调用更新函数
  await updateSelectedRecords(newNum, 'num')
  // 关闭对话框
  modifyNumDialogVisible.value = false
}

// 确认修改其他信息
const confirmmodifyAdditional = async () => {
  // 获取用户输入的新值
  const newadditional = modifyAdditional.value
  // 调用更新函数
  await updateSelectedRecords(newadditional, 'additional')
  // 关闭对话框
  modifyAdditionalDialogVisible.value = false
}

// 确认修改TGTAG
const confirmmodifyTGTAG = async () => {
  // 获取用户输入的新值
  const newtgtag = modifyTGTAG.value
  // 调用更新函数
  await updateSelectedRecords(newtgtag, 'tg_tag')
  // 关闭对话框
  modifyTGTAGDialogVisible.value = false
}

// 更新选中记录的指定字段值
const updateSelectedRecords = async (newValue, field) => {
  // 获取选中记录的ID列表
  const IDs = multipleSelection.value.map(record => record.ID)
  // 构造更新数据
  const updatedRecords = multipleSelection.value.map(record => ({
    ...record,
    [field]: newValue
  }))
  // 调用更新接口，这里调用你的更新接口函数
  const res = await Promise.all(updatedRecords.map(record => updatePartitionspage(record)))
  // 处理响应
  handleUpdateResponse(res)
}

// 处理来自API更新的响应
const handleUpdateResponse = (responses) => {
  const isSuccess = responses.every(res => res.code === 0)
  if (isSuccess) {
    ElMessage({
      type: 'success',
      message: '修改成功'
    })
    getTableData() // 刷新表格数据
  } else {
    ElMessage({
      type: 'error',
      message: '修改失败，请稍后重试'
    })
  }
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