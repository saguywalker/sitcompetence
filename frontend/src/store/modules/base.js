import router from "@/router";
import { getCiphertext, clearLoginCookie } from "@/helpers";
import { Base, Login } from "@/services";
import {
	LOAD_LOGIN_DATA,
	LOGOUT,
	LOAD_BADGES,
	LOAD_STUDENTS,
	UPDATE_NOTIFICATION
} from "../mutationTypes";

const state = {
	students: [],
	badges: [],
	user: {},
	notifications: []
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
	},
	[UPDATE_NOTIFICATION](stateData, data) {
		stateData.notifications = [
			...data
		];
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
	async doLogin({ commit, dispatch }, data) {
		if (data.username === "" && data.password === "") {
			commit(LOGOUT);
			return;
		}
		const response = await Login.login(data);

		if (response.status !== 200) {
			if (response.status === 403) {
				const notification = {
					title: "Set Public key",
					message: "Submit key successful",
					variant: "success"
				};

				dispatch("addNotification", notification);
			}
			return;
		}

		const loginDataString = JSON.stringify(response.data.user);
		const encryptedLoginData = getCiphertext(loginDataString, process.env.VUE_APP_USER_DATA_KEY);
		localStorage.setItem("user", encryptedLoginData);

		if (response.data.user.group === "inst_group") {
			if (response.data.first) {
				router.push({ name: "user-genkey" });
				return;
			}

			router.push({ name: "admin" });
			return;
		}

		commit(LOAD_LOGIN_DATA, response.data);
		router.push({ name: "student" });
	},
	logout({ commit }) {
		localStorage.removeItem("user");
		clearLoginCookie();
		commit(LOGOUT);
		router.push({ name: "login" });
	},
	addNotification({ commit, state: stateData }, data) {
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