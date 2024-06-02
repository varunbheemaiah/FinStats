import {createApp} from 'vue'
import App from './App.vue'
import router from './router';
import 'vuetify/dist/vuetify.min.css'
import './style.css'

// Vuetify
import '@mdi/font/css/materialdesignicons.css';
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import VueApexCharts from "vue3-apexcharts";

let savedTheme = localStorage.getItem('theme') || 'dark'

const vuetify = createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: savedTheme,
  }
})

const app = createApp(App)
app.use(router)
app.use(vuetify)
app.use(VueApexCharts)
app.mount('#app')
