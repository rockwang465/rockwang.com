<template>
    <div>
        <ul>
          <li v-for="(stu,index) in stus" v-bind:key="index">
            index:{{index}},stu:{{stu}}
          </li>
      
        使用对象的方式 <hr/>

           <li v-for="(value,key,index) in person" :key="index">
              value:{{value}}
              key:{{key}}
              index:{{index}}
          </li>
        </ul>
    </div>
</template>
<script>

    export default {
        data(){
            return {
                stus:[{name:'jack',score:'A'},{name:'rose',score:'B'}],
                person:{name:'mick',alise:'迈克'}
            }
        },
       
        methods:{
        }
    }
</script>
<style>
.red{
    background-color: red;
}
.green{
    background-color: green;
}
.big{
    font-size: 30px;
}
</style>