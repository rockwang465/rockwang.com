// 引入Vue
import Vue from 'vue';
import App from './app.vue';

new Vue({
    el: '#app',
    // render: function (c) {
    //     return c;
    // }
    // 简写为:
    render: c => c(App)
});