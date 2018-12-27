import Vue from 'vue'
import App from './App.vue'
import router from './router.js'
import VueMq from 'vue-mq'

Vue.use(VueMq, {
  breakpoints: {
    mobile: 450,
    tablet: 900,
    laptop: 1250,
    desktop: Infinity,
  }
})

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
