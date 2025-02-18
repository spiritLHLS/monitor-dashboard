import { login, getUserInfo, setSelfInfo } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router/index'
import { ElLoading, ElMessage } from 'element-plus'
import { defineStore } from 'pinia'
import { ref, computed, watch } from 'vue'
import { useRouterStore } from './router'
import cookie from 'js-cookie'
import { TGRRegister, RegisterWintInvite, TGRLogin } from '@/plugin/client/api/api'
import { useAppStore } from "@/pinia";

export const useUserStore = defineStore('user', () => {
  const appStore = useAppStore()
  const loadingInstance = ref(null)

  const userInfo = ref({
    uuid: '',
    nickName: '',
    headerImg: '',
    authority: {},
    authorityId: 0,
  })
  const token = ref(window.localStorage.getItem('token') || cookie.get('x-token') || '')
  const setUserInfo = (val) => {
    userInfo.value = val
    if (val.originSetting) {
      Object.keys(appStore.config).forEach((key) => {
        if (val.originSetting[key]) {
          appStore.config[key] = val.originSetting[key]
        }
      })
    }
  }

  const setToken = (val) => {
    token.value = val
  }

  const NeedInit = async () => {
    token.value = ''
    window.localStorage.removeItem('token')
    await router.push({ name: 'Init', replace: true })
  }

  const ResetUserInfo = (value = {}) => {
    userInfo.value = {
      ...userInfo.value,
      ...value
    }
  }
  /* 获取用户信息*/
  const GetUserInfo = async () => {
    const res = await getUserInfo()
    if (res.code === 0) {
      setUserInfo(res.data.userInfo)
    }
    return res
  }
  /* 登录*/
  const LoginIn = async (loginInfo) => {
    loadingInstance.value = ElLoading.service({
      fullscreen: true,
      text: '登录中，请稍候...',
    })

    const res = await login(loginInfo)

    // 登陆失败，直接返回
    if (res.code !== 0) {
      loadingInstance.value.close()
      return false
    }

    // 登陆成功，设置用户信息和权限相关信息
    setUserInfo(res.data.user)
    setToken(res.data.token)

    // 初始化路由信息
    const routerStore = useRouterStore()
    await routerStore.SetAsyncRouter()
    const asyncRouters = routerStore.asyncRouters

    // 注册到路由表里
    asyncRouters.forEach(asyncRouter => {
      router.addRoute(asyncRouter)
    })
    if (userInfo.value.authorityId === 888 || userInfo.value.authorityId === 8881) {
      await router.replace({ name: 'dashboard' })
    } else {
      if (!router.hasRoute(userInfo.value.authority.defaultRouter)) {
        ElMessage.error('请联系管理员进行授权')
      }
      await router.replace({ name: userInfo.value.authority.defaultRouter })
    }

    const isWin = ref(/windows/i.test(navigator.userAgent))
    if (isWin.value) {
      window.localStorage.setItem('osType', 'WIN')
    } else {
      window.localStorage.setItem('osType', 'MAC')
    }


    // 全部操作均结束，关闭loading并返回
    loadingInstance.value.close()
    return true
  }
  /* 登出*/
  const LoginOut = async () => {
    const res = await jsonInBlacklist()

    // 登出失败
    if (res.code !== 0) {
      return
    }

    await ClearStorage()

    // 把路由定向到登录页，无需等待直接reload
    router.push({ name: 'Login', replace: true })
    window.location.reload()
  }
  /* 清理数据 */
  const ClearStorage = async () => {
    token.value = ''
    sessionStorage.clear()
    window.localStorage.removeItem('token')
    cookie.remove('x-token')
    localStorage.removeItem('originSetting')
  }

  watch(() => token.value, () => {
    window.localStorage.setItem('token', token.value)
  })

  const UserTgRegister = async (loginInfo) => {
    loadingInstance.value = ElLoading.service({
      fullscreen: true,
      text: "注册中，请稍候...",
    });
    try {
      const res = await TGRRegister(loginInfo);
      if (res.code === 0) {
        setUserInfo(res.data.user);
        setToken(res.data.token);
        const routerStore = useRouterStore();
        await routerStore.SetAsyncRouter();
        const asyncRouters = routerStore.asyncRouters;
        asyncRouters.forEach((asyncRouter) => {
          router.addRoute(asyncRouter);
        });
        router.push({ name: userInfo.value.authority.defaultRouter });
        return true;
      }
    } catch (e) {
      loadingInstance.value.close();
    }
    loadingInstance.value.close();
  };

  const UserRegisterWithInvite = async (loginInfo) => {
    loadingInstance.value = ElLoading.service({
      fullscreen: true,
      text: "注册中，请稍候...",
    });
    try {
      const res = await RegisterWintInvite(loginInfo);
      if (res.code === 0) {
        setUserInfo(res.data.user);
        setToken(res.data.token);
        const routerStore = useRouterStore();
        await routerStore.SetAsyncRouter();
        const asyncRouters = routerStore.asyncRouters;
        asyncRouters.forEach((asyncRouter) => {
          router.addRoute(asyncRouter);
        });
        router.push({ name: userInfo.value.authority.defaultRouter });
        return true;
      }
    } catch (e) {
      loadingInstance.value.close();
    }
    loadingInstance.value.close();
  };

  const UserTgLogin = async (loginInfo) => {
    loadingInstance.value = ElLoading.service({
      fullscreen: true,
      text: '登录中，请稍候...',
    });
    try {
      // console.log("try login");
      const res = await TGRLogin(loginInfo);
      // console.log("Login response:", res);
      if (res.code === 0) {
        setUserInfo(res.data.user);
        setToken(res.data.token);
        const routerStore = useRouterStore();
        await routerStore.SetAsyncRouter();
        const asyncRouters = routerStore.asyncRouters;
        asyncRouters.forEach(asyncRouter => {
          router.addRoute(asyncRouter);
        });

        // console.log("User info:", userInfo.value);
        if (userInfo.value.authorityId === 888 || userInfo.value.authorityId === 8881) {
          await router.replace({ name: 'dashboard' })
        } else {
          if (!router.hasRoute(userInfo.value.authority.defaultRouter)) {
            ElMessage.error('请联系管理员进行授权')
          }
          await router.replace({ name: userInfo.value.authority.defaultRouter })
        }
        loadingInstance.value.close();
        const isWin = ref(/windows/i.test(navigator.userAgent));
        if (isWin.value) {
          window.localStorage.setItem('osType', 'WIN');
        } else {
          window.localStorage.setItem('osType', 'MAC');
        }
        return true;
      }
    } catch (e) {
      // console.log("Login error:", e);
      loadingInstance.value.close();
    }
    loadingInstance.value.close();
  }
  return {
    UserTgRegister,
    UserRegisterWithInvite,
    UserTgLogin,
    userInfo,
    token,
    NeedInit,
    ResetUserInfo,
    GetUserInfo,
    LoginIn,
    LoginOut,
    setToken,
    loadingInstance,
    ClearStorage
  }
})
