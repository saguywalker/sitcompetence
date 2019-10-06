import {
	UPDATE_SESSION,
	LOGOUT
} from "../mutationTypes";

const state = {
	session: null,
	role: "",
	user: {}
};

const mutations = {
	[UPDATE_SESSION](stateData, data) {
		stateData.user = {
			...data
		};
	},
	[LOGOUT](stateData) {
		stateData.session = null;
	}
};

const actions = {
	async doLogin({ commit }, data) {
		if (data.email === "" && data.password === "") {
			commit(UPDATE_SESSION, null);
			return;
		}
		const response = await authenticationAxiosInstance.post("/login", data);

		if (response.status === 200) {
			sessionStorage.setItem("loginSession", response.data.role);
			if (response.data.role === "partner") {
				sessionStorage.setItem("partner", response.data.partner);
			}
			commit(UPDATE_SESSION, response.data);
			router.push({ name: "admin" });
		}
	},
	logout({ commit }) {
		sessionStorage.clear();
		commit(LOGOUT);
		router.push({ name: "login" });
	}
};


export default {
	namespaced: true,
	state,
	mutations,
	actions
};