//引入一堆
import Vue from 'vue';
//主体
import App from './components/app.vue';

//引入
import VueResource from 'vue-resource';
//安装插件
Vue.use(VueResource); //插件都是挂载属性
//未来通过this.$http
//  Vue是 所有实例对象的构造函数
//  Vue.protptype.$http ->   实例(this)就可以.$http



//new Vue 启动
new Vue({
    el: '#app',
    render: c => c(App),
})