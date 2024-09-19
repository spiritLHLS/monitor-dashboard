<template>
  <div class="user-dashboard">
    <header class="top-bar">
      <div class="left-section">
        <img src="https://raw.githubusercontent.com/spiritlhls/pages/main/logo.png" alt="Logo" class="logo">
        <nav class="nav-links">
          <el-button type="primary" @click="openExternalLink('https://t.me/vps_reviews')">商家评价</el-button>
          <el-button type="primary" @click="openExternalLink('https://t.me/vps_spiders')">监控频道</el-button>
          <el-button type="primary" @click="openExternalLink('https://www.spiritlhl.net')">一键虚拟化项目</el-button>
          <el-button type="primary" @click="router.push('/home')">返回</el-button>
        </nav>
      </div>
    </header>

    <div class="content-wrapper">
      <div class="grid-container">
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

<script>
import { reactive, ref, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { selfModifyInfo, selfGetUserInfo } from '@/api/ecsusers/ecsusers.js'
import SelectImage from '@/components/selectImage/selectImage.vue'

export default {
  name: 'Person',
  components: {
    SelectImage
  },
  setup() {
    const router = useRouter()
    const activeName = ref('info')
    const userInfo = reactive({
      uuid: '',
      nickname: '',
      avatar: '',
      tgID: '',
      email: '',
    })
    const newPassword = ref('')

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

    onMounted(getUserInfo)

    return {
      router,
      activeName,
      userInfo,
      newPassword,
      updateAvatar,
      handleClick,
      updateUserInfo,
      openExternalLink
    }
  }
}
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
}
</style>
