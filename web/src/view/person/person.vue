<template>
  <div class="user-dashboard">
    <!-- 头部导航栏 -->
    <header class="top-bar">
      <div class="left-section">
        <img src="~@/assets/logo.png" alt="Logo" class="logo">
        <nav class="nav-links">
          <el-button type="primary" @click="openExternalLink('https://t.me/vps_reviews')">商家评价</el-button>
          <el-button type="primary" @click="openExternalLink('https://t.me/vps_spiders')">监控频道</el-button>
          <el-button type="primary" @click="openExternalLink('https://www.spiritlhl.net')">一键虚拟化项目</el-button>
          <el-button type="primary" @click="router.push('/home')">返回</el-button>
        </nav>
      </div>
    </header>

    <!-- 公告栏 -->
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
          <div v-else-if="announcement && announcement.content" v-html="announcement.content"></div>
          <div v-else>暂无公告</div>
        </div>
      </el-collapse-item>
    </el-collapse>

    <!-- 主要内容区 -->
    <div class="content-wrapper">
      <div class="grid-container">
        <!-- 用户侧边栏 -->
        <aside class="user-sidebar">
          <div class="user-card">
            <div class="avatar-container">
              <SelectImage v-model="userInfo.avatar" file-type="image" @change="updateAvatar" />
            </div>
            <h2 class="user-name">{{ userInfo.nickname }}</h2>
            <ul class="user-details">
              <li><el-icon>
                  <user />
                </el-icon>{{ userInfo.nickname }}</li>
              <li><el-icon>
                  <message />
                </el-icon>{{ userInfo.email }}</li>
              <li><el-icon><chat-dot-round /></el-icon>{{ userInfo.tgID }}</li>
            </ul>
          </div>
        </aside>

        <!-- 主要内容 -->
        <main class="main-content">
          <el-card class="info-card">
            <el-tabs v-model="activeName" @tab-click="handleClick">
              <el-tab-pane label="账号信息" name="info">
                <el-form :model="userInfo" label-width="100px">
                  <el-form-item label="昵称">
                    <el-input v-model="userInfo.nickname" />
                  </el-form-item>
                  <el-form-item label="邮箱">
                    <el-input v-model="userInfo.email" />
                  </el-form-item>
                  <el-form-item label="Telegram ID">
                    <el-input v-model="userInfo.tgID" />
                  </el-form-item>
                  <el-form-item label="新密码">
                    <el-input v-model="newPassword" type="password" show-password />
                  </el-form-item>
                  <el-form-item>
                    <el-button type="primary" @click="updateUserInfo">保存修改</el-button>
                  </el-form-item>
                </el-form>
              </el-tab-pane>
            </el-tabs>
          </el-card>
        </main>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { selfModifyInfo, selfGetUserInfo } from '@/api/ecsusers/ecsusers.js'
import SelectImage from '@/components/selectImage/selectImage.vue'
import { GetInfoPublic } from '@/plugin/announcement/api/info'

const router = useRouter()
const activeCollapse = ref(['1'])
const announcement = ref({ content: '' })
const isFetching = ref(false)
const error = ref(null)
const activeName = ref('info')

const userInfo = reactive({
  uuid: '',
  nickname: '',
  avatar: '',
  tgID: '',
  email: '',
})

const newPassword = ref('')

const fetchAnnouncement = async () => {
  isFetching.value = true
  error.value = null
  try {
    const response = await GetInfoPublic({ Title: "推送说明" })
    if (response.code === 0 && response.data) {
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

const updateAvatar = async () => {
  await updateUserInfo()
}

const handleClick = (tab, event) => {
  console.log(tab, event)
}

const updateUserInfo = async () => {
  try {
    const updateData = {
      uuid: userInfo.uuid,
      nickname: userInfo.nickname,
      avatar: userInfo.avatar,
      tgID: userInfo.tgID,
      email: userInfo.email,
    }

    if (newPassword.value) {
      updateData.password = newPassword.value
    }

    const res = await selfModifyInfo(updateData)
    if (res.code === 0) {
      ElMessage.success('更新信息成功')
      if (newPassword.value) {
        newPassword.value = ''
      }
    } else {
      ElMessage.error('更新信息失败')
    }
  } catch (error) {
    console.error('更新用户信息出错:', error)
    ElMessage.error('更新用户信息出错')
  }
}

const getUserInfo = async () => {
  try {
    const res = await selfGetUserInfo()
    if (res.code === 0 && res.data) {
      Object.assign(userInfo, res.data)
    } else {
      ElMessage.error('获取用户信息失败')
    }
  } catch (error) {
    console.error('获取用户信息出错:', error)
    ElMessage.error('获取用户信息出错')
  }
}

const openExternalLink = (url) => {
  window.open(url, '_blank')
}

onMounted(() => {
  fetchAnnouncement()
  getUserInfo()
})
</script>

<style scoped>
.user-dashboard {
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

.grid-container {
  display: grid;
  grid-template-columns: 300px 1fr;
  gap: 20px;
  width: 100%;
  padding: 20px;
}

.user-sidebar {
  background-color: #fff;
  border-radius: 4px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  padding: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.user-card {
  text-align: center;
  width: 100%;
}

.avatar-container {
  margin-bottom: 20px;
  display: flex;
  justify-content: center;
}

.user-name {
  font-size: 24px;
  color: #2f3f2f;
  margin-bottom: 30px;
}

.user-details {
  list-style-type: none;
  padding: 0;
  text-align: left;
  width: 100%;
}

.user-details li {
  display: flex;
  align-items: center;
  margin-bottom: 15px;
  color: #2f3f2f;
}

.user-details .el-icon {
  margin-right: 10px;
}

.main-content {
  flex-grow: 1;
  overflow-y: auto;
}

.info-card {
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  border-radius: 4px;
}

:deep(.el-button--primary) {
  background-color: #42b883;
  border-color: #42b883;
}

:deep(.el-button--primary:hover) {
  background-color: #33a06f;
  border-color: #33a06f;
}

.announcement-collapse {
  margin: 10px 20px;
  background-color: #f0f6f0;
  border: 1px solid #e8f5e8;
  border-radius: 4px;
}

:deep(.el-collapse-item__header) {
  background-color: #e8f5e8;
  color: #2f3f2f;
  font-weight: bold;
}

:deep(.el-collapse-item__content) {
  padding: 20px;
  background-color: #ffffff;
}

.announcement-content {
  font-size: 14px;
  line-height: 1.6;
  color: #2f3f2f;
}

.announcement-content h3 {
  margin-bottom: 10px;
  color: #42b883;
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

  .grid-container {
    grid-template-columns: 1fr;
  }

  .user-sidebar {
    margin-bottom: 20px;
  }

  .main-content {
    padding: 10px;
  }

  .announcement-collapse {
    margin: 10px;
  }
}
</style>