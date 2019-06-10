//1:引入 vue
import Vue from 'vue';

import App from './app.vue';

//引入子组件对象
import headerVue from './components/header.vue';
import bodyVue from './components/body.vue';
import footerVue from './components/footer.vue';


//声明全局组件
Vue.component('headerVue', headerVue); //注册一个组件，第一个参数是名称，在template中使用,第二个参数是实际的对象，显示成什么内容，具备什么功能
Vue.component('bodyVue', bodyVue);
Vue.component('footerVue', footerVue);


new Vue({
    el: '.app',
    // render:function(c){
    //     return c;
    // }
    render: c => c(App)
    //1:当参数是一个的时候，小括号可以省略
    //2:当代码只有一行且是返回值的时候,可以省略大括号

})