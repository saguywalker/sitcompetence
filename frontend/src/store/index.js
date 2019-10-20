import Vue from "vue";
import Vuex from "vuex";
import base from "./modules/base";
import authentication from "./modules/authentication";
import giveBadge from "./modules/admin/giveBadge";
import adminVerify from "./modules/admin/verify";
import createActivity from "./modules/admin/createActivity";
import editActivity from "./modules/admin/editActivity";
import adminActivity from "./modules/admin/activity";
import activity from "./modules/student/activity";
import portfolio from "./modules/student/portfolio";

Vue.use(Vuex);

export default new Vuex.Store({
	modules: {
		base,
		giveBadge,
		adminActivity,
		createActivity,
		editActivity,
		adminVerify,
		authentication,
		activity,
		portfolio
	}
});