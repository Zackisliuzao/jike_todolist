import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/login'
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/views/auth/Login.vue'),
      meta: { requiresAuth: false }
    },
    {
      path: '/register',
      name: 'Register',
      component: () => import('@/views/auth/Register.vue'),
      meta: { requiresAuth: false }
    },
    {
      path: '/dashboard',
      name: 'Dashboard',
      component: () => import('@/views/layout/MainLayout.vue'),
      meta: { requiresAuth: true },
      redirect: '/dashboard/tasks',
      children: [
        {
          path: 'tasks',
          name: 'Tasks',
          component: () => import('@/views/tasks/Tasks.vue'),
          meta: { requiresAuth: true }
        },
        {
          path: 'tasks/:id',
          name: 'TaskDetail',
          component: () => import('@/views/tasks/TaskDetail.vue'),
          meta: { requiresAuth: true }
        },
        {
          path: 'profile',
          name: 'Profile',
          component: () => import('@/views/user/Profile.vue'),
          meta: { requiresAuth: true }
        },
        {
          path: 'categories',
          name: 'Categories',
          component: () => import('@/views/categories/Categories.vue'),
          meta: { requiresAuth: true }
        }
      ]
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'NotFound',
      component: () => import('@/views/error/NotFound.vue')
    }
  ]
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  const requiresAuth = to.meta.requiresAuth

  // 初始化用户信息
  userStore.initUserInfo()

  if (requiresAuth && !userStore.isAuthenticated) {
    next('/login')
  } else if (!requiresAuth && userStore.isAuthenticated && (to.name === 'Login' || to.name === 'Register')) {
    next('/dashboard/tasks')
  } else {
    next()
  }
})

export default router
