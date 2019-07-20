var obj = {
        showc1: false,
        showc2: true
}
// Object.freeze(obj)

var app = new Vue({
    el: '#app',
    data: obj,
    watch: {
        question: function () {
            this.tip = 'Waiting for you to stop typing...'
            if (this.question.split('')[this.question.length-1]==='?'){
                this.tip = ''
            }
        }
    },
    methods: {
        norightclick:function (e) {
            alert('精致')
            e.preventDefault()
        }
    }

})

