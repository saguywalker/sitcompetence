import Vue from "vue";
import App from "@/App.vue";
import router from "@/router";
import store from "@/store";
import BootstrapVue from "bootstrap-vue";
import VueProgressBar from "vue-progressbar";
import BaseImage from "@/components/BaseImage.vue";
import BaseSitcomLogo from "@/components/BaseSitcomLogo.vue";
import BaseInputText from "@/components/BaseInputText.vue";
import BasePageStep from "@/components/admin/BasePageStep.vue";
import "./registerServiceWorker";
import "@/styles/global.scss";

if (process.env.NODE_ENV !== "production") {
  require("dotenv").config();
}

Vue.config.productionTip = false;

Vue.use(BootstrapVue);
Vue.use(VueProgressBar, {
	color: "#396290",
	failedColor: "#eb2e2e",
	height: "2px"
});

Vue.component("base-input-text", BaseInputText);
Vue.component("base-page-step", BasePageStep);
Vue.component("base-image", BaseImage);
Vue.component("base-sitcom-logo", BaseSitcomLogo);

new Vue({
	router,
	store,
	render: (h) => h(App)
}).$mount("#app");
