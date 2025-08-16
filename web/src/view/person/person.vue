<template>
  <div class="user-dashboard">
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
              <li><el-icon><user /></el-icon>{{ userInfo.nickname }}</li>
              <li><el-icon><message /></el-icon>{{ userInfo.email }}</li>
              <li><el-icon><chat-dot-round /></el-icon>{{ userInfo.tgID }}</li>
            </ul>
          </div>
        </aside>

        <!-- 主要内容 -->
        <main class="main-content">
          <el-card class="info-card">
            <el-tabs v-model="activeName" @tab-click="handleClick">
              <el-tab-pane label="账号信息" name="info">
                <el-form 
                  :model="userInfo" 
                  :rules="rules" 
                  ref="formRef" 
                  label-width="100px"
                >
                  <el-form-item label="昵称">
                    <el-input v-model="userInfo.nickname" />
                  </el-form-item>
                  <el-form-item label="邮箱">
                    <el-input v-model="userInfo.email" />
                  </el-form-item>
                  <el-form-item label="Telegram ID">
                    <el-input v-model="userInfo.tgID" />
                  </el-form-item>
                  <el-form-item label="新密码" prop="newPassword">
                    <el-input 
                      v-model="newPassword" 
                      :type="passwordVisible ? 'text' : 'password'" 
                      show-password
                    >
                      <template #suffix>
                        <el-icon 
                          @click="togglePasswordVisibility" 
                          class="password-toggle-icon"
                        >
                          <view v-if="passwordVisible" />
                          <hide v-else />
                        </el-icon>
                      </template>
                    </el-input>
                  </el-form-item>
                  <el-form-item label="确认新密码" prop="confirmPassword">
                    <el-input 
                      v-model="confirmPassword" 
                      :type="confirmPasswordVisible ? 'text' : 'password'" 
                      show-password
                    >
                      <template #suffix>
                        <el-icon 
                          @click="toggleConfirmPasswordVisibility" 
                          class="password-toggle-icon"
                        >
                          <view v-if="confirmPasswordVisible" />
                          <hide v-else />
                        </el-icon>
                      </template>
                    </el-input>
                  </el-form-item>
                  <el-form-item>
                    <el-button type="primary" @click="validateAndUpdateUserInfo">保存修改</el-button>
                  </el-form-item>
                  <el-form-item label="通知测试">
                    <el-button type="success" @click="testEmailNotification">测试邮件通知</el-button>
                    <el-button type="warning" @click="testTelegramNotification">测试Telegram通知</el-button>
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
import { View, Hide } from '@element-plus/icons-vue'
import { selfModifyInfo, selfGetUserInfo } from '@/api/ecsusers/ecsusers.js'
import { CheckEmail, CheckTgBot } from '@/view/person/person.js'
import SelectImage from '@/components/selectImage/selectImage.vue'
import { GetInfoPublic } from '@/plugin/announcement/api/info'

const router = useRouter()
const activeCollapse = ref(['1'])
const announcement = ref({ content: '' })
const isFetching = ref(false)
const error = ref(null)
const activeName = ref('info')
const formRef = ref(null)

const userInfo = reactive({
  uuid: '',
  nickname: '',
  avatar: '',
  tgID: '',
  email: '',
})

const newPassword = ref('')
const confirmPassword = ref('')
const passwordVisible = ref(false)
const confirmPasswordVisible = ref(false)

// 密码验证规则
const rules = {
  newPassword: [
    { 
      validator: (rule, value, callback) => {
        if (value && value.length < 6) {
          callback(new Error('密码长度不能少于6位'))
        } else {
          callback()
        }
      }
    }
  ],
  confirmPassword: [
    { 
      validator: (rule, value, callback) => {
        if (newPassword.value && value !== newPassword.value) {
          callback(new Error('两次输入的密码不一致'))
        } else {
          callback()
        }
      }
    }
  ]
}

// 切换密码可见性
const togglePasswordVisibility = () => {
  passwordVisible.value = !passwordVisible.value
}

const toggleConfirmPasswordVisibility = () => {
  confirmPasswordVisible.value = !confirmPasswordVisible.value
}

// 验证并更新用户信息
const validateAndUpdateUserInfo = async () => {
  try {
    await formRef.value.validate()

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
      newPassword.value = ''
      confirmPassword.value = ''
    } else {
      ElMessage.error('更新信息失败')
    }
  } catch (error) {
    console.error('更新用户信息出错:', error)
    ElMessage.error('请检查表单信息')
  }
}

// 获取公告
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
  await validateAndUpdateUserInfo()
}

const handleClick = (tab, event) => {
  console.log(tab, event)
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

// 测试邮件通知
const testEmailNotification = async () => {
  try {
    const res = await CheckEmail()
    if (res.code === 0) {
      ElMessage.success('邮件通知测试成功')
    } else {
      ElMessage.error('邮件通知测试失败')
    }
  } catch (error) {
    console.error('邮件通知测试出错:', error)
    ElMessage.error('邮件通知测试出错')
  }
}

// 测试Telegram通知
const testTelegramNotification = async () => {
  try {
    const res = await CheckTgBot()
    if (res.code === 0) {
      ElMessage.success('Telegram通知测试成功')
    } else {
      ElMessage.error('Telegram通知测试失败')
    }
  } catch (error) {
    console.error('Telegram通知测试出错:', error)
    ElMessage.error('Telegram通知测试出错')
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

.password-toggle-icon {
  cursor: pointer;
  color: #42b883;
  margin-left: 10px;
}

.password-toggle-icon:hover {
  color: #2c3e50;
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