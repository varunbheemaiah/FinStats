import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import HomePage from './components/Home.vue';
import NasdaqEarningsCalendar from './components/NasdaqEarningsCalendar.vue';
import Greeting from './components/Greeting.vue';
import NSEBSEEarningsCalendar from './components/NSEBSEEarningsCalendar.vue';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: HomePage
  },
  {
    path: '/nasdaq-earnings',
    name: 'NASDAQ Earnings Calendar',
    component: NasdaqEarningsCalendar
  },
  {
    path: '/nse-earnings',
    name: 'NSE/BSE Earnings Calendar',
    component: NSEBSEEarningsCalendar
  },
  {
    path: '/greeting',
    name: 'Hello',
    component: Greeting
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach((to, from, next) => {
  const pageTitle = to.name || 'Page';
  document.title = String(pageTitle);
  next();
});

export default router;