<template>
  <div class="About">
    <img :src="imageLink" id="aboutImg">
    <div>{{ aboutTxt }}</div>
  </div>
</template>

<script>
//import imageLink from "../public/about.jpeg"

export default {
  name: 'About',
  data() {
    return {
      imageLink: require('../assets/about.jpeg'),
      aboutTxt: '피아노 치는 개발자'
    }
  }

}
</script>

<style scoped>
img#aboutImg {
   width: 200px;
   height: 200px;
   border-radius: 8px;
   margin: 0 10px 10px 10px;
   box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}
</style>
