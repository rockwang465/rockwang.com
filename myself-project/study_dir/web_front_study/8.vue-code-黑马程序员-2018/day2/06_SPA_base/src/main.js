//引入一堆
import Vue from 'vue';
import VueRouter from 'vue-router';

//主体
import App from './components/app.vue';
//路由切换页面
import Music from './components/music.vue'
import Movie from './components/movie.vue'

//安装插件
Vue.use(VueRouter); //挂载属性

//创建路由对象并配置路由规则
let router = new VueRouter({
    //routes
    routes: [
        //一个个对象
        { name: 'music', path: '/music1', component: Music },
        { path: '/movie', component: Movie }

    ]
});

//new Vue 启动
new Vue({
    el: '#app',
    //让vue知道我们的路由规则
    router, //可以简写router
    render: c => c(App),
})