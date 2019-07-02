'use strict';  // "严格模式"是一种在JavaScript代码运行时自动实行更严格解析和错误处理的方法。这种模式使得Javascript在更严格的条件下运行。

//1.3 引入第三方包 开始 ++++++++++++++++++++++++++++++++++++++++
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

//1.3.4 引入mui的样式
import './static/vendor/mui/dist/css/mui.css'
//1.3.5 引入全局样式(自定义的)
import './static/css/global.css'

//1.3.3 Axios:引入axios
import Axios from 'axios';
//挂载原型
Vue.prototype.$ajax = Axios;
// Axios.prototype.$axios = Axios;  //叫$axios最好，老师习惯，叫$ajax的，所以用的上面的
//1.3.3.1 Axios的默认配置
Axios.defaults.baseURL = "http://10.5.1.80:6001/images/";

//1.3.6 Moment:引入moment.js插件
import Moment from 'moment';

//1.3.7 VuePreview:引入vue-preview
import VuePreview from 'vue-preview';
//挂载使用
// Vue.use(VuePreview);  //旧版本vue-preview引用挂载使用方法--已经不好用了
//新版本vue-preview挂载使用方法
Vue.use(VuePreview, {
  mainClass: 'pswp--minimal--dark',
  barsSize: {top: 0, bottom: 0},
  captionEl: false,
  fullscreenEl: false,
  shareEl: false,
  bgOpacity: 0.85,
  tapToClose: true,
  tapToToggleControls: false
});

//引入第三方包 结束 ++++++++++++++++++++++++++++++++++++++++++++


//1.6 引入全局组件需要的组件对象 开始 +++++++++++++++++++++++++++
import NavBar from './components/common/navBar.vue';
//引入全局组件需要的组件对象 结束 +++++++++++++++++++++++++++++++


//1.5 定义全局过滤器或全局组件，方便大家都能用  开始 +++++++++++++
Vue.filter('convertDate',function (value) {  //convertDate为过滤器名称，value为其他地方传进来的值，对这个value进行转换
    return Moment(value).format('YYYY-MM-DD');
});
Vue.component('navBar',NavBar);  //标签使用时最好用nav-bar
//定义全局过滤器或全局组件，方便大家都能用  结束 +++++++++++++++++



//1.1 引入自己的vue文件 开始 +++++++++++++++++++++++++++++++++++
import App from './app.vue';
import Home from './components/home/home.vue';
//1.1.2 引入底部按钮对应的组件文件
import Search from './components/search/search.vue';
import Vip from './components/vip/vip.vue';
import Shopcart from './components/shopcart/shopcart.vue';
//1.1.3 引入新闻列表文件
import NewsList from './components/news/newsList.vue';
//1.1.4 引入新闻详细文件
import NewsDetail from './components/news/newsDetail.vue';
//1.1.5 引入图文分享文件
import PhotoShare from './components/photo/photoShare.vue';
//1.1.6 引入图文详细文件
import PhotoDetail from './components/photo/photoDetail.vue';
//引入自己的vue文件 结束 +++++++++++++++++++++++++++++++++++++++


//1.4 VueRouter: 创建对象并配置路由规则
let router = new VueRouter({
    linkActiveClass: 'mui-active',
    routes: [
        //VueRouter: 配置路由规则
        //name为路由的名字，path为url上显示的地址，component为上面引入文件的名字。
        //默认根目录配置重定向到home
        {path: '/', redirect: {name: 'home'}},
        {name: 'home', path: '/home', component: Home},  //注意前面home和上面的home要一致
        // 1.4.2 添加底部按钮的规则
        {name: 'search', path: '/search', component: Search},
        {name: 'vip', path: '/vip', component: Vip}, //会员
        {name: 'shopcart', path: '/shopcart', component: Shopcart},
        // 1.4.3 添加新闻列表的规则
        {name: 'news.list', path: '/news/list', component: NewsList},
        // 1.4.4 添加新闻详情的规则
        {name: 'news.detail', path: '/news/detail', component: NewsDetail},
        // 1.4.5 添加图文分享的规则
        {name: 'photo.share', path: '/photo/share', component: PhotoShare},
        // 1.4.6 添加图文详细的规则(这里的:id 和 photoShare.vue中的params:{id:img.id}这个id对应)
        {name: 'photo.detail', path: '/photo/detail/:id', component: PhotoDetail},
    ]
});


//1.2 创建vue实例
new Vue({
    el: '#app',
    router, //如果这里不加，则会报错:Vue warn]: Error in render: "TypeError: Cannot read property 'matched' of undefined"
    render: c => c(App)
});


