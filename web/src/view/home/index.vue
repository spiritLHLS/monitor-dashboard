<template>
    <div class="product-dashboard">
        <header class="top-bar">
            <div class="left-section">
                <img src="~@/assets/logo.png" alt="Logo" class="logo">
                <el-dropdown trigger="click" class="merchant-dropdown">
                    <el-button type="primary" class="dropdown-button">
                        热门商家
                        <el-icon class="el-icon--right">
                            <ArrowDown />
                        </el-icon>
                    </el-button>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item v-for="merchant in hotMerchants" :key="merchant" 
                                @click="searchByMerchant(merchant)">
                                {{ merchant }}
                            </el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
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
        <el-collapse v-model="activeCollapse" class="announcement-collapse">
            <el-collapse-item name="1">
                <template #title>
                    <el-icon>
                        <InfoFilled />
                    </el-icon>
                    <span>公告</span>
                </template>
                <div class="announcement-content">
                    <div v-if="isFetching">加载中...</div>
                    <div v-else-if="error">获取公告失败: {{ error }}</div>
                    <div v-else>
                        <div v-html="announcement.content"></div>
                    </div>
                </div>
            </el-collapse-item>
        </el-collapse>
        <div class="content-wrapper">
            <aside :class="['sidebar', { 'collapsed': isCollapsed }]">
                <div class="sidebar-content" v-show="!isCollapsed">
                    <h1 class="site-title">全球VPS余量监控</h1>
                    <el-form @submit.prevent="handleSearch" class="search-form">
                    <template v-for="(field, key) in searchFields" :key="key">
                        <!-- 字符串类型字段使用普通输入框 -->
                        <el-input 
                            v-if="field.type === 'string'" 
                            v-model="searchInfo[key]"
                            :placeholder="field.placeholder" 
                            class="search-input">
                            <template #prefix>
                                <el-icon>
                                    <Search />
                                </el-icon>
                            </template>
                        </el-input>
                        <!-- 数值类型字段使用数字输入框 -->
                        <el-input-number 
                            v-else-if="field.type === 'number'"
                            v-model="searchInfo[key]"
                            :placeholder="field.placeholder"
                            class="search-input number-input"
                            :precision="2"
                            :step="0.1"
                            :min="0"
                            controls-position="right"
                            style="width: 100%"
                        />
                    </template>
                    
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
                </div>
            </aside>
            <div class="toggle-container">
                <div class="sidebar-toggle" @click="toggleSidebar">
                    <el-icon>
                        <component :is="isCollapsed ? 'ArrowRight' : 'ArrowLeft'" />
                    </el-icon>
                </div>
            </div>
            <main :class="['main-content', { 'expanded': isCollapsed }]">
                <el-card class="table-card">
                    <el-table v-loading="loading" :data="tableData" style="width: 100%"
                        :default-sort="{ prop: 'price', order: 'descending' }" @sort-change="handleSortChange">
                        <el-table-column v-for="col in tableColumns" :key="col.prop" :prop="col.prop" :label="col.label"
                            :sortable="col.sortable" :min-width="col.minWidth"
                            :show-overflow-tooltip="col.prop !== 'additional' && col.prop !== 'actions'">
                            <template #default="{ row }">
                                <template v-if="col.prop === 'actions'">
                                    <el-button type="success" size="small"
                                        @click="handleRedirectFunc(row.url)">点击购买</el-button>
                                </template>
                                <template v-else-if="col.prop === 'stock'">
                                    {{ row.stock === 1000 ? '有' : row.stock }}
                                </template>
                                <template v-else-if="col.prop === 'additional'">
                                    <span v-html="processData(row[col.prop])" style="white-space: pre-wrap;"></span>
                                </template>
                                <template v-else>
                                    {{ row[col.prop] }}
                                </template>
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
    </div>
</template>

<script setup>
import { ref, reactive, onMounted, watch, computed } from 'vue'
import { ElMessage } from 'element-plus'
import { Search, InfoFilled, ArrowLeft, ArrowRight, ArrowDown } from '@element-plus/icons-vue'
import { getProductsPublic } from '@/api/products/products'
import { useRouter } from 'vue-router'
import { handleRedirect } from '@/plugin/cryptourl/api/encryptedlink'
import { GetInfoPublic } from '@/plugin/announcement/api/info'

const isCollapsed = ref(true)
const toggleSidebar = () => {
    isCollapsed.value = !isCollapsed.value
}
const activeCollapse = ref(['1'])
const announcement = ref({})
const isFetching = ref(false)
const error = ref(null)

// 热门商家列表
const hotMerchants = ref(['racknerd', 'buyvm', 'bagevm', 'spartanhost', 'fiberstate'])

// 按商家搜索功能
const searchByMerchant = (merchant) => {
    // 重置所有搜索条件
    Object.keys(searchFields).forEach(key => {
        if (searchFields[key].type === 'number') {
            searchInfo.value[key] = null
        } else {
            searchInfo.value[key] = ''
        }
    })
    // 设置商家搜索
    searchInfo.value.tag = merchant
    displayMode.value = 'all'
    page.value = 1
    getTableData()
    ElMessage.success(`正在搜索 ${merchant} 的产品`)
}

const fetchAnnouncement = async () => {
    isFetching.value = true
    error.value = null
    try {
        const response = await GetInfoPublic({ Title: "链接说明" })
        if (response.code === 0) {
            announcement.value = response.data
        } else {
            error.value = response.msg || '获取公告失败'
        }
    } catch (err) {
        console.error('Error fetching announcement:', err)
        error.value = '网络错误，请稍后再试'
    } finally {
        isFetching.value = false
    }
}

const router = useRouter()

const searchFields = {
    tag: { placeholder: '搜索TAG 等于', type: 'string' },
    cpu: { placeholder: '搜索CPU 等于', type: 'number' },
    memory: { placeholder: '搜索内存 等于', type: 'number' },
    disk: { placeholder: '搜索磁盘 等于', type: 'number' },
    traffic: { placeholder: '搜索流量 等于', type: 'number' },
    portSpeed: { placeholder: '搜索端口 等于', type: 'number' },
    location: { placeholder: '搜索地点 等于', type: 'string' },
    price: { placeholder: '搜索价格 等于', type: 'number' },
    stock: { placeholder: '搜索库存 大于', type: 'number' },
    additional: { placeholder: '搜索其他 关键词', type: 'string' },
}

const tableColumns = [
    { label: '商家', prop: 'tag', minWidth: '120' },
    { label: 'CPU', prop: 'cpu', minWidth: '100', sortable: true },
    { label: '内存(GB)', prop: 'memory', minWidth: '120', sortable: true },
    { label: '磁盘(GB)', prop: 'disk', minWidth: '120', sortable: true },
    { label: '流量(TB)', prop: 'traffic', minWidth: '120', sortable: true },
    { label: '端口(Gbps)', prop: 'portSpeed', minWidth: '120', sortable: true },
    { label: '地点', prop: 'location', minWidth: '150' },
    { label: '价格(USD)', prop: 'price', minWidth: '120', sortable: true },
    { label: '库存', prop: 'stock', minWidth: '80', sortable: true },
    { label: '其他', prop: 'additional', minWidth: '300' },
    { label: '操作', prop: 'actions', minWidth: '120' },
]

const displayMode = ref('all')
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const sortBy = ref('stock')
const sortOrder = ref('desc')
// 初始化搜索信息，区分字符串和数值类型
const initSearchInfo = () => {
    const info = {}
    Object.keys(searchFields).forEach(key => {
        if (searchFields[key].type === 'number') {
            info[key] = null // 数值字段初始化为 null
        } else {
            info[key] = '' // 字符串字段初始化为空字符串
        }
    })
    return info
}
const searchInfo = ref(initSearchInfo())

const processData = (data) => {
    return data
        .trim()
        .replace(/\s+/g, ' ')
        .replace(/\n/g, '<br>');
}

const getTableData = async () => {
    loading.value = true
    try {
        const params = new URLSearchParams({
            page: page.value,
            pageSize: pageSize.value,
        })
        
        // 添加搜索条件，区分处理不同类型的字段
        Object.entries(searchInfo.value).forEach(([key, value]) => {
            if (value != null && value !== '') {
                // 对于数值类型，直接设置值
                if (searchFields[key].type === 'number') {
                    if (typeof value === 'number') {
                        params.set(key, value.toString())
                    }
                } else {
                    // 对于字符串类型，检查是否为有效值
                    if (typeof value === 'string' && value.trim() !== '') {
                        params.set(key, value)
                    }
                }
            }
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
    // 根据字段类型重置不同的默认值
    Object.keys(searchFields).forEach(key => {
        if (searchFields[key].type === 'number') {
            searchInfo.value[key] = null
        } else {
            searchInfo.value[key] = ''
        }
    })
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

const handleRedirectFunc = async (rowshortCode) => {
    try {
        if (rowshortCode.toLowerCase().includes('http') && rowshortCode.includes(':')) {
            await openUrl(rowshortCode)
            return
        }
        const response = await handleRedirect({ shortCode: rowshortCode })
        if (response && response.redirectUrl) {
            await openUrl(response.redirectUrl)
        } else {
            ElMessage({
                type: 'error',
                message: '服务器返回的数据无效'
            })
        }
    } catch (error) {
        console.error('重定向错误:', error)
        ElMessage({
            type: 'error',
            message: '重定向失败: ' + (error.message || '未知错误')
        })
    }
}

const openUrl = (url) => {
    return new Promise((resolve) => {
        // 尝试打开新窗口
        const newWindow = window.open(url, '_blank')
        // 检查是否成功打开
        if (newWindow && !newWindow.closed) {
            ElMessage({
                type: 'success',
                message: '新窗口已打开'
            })
            resolve()
        } else {
            // 如果无法打开新窗口，提示用户并在当前窗口打开
            ElMessage({
                type: 'info',
                message: '无法打开新窗口，将在当前页面跳转...(请浏览器允许本站点打开弹出窗口)'
            })
            setTimeout(() => {
                window.location.href = url
                resolve()
            }, 2500) // 给用户2.5秒时间看到消息
        }
    })
}

const openExternalLink = (url) => {
    window.open(url, '_blank')
}

watch([page, pageSize, sortBy, sortOrder, displayMode], () => {
    getTableData()
})

onMounted(() => {
    fetchAnnouncement()
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
    font-size: 12px;
}

.top-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 8px 16px;
    background-color: #e8f5e8;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    flex-shrink: 0;
}

.left-section {
    display: flex;
    align-items: center;
    gap: 12px;
}

.logo {
    height: 32px;
    width: auto;
}

.merchant-dropdown {
    margin-left: 4px;
}

.dropdown-button {
    font-size: 12px;
    padding: 6px 12px;
    background-color: #42b883;
    border-color: #42b883;
    border-radius: 6px;
    transition: all 0.3s ease;
}

.dropdown-button:hover {
    background-color: #33a06f;
    border-color: #33a06f;
    transform: translateY(-1px);
    box-shadow: 0 3px 8px rgba(66, 184, 131, 0.3);
}

.nav-links {
    display: flex;
    gap: 8px;
}

.nav-links .el-button,
.auth-buttons .el-button {
    font-size: 12px;
    padding: 6px 12px;
}

.auth-buttons {
    display: flex;
    gap: 8px;
}

.announcement-collapse {
    margin: 8px 16px;
    background-color: #f0f6f0;
    border: 1px solid #e8f5e8;
    border-radius: 4px;
    flex-shrink: 0;
}

.content-wrapper {
    display: flex;
    flex: 1;
    overflow: hidden;
    position: relative;
    min-height: 0;
}

.sidebar {
    position: relative;
    transition: width 0.3s ease;
    width: 260px;
    background-color: #ffffff;
    padding: 16px;
    overflow-y: auto;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    z-index: 1;
    flex-shrink: 0;
}

.sidebar.collapsed {
    width: 40px;
    padding: 16px 0;
}

.toggle-container {
    position: absolute;
    left: 260px;
    top: 50%;
    transform: translateY(-50%);
    z-index: 2;
    transition: left 0.3s ease;
}

.sidebar.collapsed+.toggle-container {
    left: 40px;
}

.sidebar-toggle {
    width: 20px;
    height: 40px;
    background-color: #42b883;
    color: white;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    border-radius: 0 4px 4px 0;
    transition: all 0.3s ease;
}

.sidebar-toggle:hover {
    background-color: #33a06f;
    transform: translateX(2px);
}

.main-content {
    flex: 1;
    padding: 16px;
    padding-left: 36px;
    overflow-y: auto;
    background-color: #f0f6f0;
    transition: padding-left 0.3s ease;
    min-width: 0;
}

.main-content.expanded {
    padding-left: 56px;
}

.site-title {
    font-size: 20px;
    color: #2f3f2f;
    margin-bottom: 16px;
    text-align: center;
    font-weight: 600;
}

.search-form {
    display: flex;
    flex-direction: column;
    gap: 10px;
}

/* 数字输入框样式 */
.number-input {
    margin-bottom: 10px;
}

.stock-toggle {
    display: flex;
    background-color: #f0f6f0;
    border-radius: 20px;
    padding: 3px;
    margin: 10px 0;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    font-size: 11px;
}

.toggle-option {
    flex: 1;
    text-align: center;
    padding: 6px 10px;
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
    margin-top: 10px;
}

.search-button,
.reset-button {
    flex: 1;
    max-width: 45%;
    font-size: 12px;
}

.table-card {
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    border-radius: 4px;
    height: 100%;
    display: flex;
    flex-direction: column;
}

.table-card :deep(.el-card__body) {
    padding: 12px;
    flex: 1;
    display: flex;
    flex-direction: column;
    overflow: hidden;
}

.pagination-wrapper {
    padding: 12px 0;
    display: flex;
    justify-content: flex-end;
    flex-shrink: 0;
}

:deep(.el-table) {
    font-size: 11px;
    flex: 1;
    background-color: #ffffff !important;
    color: #2f3f2f !important;
}

:deep(.el-table__body-wrapper) {
    overflow-x: auto;
    background-color: #ffffff !important;
}

:deep(.el-table__body) {
    background-color: #ffffff !important;
}

:deep(.el-table tbody tr) {
    background-color: #ffffff !important;
    color: #2f3f2f !important;
}

:deep(.el-table th) {
    background-color: #f0f6f0 !important;
    font-weight: 600;
    color: #2f3f2f !important;
    font-size: 11px;
    padding: 8px 6px;
}

:deep(.el-table td) {
    padding: 6px 6px;
    font-size: 11px;
    background-color: #ffffff !important;
    color: #2f3f2f !important;
}

:deep(.el-table .cell) {
    padding: 0 4px;
    line-height: 1.3;
    word-break: break-all;
    color: #2f3f2f !important;
}

:deep(.el-table--striped .el-table__body tr.el-table__row--striped td) {
    background-color: #f9fff9 !important;
}

:deep(.el-table--enable-row-hover .el-table__body tr:hover > td) {
    background-color: #e8f5e8 !important;
}

:deep(.el-card) {
    background-color: #ffffff !important;
    border: 1px solid #e8f5e8 !important;
}

:deep(.el-card__body) {
    background-color: #ffffff !important;
    color: #2f3f2f !important;
}

:deep(.el-button--small) {
    font-size: 11px;
    padding: 4px 8px;
}

:deep(.el-input) {
    font-size: 12px;
}

:deep(.el-input__inner) {
    height: 32px;
    font-size: 12px;
}

:deep(.el-input__wrapper) {
    box-shadow: 0 0 0 1px #c0cfc0 inset;
    transition: all 0.3s ease;
}

:deep(.el-input__wrapper:hover) {
    box-shadow: 0 0 0 1px #42b883 inset;
}

:deep(.el-input__wrapper.is-focus) {
    box-shadow: 0 0 0 2px rgba(66, 184, 131, 0.2);
}

/* 数字输入框深度样式 */
:deep(.el-input-number) {
    width: 100%;
}

:deep(.el-input-number .el-input__wrapper) {
    box-shadow: 0 0 0 1px #c0cfc0 inset;
    transition: all 0.3s ease;
}

:deep(.el-input-number .el-input__wrapper:hover) {
    box-shadow: 0 0 0 1px #42b883 inset;
}

:deep(.el-input-number .el-input__wrapper.is-focus) {
    box-shadow: 0 0 0 2px rgba(66, 184, 131, 0.2);
}

:deep(.el-input-number .el-input__inner) {
    height: 32px;
    font-size: 12px;
    text-align: left;
}

:deep(.el-input-number__increase),
:deep(.el-input-number__decrease) {
    background-color: #f0f6f0;
    border-color: #c0cfc0;
    color: #42b883;
}

:deep(.el-input-number__increase:hover),
:deep(.el-input-number__decrease:hover) {
    background-color: #42b883;
    border-color: #42b883;
    color: white;
}

:deep(.el-switch__core) {
    border-color: #c0cfc0;
}

:deep(.el-switch.is-checked .el-switch__core) {
    border-color: #42b883;
    background-color: #42b883;
}

:deep(.el-button--primary) {
    background-color: #42b883;
    border-color: #42b883;
    transition: all 0.3s ease;
}

:deep(.el-button--primary:hover) {
    background-color: #33a06f;
    border-color: #33a06f;
    transform: translateY(-1px);
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

:deep(.el-collapse-item__header) {
    background-color: #e8f5e8;
    color: #2f3f2f;
    font-weight: bold;
    font-size: 12px;
    padding: 10px 16px;
}

:deep(.el-collapse-item__content) {
    padding: 16px;
    background-color: #ffffff;
}

.announcement-content {
    font-size: 12px;
    line-height: 1.5;
    color: #2f3f2f;
}

.announcement-content h3 {
    margin-bottom: 8px;
    color: #42b883;
    font-size: 14px;
}

:deep(.el-pagination) {
    font-size: 12px;
}

:deep(.el-pagination .el-pager li) {
    font-size: 12px;
    min-width: 28px;
    height: 28px;
    line-height: 28px;
    transition: all 0.3s ease;
}

:deep(.el-pagination .el-pager li:hover) {
    color: #42b883;
    transform: translateY(-1px);
}

:deep(.el-pagination .el-pager li.is-active) {
    color: #42b883;
}

:deep(.el-pagination .btn-prev, .el-pagination .btn-next) {
    font-size: 12px;
    min-width: 28px;
    height: 28px;
    line-height: 28px;
    transition: all 0.3s ease;
}

:deep(.el-pagination .btn-prev:hover, .el-pagination .btn-next:hover) {
    color: #42b883;
}

:deep(.el-pagination .el-pagination__sizes .el-select .el-input) {
    width: 90px;
}

/* 下拉菜单样式 */
:deep(.el-dropdown-menu) {
    border-radius: 8px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
    border: 1px solid #e8f5e8;
    background-color: #ffffff;
    padding: 4px;
}

:deep(.el-dropdown-menu__item) {
    font-size: 12px;
    padding: 8px 12px;
    border-radius: 4px;
    margin: 2px 0;
    transition: all 0.3s ease;
    color: #2f3f2f;
}

:deep(.el-dropdown-menu__item:hover) {
    background-color: #e8f5e8;
    color: #42b883;
    transform: translateX(2px);
}

@media (max-width: 1200px) {
    .sidebar {
        width: 220px;
    }

    .toggle-container {
        left: 220px;
    }

    :deep(.el-table .cell) {
        font-size: 10px;
    }
}

@media (max-width: 768px) {
    .top-bar {
        flex-direction: column;
        align-items: flex-start;
        padding: 8px;
        gap: 8px;
    }

    .left-section {
        flex-direction: column;
        align-items: flex-start;
        width: 100%;
        gap: 8px;
    }

    .logo {
        height: 28px;
    }

    .merchant-dropdown {
        width: 100%;
    }

    .dropdown-button {
        width: 100%;
        justify-content: center;
    }

    .nav-links,
    .auth-buttons {
        width: 100%;
        flex-wrap: wrap;
    }

    .nav-links .el-button,
    .auth-buttons .el-button {
        flex: 1;
        min-width: 120px;
        font-size: 11px;
        padding: 5px 10px;
        margin: 2px;
    }

    .content-wrapper {
        flex-direction: column;
    }

    .sidebar {
        width: 100%;
        max-height: 250px;
        padding: 12px;
    }

    .toggle-container {
        display: none;
    }

    .main-content {
        height: calc(100vh - 250px - 150px);
        padding: 12px;
    }

    .button-group {
        flex-direction: row;
    }

    .search-button,
    .reset-button {
        margin: 6px 0;
        max-width: 48%;
    }

    .stock-toggle {
        font-size: 10px;
    }

    .toggle-option {
        padding: 5px 6px;
    }

    .announcement-collapse {
        margin: 8px;
    }

    :deep(.el-table .cell) {
        font-size: 9px;
        padding: 0 2px;
    }

    :deep(.el-table th) {
        padding: 6px 4px;
    }

    :deep(.el-table td) {
        padding: 4px 4px;
    }
}

@media (max-width: 480px) {
    .product-dashboard {
        font-size: 10px;
    }

    .left-section {
        gap: 6px;
    }

    .nav-links .el-button,
    .auth-buttons .el-button {
        font-size: 10px;
        padding: 4px 8px;
        min-width: 100px;
    }

    :deep(.el-table .cell) {
        font-size: 8px;
    }

    :deep(.el-button--small) {
        font-size: 9px;
        padding: 3px 6px;
    }
}
</style>