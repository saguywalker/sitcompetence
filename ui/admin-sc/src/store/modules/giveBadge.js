import {
	UPDATE_SELECT_STUDENT,
	UPDATE_GIVE_BADGE_STEP
} from "../mutationTypes";

const state = {
	steps: [],
	selectedStudents: []
};

const mutations = {
	[UPDATE_SELECT_STUDENT](stateData, data) {
		stateData.selectedStudents = [
			...data
		];
	},
	[UPDATE_GIVE_BADGE_STEP](stateData, data) {
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

		commit(UPDATE_SELECT_STUDENT, payloadWithBadge);
	},
	updateSelectedBadge({ commit }, data) {
		commit(UPDATE_SELECT_STUDENT, data);
	},
	addStep({ commit, state: stateData }, data) {
		if (stateData.steps.includes(data)) {
			return;
		}

		const payload = [
			...stateData.steps,
			data
		];

		commit(UPDATE_GIVE_BADGE_STEP, payload);
	},
	deleteStep({ commit, state: stateData }, data) {
		const payload = stateData.steps.filter((step) => step !== data);
		commit(UPDATE_GIVE_BADGE_STEP, payload);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};