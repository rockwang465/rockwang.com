<template>
    <div>
        我是body,{{texttwo}}
    </div>
</template>
<script>
    export default {
        data(){
            return {

            }
        },
        props:['texttwo']
    }
</script>
<style scoped>
     div{
    height: 100px;
    background-color: skyblue;
 }
</style>