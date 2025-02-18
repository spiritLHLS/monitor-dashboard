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
            <el-form-item label="频道TAG" prop="tag">
              <el-input v-model="searchInfo.tag" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="标志" prop="flag">
              <el-input v-model="searchInfo.flag" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="计费类型" prop="billingType">
              <el-input v-model="searchInfo.billingType" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="起始编号" prop="pdId">
              <el-input v-model.number="searchInfo.pdId" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="结束编号" prop="endId">
              <el-input v-model.number="searchInfo.endId" placeholder="搜索条件" />
            </el-form-item>
          </div>
          <div class="row">
            <el-form-item label="cpu" prop="cpu">
              <el-input v-model="searchInfo.cpu" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="内存" prop="memory">
              <el-input v-model="searchInfo.memory" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="硬盘" prop="disk">
              <el-input v-model="searchInfo.disk" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="流量" prop="traffic">
              <el-input v-model="searchInfo.traffic" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="带宽" prop="portSpeed">
              <el-input v-model="searchInfo.portSpeed" placeholder="搜索条件" />
            </el-form-item>
          </div>
          <div class="row">
            <el-form-item label="地址" prop="location">
              <el-input v-model="searchInfo.location" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="价格" prop="price">
              <el-input v-model="searchInfo.price" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="其他信息" prop="additional">
              <el-input v-model="searchInfo.additional" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="链接" prop="url">
              <el-input v-model="searchInfo.url" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="历史库存" prop="oldStock">
              <el-input v-model.number="searchInfo.oldStock" placeholder="搜索条件" />
            </el-form-item>
          </div>
          <el-form-item label="现有库存" prop="stock">
            <el-input v-model.number="searchInfo.stock" placeholder="搜索条件" />
          </el-form-item>
          <el-form-item label="消息ID" prop="messageId">
            <el-input v-model="searchInfo.messageId" placeholder="搜索条件" />
          </el-form-item>
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

      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="TAG" prop="tag" width="90" />
        <el-table-column align="left" label="标志" prop="flag" width="90" />
        <el-table-column align="left" label="计费类型" prop="billingType" width="90" />
        <el-table-column align="left" label="起始编号" prop="pdId" width="90" />
        <el-table-column align="left" label="结束编号" prop="endId" width="90" />
        <el-table-column align="left" label="CPU" prop="cpu" width="90" />
        <el-table-column align="left" label="内存" prop="memory" width="90" />
        <el-table-column align="left" label="硬盘" prop="disk" width="90" />
        <el-table-column align="left" label="流量" prop="traffic" width="90" />
        <el-table-column align="left" label="带宽" prop="portSpeed" width="90" />
        <el-table-column align="left" label="地址" prop="location" width="90" />
        <el-table-column align="left" label="价格" prop="price" width="90" />
        <el-table-column align="left" label="其他信息" prop="additional" width="90" />
        <el-table-column align="left" label="链接" prop="url" width="90" />
        <el-table-column align="left" label="历史库存" prop="oldStock" width="90" />
        <el-table-column align="left" label="现有库存" prop="stock" width="90" />
        <el-table-column align="left" label="消息ID" prop="messageId" width="90" />
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon
                style="margin-right: 5px">
                <InfoFilled />
              </el-icon>查看详情</el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateFindallpdFunc(scope.row)">变更</el-button>
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
        <el-form-item label="TAG:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true" placeholder="请输入TAG" />
        </el-form-item>
        <el-form-item label="标志:" prop="flag">
          <el-input v-model="formData.flag" :clearable="true" placeholder="请输入标志" />
        </el-form-item>
        <el-form-item label="计费类型:" prop="billingType">
          <el-input v-model="formData.billingType" :clearable="true" placeholder="请输入计费类型" />
        </el-form-item>
        <el-form-item label="起始编号:" prop="pdId">
          <el-input v-model.number="formData.pdId" :clearable="true" placeholder="请输入商品起始爬虫的编号" />
        </el-form-item>
        <el-form-item label="结束编号:" prop="endId">
          <el-input v-model.number="formData.endId" :clearable="true" placeholder="请输入商品结束爬虫的编号" />
        </el-form-item>
        <el-form-item label="CPU:" prop="cpu">
          <el-input v-model="formData.cpu" :clearable="true" placeholder="请输入CPU" />
        </el-form-item>
        <el-form-item label="内存:" prop="memory">
          <el-input v-model="formData.memory" :clearable="true" placeholder="请输入内存" />
        </el-form-item>
        <el-form-item label="硬盘:" prop="disk">
          <el-input v-model="formData.disk" :clearable="true" placeholder="请输入硬盘" />
        </el-form-item>
        <el-form-item label="流量:" prop="traffic">
          <el-input v-model="formData.traffic" :clearable="true" placeholder="请输入流量" />
        </el-form-item>
        <el-form-item label="带宽:" prop="portSpeed">
          <el-input v-model="formData.portSpeed" :clearable="true" placeholder="请输入带宽" />
        </el-form-item>
        <el-form-item label="地址:" prop="location">
          <el-input v-model="formData.location" :clearable="true" placeholder="请输入地址" />
        </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input v-model="formData.price" :clearable="true" placeholder="请输入价格" />
        </el-form-item>
        <el-form-item label="其他信息:" prop="additional">
          <el-input v-model="formData.additional" :clearable="true" placeholder="请输入其他信息" />
        </el-form-item>
        <el-form-item label="链接:" prop="url">
          <el-input v-model="formData.url" :clearable="true" placeholder="请输入链接" />
        </el-form-item>
        <el-form-item label="历史库存:" prop="oldStock">
          <el-input v-model.number="formData.oldStock" :clearable="true" placeholder="请输入历史库存" />
        </el-form-item>
        <el-form-item label="现有库存:" prop="stock">
          <el-input v-model.number="formData.stock" :clearable="true" placeholder="请输入现有库存" />
        </el-form-item>
        <el-form-item label="消息ID:" prop="messageId">
          <el-input v-model="formData.messageId" :clearable="true" placeholder="请输入消息ID" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="TAG">
          {{ detailFrom.tag }}
        </el-descriptions-item>
        <el-descriptions-item label="标志">
          {{ detailFrom.flag }}
        </el-descriptions-item>
        <el-descriptions-item label="计费类型">
          {{ detailFrom.billingType }}
        </el-descriptions-item>
        <el-descriptions-item label="起始编号">
          {{ detailFrom.pdId }}
        </el-descriptions-item>
        <el-descriptions-item label="结束编号">
          {{ detailFrom.endId }}
        </el-descriptions-item>
        <el-descriptions-item label="CPU">
          {{ detailFrom.cpu }}
        </el-descriptions-item>
        <el-descriptions-item label="内存">
          {{ detailFrom.memory }}
        </el-descriptions-item>
        <el-descriptions-item label="硬盘">
          {{ detailFrom.disk }}
        </el-descriptions-item>
        <el-descriptions-item label="流量">
          {{ detailFrom.traffic }}
        </el-descriptions-item>
        <el-descriptions-item label="带宽">
          {{ detailFrom.portSpeed }}
        </el-descriptions-item>
        <el-descriptions-item label="地址">
          {{ detailFrom.location }}
        </el-descriptions-item>
        <el-descriptions-item label="价格">
          {{ detailFrom.price }}
        </el-descriptions-item>
        <el-descriptions-item label="其他信息">
          {{ detailFrom.additional }}
        </el-descriptions-item>
        <el-descriptions-item label="链接">
          {{ detailFrom.url }}
        </el-descriptions-item>
        <el-descriptions-item label="历史库存">
          {{ detailFrom.oldStock }}
        </el-descriptions-item>
        <el-descriptions-item label="现有库存">
          {{ detailFrom.stock }}
        </el-descriptions-item>
        <el-descriptions-item label="消息ID">
          {{ detailFrom.messageId }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

  </div>
</template>

<script setup>
import {
  createFindallpd,
  deleteFindallpd,
  deleteFindallpdByIds,
  updateFindallpd,
  findFindallpd,
  getFindallpdList
} from '@/api/findallpd/findallpd'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
  name: 'Findallpd'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  tag: '',
  flag: '',
  billingType: '',
  pdId: undefined,
  endId: undefined,
  cpu: '',
  memory: '',
  disk: '',
  traffic: '',
  portSpeed: '',
  location: '',
  price: '',
  additional: '',
  url: '',
  oldStock: undefined,
  stock: undefined,
  messageId: '',
})



// 验证规则
const rule = reactive({
  tag: [{
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
  flag: [{
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
  billingType: [{
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
  endId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
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
  const table = await getFindallpdList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteFindallpdFunc(row)
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
    const res = await deleteFindallpdByIds({ IDs })
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
const updateFindallpdFunc = async (row) => {
  const res = await findFindallpd({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteFindallpdFunc = async (row) => {
  const res = await deleteFindallpd({ ID: row.ID })
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
    tag: '',
    flag: '',
    billingType: '',
    pdId: undefined,
    endId: undefined,
    cpu: '',
    memory: '',
    disk: '',
    traffic: '',
    portSpeed: '',
    location: '',
    price: '',
    additional: '',
    url: '',
    oldStock: undefined,
    stock: undefined,
    messageId: '',
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createFindallpd(formData.value)
        break
      case 'update':
        res = await updateFindallpd(formData.value)
        break
      default:
        res = await createFindallpd(formData.value)
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
  const res = await findFindallpd({ ID: row.ID })
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