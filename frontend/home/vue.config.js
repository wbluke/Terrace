/* vue.config.js */
const path = require('path');
const PrerenderSpaPlugin = require('prerender-spa-plugin');

const productionPlugins = [
  new PrerenderSpaPlugin({
    staticDir: path.join(__dirname, 'dist'),
    // __INSERTION_POSITION__ // DONT CHANGE!!
    routes: ["/","/blog/","/daily/","/blog/2019/01/29/8-5.JLIS/","/blog/2019/01/29/8-4.LIS/","/blog/2019/01/12/8-3.TRIANGLEPATH/","/blog/2019/01/12/8-2.WILDCARD/","/blog/2019/01/07/8-1.JUMPGAME/","/blog/2019/01/07/8-0.dynamic_programming/","/blog/2019/01/06/7-4.FANMEETING/","/blog/2019/01/06/7-3.FENCE/","/blog/2019/01/05/7-2.QUADTREE/","/blog/2019/01/05/7-1.KARATSUBA/","/blog/2019/01/05/7-0.Divide_Conquer/","/blog/2019/01/03/6-3.CLOCKSYNC/","/blog/2019/01/03/6-2.BOARDCOVER/","/blog/2018/12/23/6-1.PICNIC/","/blog/2018/12/18/Algorithm-Problem-Solving-Strategies/","/blog/2018/12/17/블로그-오픈/"],
    // __INSERTION_POSITION_END__ // DONT CHANGE!!
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
