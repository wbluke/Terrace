<template>
  <div id="insight">
    <div id="insight-contents" v-if="year !== undefined">
      <article>
        <div class="inner-title-container">
          <h1><a :href="address" id="insight-inner-title">{{ title }}</a></h1>
          <p class="meta">{{ year }} / {{ month }} / {{ day }}</p>
        </div>
        <div id="insight-article-content" v-html="articleHtmlSource"></div>
      </article>
      <hr>
    </div>
    <div v-else>
      <div class="insight-list-element" v-for="i in index" :key="i">
        <a v-bind:href="i.bookURL">
          <img class="insight-list-element-label" v-bind:src="'http://cdn.wbluke.com/label_'+ i.label + '.png'"/>
        </a>
        <span class="insight-list-element-title">
          <router-link :to="i.uri" class="nice-link" :class="$mq">{{ i.title }}</router-link>
        </span>
        <br>
        <span class="insight-list-element-date">
          {{ i.date }}
        </span>
        <hr class="seperate-line">
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Insight',
  metaInfo () {
    if(this.title === undefined) {
      return {
        title: 'Insight',
        meta: [
          { charset: 'utf-8' },
          {
            'property': 'og:title',
            'content': 'Insight',
            'template': chunk => `${chunk} | Luke\'s Note`,
            'vmid': 'og:title'
          },
          {
            'property': 'og:description',
            'content': 'Article List',
            'template': chunk => `${chunk}`,
            'vmid': 'og:description'
          },
        ]
      }
    } else {
      return {
        title: this.titleForMeta,
        meta: [
          { charset: 'utf-8' },
          {
            'property': 'og:title',
            'content': this.title,
            'template': chunk => `${chunk} | Luke\'s Note`,
            'vmid': 'og:title'
          },
          {
            'property': 'og:description',
            'content': 'Thank you for reading my article 🦊',
            'template': chunk => `${chunk}`,
            'vmid': 'og:description'
          }
        ]
      }
    }
  },
  methods: {
    getPage: function() {
      if (this.year === undefined) {
        return;
      }

      var htmlDocUri =
        '/insight_contents/'
        + this.year + '/'
        + this.month + '/'
        + this.day + '/'
        + this.title + '.html';

      fetch(htmlDocUri)
        .then(response => response.text())
        .then(responseText => this.articleHtmlSource = responseText);

      var uri = htmlDocUri.replace('insight_contents/', '')
                     .replace('.html', '/');
      this.address = this.domain + uri;
    },
  },
  mounted: function() {
    this.getPage();
  },
  updated: function() {
    
    // Change file name to document's title
    // When the page is an artice page, get insightContents
    var insightContents = document.querySelector('#insight-contents');
    if(insightContents == null){
      //this.toTheTop();
      return;
    }
    // Now, below scripts are for displaying an article,
    // Not the insight main page which has lists of name of articles.
    // Find all h1 tags, and choose second h1. It is real title of this doc.
    var titles = insightContents.querySelectorAll('h1');
    if(titles.length <= 1) {
      return;
    } else {
      // Title exists
      var title = insightContents.querySelector('#insight-inner-title')
      title.innerHTML = titles[1].innerHTML;
      titles[1].style.display = "none";
      // It is for og:title
      this.titleForMeta = title.innerHTML;
    }
  },
  data() {
    return {
      // __INSERTION_POSITION__ // DONT CHANGE!!
      index:
[{"title":"[브랜딩/경영] 배민다움","uri":"/insight/2019/02/07/배민다움/","date":"2019.02.07","label":"BaeMinDaum","bookURL":"https://book.naver.com/bookdb/book_detail.nhn?bid=11211590"}]
// __INSERTION_POSITION_END__ // DONT CHANGE!!
,
    year : this.$route.params.year,
    month : this.$route.params.month,
    day : this.$route.params.day,
    title : this.$route.params.title,
    articleHtmlSource : "",
    address : "",  // will have a permalink of the article
    domain : "http://wbluke.com/insight",
    }
  },
    watch: {
    '$route' (to) {
      this.year = to.params.year;
      this.month = to.params.month;
      this.day = to.params.day;
      this.title = to.params.title;
      this.getPage();

    }
  }
}
</script>

<style scoped>

div#insight {
  min-height: 300px;
  max-width: 800px;
  margin: auto;
  margin-top: 25px;
  margin-bottom: 50px;
}
#insight-contents {
  margin: auto;
  text-align: left;
}
.inner-title-container {
  text-align: center;
  text-decoration: none;
}
.inner-title-container > h1{
  margin-bottom: 10px;
}
.meta {
  margin: 0;
}

article {
  font-size: 1.08em;
}
#insight-article-content {
  text-align: left;
  color: #000;
  line-height: 1.6em;
  padding: 8px;
}
.insight-article-list {
  max-width: 800px;
  margin:auto;
  padding-bottom:1px;
}

.insight-list-element {
  /* margin-top: 25px; */
  text-decoration: none;
  max-width: 380px;
  margin: auto;
  padding: 12px;
}
.insight-list-element-label {
  display: inline;
  text-align: left;
  vertical-align: middle;
  -webkit-box-shadow: none;
  box-shadow: none;
  width: 64px;
  height: 40px;
  text-align: left;
  /* margin-right: 50px; */
}
.insight-list-element-title {
  display: inline;
  font-size: 16px;
  margin-left: 154px;
  max-width: 250px;
}
.insight-list-element-date {
  margin-left: 290px;
}
</style>
