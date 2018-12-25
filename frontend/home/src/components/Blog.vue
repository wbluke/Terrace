<template>
  <div id="blog">
    <div id="blog-contents" v-if="year !== undefined">
      <article>
        <div class="inner-title-container">
          <h1><a :href="address" id="inner-title">{{ title }}</a></h1>
          <p class="meta">{{ year }} / {{ month }} / {{ day }}</p>
        </div>
        <div id="article-content" v-html="articleHtmlSource"></div>
      </article>
      <hr>
    </div>
    <div v-else>
      <div class="blog-list-element" v-for="i in index" :key="i">
        <div class="blog-list-element-title">
          <router-link :to="i.uri" class="nice-link">{{ i.title }}</router-link>
        </div>
        {{ i.date }}
      </div>
    </div>
  </div>
</template>

<script>
import HighlightJS from 'highlight.js/lib/highlight.js'
//var HighlightJS = require("highlight.js/lib/highlight.js");
// Add languages manually to decrease size of my website
HighlightJS.registerLanguage('vim', require('highlight.js/lib/languages/vim'));
HighlightJS.registerLanguage('bash', require('highlight.js/lib/languages/bash'));
// eslint-disable-next-line
import _ from 'highlight.js/styles/xcode.css'

export default {
  name: 'blogExample',
  methods: {
    getPage: function() {
      if (this.year === undefined) {
        return;
      }

      var htmlDocUri =
        '/blog_contents/'
        + this.year + '/'
        + this.month + '/'
        + this.day + '/'
        + this.title + '.html';

      fetch(htmlDocUri)
        .then(response => response.text())
        .then(responseText => this.articleHtmlSource = responseText);

      var uri = htmlDocUri.replace('blog_contents/', '')
                     .replace('.html', '');
      this.address = this.domain + uri;
    },
  },
  mounted: function() {
    this.getPage();
  },
  updated: function() {
    // Change file name to document's title
    // When the page is an artice page, get blogContents
    var blogContents = document.querySelector('#blog-contents');
    if(blogContents == null){
      //this.toTheTop();
      return;
    }
    // Now, below scripts are for displaying an article,
    // Not the blog main page which has lists of name of articles.
    // Find all h1 tags, and choose second h1. It is real title of this doc.
    var titles = blogContents.querySelectorAll('h1');
    if(titles.length <= 1) {
      return;
    } else {
      // Title exists
      var title = blogContents.querySelector('#inner-title')
      title.innerHTML = titles[1].innerHTML;
      titles[1].style.display = "none";
      // It is for og:title
      this.titleForMeta = title.innerHTML;
    }

    // Find all code block and apply syntax highlighting
    [].forEach.call(document.querySelectorAll('code'), function(el) {
      HighlightJS.highlightBlock(el);
    });
  },
  data() {
    return {
      // __INSERTION_POSITION__ // DONT CHANGE!!
      index:
[{"title":"Testfile","uri":"/blog/2018/12/24/testpandoc/","date":"2018/12/24"},{"title":"[알고리즘] 6-1. PICNIC (소풍)","uri":"/blog/2018/12/23/6-1.PICNIC/","date":"2018/12/23"},{"title":"[알고리즘] Algorithm Problem Solving Strategies","uri":"/blog/2018/12/18/Algorithm-Problem-Solving-Strategies/","date":"2018/12/18"},{"title":"[블로그] 블로그 오픈","uri":"/blog/2018/12/17/블로그-오픈/","date":"2018/12/17"}]
// __INSERTION_POSITION_END__ // DONT CHANGE!!
,
    year : this.$route.params.year,
    month : this.$route.params.month,
    day : this.$route.params.day,
    title : this.$route.params.title,
    articleHtmlSource : "",
    address : "",  // will have a permalink of the article
    domain : "http://wbluke.com/blog",
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
@import url('https://fonts.googleapis.com/css?family=Ubuntu+Mono');
code {
  font-family: 'Ubuntu Mono', monospace;
}
div#blog {
  min-height: 300px;
  max-width: 800px;
  margin: auto;
  margin-bottom: 50px;
}
#blog-contents {
  text-align: left;
}
.inner-title-container {
  text-align: center;
  text-decoration: none;
}
.inner-title-container > h1{
  margin-bottom: 10px;
}
/* .inner-title-container > h1 > a{
  color: #04652e;
}
.inner-title-container > h1 > a:hover{
  color: #2f945a;
} */
.meta {
  margin: 0;
}
.blog-list-element {
  margin-top: 20px;
  margin-bottom: 15px;
  text-decoration: none;
}
.blog-list-element-title {
  font-size: 16px;
}

article {
  font-size: 1.08em;
}
#padding-between-title-and-article {
  height: 30px;
}
#article-content {
  text-align: left;
  color: #000;
  line-height: 1.6em;
}
.blog-article-list {
  max-width: 800px;
  margin:auto;
  padding-bottom:1px;
}
</style>
