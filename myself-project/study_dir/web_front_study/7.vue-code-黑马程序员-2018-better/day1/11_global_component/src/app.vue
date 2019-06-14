<template>
    <div>
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