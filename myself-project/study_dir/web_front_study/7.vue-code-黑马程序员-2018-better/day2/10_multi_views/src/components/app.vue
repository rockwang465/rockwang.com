<template>
    <div>
        <div>我是头部啊</div>
        <hr/>
            <!-- 留坑，非常重要 坑名-->
            <router-view class="b" name="header"></router-view>
            <router-view class="b" ></router-view>
            <router-view class="b" name="footer"></router-view>
        <hr/>
        <div>我是底部啊</div>
    </div>
</template>
<script>
    export default {
        data(){
            return {

            }
        }
    }
</script>
<style scoped>
    .h{
        height: 100px;
        background-color: yellowgreen;
    }
    .b{
        height: 100px;
        background-color: skyblue;
    }
    .f{
        height: 100px;
        background-color: hotpink;
    }
</style>