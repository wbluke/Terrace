import Vue from 'vue';
import Router from 'vue-router';

import About from './components/About.vue'
import Study from './components/Study.vue'
import Daily from './components/Daily.vue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name :'About',
      component: About
    },
    {
      path: '/study',
      name :'Study',
      component: Study
    },
    {
      path: '/daily',
      name :'Daily',
      component: Daily
    }
  ]
});
