<template>
    <div>
        HeroDetail
        <div>
            <span>{{ this.hero_info.id }} </span>
            <span>{{ this.hero_info.name }} </span>
            <span>{{ this.hero_info.country }} </span>
            <span>{{ this.hero_info.comment }} </span>
            <div>
                <img :src="this.hero_info.img" :alt="this.hero_info.name">
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "HeroDetail",
        data: function () {
            return {
                id: '',
                hero_info: ''
            }
        },
        created() {
            // console.log(this.$route.query["id"]);
            this.id = this.$route.query["id"];
            this.$axios.get('/apps/get_hero_info/' + this.id)
                .then(res => {
                    // console.log(res);
                    this.hero_info = res['data'];
                    console.log(this.hero_info)
                })
                .catch(err => {
                    console.log(err)
                })
        }
    }
</script>

<style scoped>

</style>