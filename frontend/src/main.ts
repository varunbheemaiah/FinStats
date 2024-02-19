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

const vuetify = createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'dark',
  }
})

const app = createApp(App)
app.use(router)
app.use(vuetify)
app.mount('#app')
