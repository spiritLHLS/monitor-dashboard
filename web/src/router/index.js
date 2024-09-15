import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [{
  path: '/',
  redirect: '/login'
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
  component: () => import('@/plugin/register/view/index.vue') // 有更改 原始为 @/view/login/index.vue
},
{
  path: '/register',
  name: 'Register',
  component: () => import('@/plugin/register/view/index.vue')
},
{
  path: '/resetpwd',
  name: 'Resetpwd',
  component: () => import('@/plugin/register/view/reset.vue')
},
{
  path: '/:catchAll(.*)',
  meta: {
    closeTab: true,
  },
  component: () => import('@/view/error/index.vue')
}
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router
