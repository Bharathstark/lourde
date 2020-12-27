import Vue from "vue";
import App from "./App.vue";
import axios from "./plugins/axios";
import "./registerServiceWorker";
import router from "./router";
import vuetify from "./plugins/vuetify";
import VueGtag from "vue-gtag";
import VueCookies from "vue-cookies";

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
