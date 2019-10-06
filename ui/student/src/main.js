import Vue from "vue";
import BootstrapVue from "bootstrap-vue";
import VueProgressBar from "vue-progressbar";
import App from "@/App.vue";
import router from "@/router";
import store from "@/store";
import "@/styles/global.scss";
import BaseImage from "@/components/BaseImage.vue";
import BaseInputText from "@/components/BaseInputText.vue";

Vue.config.productionTip = false;
Vue.use(BootstrapVue);
Vue.use(VueProgressBar, {
	color: "#396290",
	failedColor: "#eb2e2e",
	height: "2px"
});

Vue.component("base-input-text", BaseInputText);
Vue.component("base-image", BaseImage);

export default new Vue({
	router,
	store,
	render: (h) => h(App)
}).$mount("#app");