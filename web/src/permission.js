import { useUserStore } from '@/pinia/modules/user'
import { useRouterStore } from '@/pinia/modules/router'
import getPageTitle from '@/utils/page'
import router from '@/router'
import Nprogress from 'nprogress'
import 'nprogress/nprogress.css'
Nprogress.configure({ showSpinner: false, ease: 'ease', speed: 500 })

// 这里仅列入无需登录就能访问的页面路由 对应 src\router\index.js 的内容
const whiteList = ['Login', 'Init', 'Register', 'Admin', 'Resetpwd', 'Home', 'About', 'Redirect']

const getRouter = async (userStore) => {
  const routerStore = useRouterStore()
  await routerStore.SetAsyncRouter()
  await userStore.GetUserInfo()
  const asyncRouters = routerStore.asyncRouters
  asyncRouters.forEach(asyncRouter => {
    router.addRoute(asyncRouter)
  })
}

const removeLoading = () => {
  const element = document.getElementById('gva-loading-box');
  if (element) {
    element.remove();
  }
}

async function handleKeepAlive(to) {
  if (to.matched.some(item => item.meta.keepAlive)) {
    if (to.matched && to.matched.length > 2) {
      for (let i = 1; i < to.matched.length; i++) {
        const element = to.matched[i - 1]
        if (element.name === 'layout') {
          to.matched.splice(i, 1)
          await handleKeepAlive(to)
        }
        // 如果没有按需加载完成则等待加载
        if (typeof element.components.default === 'function') {
          await element.components.default()
          await handleKeepAlive(to)
        }
      }
    }
  }
}

router.beforeEach(async (to, from) => {
  const routerStore = useRouterStore()
  Nprogress.start()
  const userStore = useUserStore()
  to.meta.matched = [...to.matched]
  handleKeepAlive(to)
  const token = userStore.token
  // 在白名单中的判断情况
  document.title = getPageTitle(to.meta.title, to)
  if (to.meta.client) {
    return true
  }
  // 添加对重定向路由的特殊处理
  if (to.name === 'Redirect') {
    return true
  }
  if (whiteList.indexOf(to.name) > -1) {
    if (token) {
      if (!routerStore.asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
        await getRouter(userStore)
      }
      // token 可以解析但是却是不存在的用户 id 或角色 id 会导致无限调用
      if (userStore.userInfo?.authority?.defaultRouter != null) {
        if (router.hasRoute(userStore.userInfo.authority.defaultRouter)) {
          // 添加判断，防止重复跳转
          if (to.name === userStore.userInfo.authority.defaultRouter) {
            return true
          }
          return { name: userStore.userInfo.authority.defaultRouter }
        } else {
          return { path: '/layout/404' }
        }
      } else {
        // 强制退出账号
        userStore.ClearStorage()
        return {
          name: 'Login',
          query: {
            redirect: document.location.hash
          }
        }
      }
    } else {
      return true
    }
  } else {
    // 不在白名单中并且已经登录的时候
    if (token) {
      if (sessionStorage.getItem("needToHome") === 'true') {
        sessionStorage.removeItem("needToHome")
        return { path: '/' }
      }
      // 添加flag防止多次获取动态路由和栈溢出
      if (!routerStore.asyncRouterFlag && whiteList.indexOf(from.name) < 0) {
        await getRouter(userStore)
        if (userStore.token) {
          if (router.hasRoute(userStore.userInfo.authority.defaultRouter)) {
            return { ...to, replace: true }
          } else {
            return { path: '/layout/404' } // /layout/404 --> /layout/dash
          }
        } else {
          return {
            name: 'Login',
            query: { redirect: to.href }
          }
        }
      } else {
        if (to.matched.length) {
          return true
        } else {
          return { path: '/layout/404' } // /layout/404 --> /layout/person
        }
      }
    }
    // 不在白名单中并且未登录的时候
    if (!token) {
      return {
        name: 'Login',
        query: {
          redirect: document.location.hash
        }
      }
    }
  }
})

router.afterEach(() => {
  // 路由加载完成后关闭进度条
  document.getElementsByClassName('main-cont main-right')[0]?.scrollTo(0, 0)
  Nprogress.done()
})

router.onError(() => {
  // 路由发生错误后销毁进度条
  Nprogress.remove()
})

removeLoading()