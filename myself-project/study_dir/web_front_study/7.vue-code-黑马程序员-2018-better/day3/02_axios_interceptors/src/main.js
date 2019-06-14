//引入一堆
import Vue from 'vue';
//主体
import App from './components/app.vue';

//引入
import Axios from 'axios';

Axios.defaults.baseURL = 'http://182.254.146.100:8899/api/';

//默认设置
Axios.defaults.headers = {
    accept: 'defaults'
}

//拦截器
Axios.interceptors.request.use(function(config) {
    //console.log(config);
    //return false; //返回没有修改的设置,不return config 就是一个拦截
    //个性化的修改
    // config.headers.accept = 'interceptors';
    config.headers = {
        accept: 'interceptors'
    }

    return config;
})





//给Vue原型挂载一个属性
Vue.prototype.$axios = Axios;


//new Vue 启动
new Vue({
    el: '#app',
    render: c => c(App),
})