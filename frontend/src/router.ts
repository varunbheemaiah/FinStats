import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import HomePage from './components/Home.vue';
import NasdaqEarningsCalendar from './components/NasdaqEarningsCalendar.vue';
import Greeting from './components/Greeting.vue';

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