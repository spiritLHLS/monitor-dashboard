<template>
    <div class="product-dashboard">
        <header class="top-bar">
            <div class="left-section">
                <img src="~@/assets/logo.png" alt="Logo" class="logo">
                <nav class="nav-links">
                    <el-button type="primary" @click="openExternalLink('https://www.spiritlhl.net')">一键虚拟化项目</el-button>
                    <el-button type="primary" @click="openExternalLink('https://t.me/vps_spiders')">监控频道</el-button>
                    <el-button type="primary" @click="router.push('/home')">返回</el-button>
                </nav>
            </div>
        </header>
        <div class="content-wrapper">
            <aside class="sidebar">
                <h1 class="site-title">订阅配置</h1>
                <el-form @submit.prevent="handleSearch" class="search-form">
                    <template v-for="(field, key) in searchFields" :key="key">
                        <!-- 字符串类型字段 -->
                        <div v-if="field.type === 'string'" class="search-field">
                            <label class="field-label">{{ field.label }}</label>
                            <el-input 
                                v-model="searchInfo[key]"
                                :placeholder="field.placeholder" 
                                class="search-input">
                                <template #prefix>
                                    <el-icon>
                                        <Search />
                                    </el-icon>
                                </template>
                            </el-input>
                        </div>
                        <!-- 单个数值字段 -->
                        <div v-else-if="field.type === 'number'" class="search-field">
                            <label class="field-label">{{ field.label }}</label>
                            <el-input 
                                v-model="searchInfo[key]"
                                :placeholder="field.placeholder"
                                type="number"
                                class="search-input"
                            />
                        </div>
                        <!-- 下拉选择字段 -->
                        <div v-else-if="field.type === 'select'" class="search-field">
                            <label class="field-label">{{ field.label }}</label>
                            <el-select 
                                v-model="searchInfo[key]"
                                :placeholder="'请选择' + field.label"
                                class="search-select"
                                clearable
                            >
                                <el-option 
                                    v-for="option in field.options" 
                                    :key="option.value"
                                    :label="option.label" 
                                    :value="option.value"
                                />
                            </el-select>
                        </div>
                        <!-- 区间类型字段 -->
                        <div v-else-if="field.type === 'range'" class="search-field range-field">
                            <label class="field-label">{{ field.label }}</label>
                            <div class="range-inputs">
                                <el-input 
                                    v-model="searchInfo[`${key}Min`]"
                                    :placeholder="field.minPlaceholder"
                                    type="number"
                                    class="range-input"
                                />
                                <span class="range-separator">-</span>
                                <el-input 
                                    v-model="searchInfo[`${key}Max`]"
                                    :placeholder="field.maxPlaceholder"
                                    type="number"
                                    class="range-input"
                                />
                            </div>
                        </div>
                    </template>
                    <div class="display-toggle">
                        <span :class="['toggle-option', { 'active': displayMode === 'all' }]"
                            @click="handleDisplayModeChange('all')">所有商品</span>
                        <span :class="['toggle-option', { 'active': displayMode === 'subscribed' }]"
                            @click="handleDisplayModeChange('subscribed')">已订阅商品</span>
                    </div>
                    <div class="button-group">
                        <el-button type="primary" @click="handleSearch" class="search-button">搜索</el-button>
                        <el-button type="primary" plain @click="handleReset" class="reset-button">重置</el-button>
                    </div>
                </el-form>
            </aside>
            <main class="main-content">
                <el-card class="table-card">
                    <div class="batch-actions">
                        <template v-if="displayMode === 'all'">
                            <el-button type="primary" @click="handleBatchSubscribe"
                                :disabled="selectedProducts.length === 0">
                                批量订阅
                            </el-button>
                            <el-select v-model="batchNotifyChannel" placeholder="选择通知方式" style="margin-left: 10px;">
                                <el-option label="Telegram" value="telegram_bot" />
                                <el-option label="Email" value="email" />
                                <el-option label="微信" value="wechat" />
                            </el-select>
                        </template>
                        <template v-else>
                            <el-button type="primary" @click="handleBatchModifySubscribe"
                                :disabled="selectedProducts.length === 0">
                                批量修改订阅
                            </el-button>
                            <el-select v-model="batchModifyNotifyChannel" placeholder="选择通知方式"
                                style="margin-left: 10px;">
                                <el-option label="Telegram" value="telegram_bot" />
                                <el-option label="Email" value="email" />
                                <el-option label="微信" value="wechat" />
                            </el-select>
                            <el-button type="danger" @click="handleBatchUnsubscribe"
                                :disabled="selectedProducts.length === 0" style="margin-left: 10px;">
                                批量取消订阅
                            </el-button>
                            <el-button type="warning" @click="handleBatchResetStatus"
                                :disabled="selectedProducts.length === 0 || !selectedProducts.some(p => p.status === 1)"
                                style="margin-left: 10px;">
                                批量重置状态
                            </el-button>
                        </template>
                    </div>
                    <el-table v-loading="loading" :data="tableData" style="width: 100%"
                        :default-sort="{ prop: 'price', order: 'descending' }" @sort-change="handleSortChange"
                        @selection-change="handleSelectionChange">
                        <el-table-column type="selection" width="55" />
                        <el-table-column v-for="col in visibleColumns" :key="col.prop" :prop="col.prop"
                            :label="col.label" :sortable="col.sortable" :min-width="col.minWidth"
                            :show-overflow-tooltip="col.prop !== 'additional'">
                            <template #default="{ row }">
                                <template v-if="col.prop === 'stock'">
                                    {{ row.stock === 1000 ? '有' : row.stock }}
                                </template>
                                <template v-else-if="col.prop === 'notify_channel'">
                                    {{ getNotifyChannelLabel(row.notify_channel) }}
                                </template>
                                <template v-else-if="col.prop === 'additional'">
                                    <span v-html="processData(row.additional)" style="white-space: pre-wrap;"></span>
                                </template>
                                <template v-else>
                                    {{ row[col.prop] }}
                                </template>
                            </template>
                        </el-table-column>
                        <el-table-column v-if="displayMode === 'subscribed'" label="状态" :min-width="100">
                            <template #default="{ row }">
                                <template v-if="col.prop === 'stock'">
                                    {{ row.stock === 1000 ? '有' : row.stock === -1 ? '无' : row.stock }}
                                </template>
                                <template v-else-if="col.prop === 'notify_channel'">
                                    {{ getNotifyChannelLabel(row.notify_channel) }}
                                </template>
                                <template v-else-if="col.prop === 'priceUnit'">
                                    {{ getPriceUnitLabel(row.priceUnit) }}
                                </template>
                                <template v-else-if="col.prop === 'additional'">
                                    <span v-html="processData(row.additional)" style="white-space: pre-wrap;"></span>
                                </template>
                                <template v-else>
                                    {{ row[col.prop] }}
                                </template>
                            </template>
                        </el-table-column>
                        <el-table-column label="操作" :min-width="200">
                            <template #default="{ row }">
                                <el-button v-if="!row.isSubscribed" type="success" size="small"
                                    @click="handleSubscribe(row)">添加订阅</el-button>
                                <el-button v-else type="warning" size="small"
                                    @click="handleUnsubscribe(row)">取消订阅</el-button>
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
import { Search } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import { selfGetAllPd } from '@/api/subscribe/subscribe'
import { selfGetSub, selfCreateSub, selfDeleteSub, selfUpdateSub, selfBatchUpdateStatus } from '@/api/subscribe/subscribe'
const router = useRouter()
const searchFields = {
    tag: { label: '商家TAG', type: 'string', placeholder: '搜索TAG 等于' },
    cpu: { label: 'CPU', type: 'range', minPlaceholder: '最小值', maxPlaceholder: '最大值' },
    memory: { label: '内存(GB)', type: 'range', minPlaceholder: '最小值', maxPlaceholder: '最大值' },
    disk: { label: '磁盘(GB)', type: 'range', minPlaceholder: '最小值', maxPlaceholder: '最大值' },
    traffic: { label: '流量(TB)', type: 'range', minPlaceholder: '最小值', maxPlaceholder: '最大值' },
    portSpeed: { label: '端口(Gbps)', type: 'range', minPlaceholder: '最小值', maxPlaceholder: '最大值' },
    location: { label: '地点', type: 'string', placeholder: '搜索地点 等于' },
    price: { label: '价格(USD)', type: 'range', minPlaceholder: '最小值', maxPlaceholder: '最大值' },
    priceUnit: { 
        label: '计费周期', 
        type: 'select', 
        options: [
            { value: '', label: '全部' },
            { value: 'monthly', label: '月付' },
            { value: 'quarterly', label: '季付' },
            { value: 'semi-annually', label: '半年付' },
            { value: 'annually', label: '年付' },
            { value: 'biennially', label: '2年付' },
            { value: 'triennially', label: '3年付' },
        ]
    },
    stock: { label: '库存', type: 'number', placeholder: '搜索库存 大于' },
    additional: { label: '其他信息', type: 'string', placeholder: '搜索其他 关键词' },
}
const tableColumns = [
    { label: '商家', prop: 'tag', minWidth: '100' },
    { label: 'CPU', prop: 'cpu', minWidth: '120', sortable: true },
    { label: '内存', prop: 'memory', minWidth: '100', sortable: true },
    { label: '磁盘', prop: 'disk', minWidth: '100', sortable: true },
    { label: '流量', prop: 'traffic', minWidth: '100', sortable: true },
    { label: '端口', prop: 'portSpeed', minWidth: '100', sortable: true },
    { label: '地点', prop: 'location', minWidth: '150' },
    { label: '价格', prop: 'price', minWidth: '100', sortable: true },
    { label: '计费周期', prop: 'priceUnit', minWidth: '100' },
    { label: '库存', prop: 'stock', minWidth: '80', sortable: true },
    { label: '其他', prop: 'additional', minWidth: '300' },
    { label: '订阅渠道', prop: 'notify_channel', minWidth: '120' },
]
const initSearchInfo = () => {
    const info = {}
    Object.keys(searchFields).forEach(key => {
        const field = searchFields[key]
        if (field.type === 'range') {
            info[`${key}Min`] = null
            info[`${key}Max`] = null
        } else if (field.type === 'number') {
            info[key] = null
        } else if (field.type === 'select') {
            info[key] = ''
        } else {
            info[key] = ''
        }
    })
    return info
}
const getPriceUnitLabel = (value) => {
    const unitMap = {
        'monthly': '月付',
        'quarterly': '季付', 
        'semi-annually': '半年付',
        'annually': '年付',
        'biennially': '2年付',
        'triennially': '3年付'
    }
    return unitMap[value] || value
}
const searchInfo = ref(initSearchInfo())
const displayMode = ref('all')
const loading = ref(false)
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])
const sortBy = ref('stock')
const sortOrder = ref('desc')
const selectedProducts = ref([])
const subscribedProductIds = ref(new Set())
const batchNotifyChannel = ref('telegram_bot')
const batchModifyNotifyChannel = ref('telegram_bot')
const processData = (data) => {
    return data
        .trim()
        .replace(/\s+/g, ' ')
        .replace(/\n/g, '<br>');
}
const getNotifyChannelLabel = (value) => {
    const channelMap = {
        'telegram_bot': 'Telegram',
        'email': 'Email',
        'wechat': '微信'
    }
    return channelMap[value] || value
}
const visibleColumns = computed(() => {
    return tableColumns.filter(col =>
        col.prop !== 'notify_channel' || displayMode.value === 'subscribed'
    )
})
const handleDisplayModeChange = (mode) => {
    displayMode.value = mode
    page.value = 1
    getTableData()
}
const getTableData = async () => {
    loading.value = true
    try {
        const params = new URLSearchParams({
            page: page.value,
            pageSize: pageSize.value,
        })
        Object.entries(searchInfo.value).forEach(([key, value]) => {
            if (value != null && value !== '') {
                const baseKey = key.replace(/Min$|Max$/, '')
                const field = searchFields[baseKey]
                
                if (field && field.type === 'number') {
                    if (typeof value === 'number' || (typeof value === 'string' && !isNaN(Number(value)))) {
                        params.set(key, value.toString())
                    }
                } else if (field && field.type === 'select') {
                    if (typeof value === 'string' && value.trim() !== '') {
                        params.set(key, value)
                    }
                } else if (field && field.type === 'range') {
                    if (typeof value === 'number' || (typeof value === 'string' && !isNaN(Number(value)))) {
                        params.set(key, value.toString())
                    }
                } else if (typeof value === 'string' && value.trim() !== '') {
                    params.set(key, value)
                } else if (typeof value === 'number') {
                    params.set(key, value.toString())
                }
            }
        })
        if (sortBy.value) {
            params.set('sortBy', sortBy.value)
            params.set('sortOrder', sortOrder.value)
        }
        let response
        if (displayMode.value === 'all') {
            response = await selfGetAllPd(params)
        } else if (displayMode.value === 'subscribed') {
            response = await selfGetSub(params)
        }
        if (response.code === 0) {
            tableData.value = response.data.list.map(item => ({
                ...item,
                id: item.ID,
                isSubscribed: displayMode.value === 'subscribed' || subscribedProductIds.value.has(item.ID),
                status: item.status
            }))
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
const handleResetStatus = async (row) => {
    if (row.status === 0) {
        ElMessage.info('该订阅状态正常，无需重置')
        return
    }
    try {
        const response = await selfBatchUpdateStatus({ IDs: [row.id] })
        if (response.code === 0) {
            ElMessage.success('状态重置成功')
            row.status = 0
        } else {
            ElMessage.error(response.message || '状态重置失败')
        }
    } catch (error) {
        console.error('状态重置出错:', error)
        ElMessage.error('状态重置出错')
    }
}
const handleBatchResetStatus = async () => {
    const frozenProducts = selectedProducts.value.filter(product => product.status === 1)
    if (frozenProducts.length === 0) {
        ElMessage.info('没有需要重置的订阅')
        return
    }
    try {
        const IDs = frozenProducts.map(product => product.id)
        const response = await selfBatchUpdateStatus({ IDs })
        if (response.code === 0) {
            ElMessage.success('批量重置状态成功')
            frozenProducts.forEach(product => product.status = 0)
        } else {
            ElMessage.error(response.message || '批量重置状态失败')
        }
    } catch (error) {
        console.error('批量重置状态出错:', error)
        ElMessage.error('批量重置状态出错')
    }
}
const getSubscribedProducts = async () => {
    try {
        const response = await selfGetSub({ page: 1, pageSize: 1000 })
        if (response.code === 0) {
            subscribedProductIds.value = new Set(response.data.list.map(item => item.ID))
        }
    } catch (error) {
        console.error('获取已订阅商品出错:', error)
    }
}
const handleSearch = () => {
    page.value = 1
    getTableData()
}
const handleReset = () => {
    Object.keys(searchFields).forEach(key => {
        const field = searchFields[key]
        if (field.type === 'range') {
            searchInfo.value[`${key}Min`] = null
            searchInfo.value[`${key}Max`] = null
        } else if (field.type === 'number') {
            searchInfo.value[key] = null
        } else if (field.type === 'select') {
            searchInfo.value[key] = ''
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
const handleSubscribe = async (row) => {
    try {
        const response = await selfCreateSub({ product_id: row.id, notify_channel: batchNotifyChannel.value })
        if (response.code === 0) {
            ElMessage.success('订阅成功')
            subscribedProductIds.value.add(row.id)
            await getTableData() // 重新获取当前页数据
        } else {
            ElMessage.error(response.message || '订阅失败')
        }
    } catch (error) {
        console.error('订阅出错:', error)
        ElMessage.error('订阅出错')
    }
}
const handleUnsubscribe = async (row) => {
    try {
        const response = await selfDeleteSub({ product_id: row.id })
        if (response.code === 0) {
            ElMessage.success('取消订阅成功')
            subscribedProductIds.value.delete(row.id)
            if (displayMode.value === 'subscribed') {
                await getTableData() // 重新获取当前页数据
            } else {
                row.isSubscribed = false
                row.notify_channel = null
            }
        } else {
            ElMessage.error(response.message || '取消订阅失败')
        }
    } catch (error) {
        console.error('取消订阅出错:', error)
        ElMessage.error('取消订阅出错')
    }
}
const handleSelectionChange = (selection) => {
    selectedProducts.value = selection
}
const handleBatchSubscribe = async () => {
    for (const product of selectedProducts.value) {
        if (!product.isSubscribed) {
            await handleSubscribe(product)
        }
    }
    getTableData()
}
const handleBatchUnsubscribe = async () => {
    for (const product of selectedProducts.value) {
        if (product.isSubscribed) {
            await handleUnsubscribe(product)
        }
    }
    getTableData()
}
const handleBatchModifySubscribe = async () => {
    try {
        for (const product of selectedProducts.value) {
            if (product.isSubscribed) {
                const response = await selfUpdateSub({ product_id: product.id, notify_channel: batchModifyNotifyChannel.value })
                if (response.code === 0) {
                    product.notify_channel = batchModifyNotifyChannel.value
                } else {
                    ElMessage.error(`修改商品 ${product.id} 订阅失败: ${response.message}`)
                }
            }
        }
        ElMessage.success('批量修改订阅成功')
        getTableData()
    } catch (error) {
        console.error('批量修改订阅出错:', error)
        ElMessage.error('批量修改订阅出错')
    }
}
const openExternalLink = (url) => {
    window.open(url, '_blank')
}
onMounted(async () => {
    await getSubscribedProducts()
    getTableData()
})
watch([page, pageSize, sortBy, sortOrder, displayMode], () => {
    getTableData()
})
</script>

<style scoped>
.product-dashboard {
    display: flex;
    flex-direction: column;
    height: 100vh;
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

.nav-links .el-button {
    font-size: 14px;
    padding: 8px 16px;
}

.content-wrapper {
    display: flex;
    flex: 1;
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

.search-field {
    margin-bottom: 12px;
}

.field-label {
    display: block;
    margin-bottom: 4px;
    font-size: 14px;
    color: #2f3f2f;
    font-weight: 500;
}

.range-field {
    margin-bottom: 12px;
}

.range-inputs {
    display: flex;
    align-items: center;
    gap: 8px;
}

.range-input {
    flex: 1;
}

.range-separator {
    color: #666;
    font-weight: bold;
    min-width: 12px;
    text-align: center;
}

/* 下拉选择框样式 */
.search-select {
    width: 100%;
}

.search-select :deep(.el-input__wrapper) {
    box-shadow: 0 0 0 1px #c0cfc0 inset;
}

.search-select :deep(.el-input__wrapper:hover) {
    box-shadow: 0 0 0 1px #42b883 inset;
}

.search-select :deep(.el-input__wrapper.is-focus) {
    box-shadow: 0 0 0 2px rgba(66, 184, 131, 0.2);
}

/* 下拉选项样式 */
:deep(.el-select-dropdown__item.is-selected) {
    color: #42b883;
    font-weight: bold;
}

:deep(.el-select-dropdown__item:hover) {
    background-color: #e8f5e8;
}

.display-toggle {
    display: flex;
    background-color: #f0f6f0;
    border-radius: 20px;
    padding: 4px;
    margin: 12px 0;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
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
    color: #ffffff;
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

.batch-actions {
    margin-bottom: 16px;
    display: flex;
    align-items: center;
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

:deep(.el-button--danger) {
    background-color: #f56c6c;
    border-color: #f56c6c;
}

:deep(.el-button--danger:hover) {
    background-color: #f78989;
    border-color: #f78989;
}

:deep(.el-select) {
    width: 120px;
}

:deep(.el-table__body-wrapper) {
    overflow-x: auto;
}

:deep(.el-table .cell) {
    white-space: normal;
    line-height: 1.5;
}

/* 响应式设计 */
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

    .nav-links {
        margin-top: 10px;
        width: 100%;
    }

    .nav-links .el-button {
        width: 100%;
        margin-bottom: 5px;
    }

    .content-wrapper {
        flex-direction: column;
    }

    .sidebar {
        width: 100%;
        margin-bottom: 20px;
    }

    .main-content {
        padding: 10px;
    }

    .range-inputs {
        gap: 4px;
    }

    .range-separator {
        min-width: 8px;
        font-size: 12px;
    }

    /* 移动端下拉选择框调整 */
    .search-select {
        width: 100%;
    }
    
    .batch-actions {
        flex-direction: column;
        gap: 10px;
        align-items: stretch;
    }
    
    .batch-actions .el-button,
    .batch-actions .el-select {
        width: 100%;
        margin-left: 0 !important;
        margin-bottom: 5px;
    }
}

/* 小屏幕适配 */
@media (max-width: 480px) {
    .sidebar {
        padding: 15px;
    }
    
    .search-field {
        margin-bottom: 10px;
    }
    
    .field-label {
        font-size: 13px;
    }
    
    .range-inputs {
        flex-direction: column;
        gap: 6px;
    }
    
    .range-separator {
        display: none;
    }
    
    .range-input {
        width: 100%;
    }
    
    .button-group {
        flex-direction: column;
        gap: 8px;
    }
    
    .search-button,
    .reset-button {
        max-width: 100%;
        width: 100%;
    }
}
</style>