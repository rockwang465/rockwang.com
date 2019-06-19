'use strict';  // "严格模式"是一种在JavaScript代码运行时自动实行更严格解析和错误处理的方法。这种模式使得Javascript在更严格的条件下运行。

//1.3 引入第三方包 开始
import Vue from 'vue';
//1.3.1 VueRouter:引入路由对象
import VueRouter from 'vue-router';
//VueRouter
//VueRouter:安装插件
Vue.use(VueRouter);

//1.3.2 Mint:引入mint-ui
import Mint from 'mint-ui';
//Mint:引入css
import 'mint-ui/lib/style.css';
//Mint:安装插件
Vue.use(Mint);

//1.3.3 Axios:引入axios
import Axios from 'axios';
//挂载原型
Vue.prototype.$ajax = Axios;
// Axios.prototype.$axios = Axios;  //叫$axios最好，老师习惯，叫$ajax的，所以用的上面的


//引入第三方包 结束



//1.1 引入自己的vue文件 开始
import App from './app.vue';
import Home from './components/home.vue';
//引入自己的vue文件 结束


//1.4 VueRouter: 创建对象并配置路由规则
let router = new VueRouter({
    routes: [
        //VueRouter: 配置路由规则
        //默认根目录配置重定向到home
        {path: '/', redirect: {name: 'home'}},
        {name: 'Home', path: '/home', component: Home}
    ]
});


//1.2 创建vue实例
new Vue({
    el: '#app',
    router, //如果这里不加，则会报错:Vue warn]: Error in render: "TypeError: Cannot read property 'matched' of undefined"
    render: c => c(App)
});


