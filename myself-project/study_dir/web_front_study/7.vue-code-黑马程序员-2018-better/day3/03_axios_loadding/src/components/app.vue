<template>
    <div>
        <button @click="send">发请求1</button>
        <button @click="send2">发请求2</button>
    </div>
</template>
<script>
    export default {
        data(){
            return {
                data:[]
            }
        },created(){
        },
        methods:{
          send(){
             //拦截器对每一次请求都有效
            this.$axios.get('getcomments/300?pageindex=1',{
              headers:{
                accept:'get'
              }
            })
            .then(res=>{
              console.log(res);
            })
            .catch(err=>{
              console.log(err)
            })
          },
          send2(){
             //拦截器对每一次请求都有效
            this.$axios.get('getlunbo',{
              headers:{
                accept:'get'
              }
            })
            .then(res=>{
              console.log(res);
            })
            .catch(err=>{
              console.log(err)
            })
          }
        }
    }
</script>
<style scoped>
    .h{
        height: 100px;
        background-color: yellowgreen;
    }
    .b{
        height: 100px;
        background-color: skyblue;
    }
    .f{
        height: 100px;
        background-color: hotpink;
    }
</style>