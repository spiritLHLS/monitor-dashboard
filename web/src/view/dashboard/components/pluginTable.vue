<template>
  <div class="control-panel">
    <button class="control-button spider" @click="toggleSpiderStatus">
      <icon-spider :class="{ active: spiderStatus }" />
      <span>{{ spiderStatus ? '关闭爬虫' : '启用爬虫' }}</span>
    </button>
    
    <button class="control-button telegram" @click="toggleTelegramStatus">
      <icon-telegram :class="{ active: telegramStatus }" />
      <span>{{ telegramStatus ? '关闭推送' : '启用推送' }}</span>
    </button>
    
    <button class="control-button products" @click="toggleFAProductsCheckStatus">
      <icon-products :class="{ active: faProductsStatus }" />
      <span>{{ faProductsStatus ? '关闭商品爬虫' : '启用商品爬虫' }}</span>
    </button>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import {
  controlSpiders, getSpidersStatus, controlTelegramPush, getTelegramPushStatus,
  controlFAProductsCheck, getFAProductsCheckStatus
} from '@/view/controlarea/control.js'

const initializeStatus = async (statusRef, getStatusFunc) => {
  try {
    const statusData = await getStatusFunc();
    statusRef.value = statusData.data;
  } catch (error) {
    console.error('Error initializing status:', error);
  }
};

const toggleStatus = async (statusRef, controlStatusFunc, controlKey) => {
  try {
    const newStatus = !statusRef.value;
    const controlObj = { [controlKey]: newStatus };
    const statusData = await controlStatusFunc(controlObj);
    statusRef.value = statusData.data;
  } catch (error) {
    console.error('Error toggling status:', error)
  }
}

// Spider
const spiderStatus = ref(false)
const fetchSpiderStatus = () => initializeStatus(spiderStatus, getSpidersStatus)
const toggleSpiderStatus = () => toggleStatus(spiderStatus, controlSpiders, 'enable_spiders')

// Telegram
const telegramStatus = ref(false)
const fetchTelegramStatus = () => initializeStatus(telegramStatus, getTelegramPushStatus)
const toggleTelegramStatus = () => toggleStatus(telegramStatus, controlTelegramPush, 'enable_tgpush')

// FA Products Check
const faProductsStatus = ref(false)
const fetchFAProductsCheckStatus = () => initializeStatus(faProductsStatus, getFAProductsCheckStatus)
const toggleFAProductsCheckStatus = () => toggleStatus(faProductsStatus, controlFAProductsCheck, 'enable_allpdspiders')

// 初始获取状态
fetchSpiderStatus()
fetchTelegramStatus()
fetchFAProductsCheckStatus()
</script>

<style scoped lang="scss">
.control-panel {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
  padding: 30px;
  background: #f5f7fa;
  border-radius: 15px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.control-button {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 250px;
  padding: 15px 20px;
  border: none;
  border-radius: 50px;
  background: #ffffff;
  color: #333;
  font-size: 16px;
  font-weight: 600;
  transition: all 0.3s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  
  &:hover {
    transform: translateY(-2px);
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.15);
  }
  
  &:active {
    transform: translateY(1px);
  }

  span {
    margin-left: 10px;
  }
}

.spider { &:hover, &.active { background: #4CAF50; color: white; } }
.telegram { &:hover, &.active { background: #2196F3; color: white; } }
.products { &:hover, &.active { background: #FF9800; color: white; } }

// 这里假设您使用了一些图标组件，如果没有，可以使用字体图标或SVG图标替代
.icon-spider, .icon-telegram, .icon-products {
  font-size: 24px;
  transition: all 0.3s ease;
  
  &.active {
    transform: scale(1.2);
  }
}
</style>