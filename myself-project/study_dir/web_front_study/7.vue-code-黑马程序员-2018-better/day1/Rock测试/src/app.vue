<template>
    <div>
        <ul v-for="(item,index) in heros" :key="index">
            <!--<li :class="isRed?'red':'green'">{{item.id}}&#45;&#45;{{item.name}}</li>-->
            <!--1.1 不同的成绩，使用不同的颜色 --很难 -->
            <li :class="{'A':'red','B':'blue','C':'green','D':'pink'}[item.score]">{{item.id}}--{{item.name}}</li>
            <button @click="del(index)">删除</button>
        </ul>

        <br/><hr><br/>
        英雄id: <input type="text" v-model="id"><br/>
        英雄姓名: <input type="text" v-model="name"><br/>
        英雄成绩: <input type="text" v-model="score"><br/>
        <button @click="addHero">添加英雄</button>
    </div>
</template>
<script>
    export default {
        data(){
            return{
                id: '',
                name: '',
                score: '',
                heros: [
                    {id:1, name:'奥利安娜',score: 'A'},
                    {id:2, name:'崔斯特',score: 'B'},
                    {id:3, name:'赵信',score: 'C'},
                    {id:4, name:'泰达米尔',score: 'D'}
                ],
                isRed: true
            }
        },
        methods:{
            addHero(){
                this.heros.push({id: this.id, name: this.name, score: this.score});
                this.id=this.name=this.score='';
            },
            del(index){
                this.heros.splice(index,1);
            }
        }
    }

</script>
<style>
    body, li {
        margin: 0;
        padding: 0;
    }
    li {
        height: 30px;
        list-style-type: none;
        line-height: 30px;
    }
    .red {
        background-color: darkred;
    }
    .green {
        background-color: greenyellow;
    }
    .pink {
        background-color: hotpink;
    }
    .blue {
        background-color: lightskyblue;
    }
</style>