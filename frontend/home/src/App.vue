<template>
  <div id="app">
    <HelloWorld class="page_title" msg="Luke's Notes"/>
    <div class="menu effect">
      <ul>
        <li><router-link to="/">About</router-link></li>
        <li><router-link to="/blog">Blog</router-link></li>
        <li><router-link to="/daily">Daily</router-link></li>
      </ul>
    </div>
    <div>
      <img id="aboutImg" src="../public/about.jpeg">
    </div>
    <p>
      <router-view></router-view>
    </p>

  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'app',
  components: {
    HelloWorld
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

#aboutImg {
   width: 200px;
   height: 200px;
   border-radius: 8px;
   margin: 10px;
}


/* BASE TEMPLATE STYLES ========== */
@import url(https://fonts.googleapis.com/css?family=Raleway:300,700);
body {
	background: #fafafa;
	background-attachment: fixed;
	color: #555;
	font-family: 'Raleway', Helvetica, Arial, sans-serif;
	font-size: 15px;
	font-weight: 300;
	line-height: 140%;
	margin: 0;
	height: 100%;
	padding: 0;
	text-align: center;
	width: 100%;
}
.page_title {
	color: #074d11;
	font-weight: 200;
	margin: 0;
	padding: 5px 0;
	text-shadow: 0 1px 1px rgba(0,0,0,.1);
	text-transform: uppercase;
}

/* PEN STYLES ========== */
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

/* HOVER STYLE PER CLASS ========== */

/* ----- EFFECT 03 - Side Brackets */
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
	opacity: 0;
	position: absolute;
	right: 30px;
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


</style>
