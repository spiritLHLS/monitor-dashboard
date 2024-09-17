<template>
  <div class="product-dashboard">
    <aside class="sidebar">
      <h1 class="site-title">监控</h1>
      <el-form @submit.prevent="handleSearch" class="search-form">
        <el-input v-for="(field, key) in searchFields" :key="key" v-model="searchInfo[key]"
          :placeholder="field.placeholder" class="search-input">
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </el-input>
        <el-switch v-model="onlyInStock" class="stock-switch" active-text="只显示有库存" inactive-text="显示所有" />
        <div class="button-group">
          <el-button type="primary" @click="handleSearch" class="search-button">搜索</el-button>
          <el-button @click="handleReset" class="reset-button">重置</el-button>
        </div>
      </el-form>
    </aside>

    <main class="main-content">
      <el-card class="table-card">
        <el-table v-loading="loading" :data="tableData" style="width: 100%"
          :default-sort="{ prop: 'price', order: 'descending' }" @sort-change="handleSortChange">
          <el-table-column v-for="col in tableColumns" :key="col.prop" :prop="col.prop" :label="col.label"
            :sortable="col.sortable" :min-width="col.minWidth" show-overflow-tooltip>
            <template #default="{ row }" v-if="col.prop === 'actions'">
              <el-button type="primary" size="small" @click="showDetails(row)">商品详情</el-button>
              <el-button type="success" size="small" @click="openUrl(row.url)">点击购买</el-button>
            </template>
            <template #default="{ row }" v-else-if="col.prop === 'stock'">
              {{ row.stock === 1000 ? '有' : row.stock }}
            </template>
          </el-table-column>
        </el-table>

        <div class="pagination-wrapper">
          <el-pagination v-model:current-page="page" v-model:page-size="pageSize"
            :page-sizes="[10, 20, 50, 100]" :total="total" layout="total, sizes, prev, pager, next, jumper"
            @size-change="handleSizeChange" @current-change="handleCurrentChange" />
        </div>
      </el-card>
    </main>

    <el-dialog v-model="detailsVisible" title="商品详情" width="50%" :before-close="handleCloseDetails">
      <el-descriptions :column="1" border>
        <el-descriptions-item v-for="col in tableColumns" :key="col.prop" :label="col.label">
          {{ col.prop === 'stock' && selectedProduct[col.prop] === 1000 ? '有' : selectedProduct[col.prop] }}
        </el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, watch, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { Search } from '@element-plus/icons-vue'
import { getProductsPublic } from '@/api/products/products'

const searchFields = {
    tag: { placeholder: '搜索TAG' }, // Like
    cpu: { placeholder: '搜索CPU' }, // Like
    memory: { placeholder: '搜索内存' }, // Like
    disk: { placeholder: '搜索磁盘' }, // Like
    traffic: { placeholder: '搜索流量' }, // Like
    portSpeed: { placeholder: '搜索端口' }, // Like
    location: { placeholder: '搜索地点' }, // Like
    price: { placeholder: '搜索价格' }, // Like
    stock: { placeholder: '搜索库存' }, // >
}

const tableColumns = [
    { label: 'TAG', prop: 'tag', minWidth: '120', sortable: true },
    { label: 'CPU', prop: 'cpu', minWidth: '100', sortable: true },
    { label: '内存', prop: 'memory', minWidth: '100', sortable: true },
    { label: '磁盘', prop: 'disk', minWidth: '100', sortable: true },
    { label: '流量', prop: 'traffic', minWidth: '100', sortable: true },
    { label: '端口', prop: 'portSpeed', minWidth: '80', sortable: true },
    { label: '地点', prop: 'location', minWidth: '140', sortable: true },
    { label: '价格', prop: 'price', minWidth: '150', sortable: true },
    { label: '库存', prop: 'stock', minWidth: '100', sortable: true },
    { label: '其他', prop: 'additional', minWidth: '100' },
    { label: '操作', prop: 'actions', minWidth: '200' },
]

const searchInfo = ref(Object.fromEntries(Object.keys(searchFields).map(key => [key, ''])))
const onlyInStock = ref(true)
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const sortBy = ref('stock')
const sortOrder = ref('desc')
const detailsVisible = ref(false)
const selectedProduct = ref({})

const getTableData = async () => {
    loading.value = true
    try {
        const params = new URLSearchParams({
            page: page.value,
            pageSize: pageSize.value,
            sortBy: sortBy.value,
            sortOrder: sortOrder.value,
            ...Object.fromEntries(
                Object.entries(searchInfo.value).filter(([_, v]) => v != null && v !== '')
            )
        })

        if (onlyInStock.value) {
            params.set('stock', '1')  // 设置 stock 参数为 1，表示查询库存大于等于 1 的商品
        }

        const response = await getProductsPublic(params)
        if (response.code === 0) {
            tableData.value = response.data.list
            total.value = response.data.total
        } else {
            ElMessage.error(response.message || '获取数据失败')
        }
    } catch (error) {
        console.error('获取数据出错:', error)
        ElMessage.error('获取数据出错')
    } finally {
        loading.value = false
    }
}

const handleSearch = () => {
    page.value = 1
    getTableData()
}

const handleReset = () => {
    Object.keys(searchInfo.value).forEach(key => searchInfo.value[key] = '')
    onlyInStock.value = true
    page.value = 1
    pageSize.value = 10
    sortBy.value = 'stock'
    sortOrder.value = 'desc'
    getTableData()
}

const handleSizeChange = (val) => {
    pageSize.value = val
    getTableData()
}

const handleCurrentChange = (val) => {
    page.value = val
    getTableData()
}

const handleSortChange = ({ prop, order }) => {
    sortBy.value = prop
    sortOrder.value = order === 'ascending' ? 'asc' : 'desc'
    getTableData()
}

const openUrl = (url) => {
    window.open(url, '_blank')
}

const showDetails = (row) => {
    selectedProduct.value = row
    detailsVisible.value = true
}

const handleCloseDetails = () => {
    detailsVisible.value = false
}

watch([page, pageSize, sortBy, sortOrder, onlyInStock], () => {
    getTableData()
})

onMounted(() => {
    getTableData()
})
</script>

<style scoped>
.product-dashboard {
  display: flex;
  height: 100vh;
  overflow: hidden;
  font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
  background-color: #f0f2f5;
}

.sidebar {
  width: 320px;
  background-color: #ffffff;
  padding: 24px;
  overflow-y: auto;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  display: flex;
  flex-direction: column;
}

.site-title {
  font-size: 28px;
  color: #333;
  margin-bottom: 28px;
  text-align: center;
  font-weight: bold;
}

.search-form {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.search-input {
  margin-bottom: 0;
}

.stock-switch {
  margin: 8px 0;
  align-self: flex-start;
}

.button-group {
  display: flex;
  justify-content: space-between;
  margin-top: 8px;
}

.search-button,
.reset-button {
  flex: 1;
  max-width: 45%;
}

.main-content {
  flex-grow: 1;
  padding: 24px;
  overflow-y: auto;
  background-color: #f0f2f5;
}

.table-card {
  box-shadow: 0 1px 2px -2px rgba(0, 0, 0, 0.08),
              0 3px 6px 0 rgba(0, 0, 0, 0.06),
              0 5px 12px 4px rgba(0, 0, 0, 0.04);
  border-radius: 8px;
}

.pagination-wrapper {
  padding: 16px;
  display: flex;
  justify-content: flex-end;
}

:deep(.el-input__wrapper) {
  box-shadow: 0 0 0 1px rgba(0, 0, 0, 0.1) inset;
}

:deep(.el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px #1890ff inset;
}

:deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
}

:deep(.el-switch__core) {
  border-color: #d9d9d9;
}

:deep(.el-switch.is-checked .el-switch__core) {
  border-color: #1890ff;
  background-color: #1890ff;
}

:deep(.el-table th) {
  background-color: #fafafa;
  font-weight: bold;
  color: #333;
}

:deep(.el-table--striped .el-table__body tr.el-table__row--striped td) {
  background-color: #fafafa;
}

:deep(.el-table--enable-row-hover .el-table__body tr:hover > td) {
  background-color: #e6f7ff;
}

:deep(.el-button--primary) {
  background-color: #1890ff;
  border-color: #1890ff;
}

:deep(.el-button--success) {
  background-color: #52c41a;
  border-color: #52c41a;
}

@media (max-width: 768px) {
  .product-dashboard {
    flex-direction: column;
  }

  .sidebar {
    width: 100%;
    max-height: 300px;
    padding: 16px;
  }

  .main-content {
    height: calc(100vh - 300px);
  }

  .button-group {
    flex-direction: row;
  }

  .search-button,
  .reset-button {
    margin: 8px 0;
    max-width: 48%;
  }
}
</style>