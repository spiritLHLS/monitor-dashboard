<template>
  <div id="userLayout" class="w-full h-full relative">
    <div
      class="rounded-lg flex items-center justify-evenly w-full h-full md:w-screen md:h-screen md:bg-[#194bfb] bg-white">
      <div class="md:w-3/5 w-10/12 h-full flex items-center justify-evenly">
        <div class="oblique h-[130%] w-3/5 bg-white dark:bg-slate-900 transform -rotate-12 absolute -ml-52" />
        <!-- 分割斜块 -->
        <div class="z-[999] pt-12 pb-10 md:w-96 w-full rounded-lg flex flex-col justify-between box-border">
          <div>
            <div class="flex items-center justify-center">
              <img class="w-24" :src="$GIN_VUE_ADMIN.appLogo" alt />
            </div>
            <div class="mb-9">
              <p class="text-center text-4xl font-bold">
                {{ $GIN_VUE_ADMIN.appName }}
              </p>
              <p class="text-center text-sm font-normal text-gray-500 mt-2.5">
                A management platform using Golang and Vue
              </p>
            </div>
            <el-form ref="loginForm" :model="loginFormData" :rules="rules" :validate-on-rule-change="false"
              @keyup.enter="submitForm">
              <el-form-item prop="username" class="mb-6">
                <el-input v-model="loginFormData.username" size="large" placeholder="请输入用户名" suffix-icon="user" />
              </el-form-item>
              <el-form-item prop="password" class="mb-6">
                <el-input v-model="loginFormData.password" show-password size="large" type="password"
                  placeholder="请输入密码" />
              </el-form-item>
              <el-form-item v-if="loginFormData.openCaptcha" prop="captcha" class="mb-6">
                <div class="flex w-full justify-between">
                  <el-input v-model="loginFormData.captcha" placeholder="请输入验证码" size="large" class="flex-1 mr-5" />
                  <div class="w-1/3 h-11 bg-[#c3d4f2] rounded">
                    <img v-if="picPath" class="w-full h-full" :src="picPath" alt="请输入验证码" @click="loginVerify()" />
                  </div>
                </div>
              </el-form-item>
              <el-form-item class="mb-6">
                <el-button class="shadow shadow-active h-11 w-full" type="primary" size="large" @click="submitForm">登
                  录</el-button>
              </el-form-item>
              <el-form-item class="mb-6">
                <el-button class="shadow shadow-active h-11 w-full" type="primary" size="large"
                  @click="checkInit">前往初始化</el-button>
              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>
      <div class="hidden md:block w-1/2 h-full float-right bg-[#194bfb]">
        <img class="h-full" src="@/assets/login_right_banner.jpg" alt="banner" />
      </div>
    </div>

    <BottomInfo class="left-0 right-0 absolute bottom-3 mx-auto w-full z-20">
      <div class="links items-center justify-center gap-2 hidden md:flex">
        <a href="https://www.spiritlhl.net" target="_blank">
          <img src="@/assets/docs.png" class="w-8 h-8" alt="文档" />
        </a>
      </div>
    </BottomInfo>
  </div>
</template>

<script setup>
import { captcha } from '@/api/user'
import { checkDB } from '@/api/initdb'
import BottomInfo from '@/components/bottomInfo/bottomInfo.vue'
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'

defineOptions({
  name: 'Login'
})

const router = useRouter()
// 验证函数
const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error('请输入正确的用户名'))
  } else {
    callback()
  }
}
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error('请输入正确的密码'))
  } else {
    callback()
  }
}

// 获取验证码
const loginVerify = async () => {
  const ele = await captcha()
  rules.captcha.push({
    max: ele.data.captchaLength,
    min: ele.data.captchaLength,
    message: `请输入${ele.data.captchaLength}位验证码`,
    trigger: 'blur'
  })
  picPath.value = ele.data.picPath
  loginFormData.captchaId = ele.data.captchaId
  loginFormData.openCaptcha = ele.data.openCaptcha
}
loginVerify()

// 登录相关操作
const loginForm = ref(null)
const picPath = ref('')
const loginFormData = reactive({
  username: 'admin',
  password: '',
  captcha: '',
  captchaId: '',
  openCaptcha: false
})
const rules = reactive({
  username: [{ validator: checkUsername, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
  captcha: [
    {
      message: '验证码格式不正确',
      trigger: 'blur'
    }
  ]
})
const userStore = useUserStore()
const login = async () => {
  console.log('开始登录请求，登录数据:', loginFormData)
  const result = await userStore.LoginIn(loginFormData)
  console.log('登录请求结果:', result)
  return result
}
const submitForm = async () => {  // 改为 async 函数
  console.log('提交表单，开始验证')
  try {
    const valid = await loginForm.value.validate()
    console.log('表单验证结果:', valid)
    if (!valid) {
      console.log('表单验证失败')
      ElMessage({
        type: 'error',
        message: '请正确填写登录信息',
        showClose: true
      })
      await loginVerify()
      return
    }
    console.log('表单验证通过，发起登录请求')
    const flag = await login()
    console.log('登录请求结果:', flag)
    if (!flag) {
      console.log('登录失败，刷新验证码')
      await loginVerify()
      return
    }
    console.log('登录成功')
  } catch (error) {
    console.error('登录过程出错:', error)
    ElMessage({
      type: 'error',
      message: '登录过程出错，请重试',
      showClose: true
    })
    await loginVerify()
  }
}
// 跳转初始化
const checkInit = async () => {
  const res = await checkDB()
  if (res.code === 0) {
    if (res.data?.needInit) {
      userStore.NeedInit()
      await router.push({ name: 'Init' })
    } else {
      ElMessage({
        type: 'info',
        message: '已配置数据库信息，无法初始化'
      })
    }
  }
}
</script>

<style scoped>
#userLayout {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', 'Helvetica Neue', Helvetica, Arial, sans-serif;
  font-size: 12px;
}

:deep(.el-input) {
  font-size: 12px;
}

:deep(.el-input__inner) {
  font-size: 12px;
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

:deep(.el-button) {
  font-size: 12px;
}

:deep(.el-button--primary) {
  background-color: #42b883;
  border-color: #42b883;
}

:deep(.el-button--primary:hover) {
  background-color: #33a06f;
  border-color: #33a06f;
}

:deep(.el-button--large) {
  font-size: 14px;
  padding: 12px 20px;
}

:deep(.el-form-item__label) {
  font-size: 12px;
}

:deep(.el-form-item__error) {
  font-size: 11px;
}

.text-4xl {
  font-size: 2rem;
}

.text-sm {
  font-size: 12px;
}

.text-center {
  text-align: center;
}

@media (max-width: 768px) {
  #userLayout {
    font-size: 11px;
  }

  .text-4xl {
    font-size: 1.5rem;
  }

  .text-sm {
    font-size: 11px;
  }

  :deep(.el-button--large) {
    font-size: 12px;
    padding: 10px 16px;
  }

  :deep(.el-input) {
    font-size: 11px;
  }

  :deep(.el-input__inner) {
    font-size: 11px;
  }
}

@media (max-width: 480px) {
  #userLayout {
    font-size: 10px;
  }

  .text-4xl {
    font-size: 1.25rem;
  }

  .text-sm {
    font-size: 10px;
  }

  :deep(.el-button--large) {
    font-size: 11px;
    padding: 8px 14px;
  }

  :deep(.el-input) {
    font-size: 10px;
  }

  :deep(.el-input__inner) {
    font-size: 10px;
  }
}

.links a {
  transition: all 0.3s ease;
}

.links a:hover {
  transform: scale(1.1);
}

.links img {
  border-radius: 4px;
}
</style>