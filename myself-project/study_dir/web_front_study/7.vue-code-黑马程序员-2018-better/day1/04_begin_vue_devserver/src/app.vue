<template>
    <div>
           我们的第一个vue1程序<br/>
           {{text}}
           输入：<input type="text" v-model="text">
            <ul>
                <li v-for="person in list">
                    {{person.name}}
                </li>
            </ul>

    </div>
<!--     <div>
        只能包含一个根节点 2.x以后
    </div> -->
</template>
<script>
    export default {
        //配置
        // 类似$scope.xxx = '初始化'
        // data:function(){
            
        // },
        data(){
            return {  //放数据的地方
                text:'大家好',
                list:[{name:'jack'},{name:'rose'}]
            }
        }
    }
</script>
<style></style>