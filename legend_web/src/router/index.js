import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '../views/Main.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Main',
    component: Main,
    children: [
      {
        path: '/',
        name: 'home',
        component: () => import('@/views/Home/Home')
      },
      {
        path: '/shenlong',
        name: 'shenlong',
        component: () => import('@/views/ShenLong/ShenLong')
      },
      {
        path: '/shengyu',
        name: 'shengyu',
        component: () => import('@/views/ShengYu/ShengYu')
      },
      {
        path: '/zhenying',
        name: 'zhenying',
        component: () => import('@/views/ZhenYing/ZhenYing')
      },
      {
        path: '/shengxiao',
        name: 'shengxiao',
        component: () => import('@/views/ShengXiao/ShengXiao')
      },
      {
        path: '/honghuang',
        name: 'honghuang',
        component: () => import('@/views/HongHuang/HongHuang')
      },
      {
        path: '/humanoid',
        name: 'humanoid',
        component: () => import('@/views/Humanoid/Humanoid')
      },
    
    ]
  }
  
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
