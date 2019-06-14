<template>
    <div>
        我是{{textone}}头
            
            <button @click="show">从js中获取父组件值</button>
    </div>
</template>
<script>
    export default {
        data(){
            return {

            }
        },
        //接受父组件值参数的设置
        props:['textone'],
        methods:{
            show(){
                alert(this.textone)
            }
        }
    }
</script>
<style scoped>
 div{
    height: 100px;
    background-color: yellowgreen;
 }
</style>