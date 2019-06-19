<template>
    <div>
        <mt-button type="primary">测试mint-ui按钮</mt-button>
        <router-view></router-view>
    </div>
</template>

<script>
    export default {
        data() {
            return {}
        },
        created() {
            this.$ajax.get('http://10.5.1.80:6001/images/4.txt')
                .then(res => {
                    console.log(res)
                })
        }
    }
</script>

<style scoped>

</style>