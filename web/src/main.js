import './style/element_visiable.scss'
import 'element-plus/theme-chalk/dark/css-vars.css'
import { createApp } from 'vue'
import ElementPlus from 'element-plus'

import 'element-plus/dist/index.css'
// 引入gva前端初始化相关内容
import './core/gva'
// 引入封装的router
import router from '@/router/index'
import '@/permission'
import run from '@/core/gva.js'
import auth from '@/directive/auth'
import { store } from '@/pinia'
import App from './App.vue'
// 消除警告
import 'default-passive-events'

const app = createApp(App)
app.config.productionTip = false

app.use(run).use(ElementPlus).use(store).use(auth).use(router).mount('#app')
export default app
