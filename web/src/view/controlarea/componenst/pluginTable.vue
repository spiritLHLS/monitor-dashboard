<template>
  <div class="control-panel">
    <div class="control-column">
      <div v-for="(control, key) in leftControls" :key="key" class="control-item">
        <label class="switch">
          <input type="checkbox" :checked="control.status" @change="toggleStatus(key)">
          <span class="slider"></span>
        </label>
        <span class="control-label">{{ control.label }}</span>
      </div>
    </div>
    <div class="control-column">
      <div v-for="(control, key) in rightControls" :key="key" class="control-item">
        <label class="switch">
          <input type="checkbox" :checked="control.status" @change="toggleStatus(key)">
          <span class="slider"></span>
        </label>
        <span class="control-label">{{ control.label }}</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import {
  controlSpiders, getSpidersStatus, controlTelegramPush, getTelegramPushStatus,
  controlFAProductsCheck, getFAProductsCheckStatus, getPublicRegisterStatus, controlPublicRegister,
  getPublicPushStatus, controlPublicPushStatus, getTelegramBotPushStatus, controlTelegramBotPushStatus,
  getPublicInviteStatus, controlPublicInvite, getTGRegisterStatus, controlTGRegister,
  getTGLoginStatus, controlTGLogin, 
  // Import the new email push functions
  getEmailPushStatus, controlEmailPushStatus
} from '@/view/controlarea/control.js'

const controls = reactive({
  spider: { status: false, label: '商品爬虫', getStatus: getSpidersStatus, controlFunc: controlSpiders, controlKey: 'enable_spiders' },
  telegram: { status: false, label: '频道推送', getStatus: getTelegramPushStatus, controlFunc: controlTelegramPush, controlKey: 'enable_tgpush' },
  products: { status: false, label: '追新爬虫', getStatus: getFAProductsCheckStatus, controlFunc: controlFAProductsCheck, controlKey: 'enable_allpdspiders' },
  register: { status: false, label: '公开注册', getStatus: getPublicRegisterStatus, controlFunc: controlPublicRegister, controlKey: 'enable_public_register' },
  publicPush: { status: false, label: '一对一推送', getStatus: getPublicPushStatus, controlFunc: controlPublicPushStatus, controlKey: 'enable_public_push' },
  tgBot: { status: false, label: 'TGBot推送', getStatus: getTelegramBotPushStatus, controlFunc: controlTelegramBotPushStatus, controlKey: 'enable_tg_bot_push' },
  inviteCode: { status: false, label: '邀请码注册', getStatus: getPublicInviteStatus, controlFunc: controlPublicInvite, controlKey: 'enable_public_invite' },
  tgRegister: { status: false, label: '校验TG注册', getStatus: getTGRegisterStatus, controlFunc: controlTGRegister, controlKey: 'enable_tg_register' },
  tgLogin: { status: false, label: '校验TG登录', getStatus: getTGLoginStatus, controlFunc: controlTGLogin, controlKey: 'enable_tg_login' },
  // Add the new email push control
  emailPush: { status: false, label: '邮件推送', getStatus: getEmailPushStatus, controlFunc: controlEmailPushStatus, controlKey: 'enable_email_push' }
})

const leftControls = computed(() => {
  const keys = Object.keys(controls)
  return Object.fromEntries(keys.slice(0, Math.ceil(keys.length / 2)).map(key => [key, controls[key]]))
})

const rightControls = computed(() => {
  const keys = Object.keys(controls)
  return Object.fromEntries(keys.slice(Math.ceil(keys.length / 2)).map(key => [key, controls[key]]))
})

const initializeStatuses = async () => {
  const statusPromises = Object.entries(controls).map(async ([key, control]) => {
    try {
      const statusData = await control.getStatus();
      return [key, statusData.data];
    } catch (error) {
      console.error(`Error initializing ${key} status:`, error);
      return [key, false]; // 默认为 false 如果出错
    }
  });

  const results = await Promise.all(statusPromises);
  
  results.forEach(([key, status]) => {
    controls[key].status = status;
  });
}

const toggleStatus = async (key) => {
  try {
    const control = controls[key];
    const newStatus = !control.status;
    const controlObj = { [control.controlKey]: newStatus };
    const statusData = await control.controlFunc(controlObj);
    control.status = statusData.data;
  } catch (error) {
    console.error(`Error toggling ${key} status:`, error);
  }
}

// 初始获取状态
initializeStatuses()
</script>

<style scoped lang="scss">
.control-panel {
  display: flex;
  justify-content: space-between;
  gap: 30px;
  padding: 30px;
  background: #f5f7fa;
  border-radius: 15px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.control-column {
  display: flex;
  flex-direction: column;
  gap: 20px;
  flex: 1;
}

.control-item {
  display: flex;
  align-items: center;
  gap: 15px;
}

.control-label {
  font-size: 16px;
  font-weight: 600;
  color: #333;
}

/* 开关样式 */
.switch {
  position: relative;
  display: inline-block;
  width: 60px;
  height: 34px;
}

.switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  transition: .4s;
  border-radius: 34px;
}

.slider:before {
  position: absolute;
  content: "";
  height: 26px;
  width: 26px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  transition: .4s;
  border-radius: 50%;
}

input:checked + .slider {
  background-color: #4CAF50;
}

input:checked + .slider:before {
  transform: translateX(26px);
}
</style>