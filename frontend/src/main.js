import Vue from "vue";
import App from "@/App.vue";
import router from "@/router";
import store from "@/store";
import BootstrapVue from "bootstrap-vue";
import VueProgressBar from "vue-progressbar";
import VueHighlightJS from "vue-highlight.js";
import javascript from "highlight.js/lib/languages/javascript";
import BaseImage from "@/components/BaseImage.vue";
import BaseSitcomLogo from "@/components/BaseSitcomLogo.vue";
import BaseInputText from "@/components/BaseInputText.vue";
import BasePageStep from "@/components/admin/BasePageStep.vue";
import acl from "@/plugin/acl";
import "@/styles/global.scss";
import "highlight.js/styles/default.css";
import "./registerServiceWorker";

Vue.config.productionTip = false;

Vue.use(BootstrapVue);
Vue.use(VueProgressBar, {
	color: "#396290",
	failedColor: "#eb2e2e",
	height: "2px"
});
Vue.use(VueHighlightJS, {
	// Register only languages that you want
	languages: {
		javascript
	}
});

Vue.component("base-input-text", BaseInputText);
Vue.component("base-page-step", BasePageStep);
Vue.component("base-image", BaseImage);
Vue.component("base-sitcom-logo", BaseSitcomLogo);

new Vue({
	router,
	store,
	acl,
	render: (h) => h(App)
}).$mount("#app");
