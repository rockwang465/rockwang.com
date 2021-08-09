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
        // ����������1
        subOne(index) {
            // �����ǰ�鼮������С��2�������ټ��ˡ���if�жϣ����С��2������disabled���ԣ����û������ٲ����ˡ�
            // �� v-bind:disabled="item.count <= 1"
            this.books[index].count--;
        },
        // ����������1
        addOne(index) {
            this.books[index].count++;
        },
        // ɾ�����ﳵ�е��鼮
        removeRecord(index) {
            this.books.splice(index, 1);
        },
        // ����1: д����(getPrice1)
        // ����2: д������filters(������filters: getPrice2)
        getPrice1(price) {
            // toFixed(2) ��ʾ������λС����
            return "$" + price.toFixed(2)
        },
    },
    filters: {
        // ����1: д����(������methods: getPrice1)
        // ����2: д������filters(getPrice2)
        getPrice2(price) {
            return "$" + price.toFixed(2)
        }
    },
    // ����۸�һ��Ҫ�ü������� computed����ס����
    computed: {
        totalPrice() {
            // let amount = 0
            // Ҫ��let����Ҫ��varŶ��
            // ����1: forѭ���������д��
            // for (let i = 0; i < this.books.length; i++) {
            //     // �ܼ�Ϊ: �۸� x ����
            //     amount += this.books[i].price * this.books[i].count
            // }

            // ����2: forѭ������д��, iΪ����ֵ
            // for (let i in this.books){
            //     amount += this.books[i].price * this.books[i].count
            // }

            // ����3: forѭ����򵥵�д��, itemΪ����Ԫ�ص���Ϣ -- �Ƽ�
            // for (let item of this.books) {
            //     amount += item.price * item.count
            // }

            // ����4: ��reduce����forѭ��
            // return this.books.map(function (n) {
            //     return n.price * n.count
            // }).reduce(function (preValue, n) {
            //     return preValue + n
            // }, 0)
            // ����һ�������:
            // return this.books.map(n => n.price * n.count).reduce((preValue, n) => preValue + n)

            // ����5:: reduceһ���߽׺������
            // return this.books.reduce(function (preValue, n) {
            //     return preValue + n.price * n.count
            // }, 0)
            // ����һ�������:
            return this.books.reduce( (preValue, n) => preValue + n.price * n.count, 0)
        },
    }
})

const nums = [10, 40, 20, 60, 160, 70, 110, 30, 120, 5]
// 1.��������
// // ����1: ȡ��С��100������
// let newNum = nums.filter(function (n) {  // nΪ��ѯnums����
//     return n < 100  // ����С��100��ֵ�����Ǹ��жϣ�true�򷵻أ�false�򲻷��ء�
// });
// console.log(newNum);  // ���:[10, 40, 20, 60, 70, 30, 5]
// // filter�б�����д��һ������
//
// // ����2: ������С��100������ת����: ����x2
// let newNum2 = newNum.map(function (n) {
//     return n * 2  // ����x2��ֵ
// })
// console.log(newNum2);  // [20, 80, 40, 120, 140, 60, 10]
//
// // ����3: ����������2�е��������(���)
// let newNum3 = newNum2.reduce(function (preValue, n) {  // preValue��ǰһ��ֵ��nΪ��ǰ��ֵ
//     return preValue + n
// }, 0)  // ��ʼֵΪ0
// console.log(newNum3); // 470

// 2.����
// let newNums = nums.filter(function (n){
//     return n < 100
// }).map(function (n){
//     return n *2
// }).reduce(function (preValue, n){
//    return preValue + n
// },0)
// console.log(newNums); // 470

// 3.�߽��÷�(es6�﷨)
let newNums2 = nums.filter(n => n < 100).map(n => n * 2).reduce((preValue, n) => preValue + n)
console.log(newNums2);