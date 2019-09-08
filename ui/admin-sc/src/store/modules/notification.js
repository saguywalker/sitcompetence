import {
	UPDATE_NOTIFICATION
} from "../mutationTypes";

const state = {
	notifications: []
};

const mutations = {
	[UPDATE_NOTIFICATION](stateData, data) {
		stateData.notifications = [
			...data
		];
	}
};

const actions = {
	add({ commit, state: stateData }, data) {
		const payload = [
			...stateData.notifications,
			{
				...data
			}
		];

		commit(UPDATE_NOTIFICATION, payload);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};