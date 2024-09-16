<template>
  <div class="reset-password-container">
    <div class="reset-password-box">
      <div class="reset-password-header">
        <img class="reset-logo" src="~@/assets/logo.png" alt="Logo" />
        <h1 class="reset-title">{{ $GIN_VUE_ADMIN.appName }}</h1>
        <h2 class="reset-subtitle">密码重置</h2>
      </div>
      <el-form
        ref="resetPasswordForm"
        :model="resetFormData"
        :rules="resetFormRules"
        class="reset-form"
        @keyup.enter="submitForm"
      >
        <el-form-item prop="tg_id">
          <el-input 
            v-model="resetFormData.tg_id" 
            placeholder="请输入TGID"
            prefix-icon="ChatDotSquare"
          />
        </el-form-item>

        <el-form-item prop="code">
          <div class="tg-code-container">
            <el-input
              v-model="resetFormData.code"
              placeholder="请输入TG验证码"
              class="tg-code-input"
            />
            <el-button type="primary" @click="sendTGCode" class="send-tg-code-btn">
              发送验证码
            </el-button>
          </div>
        </el-form-item>

        <el-form-item prop="new_password">
          <el-input
            v-model="resetFormData.new_password"
            :type="lock === 'lock' ? 'password' : 'text'"
            placeholder="请输入新密码"
            prefix-icon="Lock"
          >
            <template #suffix>
              <el-icon @click="changeLock" class="password-icon">
                <component :is="lock" />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>

        <el-form-item prop="re_password">
          <el-input
            v-model="resetFormData.re_password"
            :type="lock === 'lock' ? 'password' : 'text'"
            placeholder="请再次输入新密码"
            prefix-icon="Lock"
          >
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
        <el-button type="text" @click="goToLoginPage" class="back-to-login">
          返回登录
        </el-button>
      </div>
      </el-form>
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
</script>

<style scoped>
.reset-password-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.reset-password-box {
  width: 400px;
  padding: 40px;
  background-color: rgba(255, 255, 255, 0.9);
  border-radius: 10px;
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);
}

.reset-password-header {
  text-align: center;
  margin-bottom: 30px;
}

.reset-logo {
  width: 80px;
  margin-bottom: 16px;
}

.reset-title {
  font-size: 24px;
  color: #333;
  margin: 0 0 10px;
}

.reset-subtitle {
  font-size: 18px;
  color: #666;
  margin: 0;
}

.reset-form :deep(.el-input__wrapper) {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  border-radius: 4px;
}

.reset-form :deep(.el-input__wrapper):hover,
.reset-form :deep(.el-input__wrapper):focus {
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

.password-icon {
  cursor: pointer;
}

.tg-code-container {
  display: flex;
  align-items: center;
}

.tg-code-input {
  flex: 1;
  margin-right: 10px;
}

.send-tg-code-btn {
  white-space: nowrap;
}

.submit-btn {
  width: 100%;
  padding: 12px 0;
  font-size: 16px;
}
</style>