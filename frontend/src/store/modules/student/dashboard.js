import { getLoginUser, getCiphertext } from "@/helpers";
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
			const user = getLoginUser();
			const userData = {
				...user,
				additional: {
					...response.data
				}
			};
			const encryptedUser = getCiphertext(JSON.stringify(userData), process.env.VUE_APP_USER_DATA_KEY);
			localStorage.setItem("user", encryptedUser);

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