import {
	UPDATE_DASHBOARD_TAB
} from "../../mutationTypes";

const state = {
	tab: 0
};

const mutations = {
	[UPDATE_DASHBOARD_TAB](stateData, data) {
		stateData.tab = data;
	}
};

const actions = {
	updateTab({ commit }, data) {
		commit(UPDATE_DASHBOARD_TAB, data);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};