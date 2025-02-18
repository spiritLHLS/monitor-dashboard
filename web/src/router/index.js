import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/home'
  },
  {
    path: '/init',
    name: 'Init',
    component: () => import('@/view/init/index.vue')
  },
  // {
  //   path: '/login',
  //   name: 'Login',
  //   component: () => import('@/view/login/index.vue')
  // },
  {
    path: '/admin',
    name: 'Admin',
    component: () => import('@/view/login/index.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/plugin/client/view/index.vue') // 有更改 原始为 @/view/login/index.vue
  },
  {
    path: '/client',
    name: 'Register',
    component: () => import('@/plugin/client/view/index.vue')
  },
  {
    path: '/resetpwd',
    name: 'Resetpwd',
    component: () => import('@/plugin/client/view/reset.vue')
  },
  {
    path: '/home',
    name: 'Home',
    component: () => import('@/view/home/index.vue')
  },
  {
    path: '/home',
    name: 'home',
    component: () => import('@/view/home/index.vue')
  },
  {
    path: '/r/:shortCode',
    name: 'Redirect',
    component: () => import('@/view/redirect/redirect.vue')
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('@/view/about/index.vue')
  },
  // 写入可以访问的路由
  {
    path: '/person',
    name: 'person',
    component: () => import('@/view/person/person.vue')
  },
  {
    path: '/dash',
    name: 'dash',
    component: () => import('@/view/dashboard/index.vue')
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: () => import('@/view/controlarea/index.vue')
  },
  // 结束
  {
    path: '/:catchAll(.*)',
    meta: {
      closeTab: true,
    },
    component: () => import('@/view/error/index.vue')
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router
