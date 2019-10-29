import router from "@/router";
import { getCiphertext, clearCookies } from "@/helpers";
import { Base, Login } from "@/services";
import {
	LOAD_LOGIN_DATA,
	LOGOUT,
	LOAD_BADGES,
	LOAD_STUDENTS
} from "../mutationTypes";

const state = {
	students: [],
	badges: [],
	user: {}
};

const mutations = {
	[LOAD_STUDENTS](stateData, data) {
		stateData.students = [
			...data
		];
	},
	[LOAD_BADGES](stateData, data) {
		stateData.badges = [
			...data
		];
	},
	[LOGOUT](stateData) {
		stateData.user = {};
	},
	[LOAD_LOGIN_DATA](stateData, data) {
		stateData.user = {
			...data
		};
	}
};

const actions = {
	async loadStudentData({ commit }) {
		const response = await Base.getStudents();

		if (response.status === 200) {
			commit(LOAD_STUDENTS, response.data);
		}

		return response;
	},
	async loadStudentDataByPage({ commit }, page) {
		const response = await Base.getStudentsPage(page);

		if (response.status === 200) {
			commit(LOAD_STUDENTS, response.data);
		}

		return response.data;
	},
	async loadBadgeData({ commit }) {
		const response = await Base.getBadges();

		if (response.status === 200) {
			commit(LOAD_BADGES, response.data);
		}

		return response;
	},
	async doLogin({ commit }, data) {
		if (data.username === "" && data.password === "") {
			commit(LOGOUT);
			return;
		}
		const response = await Login.login(data);
		if (response.status !== 200) {
			return;
		}

		const loginDataString = JSON.stringify(response.data);
		const encryptedLoginData = getCiphertext(loginDataString, process.env.VUE_APP_USER_DATA_KEY);
		localStorage.setItem("user", encryptedLoginData);

		if (response.data.group === "inst_group") {
			router.push({ name: "admin" });
			// location.href = "http://localhost:8080/admin";
			return;
		}

		commit(LOAD_LOGIN_DATA, response.data);
		router.push({ name: "student" });
		// location.href = "http://localhost:8080";
	},
	logout({ commit }) {
		localStorage.removeItem("user");
		clearCookies();
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