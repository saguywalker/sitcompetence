import { getPlainTextToken } from "@/helpers";
import { Base } from "@/services";
import {
	UPDATE_LOGIN,
	LOGOUT,
	LOAD_BADGES,
	LOAD_STUDENTS
} from "../mutationTypes";

const state = {
	students: [],
	badges: [],
	user: {},
	token: ""
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
	[UPDATE_LOGIN](stateData, data) {
		stateData.token = data;
	},
	[LOGOUT](stateData) {
		stateData.token = "";
	}
};

const actions = {
	async loadStudentData({ commit, dispatch }) {
		let response;

		try {
			response = await Base.getAllStudents();
			commit(LOAD_STUDENTS, response.data);
		} catch (err) {
			const notification = {
				title: "Fetch student data",
				message: `There was a problem fetching data: ${err.message}`,
				variant: "danger"
			};

			dispatch("notification/add", notification, { root: true });
		}
	},
	async loadBadgeData({ commit, dispatch }) {
		let response;

		try {
			response = await Base.getAllBadges();
			commit(LOAD_BADGES, response.data);
		} catch (err) {
			const notification = {
				title: "Fetch student data",
				message: `There was a problem fetching data: ${err.message}`,
				variant: "danger"
			};

			dispatch("notification/add", notification, { root: true });
		}
	},
	async doLogin({ commit }, data) {
		const token = getPlainTextToken(data);

		sessionStorage.setItem("inlog", token);
		commit(UPDATE_LOGIN, token);
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