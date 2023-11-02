import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import SvgIcon from '@/components/basic/SvgIcon.vue'
import router from './router'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import '@/assets/icon/iconfont.js'

import '@/assets/css/reset.css'

import '@/interface/index.ts'

import mitt from 'mitt'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'
const app = createApp(App)
const Mit = mitt()

declare module "vue" {
  export interface ComponentCustomProperties {
    $Bus: typeof Mit
  }
}

app.config.globalProperties.$Bus = Mit

const pinia = createPinia()

app.use(pinia)

app.use(router)

app.use(ElementPlus)

app.component('SvgIcon', SvgIcon)

app.mount('#app')


for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}
