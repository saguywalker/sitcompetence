/* eslint-disable no-console */
import Vue from "vue";
import App from "@/App.vue";
import router from "@/router";
import store from "@/store";
import BootstrapVue from "bootstrap-vue";
import BasePageStep from "@/components/BasePageStep.vue";
import BaseImage from "@/components/BaseImage.vue";
import "@/styles/global.scss";
// const prod = process.env.NODE_ENV === "production";
// const shouldSWDev = 'serviceWorker' in navigator && !prod
// const shouldSW = "serviceWorker" in navigator && prod;
// if (shouldSW) {
//   navigator.serviceWorker.register("/service-worker.js").then(() => {
//     console.log("Service Worker Registered!")
//   });
// } else if (shouldSWDev) {
//   navigator.serviceWorker.register("/service-workder-dev.js").then(() => {
//     console.log("Service Worker Registered!");
//   });
// }


Vue.config.productionTip = false;

Vue.use(BootstrapVue);

Vue.component("base-page-step", BasePageStep);
Vue.component("base-image", BaseImage);

new Vue({
	router,
	store,
	render: (h) => h(App)
}).$mount("#app");
