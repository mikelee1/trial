var obj = {
    name: '',
    age: '',
    persons: [
        {name: 'mike', age: 27}
    ]
}

var app = new Vue({
    el: '#app',
    data: obj,
    methods:{
        create:function () {
            this.persons.push({name:this.name,age:this.age})
        },
        delete1:function (index) {
            this.persons.splice(index,1)
        }
    },
})

