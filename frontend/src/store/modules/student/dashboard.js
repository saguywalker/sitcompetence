import { Dashboard } from "@/services";
import {
	UPDATE_PROFILE
} from "../../mutationTypes";

const state = {
	profile: {
		profile_picture: null,
		motto: ""
	}
};

const mutations = {
	[UPDATE_PROFILE](stateData, data) {
		stateData.profile = {
			...data
		};
	}
};

const actions = {
	async updateProfile({ commit }, data) {
		const formData = new FormData();
		Object.keys(data).forEach((item) => {
			formData.append(item, data[item]);
		});
		const response = await Dashboard.editProfile(formData);
		if (response.status === 200) {
			commit(UPDATE_PROFILE, response.data);
		}
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};