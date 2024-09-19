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
        <el-form-item label="是否冻结" prop="isFrozen">
          <el-select v-model="searchInfo.isFrozen" clearable placeholder="请选择">
            <el-option key="true" label="是" value="true">
            </el-option>
            <el-option key="false" label="否" value="false">
            </el-option>
          </el-select>
        </el-form-item>
        <template v-if="showAllQuery">
          <div class="row">
            <el-form-item label="用户名" prop="username">
              <el-input v-model="searchInfo.username" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="TGID" prop="tgID">
              <el-input v-model="searchInfo.tgID" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="昵称" prop="nickname">
              <el-input v-model="searchInfo.nickname" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="密码" prop="password">
              <el-input v-model="searchInfo.password" placeholder="搜索条件" />
            </el-form-item>
          </div>
          <div class="row">
            <el-form-item label="邮箱" prop="email">
              <el-input v-model="searchInfo.email" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="备注" prop="additional">
              <el-input v-model="searchInfo.additional" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="QQ号" prop="qqNumber">
              <el-input v-model="searchInfo.qqNumber" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="微信号" prop="weChatNumber">
              <el-input v-model="searchInfo.weChatNumber" placeholder="搜索条件" />
            </el-form-item>
            <!-- <el-form-item label="推送渠道1" prop="pushChannel1">
              <el-input v-model="searchInfo.pushChannel1" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="推送渠道2" prop="pushChannel2">
              <el-input v-model="searchInfo.pushChannel2" placeholder="搜索条件" />
            </el-form-item>
            <el-form-item label="推送渠道3" prop="pushChannel3">
              <el-input v-model="searchInfo.pushChannel3" placeholder="搜索条件" />
            </el-form-item> -->
          </div>
          <div class="row">
            <el-form-item label="用户等级" prop="level">
              <el-input v-model.number="searchInfo.startLevel" placeholder="最小值" />
              —
              <el-input v-model.number="searchInfo.endLevel" placeholder="最大值" />
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
        <el-table-column align="left" label="用户名" prop="username" width="100" />
        <el-table-column align="left" label="TGID" prop="tgID" width="100" />
        <el-table-column align="left" label="昵称" prop="nickname" width="100" />
        <el-table-column align="left" label="密码" prop="password" width="100" />
        <el-table-column align="left" label="是否冻结" prop="isFrozen" width="100">
          <template #default="scope">{{ formatBoolean(scope.row.isFrozen) }}</template>
        </el-table-column>
        <el-table-column align="left" label="邮箱" prop="email" width="100" />
        <el-table-column align="left" label="备注" prop="additional" width="100" />
        <el-table-column align="left" label="用户等级" prop="level" width="100" />
        <el-table-column align="left" label="QQ号" prop="qqNumber" width="100" />
        <el-table-column align="left" label="微信号" prop="weChatNumber" width="100" />
        <el-table-column align="left" label="推送渠道1" prop="pushChannel1" width="100" />
        <el-table-column align="left" label="推送渠道2" prop="pushChannel2" width="100" />
        <el-table-column align="left" label="推送渠道3" prop="pushChannel3" width="100" />
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column label="头像" prop="avatar" width="200">
          <template #default="scope">
            <el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.avatar)"
              fit="cover" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="303">
          <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon
                style="margin-right: 5px">
                <InfoFilled />
              </el-icon>查看</el-button>
            <el-button type="primary" link icon="edit" class="table-button"
              @click="updateEcsUsersFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            <el-button type="primary" link icon="key" @click="adminChangePasswordFunc(scope.row)">修改密码</el-button>
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
        <el-form-item label="用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true" placeholder="请输入用户名" :disabled="type === 'update'" />
        </el-form-item>
        <el-form-item label="TGID:" prop="tgID">
          <el-input v-model="formData.tgID" :clearable="true" placeholder="请输入TGID" />
        </el-form-item>
        <el-form-item label="昵称:" prop="nickname">
          <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入昵称" />
        </el-form-item>
        <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true" placeholder="请输入密码" :disabled="type === 'update'" />
        </el-form-item>
        <el-form-item label="头像:" prop="avatar">
          <SelectImage v-model="formData.avatar" file-type="image" />
        </el-form-item>
        <el-form-item label="是否冻结:" prop="isFrozen">
          <el-switch v-model="formData.isFrozen" active-color="#13ce66" inactive-color="#ff4949" active-text="是"
            inactive-text="否" clearable></el-switch>
        </el-form-item>
        <el-form-item label="邮箱:" prop="email">
          <el-input v-model="formData.email" :clearable="true" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="备注:" prop="additional">
          <el-input v-model="formData.additional" :clearable="true" placeholder="请输入备注" />
        </el-form-item>
        <el-form-item label="用户等级:" prop="level">
          <el-input v-model.number="formData.level" :clearable="false" placeholder="请输入用户等级" />
        </el-form-item>
        <el-form-item label="QQ号:" prop="qqNumber">
          <el-input v-model="formData.qqNumber" :clearable="true" placeholder="请输入QQ号" />
        </el-form-item>
        <el-form-item label="微信号:" prop="weChatNumber">
          <el-input v-model="formData.weChatNumber" :clearable="true" placeholder="请输入微信号" />
        </el-form-item>
        <el-form-item label="推送渠道1:" prop="pushChannel1">
          <el-input v-model="formData.pushChannel1" :clearable="true" placeholder="请输入推送渠道1" />
        </el-form-item>
        <el-form-item label="推送渠道2:" prop="pushChannel2">
          <el-input v-model="formData.pushChannel2" :clearable="true" placeholder="请输入推送渠道2" />
        </el-form-item>
        <el-form-item label="推送渠道3:" prop="pushChannel3">
          <el-input v-model="formData.pushChannel3" :clearable="true" placeholder="请输入推送渠道3" />
        </el-form-item>
      </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="用户名">
          {{ detailFrom.username }}
        </el-descriptions-item>
        <el-descriptions-item label="TGID">
          {{ detailFrom.tgID }}
        </el-descriptions-item>
        <el-descriptions-item label="昵称">
          {{ detailFrom.nickname }}
        </el-descriptions-item>
        <el-descriptions-item label="头像">
          <el-image style="width: 50px; height: 50px" :preview-src-list="returnArrImg(detailFrom.avatar)"
            :src="getUrl(detailFrom.avatar)" fit="cover" />
        </el-descriptions-item>
        <el-descriptions-item label="密码">
          {{ detailFrom.password }}
        </el-descriptions-item>
        <el-descriptions-item label="是否冻结">
          {{ detailFrom.isFrozen }}
        </el-descriptions-item>
        <el-descriptions-item label="邮箱">
          {{ detailFrom.email }}
        </el-descriptions-item>
        <el-descriptions-item label="备注">
          {{ detailFrom.additional }}
        </el-descriptions-item>
        <el-descriptions-item label="用户等级">
          {{ detailFrom.level }}
        </el-descriptions-item>
        <el-descriptions-item label="QQ号">
          {{ detailFrom.qqNumber }}
        </el-descriptions-item>
        <el-descriptions-item label="微信号">
          {{ detailFrom.weChatNumber }}
        </el-descriptions-item>
        <el-descriptions-item label="推送渠道1">
          {{ detailFrom.pushChannel1 }}
        </el-descriptions-item>
        <el-descriptions-item label="推送渠道2">
          {{ detailFrom.pushChannel2 }}
        </el-descriptions-item>
        <el-descriptions-item label="推送渠道3">
          {{ detailFrom.pushChannel3 }}
        </el-descriptions-item>
      </el-descriptions>
    </el-drawer>

  </div>
</template>

<script setup>
import {
  createEcsUsers,
  deleteEcsUsers,
  deleteEcsUsersByIds,
  updateEcsUsers,
  findEcsUsers,
  getEcsUsersList,
  adminChangePassword
} from '@/api/ecsusers/ecsusers'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'EcsUsers'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  username: '',
  nickname: '',
  avatar: "",
  password: '',
  isFrozen: false,
  pushChannel1: '',
  pushChannel2: '',
  pushChannel3: '',
  tgID: '',
  qqNumber: '',
  weChatNumber: '',
  email: '',
  additional: '',
  level: undefined,
})



// 验证规则
const rule = reactive({
  username: [{
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
  nickname: [{
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
  password: [{
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
  tgID: [{
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
    if (searchInfo.value.isFrozen === "") {
      searchInfo.value.isFrozen = null
    }
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
  const table = await getEcsUsersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteEcsUsersFunc(row)
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
    const res = await deleteEcsUsersByIds({ IDs })
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
const updateEcsUsersFunc = async (row) => {
  const res = await findEcsUsers({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteEcsUsersFunc = async (row) => {
  const res = await deleteEcsUsers({ ID: row.ID })
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
    username: '',
    nickname: '',
    avatar: "",
    password: '',
    isFrozen: false,
    pushChannel1: '',
    pushChannel2: '',
    pushChannel3: '',
    tgID: '',
    qqNumber: '',
    weChatNumber: '',
    email: '',
    additional: '',
    level: undefined,
  }
}
// 弹窗确定
const enterDialog = async () => {
  elFormRef.value?.validate(async (valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createEcsUsers(formData.value)
        break
      case 'update':
        res = await updateEcsUsers(formData.value)
        break
      default:
        res = await createEcsUsers(formData.value)
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
  const res = await findEcsUsers({ ID: row.ID })
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

const adminChangePasswordFunc = (row) => {
  ElMessageBox.prompt('请输入你要修改的密码', '修改密码', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async ({ value }) => {
    const res = await adminChangePassword({ userID: row.ID, password: value})
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '修改成功'
      })
      getTableData()
    }
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