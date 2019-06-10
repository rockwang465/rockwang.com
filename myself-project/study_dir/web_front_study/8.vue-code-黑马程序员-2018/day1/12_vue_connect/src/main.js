//1:引入 vue
import Vue from 'vue';

import App from './app.vue';

//引入子组件对象
import sub from './components/sub.vue';

Vue.component('subVue', sub);


new Vue({
    el: '.app',
    // render:function(c){
    //     return c;
    // }
    render: c => c(App)
    //1:当参数是一个的时候，小括号可以省略
    //2:当代码只有一行且是返回值的时候,可以省略大括号

})