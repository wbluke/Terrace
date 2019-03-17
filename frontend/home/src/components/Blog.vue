<template>
  <div id="blog">
    <div id="blog-contents" v-if="year !== undefined">
      <article>
        <div class="inner-title-container">
          <h1><a :href="address" id="inner-title">{{ title }}</a></h1>
          <p class="meta">{{ year }}. {{ month }}. {{ day }}. </p>
        </div>
        <div id="article-content" v-html="articleHtmlSource"></div>
      </article>
      <hr>
      <div id="disqus_thread"></div>
    </div>
    <div v-else>
      <div class="blog-list-element" v-for="i in index" :key="i">
        <div class="blog-list-element-title">
          <router-link :to="i.uri" class="nice-link" :class="$mq">{{ i.title }}</router-link>
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
HighlightJS.registerLanguage('bash', require('highlight.js/lib/languages/bash'));
HighlightJS.registerLanguage('cpp', require('highlight.js/lib/languages/cpp'));
HighlightJS.registerLanguage('markdown', require('highlight.js/lib/languages/markdown'));
HighlightJS.registerLanguage('java', require('highlight.js/lib/languages/java'));
HighlightJS.registerLanguage('javascript', require('highlight.js/lib/languages/javascript'));
HighlightJS.registerLanguage('css', require('highlight.js/lib/languages/css'));
//HighlightJS.registerLanguage('html', require('highlight.js/lib/languages/html'));
// eslint-disable-next-line
import _ from 'highlight.js/styles/xcode.css'

export default {
  name: 'Blog',
  metaInfo () {
    if(this.title === undefined) {
      return {
        title: 'Blog',
        meta: [
          { charset: 'utf-8' },
          {
            'property': 'og:title',
            'content': 'Blog',
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
    initDisqus: function(shortname, identifier, title, url) {
      //config
      if (typeof DISQUS === "undefined") {
        (async () => {
          var vars_text =
            'var disqus_shortname  = "' +
            shortname +
            '";\n' +
            'var disqus_title      = "' +
            title +
            '";\n' +
            'var disqus_identifier = "' +
            identifier +
            '";\n' +
            'var disqus_url        = "' +
            url +
            '";\n';
          var vars_obj = document.createElement("script");
          vars_obj.type = "text/javascript";
          vars_obj.async = true;
          vars_obj.text = vars_text;
          (
            document.getElementsByTagName("head")[0] ||
            document.getElementsByTagName("body")[0]
          ).appendChild(vars_obj);
          var dsq = document.createElement("script");
          dsq.type = "text/javascript";
          dsq.async = true;
          dsq.src = "//" + shortname + ".disqus.com/embed.js";
          (
            document.getElementsByTagName("head")[0] ||
            document.getElementsByTagName("body")[0]
          ).appendChild(dsq);
        })();
      } else {
        // eslint-disable-next-line
        DISQUS.reset({
          reload: true,
          config: function() {
            this.page.identifier = identifier;
            this.page.url = url;
            this.page.title = title;
          }
        });
      }
    },
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
                     .replace('.html', '/');
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

    this.initDisqus(
      "wbluke-com",
      this.address.replace(this.domain, ""), // uri as an identifier
      this.title,
      this.address
    );

    // Find all code block and apply syntax highlighting
    [].forEach.call(document.querySelectorAll('code'), function(el) {
      HighlightJS.highlightBlock(el);
    });
  },
  data() {
    return {
      // __INSERTION_POSITION__ // DONT CHANGE!!
      index:
[{"title":"[알고리즘] 21-1. TRAVERSAL (트리 순회 순서 변경)","uri":"/blog/2019/03/17/21-1.TRAVERSAL/","date":"2019.03.17."},{"title":"[알고리즘] 19-2. ITES (외계 신호 분석)","uri":"/blog/2019/03/05/19-2.ITES/","date":"2019.03.05."},{"title":"[알고리즘] 19-1. BRACKETS2 (짝이 맞지 않는 괄호)","uri":"/blog/2019/03/03/19-1.BRACKETS2/","date":"2019.03.03."},{"title":"[알고리즘] 18-1. JOSEPHUS (조세푸스 문제)","uri":"/blog/2019/03/03/18-1.JOSEPHUS/","date":"2019.03.03."},{"title":"[알고리즘] 8-13. NUMB3RS (두니발 박사의 탈옥)","uri":"/blog/2019/02/18/8-13.NUMB3RS/","date":"2019.02.18."},{"title":"[알고리즘] 8-12. POLY (폴리오미노)","uri":"/blog/2019/02/18/8-12.POLY/","date":"2019.02.18."},{"title":"[Java] 컬렉션 프레임워크 (Collection Framework)","uri":"/blog/2019/02/16/자바_컬렉션_프레임워크/","date":"2019.02.16."},{"title":"[알고리즘] 8-11. ASYMTILING (비대칭 타일링)","uri":"/blog/2019/02/13/8-11.ASYMTILING/","date":"2019.02.13."},{"title":"[알고리즘] 8-10. SNAIL (달팽이)","uri":"/blog/2019/02/13/8-10.SNAIL/","date":"2019.02.13."},{"title":"[알고리즘] 8-9. TRIPATHCNT (삼각형 위의 최대 경로 수 세기)","uri":"/blog/2019/02/13/8-09.TRIPATHCNT/","date":"2019.02.13."},{"title":"[알고리즘] 8-8. TILING2 (타일링)","uri":"/blog/2019/02/12/8-8.TILING2/","date":"2019.02.12."},{"title":"[알고리즘] 8-7. QUANTIZE (양자화)","uri":"/blog/2019/02/12/8-7.QUANTIZE/","date":"2019.02.12."},{"title":"[Java] 제네릭2_와일드카드(Generic)","uri":"/blog/2019/02/11/자바-제네릭2/","date":"2019.02.11."},{"title":"[Java] 제네릭1_기본 문법(Generic)","uri":"/blog/2019/02/11/자바-제네릭1/","date":"2019.02.11."},{"title":"[Vue.js] Vue.js에서 미디어쿼리 적용하기","uri":"/blog/2019/02/09/vuejs-미디어쿼리/","date":"2019.02.09."},{"title":"[알고리즘] 알고리즘 풀이 언어 변경(Cpp to Java)","uri":"/blog/2019/02/02/알고리즘_언어_변경/","date":"2019.02.02."},{"title":"[알고리즘] 8-6. PI (원주율 외우기)","uri":"/blog/2019/02/01/8-6.PI/","date":"2019.02.01."},{"title":"[알고리즘] 8-5. JLIS (합친 LIS)","uri":"/blog/2019/01/29/8-5.JLIS/","date":"2019.01.29."},{"title":"[알고리즘] 8-4. LIS (Longest Increasing Sequence)","uri":"/blog/2019/01/29/8-4.LIS/","date":"2019.01.29."},{"title":"[알고리즘] 8-3. TRIANGLEPATH (삼각형 위의 최대 경로)","uri":"/blog/2019/01/12/8-3.TRIANGLEPATH/","date":"2019.01.12."},{"title":"[알고리즘] 8-2. WILDCARD (와일드카드)","uri":"/blog/2019/01/12/8-2.WILDCARD/","date":"2019.01.12."},{"title":"[알고리즘] 8-1. JUMPGAME (외발 뛰기)","uri":"/blog/2019/01/07/8-1.JUMPGAME/","date":"2019.01.07."},{"title":"[알고리즘] 8-0. 동적 계획법 (Dynamic Programming)","uri":"/blog/2019/01/07/8-0.dynamic_programming/","date":"2019.01.07."},{"title":"[알고리즘] 7-4. FANMEETING (팬미팅)","uri":"/blog/2019/01/06/7-4.FANMEETING/","date":"2019.01.06."},{"title":"[알고리즘] 7-3. FENCE (울타리 잘라내기)","uri":"/blog/2019/01/06/7-3.FENCE/","date":"2019.01.06."},{"title":"[알고리즘] 7-2. QUADTREE (쿼드 트리 뒤집기)","uri":"/blog/2019/01/05/7-2.QUADTREE/","date":"2019.01.05."},{"title":"[알고리즘] 7-1. KARATSUBA (카라츠바 알고리즘)","uri":"/blog/2019/01/05/7-1.KARATSUBA/","date":"2019.01.05."},{"title":"[알고리즘] 7-0. 분할 정복 (Divide and Conquer)","uri":"/blog/2019/01/05/7-0.Divide_Conquer/","date":"2019.01.05."},{"title":"[알고리즘] 6-3. CLOCKSYNC (시계 맞추기)","uri":"/blog/2019/01/03/6-3.CLOCKSYNC/","date":"2019.01.03."},{"title":"[알고리즘] 6-2. BOARDCOVER (게임판 덮기)","uri":"/blog/2019/01/03/6-2.BOARDCOVER/","date":"2019.01.03."},{"title":"[알고리즘] 6-1. PICNIC (소풍)","uri":"/blog/2018/12/23/6-1.PICNIC/","date":"2018.12.23."},{"title":"[알고리즘] Algorithm Problem Solving Strategies","uri":"/blog/2018/12/18/Algorithm-Problem-Solving-Strategies/","date":"2018.12.18."},{"title":"[블로그] 블로그 오픈","uri":"/blog/2018/12/17/블로그-오픈/","date":"2018.12.17."}]
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

div#blog {
  min-height: 300px;
  max-width: 800px;
  margin: auto;
  margin-top: 25px;
  margin-bottom: 50px;
}
#blog-contents {
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
.blog-list-element {
  /* margin-top: 25px; */
  text-decoration: none;
  padding: 12px;
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
  line-height: 1.8em;
  padding: 8px;
}
.blog-article-list {
  max-width: 800px;
  margin:auto;
  padding-bottom:1px;
}
</style>
