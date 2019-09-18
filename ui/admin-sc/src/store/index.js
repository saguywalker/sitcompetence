import Vue from "vue";
import Vuex from "vuex";
import giveBadge from "./modules/giveBadge";
import notification from "./modules/notification";
import verify from "./modules/verify";
import createActivity from "./modules/createActivity";
import base from "./modules/base";

Vue.use(Vuex);

export default new Vuex.Store({
	modules: {
		base,
		giveBadge,
		notification,
		verify,
		createActivity
	}
});