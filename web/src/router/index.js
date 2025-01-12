import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    redirect: '/login'
  },
  {
    path: '/init',
    name: 'Init',
    component: () => import('@/view/init/index.vue')
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/view/login/index.vue')
  },
  {
    path: '/:catchAll(.*)',
    meta: {
      closeTab: true
    },
    component: () => import('@/view/error/index.vue')
  },
  {
    path: '/h5/member/detail',
    name: 'memberH5Detail',
    component: () => import('@/view/member/mMemberInfo/h5Detail.vue'),
    meta: {
      title: '会员详情',
    }
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
