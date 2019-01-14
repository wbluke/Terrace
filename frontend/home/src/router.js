import Vue from 'vue';
import Router from 'vue-router';
import Meta from 'vue-meta';

import About from './components/About.vue'
import Blog from './components/Blog.vue'
import Daily from './components/Daily.vue'

Vue.use(Router)
Vue.use(Meta);

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name :'About',
      component: About
    },
    {
      path: '/blog',
      name :'Blog',
      component: Blog,
      children: [
        {
          //send this info. to blog component
          path: ':year/:month/:day/:title', 
        }
      ]
    },
    {
      path: '/daily',
      name :'Daily',
      component: Daily
    }
  ]
});
