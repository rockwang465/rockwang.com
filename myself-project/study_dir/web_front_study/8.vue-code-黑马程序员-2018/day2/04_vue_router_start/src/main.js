// 1.1 引入Vue
import Vue from 'vue';
// 1.2 导入App
import App from './app.vue';

// 2.4 导入Home组件
import Home from './home.vue';

// 2.1 导入VueRouter
import VueRouter from 'vue-router';

// 2.2 安装插件
Vue.use(VueRouter); //挂载属性

// 2.3 创建路由对象并配置路由规则
let router = new VueRouter({
    //routes
    routes: [
        // 一个个对象
        {path: '/home', component: Home}
    ]
});

// 1.3 使用Vue组件功能
new Vue({
    el: '#app',
    router: router,  //简写 router
    render: c => c(App),
    // 2.5 让vue知道我们的路由规则
});