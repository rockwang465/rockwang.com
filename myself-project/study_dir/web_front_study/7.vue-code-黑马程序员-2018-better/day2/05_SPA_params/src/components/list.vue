<template>
    <div>
        查询字符串列表:
        <ul>
            <li v-for="(hero,index) in heros" :key="index">
                    {{hero.name}}  
                    <!-- 1:去哪里  /detail?id=12 -->
                    <router-link :to="{name:'detail',query:{id:index}  }">查看</router-link>
            </li>
        </ul>
       <!--   path列表:
        <ul>
            <li v-for="(hero,index) in heros" :key="index">
                    {{hero.name}}  
                    1:去哪里  /detail/12
                    <router-link :to="{name:'detail',params:{id:index}  }">查看</router-link>
            </li>
        </ul> -->
    </div>
</template>
<script>
    export default {
        data(){
            return {
                heros:[{
                    name:'李白'
                },{
                    name:'王昭君'
                }]
            }
        }
    }
</script>
<style>
    
</style>