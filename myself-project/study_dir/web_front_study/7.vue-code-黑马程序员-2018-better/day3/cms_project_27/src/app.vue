<template>
    <div>
        <!--头部内容-->
        <mt-header title="传智27期信息管理系统"></mt-header>

        <!--中间主体内容-->
        <router-view></router-view>

        <!--底部内容-->
        <nav class="mui-bar mui-bar-tab">
            <router-link class="mui-tab-item" :to="{name:'home'}">
                <span class="mui-icon icon-ai-home"></span>
                <span class="mui-tab-label">首页</span>
            </router-link>
            <router-link class="mui-tab-item" :to="{name:'vip'}">
                <span class="mui-icon icon-huiyuan3"></span>
                <span class="mui-tab-label">会员</span>
            </router-link>
            <router-link class="mui-tab-item" :to="{name:'shopcart'}">
                <span class="mui-icon icon-gouwucheman"><span class="mui-badge">9</span></span>
                <span class="mui-tab-label">购物车</span>
            </router-link>
            <router-link class="mui-tab-item" :to="{name:'search'}">
                <span class="mui-icon icon-sousuo"></span>
                <span class="mui-tab-label">查找</span>
            </router-link>
        </nav>
    </div>
</template>

<script>
    export default {
        data() {
            return {}
        },
        created() {
            // this.$ajax.get('4.txt')
            //     .then(res => {
            //         console.log(res)
            //     })
        }
    }
</script>

<style scoped>

</style>