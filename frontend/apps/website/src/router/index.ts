import { createRouter, createWebHistory } from 'vue-router'

import { rootRoute } from './root'
import { mainRoute } from './main'
import { projectDetailRoute } from './project.detail'
import { loginRoute, registerRoute, notFoundRoute } from './base'

export const router = createRouter({
  history: createWebHistory(),
  routes: [rootRoute, loginRoute, registerRoute, mainRoute, projectDetailRoute, notFoundRoute],
})

export * from './base'
export * from './root'
export * from './main'
export * from './project.detail'

// 路由拦截器
export * from './filter'

export default router
