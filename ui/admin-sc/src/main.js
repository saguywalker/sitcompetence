import Vue from "vue";
import App from "@/App.vue";
import router from "@/router";
import store from "@/store";
import BootstrapVue from "bootstrap-vue";
import BasePageStep from "@/components/BasePageStep.vue";
import BaseImage from "@/components/BaseImage.vue";
import "@/styles/global.scss";


Vue.config.productionTip = false;

Vue.use(BootstrapVue);

Vue.component("base-page-step", BasePageStep);
Vue.component("base-image", BaseImage);

new Vue({
	router,
	store,
	render: (h) => h(App)
}).$mount("#app");
