# 使用说明

拖拽插件压缩包进入插件安装页面进行安装再进行后续操作

由于有TG发信和获取频道用户信息的需求，所以需要事先安装GVA插件市场中的```灰机消息发送插件```(也即是 https://github.com/spiritysdx/tgm )

## 具体说明

### 后台设置

#### 数据库相关

由于有TG验证码时限删除的需求，需要自行安装并启用Redis数据库

在使用本插件时务必自行安装并配置gva的server目录下的```config.yaml```的

```
redis:
    addr: 127.0.0.1:6379
    password: ""
    db: 0
```

和

```
system:
    ....
    use-redis: true
    ....
```

或

使用管理员权限查看 系统工具 -- 系统配置 -- Redis admin数据库配置 进行配置。

![image](https://github.com/spiritysdx/tgr/assets/97792170/4122a892-0b34-4366-906f-aeb776ca99eb)

务必记得启用，然后在

![image](https://github.com/spiritysdx/tgr/assets/97792170/2ea797dc-59a1-4f2c-8f69-a7a571048370)

这里启用，否则仅配置不启动也没有用

### server

查看 ```server/initialize/plugin.go``` 文件中是否已注册插件，如若未注册，在```import```中插入

```
"server/plugin/register"
```

在函数

```
func InstallPlugin
```

中插入

```
  // 8881 为普通子用户ID，可自行更改替换注册的角色
  PluginInit(PublicGroup, register.CreateRegisterPlug("角色ID", "tgbot的token", "验证码的长度", "频道的chat_id"))
  // 示例
  // PluginInit(PublicGroup, register.CreateRegisterPlug(8881, "7009xxxx:AAExxxxx", 6, "-100197xxxxx"))
```

这里的bot务必已加入到频道中并给予了管理员权限

### web

#### 1

修改 ```web/src/pinia/modules/user.js``` 在 ```import```那块加上

```
import { TGRRegister, TGRLogin } from '@/plugin/register/api/api'
```

在最后的return前后加上

```
...
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
                router.push({name: userInfo.value.authority.defaultRouter});
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
        console.log("try login");
        const res = await TGRLogin(loginInfo);
        console.log("Login response:", res);
        if (res.code === 0) {
            setUserInfo(res.data.user);
            setToken(res.data.token);
            const routerStore = useRouterStore();
            await routerStore.SetAsyncRouter();
            const asyncRouters = routerStore.asyncRouters;
            asyncRouters.forEach(asyncRouter => {
            router.addRoute(asyncRouter);
            });

            console.log("User info:", userInfo.value);
            if (!router.hasRoute(userInfo.value.authority.defaultRouter)) {
            ElMessage.error('请联系管理员进行授权');
            } else {
            console.log("Redirecting to:", userInfo.value.authority.defaultRouter);
            await router.replace({ name: userInfo.value.authority.defaultRouter });
            console.log("Redirected successfully.");
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
        console.log("Login error:", e);
        loadingInstance.value.close();
        }
        loadingInstance.value.close();
    }
    return {
        UserTgRegister,
        UserTgLogin,
        ...
    }
```

#### 2

查看 ```web/src/permission.js``` 并参照

```
const whiteList = ['Login', 'Init', 'Register', 'Admin', 'Resetpwd']
```

添加对应的路由名字

#### 3

查看 ```web/src/router/index.js``` 在

```
const routes
```

中更改和添加对应内容

```
// 以下为有更改的部分
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
// 以下为新增的部分
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
```

定义用户进入的首页是插件中定义的首页，而不是官方GVA定义的初始化页面，login路由你也可以不更改，登录依然使用官方界面

## 后言

所有更改成功后务必保存所有文件再重启前后端，才能正常使用
