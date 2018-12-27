<template>
  <div id="app">
    <div id="site-title">
      <router-link to="/" class="page_title">
        <h1>
          <img :src="icoLink" width="24" height="24">
          Luke's Note
        </h1>
      </router-link>
    </div>
    <div class="menu effect">
      <ul>
        <li><router-link to="/">About</router-link></li>
        <li><router-link to="/blog">Blog</router-link></li>
        <li><router-link to="/daily">Daily</router-link></li>
      </ul>
    </div>

    <p>
      <router-view></router-view>
    </p>

  </div>
</template>


<script>

export default {
  name: 'app',
  methods: {
    applyNiceLink: function() {
      var elements = document.querySelectorAll('a.nice-link');
      Array.prototype.forEach.call(elements, function(el, i){
        el.setAttribute('data-text', el.textContent)
      });
    },
    // cancelHoverCssOnMobile: function() {
    //   [].forEach.call( document.querySelectorAll('a'), function(el) {
    //     el.addEventListener('click touchend', function(e) {
    //       var el = document.querySelector(this);
    //       var link = el.setAttribute('href');
    //       window.location = link;
    //     }, false);
    //   });
    // }
      //have to change
      // $(document).ready(function() {
      //   $("a").on("click touchend", function(e) {
      //     var el = $(this);
      //     var link = el.attr("href");
      //     window.location = link;
      //   });
      // });
  },
  mounted: function() {
    this.applyNiceLink()
    // this.cancelHoverCssOnMobile()
  },
  updated: function() { this.applyNiceLink() },
  data() {
    return {
      icoLink: require('./assets/favicon.png')
    }
  },
}
</script>

<style>
/* BASE TEMPLATE STYLES ========== */
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  letter-spacing: 0.5px;
  color: #2c3e50;
  margin-top: 60px;
  font-size: 15px;
	font-weight: 300;
	line-height: 140%;
}

body {
	background: #fafafa;
	background-attachment: fixed;
	margin: 0;
	height: 100%;
	padding: 0;
	width: 100%;
}
.page_title {
	color: #074d11;
	font-weight: 200;
	margin: 0;
	padding: 5px 0;
	text-shadow: 0 1px 1px rgba(0,0,0,.1);
	text-transform: uppercase;
  text-decoration: none;
}




/* MENU PEN STYLES ========== */
.menu,
.menu .menu_title,
.menu ul li a,
.menu ul li a:after,
.menu ul li a:before {
  -webkit-transition: all 300ms ease-in-out;
  transition: all 300ms ease-in-out;
}
.menu {
  border-bottom: 1px solid rgba(0,0,0,.05);
  border-top: 1px solid #fff;
  margin: 0;
  padding: 0px 0;
  max-width: 400px;
  margin: auto auto 0px;
  border-radius: 4px;
}
.menu:hover {
  background: #fff;
  border-bottom: 1px solid rgba(0,0,0,.2);
}
.menu .menu_title {
  color: #ccc;
  font-size: 1.5em;
  font-weight: 300;
  margin: 0 0 20px;
}
.menu:hover .menu_title {
  color: #0af;
  text-shadow: 0 1px 1px rgba(0,0,0,.1);
}
.menu ul {
  display: block;
  list-style: none;
  margin: 0;
  padding: 0;
}
.menu ul li {
  display: inline-block;
}
.menu ul li a {
  color: #555;
  display: block;
  line-height: 35px;
  padding: 0 30px;
  text-decoration: none;
}
/* PEN STYLES:: EFFECT 03 - Side Brackets */
.menu.effect ul li a {
  position: relative;
}
.menu.effect ul li a:before {
  color: #57bf3d;
  content: '[';
  display: block;
  font-size: 1.5em;
  left: 30px;
  opacity: 0;
  position: absolute;
  top: 0;
}
.menu.effect ul li a:after {
  color: #57bf3d;
  content: ']';
  display: block;
  font-size: 1.5em;
  right: 30px;
  opacity: 0;
  position: absolute;
  top: 0;
}
.menu.effect ul li a:hover:before {
  left: 15px;
  opacity: 1;
}
.menu.effect ul li a:hover:after {
  opacity: 1;
  right: 15px;
}
.menu.effect ul li a:hover {
  background: trasnparent;
}




/* BASIC CODE BOX STYLES ========== */
@import url('https://fonts.googleapis.com/css?family=PT+Mono');
code.hljs{
  display: inline;
  background-color: #F0F0F0;
  box-shadow: none;
  padding: 1px;
}


code, pre>code.hljs {
  font-family: 'PT Mono', monospace;
  background-color: #f1f1f1;
  border-radius: 1px;
  -webkit-box-shadow: 3px 3px 7px 0px #ccc;  /* Safari 3-4, iOS 4.0.2 - 4.2, Android 2.3+ */
  -moz-box-shadow:    3px 3px 7px 0px #ccc;  /* Firefox 3.5 - 3.6 */
  box-shadow:         3px 3px 7px 0px #ccc;  /* Opera 10.5, IE 9, Firefox 4+, Chrome 6+, iOS 5 */
  line-height: 1.6em;
  padding: 0.7em;
  padding-left: 1.6em;
  display: block;
  overflow: auto;
  font-size: 14px;
}


/* BASIC LINK STYLES ========== */
a {
  text-decoration: none;
  color: #04652e;
}
a:hover {
  color: #2f945a;
}




/* DETAILS & SUMMARY LINK STYLES ========== */
details {
  border-radius: 3px;
  background: #EEE;
  padding-left: 5px;
  padding-right: 5px;
}
details summary {
  font-size: 17px;
  vertical-align: top;
  background: #333;
  color: #FFF;
  border-radius: 3px;
  padding: 5px 10px;
  outline: none;
  cursor: pointer;
}
details summary:hover {
  background: #04652e;
}
details[open] summary {
  background: #04652e;
  color: #FFF;
}
details[open] summary:hover {
  background: #333;
  color: #FFF;
}

/* Hide the default marker. */
details summary::-webkit-details-marker {
  display: none;
}

/* Add the custom marker in the default state. */
details summary:before {
  display: inline-block;
  width: 10px;
  height: 18px;
  content: "";
  background-repeat: no-repeat;
  background-position: 0 0;
}

/* Move the sprite image when the details box is open. */
details[open] summary:before {
  background-position: -18px 0;
}




/* NICE-LINK STYLES ========== */
a.nice-link {
  position: relative;
  color:  #04652e;
}
.blog-list-element a.nice-link:after {
  border-bottom: 1px solid #53cf44; /* Underline */
}
a.nice-link:after {
  text-align: justify;
  display: inline-block;
  content: attr(data-text);
  position: absolute;
  left: 0;
  /* top: 0; */
  white-space: nowrap;
  overflow: hidden;
  color: #53cf44;
  min-height: 100%;
  width: 0;
  max-width: 100%; /* 'cause of IE bug */
  /* background: #fafafa; */
  -moz-transition: .3s;
  -o-transition: .3s;
  -webkit-transition: .3s;
  transition: .3s;
}
@media only screen and (min-width: 768px) {
  a.nice-link:hover {
    color: #04652e; /* To override default hover color */
  }
  a.nice-link:hover:after {
    width: 100%;
  }
};

</style>
