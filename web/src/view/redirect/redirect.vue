<template>
    <div>
        <h1>正在重定向...</h1>
        <p v-if="error">{{ error }}</p>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { handleRedirect } from '@/plugin/cryptourl/api/encryptedlink'
import { ElMessage } from 'element-plus'

const route = useRoute()
const router = useRouter()
const error = ref('')

const redirect = async () => {
    try {
        const shortCode = route.params.shortCode
        const response = await handleRedirect({ shortCode })
        if (response && response.redirectUrl) {
            // 直接在当前窗口打开链接
            window.location.href = response.redirectUrl
        } else {
            error.value = '无效的重定向链接'
            ElMessage.error('无效的重定向链接')
        }
    } catch (err) {
        console.error('重定向错误:', err)
        error.value = '重定向失败: ' + (err.message || '未知错误')
        ElMessage.error(error.value)
    }
}

onMounted(() => {
    redirect()
})
</script>