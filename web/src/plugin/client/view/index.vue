<template>
  <Transition name="fade" mode="out-in">
    <div class="app-container">
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
            <div v-else>
              <div v-html="announcement.content"></div>
            </div>
          </div>
        </el-collapse-item>
      </el-collapse>
      <div class="content-wrapper">
        <div class="login-container">
          <div class="login-box">
            <div class="login-header">
              <img class="login-logo" src="~@/assets/logo.png" alt="Logo" />
              <h1 class="login-title">ECS-SPIDERS</h1>
            </div>
            <el-form ref="loginForm" :model="currentFormData" :rules="formRules" class="login-form"
              @keyup.enter="submitForm">
              <el-form-item prop="username">
                <el-input v-model="currentFormData.username" placeholder="请输入用户名" prefix-icon="User"
                  @keyup.enter="focusNextInput($event, passwordInput)" />
              </el-form-item>

              <el-form-item prop="password">
                <el-input ref="passwordInput" v-model="currentFormData.password"
                  :type="lock === 'lock' ? 'password' : 'text'" placeholder="请输入密码" prefix-icon="Lock"
                  @keyup.enter="focusNextInput($event, confirmPasswordInput)">
                  <template #suffix>
                    <el-icon @click="changeLock" class="password-icon">
                      <component :is="lock" />
                    </el-icon>
                  </template>
                </el-input>
              </el-form-item>

              <el-form-item v-if="registerType" prop="confirmPassword">
                <el-input ref="confirmPasswordInput" v-model="currentFormData.confirmPassword"
                  :type="lock === 'lock' ? 'password' : 'text'" placeholder="请再次输入密码" prefix-icon="Lock"
                  @keyup.enter="focusNextInput($event, captchaInput)">
                  <template #suffix>
                    <el-icon @click="changeLock" class="password-icon">
                      <component :is="lock" />
                    </el-icon>
                  </template>
                </el-input>
              </el-form-item>

              <el-form-item prop="captcha">
                <div class="captcha-container">
                  <el-input ref="captchaInput" v-model="currentFormData.captcha" placeholder="请输入验证码"
                    class="captcha-input" />
                  <img v-if="picPath" :src="picPath" alt="验证码" @click="loginVerify" class="captcha-image" />
                </div>
              </el-form-item>

              <template v-if="registerType">
                <template v-if="showTGFields">
                  <el-form-item prop="tg_id">
                    <el-input v-model="currentFormData.tg_id" placeholder="请输入TGID" prefix-icon="ChatDotSquare" />
                  </el-form-item>

                  <el-form-item prop="code">
                    <div class="tg-code-container">
                      <el-input v-model="currentFormData.code" placeholder="请输入TG验证码" class="tg-code-input" />
                      <el-button type="primary" @click="sendTGCode" class="send-tg-code-btn">
                        发送验证码
                      </el-button>
                    </div>
                  </el-form-item>
                </template>

                <el-form-item v-if="showInviteCodeField" prop="invite_code">
                  <el-input v-model="currentFormData.invite_code" placeholder="请输入邀请码" prefix-icon="Ticket" />
                </el-form-item>
              </template>

              <el-form-item>
                <el-checkbox v-model="rememberMe">记住密码</el-checkbox>
              </el-form-item>

              <el-form-item>
                <el-button type="primary" @click="submitForm" class="submit-btn" :loading="loading">
                  {{ registerType ? '注册' : '登录' }}
                </el-button>
              </el-form-item>

              <div class="form-footer">
                <div class="toggle-container">
                  <span :class="['toggle-option', { active: !registerType }]" @click="registerType = false">
                    登录
                  </span>
                  <span :class="['toggle-option', { active: registerType }]" @click="registerType = true">
                    注册
                  </span>
                </div>
                <el-button link @click="goToResetPage" class="forgot-password">
                  忘记密码？
                </el-button>
              </div>
            </el-form>
          </div>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup>
import { captcha } from "@/api/user";
import { TGRGetCode, getTGRegisterStatus, getPublicInviteStatus } from "@/plugin/client/api/api";
import { reactive, ref, watch, onMounted, computed } from "vue";
import { ElMessage, ElLoading } from "element-plus";
import { useRouter } from "vue-router";
import { useUserStore } from "@/pinia/modules/user";
import { useRouterStore } from "@/pinia/modules/router";
import { GetInfoPublic } from '@/plugin/announcement/api/info'

const activeCollapse = ref(['1'])
const announcement = ref({})
const isFetching = ref(false)
const error = ref(null)
const router = useRouter();
const userStore = useUserStore();
const routerStore = useRouterStore();

const loginForm = ref(null);
const passwordInput = ref(null);
const confirmPasswordInput = ref(null);
const captchaInput = ref(null);

const loginFormData = reactive({
  username: "",
  password: "",
  captcha: "",
  captchaId: "",
});

const registerFormData = reactive({
  username: "",
  password: "",
  confirmPassword: "",
  captcha: "",
  captchaId: "",
  tg_id: "",
  code: "",
  invite_code: "",
});

const lock = ref("lock");
const picPath = ref("");
const registerType = ref(false);
const currentFormData = ref(loginFormData);
const loading = ref(false);
const rememberMe = ref(false);
const showTGFields = ref(false);
const showInviteCodeField = ref(false);

const fetchAnnouncement = async () => {
  isFetching.value = true
  error.value = null
  try {
    const response = await GetInfoPublic({ Title: "登录注册通知" })
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

const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error("用户名必须大于或等于5个字符"));
  } else {
    callback();
  }
};

const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error("密码必须大于或等于6个字符"));
  } else {
    callback();
  }
};

const validateConfirmPassword = (rule, value, callback) => {
  if (value !== currentFormData.value.password) {
    callback(new Error('两次输入的密码不一致'));
  } else {
    callback();
  }
};

const formRules = computed(() => ({
  username: [{ validator: checkUsername, trigger: "blur" }],
  password: [{ validator: checkPassword, trigger: "blur" }],
  confirmPassword: [{ validator: validateConfirmPassword, trigger: "blur" }],
  captcha: [
    { required: true, message: "请输入验证码", trigger: "blur" },
    { message: "验证码格式不正确", trigger: "blur" },
  ],
  ...(registerType.value && showTGFields.value
    ? {
      tg_id: [{ required: true, message: "请输入TG ID", trigger: "blur" }],
      code: [{ required: true, message: "请输入TG验证码", trigger: "blur" }],
    }
    : {}),
  ...(registerType.value && showInviteCodeField.value
    ? {
      invite_code: [{ required: true, message: "请输入邀请码", trigger: "blur" }],
    }
    : {}),
}));

const loginVerify = async () => {
  try {
    const ele = await captcha({});
    formRules.value.captcha[1].max = ele.data.captchaLength;
    formRules.value.captcha[1].min = ele.data.captchaLength;
    picPath.value = ele.data.picPath;
    currentFormData.value.captchaId = ele.data.captchaId;
    registerFormData.captchaId = ele.data.captchaId;
  } catch (error) {
    console.error('Failed to fetch captcha:', error);
    ElMessage.error('获取验证码失败，请刷新页面重试');
  }
};

const changeLock = () => {
  lock.value = lock.value === "lock" ? "unlock" : "lock";
};

const login = async () => {
  try {
    const res = await userStore.UserTgLogin(loginFormData);
    if (res) {
      if (rememberMe.value) {
        localStorage.setItem('rememberedUser', JSON.stringify({
          username: loginFormData.username,
          password: loginFormData.password
        }));
      } else {
        localStorage.removeItem('rememberedUser');
      }
      await routerStore.SetAsyncRouter();
      const asyncRouters = routerStore.asyncRouters;
      asyncRouters.forEach((asyncRouter) => {
        router.addRoute(asyncRouter);
      });
      router.push({ name: userStore.userInfo.authority.defaultRouter });
      return true;
    }
    return false;
  } catch (error) {
    return false;
  }
};

const register = async () => {
  try {
    const dataToSend = { ...registerFormData };
    if (!showTGFields.value) {
      dataToSend.tg_id = "";
      dataToSend.code = "";
    }
    if (!showInviteCodeField.value) {
      delete dataToSend.invite_code;
    }
    delete dataToSend.confirmPassword;

    let res;
    if (showInviteCodeField.value) {
      res = await userStore.UserRegisterWithInvite(dataToSend);
    } else {
      res = await userStore.UserTgRegister(dataToSend);
    }

    if (res) {
      await routerStore.SetAsyncRouter();
      const asyncRouters = routerStore.asyncRouters;
      asyncRouters.forEach((asyncRouter) => {
        router.addRoute(asyncRouter);
      });
      router.push({ name: userStore.userInfo.authority.defaultRouter });
      return true;
    }
    return false;
  } catch (error) {
    return false;
  }
};

const submitForm = async () => {
  loading.value = true;
  try {
    const form = loginForm.value;
    const validationResult = await new Promise((resolve) => {
      form.validate((valid) => {
        resolve(valid);
      });
    });
    if (validationResult) {
      let flag;
      if (registerType.value) {
        flag = await register();
        if (flag) {
          ElMessage({
            type: "success",
            message: "注册成功，正在跳转后台界面...",
            showClose: true,
          });
        }
      } else {
        flag = await login();
      }
      if (!flag) {
        ElMessage({
          message: registerType.value ? '注册失败，请稍后重试' : '登录失败，请检查用户名和密码',
          type: 'error',
          duration: 3000
        });
        loginVerify();
        registerType.value = false;
      }
    } else {
      ElMessage({
        type: "error",
        message: "请正确填写登录信息",
        showClose: true,
      });
      loginVerify();
    }
  } finally {
    loading.value = false;
  }
};

const sendTGCode = () => {
  const tg_id = registerFormData.tg_id;
  TGRGetCode({ tg_id })
    .then((response) => {
      ElMessage({
        type: "success",
        message: "TG验证码已发送，请查收",
        showClose: true,
      });
    })
    .catch((error) => {
      console.error("Error sending TG code:", error);
      ElMessage({
        type: "error",
        message: "发送TG验证码时出错，请稍后再试",
        showClose: true,
      });
    });
};

const goToResetPage = () => {
  router.push("/resetpwd");
};

const openExternalLink = (url) => {
  window.open(url, '_blank');
};

const focusNextInput = (event, nextRef) => {
  if (event.key === 'Enter') {
    nextRef.value.focus();
  }
};

watch(registerType, (newValue) => {
  currentFormData.value = newValue ? registerFormData : loginFormData;
});

onMounted(async () => {
  fetchAnnouncement();
  loginVerify();
  const rememberedUser = JSON.parse(localStorage.getItem('rememberedUser'));
  if (rememberedUser) {
    loginFormData.username = rememberedUser.username;
    loginFormData.password = rememberedUser.password;
    rememberMe.value = true;
  }

  try {
    const [tgResponse, inviteResponse] = await Promise.all([
      getTGRegisterStatus(),
      getPublicInviteStatus()
    ]);

    if (tgResponse.code === 0) {
      showTGFields.value = tgResponse.data;
    } else {
      console.error('Failed to get TG register status:', tgResponse.msg);
    }

    if (inviteResponse.code === 0) {
      showInviteCodeField.value = inviteResponse.data;
    } else {
      console.error('Failed to get invite code status:', inviteResponse.msg);
    }
  } catch (error) {
    console.error('Error fetching registration settings:', error);
    ElMessage.error('获取注册设置失败');
  }
});

defineExpose({
  submitForm,
  loginVerify,
  goToResetPage
});
</script>

<style scoped>
.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  max-height: 100vh; /* 限制最大高度为视窗高度 */
  overflow-y: auto; /* 启用垂直滚动 */
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', 'Helvetica Neue', Helvetica, Arial, sans-serif;
  background-color: #f0f6f0;
  color: #2c3e50;
}

.top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px 30px;
  background-color: #e8f5e8;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.left-section {
  display: flex;
  align-items: center;
}

.logo {
  height: 45px;
  width: auto;
  margin-right: 25px;
}

.nav-links {
  display: flex;
  gap: 15px;
}

.nav-links .el-button {
  font-size: 14px;
  padding: 10px 20px;
  border-radius: 20px;
  transition: all 0.3s ease;
}

.content-wrapper {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f0f6f0;
  padding: 40px 20px;
}

.login-container {
  width: 100%;
  max-width: 420px;
}

.login-box {
  background-color: #ffffff;
  border-radius: 12px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
  padding: 40px;
}

.login-header {
  text-align: center; /* 让文本居中 */
  margin-bottom: 35px;
  display: flex; /* 使用flex布局 */
  flex-direction: column; /* 垂直排列 */
  align-items: center; /* 水平居中 */
}


.login-logo {
  width: 90px;
  margin-bottom: 20px;
}

.login-title {
  font-size: 28px;
  color: #42b883;
  margin: 0;
  font-weight: 600;
}

.login-form :deep(.el-input__wrapper) {
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.05);
  border-radius: 8px;
  transition: all 0.3s ease;
}

.login-form :deep(.el-input__wrapper):hover,
.login-form :deep(.el-input__wrapper):focus {
  box-shadow: 0 4px 15px rgba(66, 184, 131, 0.1);
}

.password-icon {
  cursor: pointer;
  color: #42b883;
}

.captcha-container,
.tg-code-container {
  display: flex;
  align-items: center;
  gap: 10px;
}

.captcha-input,
.tg-code-input {
  flex: 1;
}

.captcha-image {
  height: 40px;
  cursor: pointer;
  border-radius: 8px;
  transition: opacity 0.3s ease;
}

.captcha-image:hover {
  opacity: 0.8;
}

.send-tg-code-btn {
  white-space: nowrap;
  border-radius: 8px;
}

.submit-btn {
  width: 100%;
  padding: 12px 0;
  font-size: 16px;
  border-radius: 8px;
  transition: all 0.3s ease;
}

.form-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 25px;
}

.toggle-container {
  display: flex;
  background-color: #f0f0f0;
  border-radius: 20px;
  overflow: hidden;
}

.toggle-option {
  padding: 8px 16px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 14px;
}

.toggle-option.active {
  background-color: #42b883;
  color: white;
}

.forgot-password {
  font-size: 14px;
  color: #42b883;
  text-decoration: none;
  transition: color 0.3s ease;
}

.forgot-password:hover {
  color: #33a06f;
}

:deep(.el-button.is-link) {
  color: #42b883;
}

:deep(.el-button--primary) {
  background-color: #42b883;
  border-color: #42b883;
}

:deep(.el-button--primary:hover) {
  background-color: #33a06f;
  border-color: #33a06f;
}

:deep(.el-checkbox__input.is-checked .el-checkbox__inner) {
  background-color: #42b883;
  border-color: #42b883;
}

:deep(.el-checkbox__label) {
  color: #2c3e50;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
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
    padding: 15px;
  }

  .left-section {
    flex-direction: column;
    align-items: flex-start;
    width: 100%;
  }

  .logo {
    margin-bottom: 15px;
  }

  .nav-links {
    flex-wrap: wrap;
    gap: 10px;
    width: 100%;
  }

  .nav-links .el-button {
    flex: 1 0 calc(50% - 5px);
    margin-bottom: 5px;
  }

  .login-container {
    padding: 20px 15px;
  }

  .login-box {
    padding: 30px 20px;
  }

  .announcement-collapse {
    margin: 10px;
  }
}
</style>
