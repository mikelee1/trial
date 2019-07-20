const a = Vue.component('todo-item', {
    props: ['todo'],
    template: '<li>{{todo.title}}</li>'
})
const b = Vue.component('solved-item', {
    props: ['solved'],
    template: '<li>{{solved.title}}</li>'
})


var app = new Vue({
    el: '#app',
    data: {
        todos: [
            {key: 1, title: 'todo1'},
            {key: 2, title: 'todo2'}
        ],
        solveds: [
            {key:1 ,title: 'solved1'}
        ]
    },
    components: [a, b],
})

