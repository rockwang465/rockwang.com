<template>
    <div id="app">
        <div id="nav">
            <router-view></router-view> <!--必须用router-view,否则无法使用router的跳转页面功能，会一直卡在原来页面，url也不变化-->
        </div>
    </div>
</template>

<script>
    import HeroList from './components/HeroList'

    export default {
        name: 'App',
        components: {
            HeroList,
        }
    }

</script>

<style>
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
    }

    /*#nav a.router-link-exact-active {*/
    /*color: #42b983;*/
    /*}*/
</style>
