import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
// 添加vue-resource，或者可以用axios也行，用于http请求的，老师用的vue-resource
import VueResource from 'vue-resource'
// 添加axios用于http请求的，rock另外学习测试下
import Axios from 'axios'

// 使用vue-resource模块
Vue.use(VueResource);
//挂载原型(vue-resource和axios挂载方式不同)
Vue.prototype.$axios = Axios;


Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app');
