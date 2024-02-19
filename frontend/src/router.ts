import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import NasdaqEarningsCalendar from './components/NasdaqEarningsCalendar.vue';
import Greeting from './components/Greeting.vue';

const routes: RouteRecordRaw[] = [
  { path: '/nasdaq-earnings', component: NasdaqEarningsCalendar },
  { path: '/greeting', component: Greeting }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;