var obj = {
    message: 'hello',
    rawHtml: '<h1>big</h1>',
    seen: true,
    firstname: 'mike',
    lastname: 'lee'
}
// Object.freeze(obj)

var app = new Vue({
    el: '#app',
    data: obj,
    methods:{
        doSomething () {
            alert('click')
        }
    },
    computed:{
        name: {
            get: function () {
                return this.firstname + ' '+ this.lastname
            },
            set: function (newvalue) {
                const names = newvalue.split(' ')
                this.firstname = names[0]
                this.lastname = names[names.length-1]
            }

        },
        reversedmessage: function () {
            return this.message.split('').reverse().join('')
        }
    }
})

