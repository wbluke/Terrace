<template>
  <div id="app">
    <HelloWorld msg="Hello! My name is Luke :D"/>
    <p>
      <router-link to="/">About</router-link>
      &nbsp;
      <router-link to="/study">Study</router-link>
      &nbsp;
      <router-link to="/daily">Daily</router-link>
      &nbsp;
    </p>
    
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
</style>
