import { createRouter, createWebHistory } from 'vue-router'
const routes = [
  { path: '/', redirect: '/assets' },
  { path: '/login', component: () => import('../views/Login.vue') },
  { path: '/assets', component: () => import('../views/AssetList.vue'), meta:{auth:true} },
  { path: '/assets/:id', component: () => import('../views/AssetDetail.vue'), meta:{auth:true} },
  { path: '/assignments', component: () => import('../views/Assignments.vue'), meta:{auth:true} },
  { path: '/maintenance', component: () => import('../views/Maintenance.vue'), meta:{auth:true} },
]
const router = createRouter({ history: createWebHistory(), routes })
router.beforeEach(to => {
  if (to.meta.auth && !localStorage.getItem('token')) return '/login'
})

export default router