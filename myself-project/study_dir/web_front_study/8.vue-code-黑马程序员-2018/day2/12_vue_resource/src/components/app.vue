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
            // console.log(this.$http);
            // this.$http.get('http://182.254.146.100:8899/api/getlunbo')
            // .then(res=>{
            //     this.data = res.body.message;
            // },err=>{
            //     console.log(err)
            // })
            
            //post请求
            this.$http.post('http://182.254.146.100:8899/api/postcomment/300',{
                content:'传智27期，霸气的一批'
            },{
                emulateJSON:true
            })
            .then(res=>{
                this.data = res.body.message;
            },err=>{
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