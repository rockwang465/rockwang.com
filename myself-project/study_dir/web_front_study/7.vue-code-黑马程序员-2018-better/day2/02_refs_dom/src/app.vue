<template>
    <div>
    
        <sub-vue ref="sub"></sub-vue>

       <!--  <div ref="myDiv">
            {{text}}
        </div> -->
    </div>
</template>
<script>
    import SubVue from './sub.vue';
    export default {
        data(){
            return {
                text:'123'
            }
        },
        components:{
            subVue:SubVue
        },
        //组件创建后,  数据已经完成初始化，但是DOM还未生成
        created(){  //  事件的处理函数(created)
            console.log('created:',this.$refs.myDiv);//获取不到
        },
        //数据装载DOM上后，各种数据已经就位,将数据渲染到DOM上，DOM已经生成
        mounted(){
            // console.log('sub:',this.$refs.sub.$el);
            
            //获取组件对象，并获取到其的DOM对象
            this.$refs.sub.$el.innerHTML = '嘻嘻';

            // console.log('mounted:',this.$refs.myDiv);

            // 涉及DOM类的操作
            // this.$refs.myDiv.innerHTML = '哈哈呵呵';
            // 涉及到数据的操作
            this.text = '嘻嘻呵呵'
        }

    }
</script>
<style>
    
</style>