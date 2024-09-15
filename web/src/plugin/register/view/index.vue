<template>
  <div id="userLayout">
    <div class="login-panel">
      <div class="login-panel-form">
        <div class="login-panel-form-title">
          <img class="login-panel-form-title-logo" src="~@/assets/logo.png" alt />
          <p class="login-panel-form-title-p">{{ $GIN_VUE_ADMIN.appName }}</p>
        </div>
        <el-form
          ref="loginForm"
          :model="currentFormData"
          :rules="formRules"
          @keyup.enter="submitForm"
        >
          <el-form-item prop="username">
            <el-input v-model="currentFormData.username" placeholder="请输入用户名">
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <user />
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              v-model="currentFormData.password"
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
            >
              <template #suffix>
                <span class="input-icon">
                  <el-icon>
                    <component :is="lock" @click="changeLock" />
                  </el-icon>
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="captcha">
            <div class="captcha-input-container">
              <el-input
                v-model="currentFormData.captcha"
                placeholder="请输入验证码"
                class="captcha-input"
                style="width: calc(100% - 150px)"
              />
              <img
                v-if="picPath"
                :src="picPath"
                alt="请输入验证码"
                @click="loginVerify()"
                class="captcha-image"
              />
            </div>
          </el-form-item>
          <el-form-item v-if="registerType" prop="tg_id">
            <el-input v-model="currentFormData.tg_id" placeholder="请输入TGID"></el-input>
          </el-form-item>
          <el-form-item v-if="registerType" prop="code" class="form-item-inline">
            <el-input
              v-model="currentFormData.code"
              placeholder="请输入TG验证码"
            ></el-input>
            <div class="send-tg-code-container">
              <el-button type="primary" size="default" @click="sendTGCode"
                >发送TG验证码</el-button
              >
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" size="large" @click="submitForm">
              <div v-if="registerType">注册</div>
              <div v-else>登录</div>
            </el-button>
            <el-button type="primary" size="large" @click="goToResetPage"
              >找回密码</el-button
            >
          </el-form-item>
          <el-form-item>
            <el-switch v-model="registerType" />
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { captcha } from "@/api/user";
import { TGRGetCode } from "@/plugin/register/api/api";
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
#userLayout {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.login-panel-form {
  max-width: 400px;
  width: 100%;
  padding: 20px;
  background-color: #f5f5f5;
  border-radius: 8px;
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
}

.login-panel-form-title {
  text-align: center;
  margin-bottom: 20px;
}

.login-panel-form-title-logo {
  width: 100px;
}

.login-panel-form-title-p {
  font-size: 24px;
  font-weight: bold;
  margin: 10px 0;
}

.el-input {
  width: 100%;
}

.input-icon {
  position: absolute;
  right: 10px;
  top: 50%;
  transform: translateY(-50%);
}

.captcha-input-container {
  position: relative;
  display: flex;
  align-items: center;
}

.captcha-input {
  width: calc(100% - 150px);
}

.captcha-image {
  margin-left: 10px;
  cursor: pointer;
}

.form-item-inline .el-input {
  display: inline-block;
  width: calc(100% - 120px);
}

.send-tg-code-container {
  display: inline-block;
  width: 120px;
  margin-top: 0;
  text-align: right;
}

.send-tg-code-button {
  max-width: 100%;
  white-space: nowrap;
}
</style>
