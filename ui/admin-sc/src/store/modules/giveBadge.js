import {
	SET_SELECT_STUDENTS,
	SET_GIVE_BADGE_STEP
} from "../mutationTypes";

const state = {
	steps: [],
	selectedStudents: []
};

const mutations = {
	[SET_SELECT_STUDENTS](stateData, data) {
		stateData.selectedStudents = [
			...data
		];
	},
	[SET_GIVE_BADGE_STEP](stateData, data) {
		stateData.steps = [
			...data
		];
	}
};

const actions = {
	updateSelectedStudents({ commit }, data) {
		const payloadWithBadge = data.map((item) => {
			return {
				...item,
				badges: [],
				show: true
			};
		});

		commit(SET_SELECT_STUDENTS, payloadWithBadge);
	},
	updateStep({ commit }, data) {
		commit(SET_GIVE_BADGE_STEP, data);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};