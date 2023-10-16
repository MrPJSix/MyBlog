import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import SvgIcon from '@/components/basic/SvgIcon.vue'
import router from './router'

import '@/assets/icon/iconfont.js'

import '@/assets/css/reset.css'

import '@/interface/index.ts'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.component('SvgIcon', SvgIcon)

app.mount('#app')
