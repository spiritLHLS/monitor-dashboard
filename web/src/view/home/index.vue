<template>
    <div class="product-dashboard">
        <header class="top-bar">
            <div class="left-section">
                <img src="https://raw.githubusercontent.com/spiritlhls/pages/main/logo.png" alt="Logo" class="logo">
                <nav class="nav-links">
                    <el-button type="primary" @click="openExternalLink('https://t.me/vps_reviews')">商家评价</el-button>
                    <el-button type="primary" @click="openExternalLink('https://t.me/vps_spiders')">监控频道</el-button>
                    <el-button type="primary" @click="openExternalLink('https://www.spiritlhl.net')">一键虚拟化项目</el-button>
                    <el-button type="primary" @click="router.push('/about')">关于</el-button>
                </nav>
            </div>
            <div class="auth-buttons">
                <el-button type="primary" @click="router.push('/login')">登录/注册</el-button>
            </div>
        </header>

        <div class="content-wrapper">
            <aside class="sidebar">
                <h1 class="site-title">全球VPS余量监控</h1>
                <el-form @submit.prevent="handleSearch" class="search-form">
                    <el-input v-for="(field, key) in searchFields" :key="key" v-model="searchInfo[key]"
                        :placeholder="field.placeholder" class="search-input">
                        <template #prefix>
                            <el-icon>
                                <Search />
                            </el-icon>
                        </template>
                    </el-input>
                    <div class="stock-toggle">
                        <span :class="['toggle-option', { 'active': displayMode === 'all' }]"
                            @click="displayMode = 'all'">显示所有</span>
                        <span :class="['toggle-option', { 'active': displayMode === 'inStock' }]"
                            @click="displayMode = 'inStock'">只显示有库存</span>
                        <span :class="['toggle-option', { 'active': displayMode === 'dedicatedOnly' }]"
                            @click="displayMode = 'dedicatedOnly'">仅独立服务器</span>
                    </div>
                    <div class="button-group">
                        <el-button type="primary" @click="handleSearch" class="search-button">搜索</el-button>
                        <el-button type="primary" plain @click="handleReset" class="reset-button">重置</el-button>
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
                            :page-sizes="[10, 20, 50, 100]" :total="total"
                            layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange"
                            @current-change="handleCurrentChange" />
                    </div>
                </el-card>
            </main>
        </div>

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
import { useRouter } from 'vue-router'

const router = useRouter()

const searchFields = {
    tag: { placeholder: '搜索TAG 等于' },
    cpu: { placeholder: '搜索CPU 等于' },
    memory: { placeholder: '搜索内存 等于' },
    disk: { placeholder: '搜索磁盘 等于' },
    traffic: { placeholder: '搜索流量 等于' },
    portSpeed: { placeholder: '搜索端口 等于' },
    location: { placeholder: '搜索地点 等于' },
    price: { placeholder: '搜索价格 等于' },
    stock: { placeholder: '搜索库存 大于' },
    additional: { placeholder: '搜索其他 关键词' },
}

const tableColumns = [
    { label: '商家', prop: 'tag', minWidth: '120', sortable: true },
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
const displayMode = ref('all')
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const sortBy = ref('stock')
const sortOrder = ref('desc')
const detailsVisible = ref(false)
const selectedProduct = ref({})

const processData = (data) => {
    return data
        .replace(/<[^>]+>/g, '')
        .trim()
        .replace(/\s+/g, ' ');
}

const getTableData = async () => {
    loading.value = true
    try {
        const params = new URLSearchParams({
            page: page.value,
            pageSize: pageSize.value,
            ...Object.fromEntries(
                Object.entries(searchInfo.value).filter(([_, v]) => v != null && v !== '')
            )
        })
        if (displayMode.value === 'inStock') {
            params.set('stock', '1')
        } else if (displayMode.value === 'dedicatedOnly') {
            params.append('additional', '独服')
        }
        if (sortBy.value) {
            params.set('sortBy', sortBy.value)
            params.set('sortOrder', sortOrder.value)
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
    displayMode.value = 'all'
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
    selectedProduct.value = {
        ...row,
        additional: processData(row.additional),
    }
    detailsVisible.value = true
}

const handleCloseDetails = () => {
    detailsVisible.value = false
}

const openExternalLink = (url) => {
    window.open(url, '_blank')
}

watch([page, pageSize, sortBy, sortOrder, displayMode], () => {
    getTableData()
})

onMounted(() => {
    getTableData()
})
</script>

<style scoped>
.product-dashboard {
    display: flex;
    flex-direction: column;
    height: 100vh;
    overflow: hidden;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', 'Helvetica Neue', Helvetica, Arial, sans-serif;
    background-color: #f0f6f0;
}

.top-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px 20px;
    background-color: #e8f5e8;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.left-section {
    display: flex;
    align-items: center;
}

.logo {
    height: 40px;
    width: auto;
    margin-right: 20px;
}

.nav-links {
    display: flex;
    gap: 10px;
}

.nav-links .el-button,
.auth-buttons .el-button {
    font-size: 14px;
    padding: 8px 16px;
}

.auth-buttons {
    display: flex;
    gap: 10px;
}

.content-wrapper {
    display: flex;
    flex: 1;
    overflow: hidden;
}

.sidebar {
    width: 280px;
    background-color: #ffffff;
    padding: 20px;
    overflow-y: auto;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

.site-title {
    font-size: 24px;
    color: #2f3f2f;
    margin-bottom: 20px;
    text-align: center;
    font-weight: 600;
}

.search-form {
    display: flex;
    flex-direction: column;
    gap: 12px;
}

.stock-toggle {
    display: flex;
    background-color: #f0f6f0;
    border-radius: 20px;
    padding: 4px;
    margin: 12px 0;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    font-size: 14px;
}

.toggle-option {
    flex: 1;
    text-align: center;
    padding: 8px 12px;
    cursor: pointer;
    border-radius: 16px;
    transition: all 0.3s ease;
    color: #2f3f2f;
}

.toggle-option.active {
    background-color: #42b883;
    color: white;
    font-weight: 500;
    box-shadow: 0 2px 4px rgba(66, 184, 131, 0.2);
}

.button-group {
    display: flex;
    justify-content: space-between;
    margin-top: 12px;
}

.search-button,
.reset-button {
    flex: 1;
    max-width: 45%;
}

.main-content {
    flex-grow: 1;
    padding: 20px;
    overflow-y: auto;
    background-color: #f0f6f0;
}

.table-card {
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    border-radius: 4px;
}

.pagination-wrapper {
    padding: 16px;
    display: flex;
    justify-content: flex-end;
}

:deep(.el-input__wrapper) {
    box-shadow: 0 0 0 1px #c0cfc0 inset;
}

:deep(.el-input__wrapper:hover) {
    box-shadow: 0 0 0 1px #42b883 inset;
}

:deep(.el-input__wrapper.is-focus) {
    box-shadow: 0 0 0 2px rgba(66, 184, 131, 0.2);
}

:deep(.el-switch__core) {
    border-color: #c0cfc0;
}

:deep(.el-switch.is-checked .el-switch__core) {
    border-color: #42b883;
    background-color: #42b883;
}

:deep(.el-table th) {
    background-color: #f0f6f0;
    font-weight: 600;
    color: #2f3f2f;
}

:deep(.el-table--striped .el-table__body tr.el-table__row--striped td) {
    background-color: #f9fff9;
}

:deep(.el-table--enable-row-hover .el-table__body tr:hover > td) {
    background-color: #e8f5e8;
}

:deep(.el-button--primary) {
    background-color: #42b883;
    border-color: #42b883;
}

:deep(.el-button--primary:hover) {
    background-color: #33a06f;
    border-color: #33a06f;
}

:deep(.el-button--primary) {
    background-color: #42b883;
    border-color: #42b883;
}

:deep(.el-button--primary:hover) {
    background-color: #33a06f;
    border-color: #33a06f;
}

:deep(.el-button--primary.is-plain) {
    color: #42b883;
    background: #ecf5f0;
    border-color: #42b883;
}

:deep(.el-button--primary.is-plain:hover) {
    background: #42b883;
    color: #ffffff;
}

:deep(.el-button--success) {
    background-color: #42b883;
    border-color: #42b883;
}

:deep(.el-button--success:hover) {
    background-color: #33a06f;
    border-color: #33a06f;
}

@media (max-width: 768px) {
    .top-bar {
        flex-direction: column;
        align-items: flex-start;
        padding: 10px;
    }

    .left-section {
        flex-direction: column;
        align-items: flex-start;
    }

    .logo {
        margin-bottom: 10px;
    }

    .nav-links,
    .auth-buttons {
        margin-top: 10px;
        width: 100%;
    }

    .nav-links .el-button,
    .auth-buttons .el-button {
        width: 100%;
        margin-bottom: 5px;
    }

    .content-wrapper {
        flex-direction: column;
    }

    .sidebar {
        width: 100%;
        max-height: 300px;
        padding: 16px;
    }

    .main-content {
        height: calc(100vh - 300px - 120px);
    }

    .button-group {
        flex-direction: row;
    }

    .search-button,
    .reset-button {
        margin: 8px 0;
        max-width: 48%;
    }

    .stock-toggle {
        font-size: 12px;
    }

    .toggle-option {
        padding: 6px 8px;
    }
}
</style>