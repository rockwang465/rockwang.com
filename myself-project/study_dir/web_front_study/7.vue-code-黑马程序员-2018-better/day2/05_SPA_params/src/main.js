//引入一堆
import Vue from 'vue';
import VueRouter from 'vue-router';

//主体
import App from './components/app.vue';
//路由切换页面
import List from './components/list.vue'
import Detail from './components/detail.vue'

//安装插件
Vue.use(VueRouter); //挂载属性

//创建路由对象并配置路由规则
let router = new VueRouter({
    //routes
    routes: [
        //一个个对象
        { name: 'list', path: '/list', component: List },
        //以下规则匹配  /detail? xxx = xx & xxx = xxx 多少个查询字符串都不影响
        //查询字符串path不用改
        { name: 'detail', path: '/detail', component: Detail },
        //
        //  {name:'detail',params:{id:index}  } -> /detail/12
        { name: 'detail', path: '/detail/:id', component: Detail }

    ]
});

//new Vue 启动
new Vue({
    el: '#app',
    //让vue知道我们的路由规则
    router, //可以简写router
    render: c => c(App),
})