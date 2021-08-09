const app = new Vue({
    el: '#app',
    data: {
        books: [
            {
                id: 1,
                name: "Python",
                date: "2006-9",
                price: 85.00,
                count: 1
            },
            {
                id: 2,
                name: "Unix",
                date: "2006-2",
                price: 59.00,
                count: 1
            },
            {
                id: 3,
                name: "C++/C",
                date: "2008-10",
                price: 39.00,
                count: 1
            },
            {
                id: 4,
                name: "Golang",
                date: "2006-3",
                price: 128.00,
                count: 1
            }
        ],
        amount: 0,
    },
    methods: {
        // 购买数量减1
        subOne(index) {
            // 如果当前书籍的数量小于2，则不能再减了。先if判断，如果小于2则增加disabled属性，让用户不能再操作了。
            // 即 v-bind:disabled="item.count <= 1"
            this.books[index].count--;
        },
        // 购买数量加1
        addOne(index) {
            this.books[index].count++;
        },
        // 删除购物车中的书籍
        removeRecord(index) {
            this.books.splice(index, 1);
        },
        // 方法1: 写函数(getPrice1)
        // 方法2: 写过滤器filters(见下面filters: getPrice2)
        getPrice1(price) {
            // toFixed(2) 表示保留两位小数点
            return "$" + price.toFixed(2)
        },
    },
    filters: {
        // 方法1: 写函数(见上面methods: getPrice1)
        // 方法2: 写过滤器filters(getPrice2)
        getPrice2(price) {
            return "$" + price.toFixed(2)
        }
    },
    // 计算价格，一定要用计算属性 computed，记住！！
    computed: {
        totalPrice() {
            // let amount = 0
            // 要用let，不要用var哦！
            // 方法1: for循环最基础的写法
            // for (let i = 0; i < this.books.length; i++) {
            //     // 总价为: 价格 x 数量
            //     amount += this.books[i].price * this.books[i].count
            // }

            // 方法2: for循环进阶写法, i为索引值
            // for (let i in this.books){
            //     amount += this.books[i].price * this.books[i].count
            // }

            // 方法3: for循环最简单的写法, item为单个元素的信息 -- 推荐
            // for (let item of this.books) {
            //     amount += item.price * item.count
            // }

            // 方法4: 用reduce代替for循环
            // return this.books.map(function (n) {
            //     return n.price * n.count
            // }).reduce(function (preValue, n) {
            //     return preValue + n
            // }, 0)
            // 或者一条语句解决:
            // return this.books.map(n => n.price * n.count).reduce((preValue, n) => preValue + n)

            // 方法5:: reduce一个高阶函数解决
            // return this.books.reduce(function (preValue, n) {
            //     return preValue + n.price * n.count
            // }, 0)
            // 或者一条语句解决:
            return this.books.reduce( (preValue, n) => preValue + n.price * n.count, 0)
        },
    }
})

const nums = [10, 40, 20, 60, 160, 70, 110, 30, 120, 5]
// 1.基础入门
// // 需求1: 取出小于100的数字
// let newNum = nums.filter(function (n) {  // n为轮询nums数组
//     return n < 100  // 返回小于100的值，这是个判断，true则返回，false则不返回。
// });
// console.log(newNum);  // 结果:[10, 40, 20, 60, 70, 30, 5]
// // filter中必须先写上一个函数
//
// // 需求2: 将所有小于100的数字转换成: 数字x2
// let newNum2 = newNum.map(function (n) {
//     return n * 2  // 返回x2的值
// })
// console.log(newNum2);  // [20, 80, 40, 120, 140, 60, 10]
//
// // 需求3: 将所有需求2中的数字相加(求和)
// let newNum3 = newNum2.reduce(function (preValue, n) {  // preValue是前一个值，n为当前的值
//     return preValue + n
// }, 0)  // 初始值为0
// console.log(newNum3); // 470

// 2.进阶
// let newNums = nums.filter(function (n){
//     return n < 100
// }).map(function (n){
//     return n *2
// }).reduce(function (preValue, n){
//    return preValue + n
// },0)
// console.log(newNums); // 470

// 3.高阶用法(es6语法)
let newNums2 = nums.filter(n => n < 100).map(n => n * 2).reduce((preValue, n) => preValue + n)
console.log(newNums2);