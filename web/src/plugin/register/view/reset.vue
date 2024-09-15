<template>
  <div id="resetPasswordLayout">
    <div class="reset-password-panel">
      <div class="reset-password-panel-form">
        <div class="reset-password-panel-form-title">
          <img
            class="reset-password-panel-form-title-logo"
            src="~@/assets/logo.png"
            alt
          />
          <p class="reset-password-panel-form-title-p">
            {{ $GIN_VUE_ADMIN.appName }}
          </p>
          <p class="reset-password-panel-form-title-p">密码重置</p>
        </div>
        <el-form
          ref="resetPasswordForm"
          :model="resetFormData"
          :rules="resetFormRules"
          @keyup.enter="submitForm"
        >
          <el-form-item prop="tg_id">
            <el-input v-model="resetFormData.tg_id" placeholder="请输入TGID"></el-input>
          </el-form-item>
          <el-form-item prop="code">
            <div class="code-input-container">
              <el-input
                v-model="resetFormData.code"
                placeholder="请输入TG验证码"
                class="code-input"
                style="width: calc(100% - 150px)"
              />
              <div class="send-tg-code-container">
                <el-button type="primary" size="default" @click="sendTGCode"
                  >发送TG验证码</el-button
                >
              </div>
            </div>
          </el-form-item>
          <el-form-item prop="new_password">
          <el-input
            v-model="resetFormData.new_password"
            :type="lock === 'lock' ? 'password' : 'text'"
            placeholder="请输入新密码"
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
        <el-form-item prop="re_password">
          <el-input
            v-model="resetFormData.re_password"
            :type="lock === 'lock' ? 'password' : 'text'"
            placeholder="请再次输入新密码"
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
          <el-form-item>
            <el-button type="primary" size="large" @click="submitForm"
              >确认重置</el-button
            >
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { TGRGetCode, TGRChangePassword } from "@/plugin/register/api/api";
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
</script>

<style scoped>
#resetPasswordLayout {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.reset-password-panel-form {
  max-width: 400px;
  width: 100%;
  padding: 20px;
  background-color: #f5f5f5;
  border-radius: 8px;
  box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
}

.reset-password-panel-form-title {
  text-align: center;
  margin-bottom: 20px;
}

.reset-password-panel-form-title-logo {
  width: 100px;
}

.reset-password-panel-form-title-p {
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

.code-input-container {
  display: flex;
  align-items: center;
}

.code-input {
  width: calc(100% - 150px);
}

.send-tg-code-container {
  margin-left: 10px;
}

.el-button--primary {
  width: 100%;
}
</style>
