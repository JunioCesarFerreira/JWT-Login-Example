import { createRouter, createWebHistory } from 'vue-router'
import Login from '@/views/Login.vue';
import Protected from '@/views/Protected.vue';

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/protected',
    name: 'Protected',
    component: Protected
  }
];


const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
