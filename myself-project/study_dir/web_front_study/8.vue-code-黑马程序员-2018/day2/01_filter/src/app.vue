<template>
    <div>
       <!--  请输入内容
        <input type="text" name="" v-model="text">
        显示：{{text | myFilter}} -->

        <!-- 测试组件的范围 -->
        <sub-vue></sub-vue>

        <div></div>
    </div>
</template>
<script>
    import SubVue from './sub.vue';
    export default {
        filters:{
            myFilter:function(value){  //value 就是-> text
                //输入的内容帮我做一个反转
                //转换成数组，反转数组，转换成字符串
                
                return value.split('').reverse().join('');
            }
        },
        data(){
            return {
                text:''
            }
        },
        components:{
            subVue:SubVue
        }
    }
</script>
<style>
    
</style>