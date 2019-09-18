import { Base } from "@/services";
import {
	LOAD_BADGES,
	LOAD_STUDENTS
} from "../mutationTypes";

const state = {
	students: [],
	badges: []
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
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};