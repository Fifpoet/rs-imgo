import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import utils from './utils'

const app = createApp(App);
app.config.globalProperties.$utils = utils;
app.mount('#app');