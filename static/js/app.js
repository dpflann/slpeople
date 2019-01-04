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
                     this.people = response.body.people;
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

var possibleDuplicates = new Vue({
      el: '#possible-duplicates',
      data: {
          possibleDuplicates: []
     },
     methods: {
         listPossibleDuplicates(resource) {
             this.$http.get('/people/duplicates').then(response => {
                     this.possibleDuplicates = response.body.possibleDuplicates
           }, response => {
           });
         }
     }
});
