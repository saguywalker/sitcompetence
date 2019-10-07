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
			return;
		}

		sessionStorage.setItem("inlog", response.data.token);
		sessionStorage.setItem("user", response.data.user.uid);
		commit(UPDATE_LOGIN, response.data);
		location.href = "http://localhost:8082/dashboard";
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