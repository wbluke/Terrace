<template>
  <div id="daily">Coming Soon...!</div>
</template>

<script>
export default {
  name: 'Daily',
  metaInfo: {
    title: 'Daily',
    titleTemplate: '%s | Luke\'s Note',
    meta: [
      { charset: 'utf-8' },
      {
        'property': 'og:title',
        'content': 'Daily',
        'template': chunk => `${chunk} | Luke\'s Note`,
        'vmid': 'og:title'
      },
      {
        'property': 'og:description',
        'content': 'Daily Story',
        'template': chunk => `${chunk}`,
        'vmid': 'og:description'
      }
    ]
  },
}
</script>

<style scoped>
@import url(https://fonts.googleapis.com/css?family=Raleway:300,700);
#daily {
  font-family: 'Raleway', Helvetica, Arial, sans-serif;
  min-height: 300px;
  max-width: 800px;
  margin: auto;
  margin-bottom: 50px;
}
</style>
