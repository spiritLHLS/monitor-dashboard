<template>
  <div class="app-container">
    <header class="top-bar">
      <div class="left-section">
        <img src="~@/assets/logo.png" alt="Logo" class="logo">
        <nav class="nav-links">
          <el-button type="primary" @click="openExternalLink('https://paste.spiritlhl.net')">商家评价</el-button>
          <el-button type="primary" @click="openExternalLink('https://t.me/vps_spiders')">监控频道</el-button>
          <el-button type="primary" @click="openExternalLink('https://www.spiritlhl.net')">一键虚拟化项目</el-button>
          <el-button type="primary" @click="router.push('/home')">回到监控页面</el-button>
        </nav>
      </div>
    </header>

    <div class="content-wrapper">
      <div class="reset-password-container">
        <div class="reset-password-box">
          <div class="reset-password-header">
            <img class="reset-logo" src="~@/assets/logo.png" alt="Logo" />
            <h1 class="reset-title">ECS-SPIDERS</h1>
            <h2 class="reset-subtitle">密码重置</h2>
          </div>
          <el-form ref="resetPasswordForm" :model="resetFormData" :rules="resetFormRules" class="reset-form"
            @keyup.enter="submitForm">
            <el-form-item prop="tg_id">
              <el-input v-model="resetFormData.tg_id" placeholder="请输入TGID" prefix-icon="ChatDotSquare" />
            </el-form-item>

            <el-form-item prop="code">
              <div class="tg-code-container">
                <el-input v-model="resetFormData.code" placeholder="请输入TG验证码" class="tg-code-input" />
                <el-button type="primary" @click="sendTGCode" class="send-tg-code-btn">
                  发送验证码
                </el-button>
              </div>
            </el-form-item>

            <el-form-item prop="new_password">
              <el-input v-model="resetFormData.new_password" :type="lock === 'lock' ? 'password' : 'text'"
                placeholder="请输入新密码" prefix-icon="Lock">
                <template #suffix>
                  <el-icon @click="changeLock" class="password-icon">
                    <component :is="lock" />
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>

            <el-form-item prop="re_password">
              <el-input v-model="resetFormData.re_password" :type="lock === 'lock' ? 'password' : 'text'"
                placeholder="请再次输入新密码" prefix-icon="Lock">
                <template #suffix>
                  <el-icon @click="changeLock" class="password-icon">
                    <component :is="lock" />
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>

            <el-form-item>
              <el-button type="primary" @click="submitForm" class="submit-btn">
                确认重置
              </el-button>
            </el-form-item>
            <div class="form-footer">
              <el-button link @click="goToLoginPage" class="back-to-login">
                返回登录
              </el-button>
            </div>
          </el-form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { TGRGetCode, TGRChangePassword } from "@/plugin/client/api/api";
import { reactive, ref } from "vue";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";

const router = useRouter();
const resetFormData = reactive({
  tg_id: "",
  new_password: "",
  re_password: "",
  code: "",
});
const resetPasswordForm = ref(null);
const lock = ref("lock");

const confirmPasswordValidator = (rule, value, callback) => {
  if (value !== resetFormData.new_password) {
    callback(new Error("两次输入的密码不一致"));
  } else {
    callback();
  }
};

const resetFormRules = reactive({
  tg_id: [{ required: true, message: "请输入TG ID", trigger: "blur" }],
  new_password: [{ required: true, message: "请输入新密码", trigger: "blur" }],
  re_password: [
    { required: true, message: "请再次输入新密码", trigger: "blur" },
    { validator: confirmPasswordValidator, trigger: "blur" }
  ],
  code: [{ required: true, message: "请输入TG验证码", trigger: "blur" }],
});

const changeLock = () => {
  lock.value = lock.value === "lock" ? "unlock" : "lock";
};

const submitForm = () => {
  if (resetPasswordForm.value) {
    resetPasswordForm.value.validate((valid) => {
      if (valid) {
        const { re_password, ...requestData } = resetFormData;
        TGRChangePassword(requestData)
          .then(() => {
            ElMessage({
              type: "success",
              message: "密码重置成功，正在跳转登录界面...",
              showClose: true,
            });
            router.push("/login");
          })
          .catch((error) => {
            console.error("Error resetting password:", error);
            ElMessage({
              type: "error",
              message: "密码重置失败，请稍后再试",
              showClose: true,
            });
          });
      } else {
        ElMessage({
          type: "error",
          message: "填写有误，请检查输入",
          showClose: true,
        });
      }
    });
  }
};

const sendTGCode = () => {
  const tg_id = resetFormData.tg_id;
  TGRGetCode({ tg_id })
    .then(() => {
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

const goToLoginPage = () => {
  router.push("/login");
};

const openExternalLink = (url) => {
  window.open(url, '_blank');
};
</script>

<style scoped>
.app-container {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', 'Helvetica Neue', Helvetica, Arial, sans-serif;
  background-color: #f0f6f0;
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
  padding: 40px 20px;
}

.reset-password-container {
  width: 100%;
  max-width: 420px;
}

.reset-password-box {
  background-color: #ffffff;
  border-radius: 12px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
  padding: 40px;
}

.reset-password-header {
  text-align: center;
  margin-bottom: 30px;
}

.reset-logo {
  width: 90px;
  margin-bottom: 20px;
}

.reset-title {
  font-size: 28px;
  color: #42b883;
  margin: 0 0 10px;
  font-weight: 600;
}

.reset-subtitle {
  font-size: 18px;
  color: #666;
  margin: 0;
}

.reset-form :deep(.el-input__wrapper) {
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.05);
  border-radius: 8px;
  transition: all 0.3s ease;
}

.reset-form :deep(.el-input__wrapper):hover,
.reset-form :deep(.el-input__wrapper):focus {
  box-shadow: 0 4px 15px rgba(66, 184, 131, 0.1);
}

.password-icon {
  cursor: pointer;
  color: #42b883;
}

.tg-code-container {
  display: flex;
  align-items: center;
  gap: 10px;
}

.tg-code-input {
  flex: 1;
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
  margin-top: 20px;
  text-align: center;
}

.back-to-login {
  font-size: 14px;
  color: #42b883;
  text-decoration: none;
  transition: color 0.3s ease;
}

.back-to-login:hover {
  color: #33a06f;
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

  .reset-password-container {
    padding: 20px 15px;
  }

  .reset-password-box {
    padding: 30px 20px;
  }
}
</style>
