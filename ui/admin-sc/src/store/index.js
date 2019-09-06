import Vue from "vue";
import Vuex from "vuex";
import events from "./modules/events";
import giveBadge from "./modules/giveBadge";

Vue.use(Vuex);

export default new Vuex.Store({
	modules: {
		events,
		giveBadge
	}
});