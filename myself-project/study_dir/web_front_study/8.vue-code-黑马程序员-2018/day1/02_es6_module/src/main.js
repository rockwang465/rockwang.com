//默认导入
// import defaultObj from './cal.js';
// console.log(defaultObj);


//声明式导入(按需加载)

// function test(){//不能包含在函数内部，只能在最外层作用域import和export
//     import {obj1,obj2,obj4} from './cal.js'
//     console.log(obj1,obj2,obj4);
// }
// test();

// import {obj1,obj2,obj4} from './cal.js'
// console.log(obj1,obj2,obj4);
// 

//全体导入
import * as obj from './cal.js';
console.log(obj)