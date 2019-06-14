<template>
    <div>
        <!-- <input type="text" name="" v-model="text">
        {{text}} -->

        <button @click="changeValue">改变值</button>
    </div>
</template>
<script>
    export default {
        data(){
            return {
                text:'',
                persons:[{
                  name:'jack'
                },{
                  name:'rose'
                }]
            }
        },created(){},
        methods:{
          changeValue(){
            // this.text = 'abc';
            this.persons[0].name = 'mick';
        }
      },
      watch:{
        text:function(newV,oldV){
          console.log('值改变了，很开心')
        },
        persons:{
          handler: function (val, oldVal) { 
              console.log('person改变了。很开心')
           },
          deep: true
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