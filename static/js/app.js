var people = new Vue({
      el: '#people',
      created() {
          this.listPeople();
      },
      data: {
          people: []
     },
     methods: {
         listPeople(resource) {
             this.$http.get('/people').then(response => {
                     this.people = response.body.People;
           }, response => {
           });
         }
     }
});
