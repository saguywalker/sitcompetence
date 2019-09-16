import {
	UPDATE_VERIFY_HASH
} from "../mutationTypes";

const state = {
	hashId: ""
};

const mutations = {
	[UPDATE_VERIFY_HASH](stateData, data) {
		stateData.hashId = data;
	}
};

const actions = {
	updateHashId({ commit }, data) {
		commit(UPDATE_VERIFY_HASH, data);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};