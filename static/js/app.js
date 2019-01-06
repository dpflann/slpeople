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
             this.$http.get('/people/emails/char-frequencies').then(response => {
                     this.frequencies = response.body.frequencies;
                     var data = this.frequencies.map(function (kv) { return kv.value });
                     var x = d3.scale.linear()
                        .domain([0, d3.max(data)])
                        .range([0, d3.max(data) - 50]);

                     d3.select(".chart-data")
                      .selectAll("div")
                        .data(this.frequencies)
                      .enter().append("div")
                        .style("width", function(fr) { return x(fr.value) + 45 + "px"; })
                        .style("height", "24px")
                        .text(function(fr) { return fr.key + " = "; })
                        .append("span").text(function(fr) { return fr.value; });
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
             this.$http.get('/people/emails/duplicates').then(response => {
                     this.possibleDuplicates = response.body.possibleDuplicates
           }, response => {
           });
         }
     }
});
