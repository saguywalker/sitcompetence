import Vue from "vue";
import Vuex from "vuex";
import authentication from "./modules/authentication";
import activity from "./modules/activity";
import portfolio from "./modules/portfolio";

Vue.use(Vuex);

export default new Vuex.Store({
	modules: {
		activity,
		authentication,
		portfolio
	}
});