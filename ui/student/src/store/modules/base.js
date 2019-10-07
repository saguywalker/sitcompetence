import { Base } from "@/services";
import {
	LOAD_LOGIN_USER
} from "../mutationTypes";

const state = {
	user: {}
};

const mutations = {
	[LOAD_LOGIN_USER](stateData, data) {
		stateData.user = {
			...data
		};
	}
};

const actions = {
	async loadUserDetail({ commit }) {
		const response = await Base.getUserDetail();

		if (response.status === 200) {
			commit(LOAD_LOGIN_USER, response.data);
		}
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};