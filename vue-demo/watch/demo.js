var obj = {
    question: '',
    tip: ''
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
    created: function () {
        // `_.debounce` 是一个通过 Lodash 限制操作频率的函数。
        // 在这个例子中，我们希望限制访问 yesno.wtf/api 的频率
        // AJAX 请求直到用户输入完毕才会发出。想要了解更多关于
        // `_.debounce` 函数 (及其近亲 `_.throttle`) 的知识，
        // 请参考：https://lodash.com/docs#debounce
        this.debouncedGetAnswer = _.debounce(this.getAnswer, 500)
    },
})

