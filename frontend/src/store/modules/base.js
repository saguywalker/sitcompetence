import router from "@/router";
import { getCiphertext, getEncryptedHex, clearLoginState } from "@/helpers";
import { Base, Login, Portfolio } from "@/services";
import {
	LOAD_LOGIN_DATA,
	LOGOUT,
	LOAD_BADGES,
	LOAD_STUDENTS,
	UPDATE_NOTIFICATION,
	LOAD_SHARE_PORTFOLIO
} from "../mutationTypes";

const state = {
	students: [],
	badges: [],
	user: {},
	notifications: [],
	sharePortfolio: {}
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
	},
	[LOAD_SHARE_PORTFOLIO](stateData, data) {
		stateData.sharePortfolio = {
			...data
		};
	}
};

const actions = {
	async loadSharePortfolio({ commit }, urlKey) {
		const response = await Portfolio.getBadgeWithToken(urlKey);

		if (response.status === 200) {
			commit(LOAD_SHARE_PORTFOLIO, response.data);
		}

		return response;
	},
	async loadStudentData({ commit }) {
		const response = await Base.getStudents();

		if (response.status === 200) {
			commit(LOAD_STUDENTS, response.data);
		}

		return response;
	},
	async loadSearchTable({ commit }, data) {
		const response = await Base.getStudentsBySearch(data.key, data.value);

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
	async doLogin({ dispatch, commit }, data) {
		if (data.username === "" && data.password === "") {
			commit(LOGOUT);
			return;
		}
		const payload = {
			username: data.username,
			password: getEncryptedHex(data.password)
		};
		const response = await Login.login(payload);
		if (response.status !== 200) {
			if (response.status === 401) {
				const notification = {
					title: "Login Failed",
					message: "Invalid username or password",
					variant: "danger"
				};
				dispatch("addNotification", notification);
			} else if (response.status === 403) {
				const notification = {
					title: "Login Blocked",
					message: "You have been blocked for 3 minutes",
					variant: "danger"
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
		clearLoginState();
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