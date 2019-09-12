import {
	GIVE_BADGE_SELECT_STUDENT,
	GIVE_BADGE_STEP,
	GIVE_BADGE_CONFIRM_BADGE,
	GIVE_BADGE_SELECT_BADGE,
	GIVE_BADGE_SUCCESS
} from "../mutationTypes";

const state = {
	steps: [],
	selectedStudents: []
};

const mutations = {
	[GIVE_BADGE_SELECT_STUDENT](stateData, data) {
		stateData.selectedStudents = [
			...data
		];
	},
	[GIVE_BADGE_SELECT_BADGE](stateData, data) {
		stateData.selectedStudents = [
			...data
		];
	},
	[GIVE_BADGE_CONFIRM_BADGE](stateData, data) {
		stateData.selectedStudents = [
			...data
		];
	},
	[GIVE_BADGE_SUCCESS](stateData) {
		stateData.selectedStudents = [];
	},
	[GIVE_BADGE_STEP](stateData, data) {
		stateData.steps = [
			...data
		];
	}
};

const actions = {
	updateSelectedStudents({ commit }, data) {
		const payloadWithBadge = data.map((item) => {
			if (item.badges) {
				return {
					...item
				};
			}

			return {
				...item,
				badges: [],
				show: true
			};
		});

		commit(GIVE_BADGE_SELECT_STUDENT, payloadWithBadge);
	},
	updateSelectedBadge({ commit }, data) {
		commit(GIVE_BADGE_SELECT_BADGE, data);
	},
	confirmGiveBadge({ commit, state: stateData }, data) {
		const payload = stateData.selectedStudents.map((student) => {
			return {
				...student,
				giver: data
			};
		});

		commit(GIVE_BADGE_CONFIRM_BADGE, payload);
	},
	giveBadgeSuccess({ commit }) {
		commit(GIVE_BADGE_SUCCESS);
	},
	addStep({ commit, state: stateData }, data) {
		if (stateData.steps.includes(data)) {
			return;
		}

		const payload = [
			...stateData.steps,
			data
		];

		commit(GIVE_BADGE_STEP, payload);
	},
	deleteStep({ commit, state: stateData }, data) {
		const payload = stateData.steps.filter((step) => step !== data);
		commit(GIVE_BADGE_STEP, payload);
	},
	clearStep({ commit }) {
		commit(GIVE_BADGE_STEP, []);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};