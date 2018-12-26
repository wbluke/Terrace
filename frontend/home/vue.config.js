/* vue.config.js */
const path = require('path');
const PrerenderSpaPlugin = require('prerender-spa-plugin');

const productionPlugins = [
  new PrerenderSpaPlugin({
    staticDir: path.join(__dirname, 'dist'),
    routes: [
        "/",
        "/blog",
        "/daily",
        "/blog/2018/12/17/블로그-오픈",
        "/blog/2018/12/18/Algorithm-Problem-Solving-Strategies",
        "/blog/2018/12/23/6-1.PICNIC/"
    ],
    renderer: new PrerenderSpaPlugin.PuppeteerRenderer({
      renderAfterElementExists: '#app'
    }),
  }),
];

module.exports = {
  lintOnSave: false,
  configureWebpack: (config) => {
    if (process.env.NODE_ENV === 'production') {
      config.plugins.push(...productionPlugins);
    }
  },
};
