<template>
    <div>
        我是底部
        <button @click="callDaddy">打电话给爸爸</button>
    </div>
</template>
<script>
import connect from '../connector.js';
    export default {
        data(){
            return {
            }
        },methods:{
            callDaddy(){
                //发射信号
                connect.$emit('phone','62分钟来');
            }
        }
    }
</script>
<style scoped>
     div{
    height: 100px;
    background-color: hotpink;
 }
</style>