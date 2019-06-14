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
           
           //将两个请求一起发送，只要有一个失败，就算失败，成功只有是全体成功
           
           function getMsg(res1,res2){
             console.log('成功啦');
             console.log(res1);
             console.log(res2);
           }
           // 获取省市数据的需求
           this.$axios.all([ 
            this.$axios.post('postcomment1/300','content=123'),
            this.$axios.get('getcomments/300?pageindex=1')
           ])
           //分发响应
           .then(this.$axios.spread(getMsg))
           .catch(err=>{
            console.log(err);
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