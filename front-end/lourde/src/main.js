import "@mdi/font/css/materialdesignicons.css";
import "./registerServiceWorker";

import App from "./App.vue";
import Vue from "vue";
import VueCookies from "vue-cookies";
import VueGtag from "vue-gtag";
import axios from "./plugins/axios";
import router from "./router";
import vuetify from "./plugins/vuetify";

Vue.config.productionTip = false;
Vue.use(VueCookies);
Vue.use(
  VueGtag,
  {
    config: {
      id: "G-5NN3YLTQR2"
    }
  },
  router
);

new Vue({
  router,
  vuetify,
  axios,
  render: h => h(App)
}).$mount("#app");
