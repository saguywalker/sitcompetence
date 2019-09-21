import Vue from "vue";
import BootstrapVue from "bootstrap-vue";
import App from "@/App.vue";
import router from "@/router";
import store from "@/store";
import "@/styles/global.scss";
import BaseInputText from "@/components/BaseInputText.vue";

Vue.config.productionTip = false;
Vue.use(BootstrapVue);
Vue.component("base-input-text", BaseInputText);

new Vue({
	router,
	store,
	render: (h) => h(App)
}).$mount("#app");
