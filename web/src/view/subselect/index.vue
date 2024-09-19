<template>
    <div class="product-dashboard">
        <div class="content-wrapper">
            <aside class="sidebar">
                <h1 class="site-title">订阅配置</h1>
                <el-form @submit.prevent="handleSearch" class="search-form">
                    <el-input v-for="(field, key) in searchFields" :key="key" v-model="searchInfo[key]"
                        :placeholder="field.placeholder" class="search-input">
                        <template #prefix>
                            <el-icon>
                                <Search />
                            </el-icon>
                        </template>
                    </el-input>
                    <div class="display-toggle">
                        <el-radio-group v-model="displayMode" @change="handleDisplayModeChange">
                            <el-radio-button label="all">所有商品</el-radio-button>
                            <el-radio-button label="subscribed">已订阅商品</el-radio-button>
                        </el-radio-group>
                    </div>
                    <div class="button-group">
                        <el-button type="primary" @click="handleSearch" class="search-button">搜索</el-button>
                        <el-button type="primary" plain @click="handleReset" class="reset-button">重置</el-button>
                    </div>
                </el-form>
            </aside>

            <main class="main-content">
                <el-card class="table-card">
                    <div class="batch-actions" v-if="displayMode === 'all'">
                        <el-button type="primary" @click="handleBatchSubscribe" :disabled="selectedProducts.length === 0">
                            批量订阅
                        </el-button>
                    </div>
                    <el-table v-loading="loading" :data="tableData" style="width: 100%"
                        :default-sort="{ prop: 'price', order: 'descending' }" @sort-change="handleSortChange"
                        @selection-change="handleSelectionChange">
                        <el-table-column type="selection" width="55" v-if="displayMode === 'all'" />
                        <el-table-column v-for="col in tableColumns" :key="col.prop" :prop="col.prop" :label="col.label"
                            :sortable="col.sortable" :min-width="col.minWidth" show-overflow-tooltip>
                            <template #default="{ row }" v-if="col.prop === 'actions'">
                                <el-button type="primary" size="small" @click="showDetails(row)">查看详情</el-button>
                                <el-button v-if="!row.isSubscribed" type="success" size="small"
                                    @click="handleSubscribe(row)">添加订阅</el-button>
                                <el-button v-else type="warning" size="small"
                                    @click="handleUnsubscribe(row)">取消订阅</el-button>
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
import { ref, reactive, onMounted, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Search } from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'
import { selfGetAllPd } from '@/api/subscribe/subscribe'
import { selfGetSub, selfCreateSub, selfDeleteSub } from '@/api/subscribe/subscribe'

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
    { label: '操作', prop: 'actions', minWidth: '220' },
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
const selectedProducts = ref([])
const subscribedProductIds = ref(new Set())

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
        if (sortBy.value) {
            params.set('sortBy', sortBy.value)
            params.set('sortOrder', sortOrder.value)
        }

        let response
        if (displayMode.value === 'all') {
            response = await selfGetAllPd(params)
        } else {
            response = await selfGetSub(params)
        }

        if (response.code === 0) {
            tableData.value = response.data.list.map(item => ({
                id: item.ID,
                createdAt: item.CreatedAt,
                updatedAt: item.UpdatedAt,
                tag: item.tag,
                cpu: item.cpu,
                memory: item.memory,
                disk: item.disk,
                traffic: item.traffic,
                portSpeed: item.portSpeed,
                location: item.location,
                price: item.price,
                additional: item.additional,
                oldStock: item.oldStock,
                stock: item.stock,
                isSubscribed: displayMode.value === 'subscribed' || subscribedProductIds.value.has(item.ID)
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

const handleDisplayModeChange = () => {
    page.value = 1
    getTableData()
}

const handleSubscribe = async (row) => {
    try {
        const response = await selfCreateSub({product_id: row.id, notify_channel: 'telegram_bot' })
        if (response.code === 0) {
            ElMessage.success('订阅成功')
            row.isSubscribed = true
            subscribedProductIds.value.add(row.id)
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
            if (displayMode.value === 'subscribed') {
                tableData.value = tableData.value.filter(item => item.id !== row.id)
            } else {
                row.isSubscribed = false
            }
            subscribedProductIds.value.delete(row.id)
        } else {
            ElMessage.error(response.message || '取消订阅失败')
        }
    } catch (error) {
        console.error('取消订阅出错:', error)
        ElMessage.error('取消订阅出错')
    }
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

watch([page, pageSize, sortBy, sortOrder, displayMode], () => {
    getTableData()
})

onMounted(async () => {
    await getSubscribedProducts()
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

.display-toggle {
    margin: 12px 0;
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

@media (max-width: 768px) {
    .content-wrapper {
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