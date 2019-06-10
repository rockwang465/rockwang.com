<template>
    <div>
        <div class="h">头部
       <!--  <a href="#/mymusic">进入音乐1</a>
        <a href="#/movie">进入电影1</a>
        <a href="#/mymusic">进入音乐2</a>
        <a href="#/movie">进入电影2</a>
        <a href="#/mymusic">进入音乐3</a>
        <a href="#/movie">进入电影3</a> -->

    <!-- 1:去哪里 -->
    <router-link :to="{name:'music'}">进入音乐1</router-link>
    <router-link :to="{name:'music'}">进入音乐2</router-link>
    <router-link :to="{name:'music'}">进入音乐3</router-link>

    <!-- 也可以跟路径使用 -->
    <router-link to="/movie">进入电影</router-link>

        </div>
        <!-- 留坑，非常重要 -->
        <router-view class="b"></router-view>
        <div class="f">底部</div>
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