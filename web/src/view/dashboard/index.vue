<template>
  <div class="announcement-dashboard">
    <header class="top-bar">
      <div class="left-section">
        <img src="~@/assets/logo.png" alt="Logo" class="logo">
        <nav class="nav-links">
          <el-button type="primary" @click="openExternalLink('https://paste.spiritlhl.net')">商家评价</el-button>
          <el-button type="primary" @click="openExternalLink('https://t.me/vps_spiders')">监控频道</el-button>
          <el-button type="primary" @click="openExternalLink('https://www.spiritlhl.net')">一键虚拟化项目</el-button>
        </nav>
      </div>
    </header>

    <div class="content-wrapper">
      <main class="main-content">
        <el-card class="announcement-card">
          <div class="announcement">
            <div v-if="isFetching">加载中...</div>
            <div v-else-if="error">获取公告失败: {{ error }}</div>
            <div v-else>
              <h2>{{ announcement.title }}</h2>
              <div v-html="announcement.content"></div>
            </div>
          </div>
        </el-card>
      </main>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { GetInfoPublic } from '@/plugin/announcement/api/info'

export default {
  setup() {
    const router = useRouter()
    const isFetching = ref(false)
    const error = ref(null)
    const announcement = ref({})

    const fetchAnnouncement = async () => {
      isFetching.value = true
      error.value = null
      try {
        const response = await GetInfoPublic({ Title: "首页通知" })
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

    const openExternalLink = (url) => {
      window.open(url, '_blank')
    }

    onMounted(() => {
      fetchAnnouncement()
    })

    return {
      router,
      isFetching,
      error,
      announcement,
      openExternalLink
    }
  }
}
</script>

<style scoped>
.announcement-dashboard {
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

.nav-links .el-button {
  font-size: 14px;
  padding: 8px 16px;
}

.content-wrapper {
  display: flex;
  flex: 1;
  overflow: hidden;
}

.main-content {
  flex-grow: 1;
  padding: 20px;
  overflow-y: auto;
  background-color: #f0f6f0;
}

.announcement-card {
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  border-radius: 4px;
}

.announcement {
  padding: 20px;
}

.announcement h2 {
  margin-top: 0;
  color: #2f3f2f;
  font-size: 24px;
  margin-bottom: 20px;
}

.announcement div {
  padding: 10px 0;
}

.announcement p {
  margin-bottom: 10px;
  line-height: 1.6;
  color: #2f3f2f;
}

:deep(.el-button--primary) {
  background-color: #42b883;
  border-color: #42b883;
}

:deep(.el-button--primary:hover) {
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

  .nav-links {
    margin-top: 10px;
    width: 100%;
  }

  .nav-links .el-button {
    width: 100%;
    margin-bottom: 5px;
  }

  .main-content {
    padding: 10px;
  }

  .announcement h2 {
    font-size: 20px;
  }
}
</style>
