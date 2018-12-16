import Vue from 'vue';
import VueRouter from 'vue-router';
import About from '../components/About';
import Documentation from '../components/Documentation';
import Home from '../components/Home';
import NotFound from '../components/NotFound.vue';

Vue.use(VueRouter);

const router = new VueRouter({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: Home
    },
    {
      path: '/about',
      component: About
    },
    {
      path: '/documentation',
      component: Documentation
    },
    {
      path: '*',
      component: NotFound
    }
  ]
});

export default router;
