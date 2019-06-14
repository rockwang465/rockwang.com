<template>
    <div>
        {{data}}
    </div>
</template>
<script>
    export default {
        data(){
            return {
                data:[]
            }
        },created(){
           
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
          


          this.$axios.get('getcomments/300?pageindex=1')
          .then(res=>{
            console.log(res);
          })
          .catch(err=>{
            console.log(err)
          })
          

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