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

var charFrequences = new Vue({
      el: '#character-frequencies',
      data: {
          frequencies: []
     },
     methods: {
         listCharacterFrequencies(resource) {
             this.$http.get('/people/char_frequencies').then(response => {
                     this.frequencies = response.body.frequencies
           }, response => {
           });
         }
     }
});
