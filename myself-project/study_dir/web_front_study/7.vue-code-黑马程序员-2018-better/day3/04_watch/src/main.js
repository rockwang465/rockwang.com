//引入一堆
import Vue from 'vue';
//主体
import App from './components/app.vue';


new Vue({
    el: '#app',
    render: c => c(App),
})