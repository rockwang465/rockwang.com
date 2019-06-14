//引入一堆
import Vue from 'vue';
//主体
import App from './components/app.vue';

//引入
import Axios from 'axios';

//引入mint-ui
import Mint from 'mint-ui'; //  export default 整个对象
// import { Indicator } from 'mint-ui'; //export 整个对象.Indicator -> {Indicator}
//引入css
import 'mint-ui/lib/style.css';
//安装插件，注册一堆全局组件
Vue.use(Mint);




Axios.defaults.baseURL = 'http://182.254.146.100:8899/api/';

//默认设置
Axios.defaults.headers = {
    accept: 'defaults'
}

//拦截器
Axios.interceptors.request.use(function(config) {
    Mint.Indicator.open();
    //请求发起之前  显示loadding
    return config;
})

Axios.interceptors.response.use(function(config) {
    //在响应回来之后，隐藏loadding
    Mint.Indicator.close();
    // console.log(config);
    return config;
})



//给Vue原型挂载一个属性
Vue.prototype.$axios = Axios;


//new Vue 启动
new Vue({
    el: '#app',
    render: c => c(App),
})