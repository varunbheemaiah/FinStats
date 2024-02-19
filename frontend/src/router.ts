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
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach((to, from, next) => {
  // Assuming each route has a meta field with pageTitle
  const pageTitle = to.name || 'Page';
  next();
});

export default router;