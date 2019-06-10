//引入vue
import Vue from 'vue';

//引入app.vue
import App from './app.vue';

//new Vue
new Vue({
    el: '#app',
    render: c => c(App)
})