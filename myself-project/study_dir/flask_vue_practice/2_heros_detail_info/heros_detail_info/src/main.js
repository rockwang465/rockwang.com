import Vue from 'vue'
import App from './App.vue'
// import router from './router'
import store from './store'
import axios from 'axios'
import VueRouter from 'vue-router'
import HeroList from "./components/HeroList";
import HeroDetail from "./components/HeroDetail";

Vue.prototype.$axios = axios;  //挂载原形
// if (process.env.NODE_ENV !== 'production') require('@/mock')
axios.defaults.baseURL = "http://10.151.106.247:5001/v1";  //设置基础url,解决跨域问题
// axios.defaults.baseURL = "";  //设置基础url,解决跨域问题
Vue.use(VueRouter);

Vue.config.productionTip = false;

let router = new VueRouter({
    mode: 'history',  // 用于解决 url后有个 "#/" 存在的问题，加了这个，则不显示此锚点了
    linkActiveClass: 'mui-active',
    routes: [
        //VueRouter: 配置路由规则
        //name为路由的名字，path为url上显示的地址，component为上面引入文件的名字。
        //默认根目录配置重定向到home
        {path: '/', redirect: {name: 'herolist'}},
        {name: 'herolist', path: '/herolist', component: HeroList},
        {name: 'herodetail', path: '/herodetail', component: HeroDetail},

        // {name: 'photo.share', path: '/photo/share', component: PhotoShare},
        // 1.4.6 添加图文详细的规则(这里的:id 和 photoShare.vue中的params:{id:img.id}这个id对应)
        // {name: 'photo.detail', path: '/photo/detail/:id', component: PhotoDetail},
    ]
});

//1.2 创建vue实例
new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app');
