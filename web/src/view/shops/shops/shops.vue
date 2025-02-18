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
            <el-form-item label="消息ID" prop="mid">
              <el-input v-model="searchInfo.mid" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="商家类型" prop="type">
              <el-input v-model="searchInfo.type" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="推广链接" prop="affLink">
              <el-input v-model="searchInfo.affLink" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="商家链接" prop="shopLink">
              <el-input v-model="searchInfo.shopLink" placeholder="搜索条件" />
            </el-form-item>
          </div>
          <div class="row">
            <el-form-item label="额外标签" prop="additionalTags">
              <el-input v-model="searchInfo.additionalTags" placeholder="搜索条件" />
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

      </div>
      <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
        @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" />


        <el-table-column align="left" label="频道TAG" prop="tag" width="90" />
        <el-table-column align="left" label="消息编号" prop="mid" width="90" />
        <el-table-column align="left" label="商家类型" prop="type" width="90" />
        <el-table-column align="left" label="推广链接" prop="affLink" width="90" />
        <el-table-column align="left" label="商家链接" prop="shopLink" width="90" />
        <el-table-column align="left" label="额外标签" prop="additionalTags" width="90" />
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
              @click="updateShopsFunc(scope.row)">变更</el-button>
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
        <el-form-item label="频道TAG:" prop="tag">
          <el-input v-model="formData.tag" :clearable="true" placeholder="请输入频道TAG" />
        </el-form-item>
        <el-form-item label="消息编号:" prop="mid">
          <el-input v-model="formData.mid" :clearable="true" placeholder="请输入消息编号" />
        </el-form-item>
        <el-form-item label="商家类型:" prop="type">
          <el-select v-model="formData.type" placeholder="请选择商家类型" clearable>
            <el-option label="whmcs" value="whmcs"></el-option>
            <el-option label="hostbill" value="hostbill"></el-option>
            <el-option label="blesta" value="blesta"></el-option>
            <el-option label="other" value="other"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="推广链接:" prop="affLink">
          <el-input v-model="formData.affLink" :clearable="true" placeholder="请输入推广链接" />
        </el-form-item>
        <el-form-item label="商家链接:" prop="shopLink">
          <el-input v-model="formData.shopLink" :clearable="true" placeholder="请输入商家链接" />
        </el-form-item>
        <el-form-item label="额外标签:" prop="additionalTags">
          <el-input v-model="formData.additionalTags" :clearable="true" placeholder="请输入额外标签" />
        </el-form-item>
        <!-- <el-form-item label="创建者:"  prop="createdBy" >
              <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入创建者" />
            </el-form-item>
            <el-form-item label="更新者:"  prop="updatedBy" >
              <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入更新者" />
            </el-form-item>
            <el-form-item label="删除者:"  prop="deletedBy" >
              <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入删除者" />
            </el-form-item> -->
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="频道TAG">
          {{ detailFrom.tag }}
        </el-descriptions-item>
        <el-descriptions-item label="消息编号">
          {{ detailFrom.mid }}
        </el-descriptions-item>
        <el-descriptions-item label="商家类型">
          {{ detailFrom.type }}
        </el-descriptions-item>
        <el-descriptions-item label="推广链接">
          {{ detailFrom.affLink }}
        </el-descriptions-item>
        <el-descriptions-item label="商家链接">
          {{ detailFrom.shopLink }}
        </el-descriptions-item>
        <el-descriptions-item label="额外标签">
          {{ detailFrom.additionalTags }}
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
  createShops,
  deleteShops,
  deleteShopsByIds,
  updateShops,
  findShops,
  getShopsList
} from '@/api/shops/shops'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'




defineOptions({
  name: 'Shops'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
  const table = await getShopsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteShopsFunc(row)
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
    const res = await deleteShopsByIds({ IDs })
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
const updateShopsFunc = async (row) => {
  const res = await findShops({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteShopsFunc = async (row) => {
  const res = await deleteShops({ ID: row.ID })
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
    mid: '',
    type: '',
    affLink: '',
    shopLink: '',
    additionalTags: '',
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
  const res = await findShops({ ID: row.ID })
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