//2:引入 Vue
import Vue from 'vue';
import App from './app.vue';



//3:创建一个Vue的实例，一般一个项目，大多就是一个vue实例来完成显示
new Vue({
    // el:'#app', //目的地
    // render:'dom结构'//渲染的内容
    el:'#app',
    render:function(creater){
        return creater(App); //App包含 template/script/style,最终生成DOM结构
    }
})