<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-header">
        <img class="login-logo" src="~@/assets/logo.png" alt="Logo" />
        <h1 class="login-title">{{ $GIN_VUE_ADMIN.appName }}</h1>
      </div>
      <el-form
        ref="loginForm"
        :model="currentFormData"
        :rules="formRules"
        class="login-form"
        @keyup.enter="submitForm"
      >
        <el-form-item prop="username">
          <el-input 
            v-model="currentFormData.username" 
            placeholder="请输入用户名"
            prefix-icon="User"
          />
        </el-form-item>

        <el-form-item prop="password">
          <el-input
            v-model="currentFormData.password"
            :type="lock === 'lock' ? 'password' : 'text'"
            placeholder="请输入密码"
            prefix-icon="Lock"
          >
            <template #suffix>
              <el-icon @click="changeLock" class="password-icon">
                <component :is="lock" />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>

        <el-form-item prop="captcha">
          <div class="captcha-container">
            <el-input
              v-model="currentFormData.captcha"
              placeholder="请输入验证码"
              class="captcha-input"
            />
            <img
              v-if="picPath"
              :src="picPath"
              alt="验证码"
              @click="loginVerify"
              class="captcha-image"
            />
          </div>
        </el-form-item>

        <template v-if="registerType">
          <el-form-item prop="tg_id">
            <el-input 
              v-model="currentFormData.tg_id" 
              placeholder="请输入TGID"
              prefix-icon="ChatDotSquare"
            />
          </el-form-item>

          <el-form-item prop="code">
            <div class="tg-code-container">
              <el-input
                v-model="currentFormData.code"
                placeholder="请输入TG验证码"
                class="tg-code-input"
              />
              <el-button type="primary" @click="sendTGCode" class="send-tg-code-btn">
                发送验证码
              </el-button>
            </div>
          </el-form-item>
        </template>

        <el-form-item>
          <el-button type="primary" @click="submitForm" class="submit-btn">
            {{ registerType ? '注册' : '登录' }}
          </el-button>
        </el-form-item>

        <div class="form-footer">
          <el-switch
            v-model="registerType"
            active-text="注册"
            inactive-text="登录"
            class="register-switch"
          />
          <el-button type="text" @click="goToResetPage" class="forgot-password">
            忘记密码？
          </el-button>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import { captcha } from "@/api/user";
import { TGRGetCode } from "@/plugin/client/api/api";
import { reactive, ref, watch } from "vue";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";
import { useUserStore } from "@/pinia/modules/user";
const router = useRouter();
// 对用户的输入强制要求符合要求
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
const loginForm = ref(null);
const loginFormData = reactive({
  username: "admin",
  password: "123456",
  captcha: "",
  captchaId: "",
});
const registerFormData = reactive({
  username: "admin",
  password: "123456",
  captcha: "",
  captchaId: "",
  tg_id: "",
  code: "",
});
const lock = ref("lock");
const picPath = ref("");
const registerType = ref(false);
const currentFormData = ref(loginFormData);
const rules = reactive({
  username: [{ validator: checkUsername, trigger: "blur" }],
  password: [{ validator: checkPassword, trigger: "blur" }],
  captcha: [
    { required: true, message: "请输入验证码", trigger: "blur" },
    { message: "验证码格式不正确", trigger: "blur" },
  ],
  tg_id: [{ required: true, message: "请输入TG ID", trigger: "blur", visible: false }],
  code: [{ required: true, message: "请输入TG验证码", trigger: "blur", visible: false }],
});

const formRules = ref(rules);
const userStore = useUserStore();
const loginVerify = () => {
  captcha({}).then((ele) => {
    rules.captcha[1].max = ele.data.captchaLength;
    rules.captcha[1].min = ele.data.captchaLength;
    picPath.value = ele.data.picPath;
    currentFormData.value.captchaId = ele.data.captchaId;
    registerFormData.captchaId = ele.data.captchaId;
  });
};
loginVerify();
const changeLock = () => {
  lock.value = lock.value === "lock" ? "unlock" : "lock";
};

const login = async () => {
  try {
    await userStore.UserTgLogin(loginFormData);
    return true;
  } catch (error) {
    return false;
  }
};
const register = async () => {
  try {
    await userStore.UserTgRegister(registerFormData);
    return true;
  } catch (error) {
    return false;
  }
};
const submitForm = async () => {
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
};
// TG验证码发送
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
// 跳转密码重置界面
const goToResetPage = () => {
  router.push("/resetpwd");
};
// 监听注册类型的变化，切换表单数据
watch(registerType, (newValue, oldValue) => {
  if (newValue === true && oldValue === false) {
    currentFormData.value = registerFormData;
  }
});
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.login-box {
  width: 400px;
  padding: 40px;
  background-color: rgba(255, 255, 255, 0.9);
  border-radius: 10px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
}

.login-header {
  text-align: center;
  margin-bottom: 30px;
}

.login-logo {
  width: 80px;
  margin-bottom: 16px;
}

.login-title {
  font-size: 24px;
  color: #333;
  margin: 0;
}

.login-form :deep(.el-input__wrapper) {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  border-radius: 4px;
}

.login-form :deep(.el-input__wrapper):hover,
.login-form :deep(.el-input__wrapper):focus {
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

.password-icon {
  cursor: pointer;
}

.captcha-container,
.tg-code-container {
  display: flex;
  align-items: center;
}

.captcha-input,
.tg-code-input {
  flex: 1;
  margin-right: 10px;
}

.captcha-image {
  height: 40px;
  cursor: pointer;
  border-radius: 4px;
}

.send-tg-code-btn {
  white-space: nowrap;
}

.submit-btn {
  width: 100%;
  padding: 12px 0;
  font-size: 16px;
}

.form-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 20px;
}

.register-switch {
  font-size: 14px;
}

.forgot-password {
  font-size: 14px;
}

:deep(.el-switch__label) {
  color: #606266;
}

:deep(.el-button--text) {
  color: #409EFF;
}
</style>