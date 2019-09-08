import Vue from "vue";
import Vuex from "vuex";
import events from "./modules/events";
import giveBadge from "./modules/giveBadge";
import notification from "./modules/notification";

Vue.use(Vuex);

export default new Vuex.Store({
	modules: {
		events,
		giveBadge,
		notification
	}
});