import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'

import './assets/main.css'

// main.ts

// Create Vue app instance and mount it to #app
const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
