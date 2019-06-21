<template>
    <div>
        <!--父组件: 将值放在定义的 组件名称 的标签中，用于传递(可以最后写父组件，这样标签名好确认)-->
        <header-vue textone="大"></header-vue>
        <body-vue v-bind:texttwo="text2"></body-vue>
        <footer-vue></footer-vue>
    </div>
</template>
<script>

    export default {
        data(){
            return {
                text2:'哈哈呵呵'   
            }
        },
        methods:{},
    }
</script>
<style>
.red{
    background-color: red;
}
.green{
    background-color: yellowgreen;
}
.blue{
    background-color: skyblue;
}
.pink{
    background-color: hotpink;
}
</style>