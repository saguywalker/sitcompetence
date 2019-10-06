import router from "@/router";
import { getCiphertext } from "@/helpers";
import { Login } from "@/services";
import {
	UPDATE_LOGIN,
	LOGOUT
} from "../mutationTypes";

const state = {
	role: "",
	user: {}
};

const mutations = {
	[UPDATE_LOGIN](stateData, data) {
		stateData.user = {
			...data
		};
	},
	[LOGOUT](stateData) {
		stateData.user = {};
	}
};

const actions = {
	async doLogin({ commit }, data) {
		if (data.username === "" && data.password === "") {
			commit(LOGOUT);
			return;
		}
		const response = await Login.login(data);
		if (response.status !== 200) {
			return;
		}

		if (response.data.user.group === "inst_group") {
			const hash = getCiphertext(response.data.token);
			location.href = `http://localhost:8080/admin/login/${hash}`;
		} else {
			sessionStorage.setItem("inlog", true);
			commit(UPDATE_LOGIN, response.data);
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