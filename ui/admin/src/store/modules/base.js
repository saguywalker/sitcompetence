import { getPlainTextToken } from "@/helpers";
import { Base } from "@/services";
import {
	LOAD_LOGIN_USER,
	LOGOUT,
	NOTHING,
	LOAD_BADGES,
	LOAD_STUDENTS
} from "../mutationTypes";

const state = {
	students: [],
	badges: [],
	user: {},
	nothing: ""
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
	[NOTHING](stateData) {
		stateData.nothing = "";
	},
	[LOGOUT](stateData) {
		stateData.user = {};
	},
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
			sessionStorage.setItem("user", response.data.uid);
			commit(LOAD_LOGIN_USER, response.data);
		}

		return response;
	},
	async loadStudentData({ commit }) {
		const response = await Base.getAllStudents();

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
		const response = await Base.getAllBadges();

		if (response.status === 200) {
			commit(LOAD_BADGES, response.data);
		}

		return response;
	},
	async doLogin({ commit }, data) {
		const token = getPlainTextToken(data);
		sessionStorage.setItem("inlog", token);
		commit(NOTHING);
		location.href = "http://localhost:8080/admin/dashboard";
	},
	logout({ commit }) {
		sessionStorage.clear();
		commit(LOGOUT);
		location.href = "http://localhost:8082/login";
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};