<template>
  <div class="about">
    <img :src="aboutMyImageLink" id="aboutImg">
    <h1 id="aboutMyName">{{ aboutMyName }}</h1>
    <div class="personInfo">
      <div class="personInfoRecord">
        <img :src="aboutMyHobbyLink" class="personInfoIcon">
        <div class="personInfoText">{{ aboutMyHobby }}</div>
      </div>
      <div class="personInfoRecord">
        <img :src="aboutMyJobLink" class="personInfoIcon">
        <div class="personInfoText">{{ aboutMyJob }}</div>
      </div>
      <div class="personInfoRecord">
        <img :src="aboutMyMajorLink" class="personInfoIcon">
        <div class="personInfoText">{{ aboutMyMajor }}</div>
      </div>
      <div class="personInfoRecord">
        <img :src="aboutMyGithubLink" class="personInfoIcon">
        <div class="personInfoText">
          <a href="https://github.com/wbluke">github.com/wbluke</a>
        </div>
      </div>
      <div class="personInfoRecord">
        <img :src="aboutMyEmailLink" class="personInfoIcon">
        <div class="personInfoText">
          <a href="mailto:dev@wbluke.com">dev@wbluke.com</a>
        </div>
      </div>
      <div class="personInfoRecord">
        <img :src="aboutMyLocationLink" class="personInfoIcon">
        <div class="personInfoText">{{ aboutMyLocation }}</div>
      </div>

    </div>
  </div>
</template>

<script>
//import imageLink from "../public/about.jpeg"

export default {
  name: 'About',
  data() {
    return {
      aboutMyImageLink: require('../assets/about.jpeg'),
      aboutMyHobbyLink: require('../assets/pianoIcon.png'),
      aboutMyJobLink: require('../assets/devIcon.png'),
      aboutMyMajorLink: require('../assets/studyIcon.png'),
      aboutMyGithubLink: require('../assets/gitIcon.png'),
      aboutMyEmailLink: require('../assets/mailIcon.png'),
      aboutMyLocationLink: require('../assets/locateIcon.png'),
      aboutMyName: 'Park WooBin',
      aboutMyHobby: '피아노 치는 개발자',
      aboutMyJob: 'Software Engineer',
      aboutMyMajor: 'Mathematics, Hanyang University',
      aboutMyLocation: 'Seoul, Korea'
    }
  }

}
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css?family=Indie+Flower');
@import url('https://fonts.googleapis.com/css?family=Gothic+A1');

.about {
  margin-top: 20px;
}
#aboutImg {
   width: 200px;
   height: 200px;
   border-radius: 8px;
   margin: 0 10px 3px 10px;
   box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}
#aboutMyName {
  font-family: 'Indie Flower', cursive;
  font-size: 32px;
  padding: 10px;
  margin: 5px;
  font-weight: normal;
}
.personInfo {
  font-family: 'Gothic A1', sans-serif;
  font-size: 13px;
  text-align: left;
  margin: auto;
  width: 260px;
  margin-top: 15px;
}
.personInfoRecord {
  margin-bottom: 1px;
}
.personInfoIcon {
  width: 16px;
  height: 16px;
  display: inline;
  text-align: center;
  vertical-align: top;
  margin-left: 20px;
  -webkit-box-shadow: none;  
  -moz-box-shadow:    none;
  box-shadow:         none;
  border-radius: 0px;
}
.personInfoText {
  display: inline;
  margin-left: 7px;
  vertical-align: top;
}
</style>
