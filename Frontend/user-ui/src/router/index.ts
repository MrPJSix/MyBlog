import { createRouter, createWebHistory } from 'vue-router'
import { useUserDataStore } from '@/stores/modules/userData'
import { ElMessage } from 'element-plus'
const Login = () => import('@/views/Login.vue')

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/home',
      component: () => import('@/views/home/Home.vue'),
    },
    {
      path: '/login',
      component: Login,
    },
    {
      path: '/',
      component: () => import('@/views/Index.vue'),
      children: [
        {
          path: 'article/:id',
          component: () => import('@/views/Main.vue'),
        },
        {
          path: 'usercenter',
          component: () => import('@/views/user/UserCenter.vue'),
          children: [
            {
              path: '',
              component: () => import('@/views/user/user-center/Home.vue'),
            },
            {
              path: 'config',
              component: () => import('@/views/user/user-center/Config.vue'),
            },
            {
              path: 'mytopics',
              component: () => import('@/views/user/user-center/MyTopics.vue'),
            },
            {
              path: 'myreplies',
              component: () => import('@/views/user/user-center/MyReplies.vue'),
            },
            {
              path: 'myfavorites',
              component: () =>
                import('@/views/user/user-center/MyFavorites.vue'),
            },
            {
              path: 'myfollowings',
              component: () =>
                import('@/views/user/user-center/MyFollowings.vue'),
            },
            {
              path: 'myfans',
              component: () => import('@/views/user/user-center/MyFans.vue'),
            },
          ],
        },
        {
          path: 'post',
          component: () => import('@/views/post/Post.vue'),
        },
        // {
        //   path: 'postMd',
        //   component: () => import('@/components/post/MdEditor.vue'),
        // }
      ],
    },
  ],
})

router.beforeEach((to, from, next) => {
  const userDataStore = useUserDataStore()
  if (localStorage.getItem('token')) {
    userDataStore.updateUserData()
    console.log(userDataStore.avatar_url);
    next()
  } else {
    if (to.path === '/login' || to.path === '/home') {
      next()
    } else {
      ElMessage({
        message: '登录后才能访问哦^0^',
        type: 'error',
      })
      next('/login')
    }
  }
})

export default router
