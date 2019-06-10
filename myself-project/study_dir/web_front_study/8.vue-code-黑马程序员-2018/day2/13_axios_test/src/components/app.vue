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
           this.$axios.get('getlunbo',{
                params:{
                    id:'123'
                },
                baseURL:'http://www.baidu.com'
           })
           .then(res=>{
            this.data = res.data.message;
           })
           .catch(err=>{
            console.log(err);
           })
           // .finally(()=>{
           //  console.log('最终');
           // })
           // 
           
           //post
           this.$axios.post('postcomment/300',
            'content=传智27期，真正霸气的一批'
           // {
           //  content:'传智27期，再次霸气的一批'
           // },
           // {
           //  //设置
           //  headers:{
           //      'content-type':'application/x-www-form-urlencoded'
           //  }
           // }
           )
           .then(res=>{
                this.data = res.data.message;
           })
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