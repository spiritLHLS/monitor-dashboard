
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
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
      

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button v-auth="btnAuth.add" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button v-auth="btnAuth.batchDelete" icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
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
        
          <el-table-column align="left" label="tag字段" prop="tag" width="120" />
          <el-table-column align="left" label="cpu字段" prop="cpu" width="120" />
          <el-table-column align="left" label="memory字段" prop="memory" width="120" />
          <el-table-column align="left" label="disk字段" prop="disk" width="120" />
          <el-table-column align="left" label="traffic字段" prop="traffic" width="120" />
          <el-table-column align="left" label="portSpeed字段" prop="portSpeed" width="120" />
          <el-table-column align="left" label="location字段" prop="location" width="120" />
          <el-table-column align="left" label="price字段" prop="price" width="120" />
                      <el-table-column label="additional字段" prop="additional" width="200">
                         <template #default="scope">
                            [富文本内容]
                         </template>
                      </el-table-column>
          <el-table-column align="left" label="url字段" prop="url" width="120" />
          <el-table-column align="left" label="billingType字段" prop="billingType" width="120" />
          <el-table-column align="left" label="oldStock字段" prop="oldStock" width="120" />
          <el-table-column align="left" label="stock字段" prop="stock" width="120" />
          <el-table-column align="left" label="是否3次检测再推送" prop="multiCheck" width="120" />
          <el-table-column align="left" label="间隔为-1时不进行爬虫" prop="intervals" width="120" />
          <el-table-column align="left" label="自动填充-帖子ID" prop="messageId" width="120" />
          <el-table-column align="left" label="间隔为-1时不进行推送" prop="pushIntervals" width="120" />
         <el-table-column align="left" label="自动填充-最新推送的时间" prop="pushTime" width="180">
            <template #default="scope">{{ formatDate(scope.row.pushTime) }}</template>
         </el-table-column>
          <el-table-column align="left" label="创建者" prop="createdBy" width="120" />
          <el-table-column align="left" label="更新者" prop="updatedBy" width="120" />
          <el-table-column align="left" label="删除者" prop="deletedBy" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button v-auth="btnAuth.info" type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看详情</el-button>
            <el-button v-auth="btnAuth.edit" type="primary" link icon="edit" class="table-button" @click="updateProductsFunc(scope.row)">变更</el-button>
            <el-button v-auth="btnAuth.delete" type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="tag字段:"  prop="tag" >
              <el-input v-model="formData.tag" :clearable="true"  placeholder="请输入tag字段" />
            </el-form-item>
            <el-form-item label="cpu字段:"  prop="cpu" >
              <el-input v-model="formData.cpu" :clearable="true"  placeholder="请输入cpu字段" />
            </el-form-item>
            <el-form-item label="memory字段:"  prop="memory" >
              <el-input v-model="formData.memory" :clearable="true"  placeholder="请输入memory字段" />
            </el-form-item>
            <el-form-item label="disk字段:"  prop="disk" >
              <el-input v-model="formData.disk" :clearable="true"  placeholder="请输入disk字段" />
            </el-form-item>
            <el-form-item label="traffic字段:"  prop="traffic" >
              <el-input v-model="formData.traffic" :clearable="true"  placeholder="请输入traffic字段" />
            </el-form-item>
            <el-form-item label="portSpeed字段:"  prop="portSpeed" >
              <el-input v-model="formData.portSpeed" :clearable="true"  placeholder="请输入portSpeed字段" />
            </el-form-item>
            <el-form-item label="location字段:"  prop="location" >
              <el-input v-model="formData.location" :clearable="true"  placeholder="请输入location字段" />
            </el-form-item>
            <el-form-item label="price字段:"  prop="price" >
              <el-input v-model="formData.price" :clearable="true"  placeholder="请输入price字段" />
            </el-form-item>
            <el-form-item label="additional字段:"  prop="additional" >
              <RichEdit v-model="formData.additional"/>
            </el-form-item>
            <el-form-item label="url字段:"  prop="url" >
              <el-input v-model="formData.url" :clearable="true"  placeholder="请输入url字段" />
            </el-form-item>
            <el-form-item label="billingType字段:"  prop="billingType" >
              <el-input v-model="formData.billingType" :clearable="true"  placeholder="请输入billingType字段" />
            </el-form-item>
            <el-form-item label="oldStock字段:"  prop="oldStock" >
              <el-input v-model.number="formData.oldStock" :clearable="true" placeholder="请输入oldStock字段" />
            </el-form-item>
            <el-form-item label="stock字段:"  prop="stock" >
              <el-input v-model.number="formData.stock" :clearable="true" placeholder="请输入stock字段" />
            </el-form-item>
            <el-form-item label="是否3次检测再推送:"  prop="multiCheck" >
              <el-input v-model.number="formData.multiCheck" :clearable="true" placeholder="请输入是否3次检测再推送" />
            </el-form-item>
            <el-form-item label="间隔为-1时不进行爬虫:"  prop="intervals" >
              <el-input v-model.number="formData.intervals" :clearable="true" placeholder="请输入间隔为-1时不进行爬虫" />
            </el-form-item>
            <el-form-item label="自动填充-帖子ID:"  prop="messageId" >
              <el-input v-model="formData.messageId" :clearable="true"  placeholder="请输入自动填充-帖子ID" />
            </el-form-item>
            <el-form-item label="间隔为-1时不进行推送:"  prop="pushIntervals" >
              <el-input v-model.number="formData.pushIntervals" :clearable="true" placeholder="请输入间隔为-1时不进行推送" />
            </el-form-item>
            <el-form-item label="自动填充-最新推送的时间:"  prop="pushTime" >
              <el-date-picker v-model="formData.pushTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="创建者:"  prop="createdBy" >
              <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入创建者" />
            </el-form-item>
            <el-form-item label="更新者:"  prop="updatedBy" >
              <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入更新者" />
            </el-form-item>
            <el-form-item label="删除者:"  prop="deletedBy" >
              <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入删除者" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="tag字段">
                        {{ detailFrom.tag }}
                    </el-descriptions-item>
                    <el-descriptions-item label="cpu字段">
                        {{ detailFrom.cpu }}
                    </el-descriptions-item>
                    <el-descriptions-item label="memory字段">
                        {{ detailFrom.memory }}
                    </el-descriptions-item>
                    <el-descriptions-item label="disk字段">
                        {{ detailFrom.disk }}
                    </el-descriptions-item>
                    <el-descriptions-item label="traffic字段">
                        {{ detailFrom.traffic }}
                    </el-descriptions-item>
                    <el-descriptions-item label="portSpeed字段">
                        {{ detailFrom.portSpeed }}
                    </el-descriptions-item>
                    <el-descriptions-item label="location字段">
                        {{ detailFrom.location }}
                    </el-descriptions-item>
                    <el-descriptions-item label="price字段">
                        {{ detailFrom.price }}
                    </el-descriptions-item>
                    <el-descriptions-item label="additional字段">
                        {{ detailFrom.additional }}
                    </el-descriptions-item>
                    <el-descriptions-item label="url字段">
                        {{ detailFrom.url }}
                    </el-descriptions-item>
                    <el-descriptions-item label="billingType字段">
                        {{ detailFrom.billingType }}
                    </el-descriptions-item>
                    <el-descriptions-item label="oldStock字段">
                        {{ detailFrom.oldStock }}
                    </el-descriptions-item>
                    <el-descriptions-item label="stock字段">
                        {{ detailFrom.stock }}
                    </el-descriptions-item>
                    <el-descriptions-item label="是否3次检测再推送">
                        {{ detailFrom.multiCheck }}
                    </el-descriptions-item>
                    <el-descriptions-item label="间隔为-1时不进行爬虫">
                        {{ detailFrom.intervals }}
                    </el-descriptions-item>
                    <el-descriptions-item label="自动填充-帖子ID">
                        {{ detailFrom.messageId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="间隔为-1时不进行推送">
                        {{ detailFrom.pushIntervals }}
                    </el-descriptions-item>
                    <el-descriptions-item label="自动填充-最新推送的时间">
                        {{ detailFrom.pushTime }}
                    </el-descriptions-item>
                    <el-descriptions-item label="创建者">
                        {{ detailFrom.createdBy }}
                    </el-descriptions-item>
                    <el-descriptions-item label="更新者">
                        {{ detailFrom.updatedBy }}
                    </el-descriptions-item>
                    <el-descriptions-item label="删除者">
                        {{ detailFrom.deletedBy }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createProducts,
  deleteProducts,
  deleteProductsByIds,
  updateProducts,
  findProducts,
  getProductsList
} from '@/api/products/products'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
// 引入按钮权限标识
import { useBtnAuth } from '@/utils/btnAuth'




defineOptions({
    name: 'Products'
})
// 按钮权限实例化
    const btnAuth = useBtnAuth()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
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
            oldStock: undefined,
            stock: undefined,
            multiCheck: undefined,
            intervals: undefined,
            messageId: '',
            pushIntervals: undefined,
            pushTime: new Date(),
            createdBy: undefined,
            updatedBy: undefined,
            deletedBy: undefined,
        })



// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
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
  elSearchFormRef.value?.validate(async(valid) => {
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
const getTableData = async() => {
  const table = await getProductsList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const setOptions = async () =>{
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
            deleteProductsFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
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

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateProductsFunc = async(row) => {
    const res = await findProducts({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
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
        oldStock: undefined,
        stock: undefined,
        multiCheck: undefined,
        intervals: undefined,
        messageId: '',
        pushIntervals: undefined,
        pushTime: new Date(),
        createdBy: undefined,
        updatedBy: undefined,
        deletedBy: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createProducts(formData.value)
                  break
                case 'update':
                  res = await updateProducts(formData.value)
                  break
                default:
                  res = await createProducts(formData.value)
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
  const res = await findProducts({ ID: row.ID })
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

</style>