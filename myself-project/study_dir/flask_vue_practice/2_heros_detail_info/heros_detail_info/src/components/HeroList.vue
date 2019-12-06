<template>
    <div>
        <div v-for="(info,id) in this.all_heros_info" :key="id">
            <span>{{info.id}} {{info.name}} </span>
            <router-link :to="{name: 'herodetail', query: {id: info.id}}">查看</router-link>
        </div>
    </div>
</template>

<script>

    export default {
        name: "HeroList",
        data: function () {
            return {
                all_heros_info: ''
            }
        },
        created() {
            this.$axios.get("/apps/get_all_heros_data")
                .then(res => {
                    // console.log(res['data']['hero_detail_info']);
                    this.all_heros_info = res['data']['hero_detail_info'];
                })
                .catch(err => {
                    console.log(err)
                })
        }
    }
</script>

<style scoped>

</style>