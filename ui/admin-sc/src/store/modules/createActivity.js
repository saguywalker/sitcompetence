import {
	CREATE_ACTIVITY_STEP,
	CREATE_ACTIVITY_DETAIL,
	CREATE_ACTIVITY_COMPETENCE
} from "../mutationTypes";

const state = {
	detailInput: {
		name: "",
		description: "",
		img: null,
		openRegistDate: "",
		openRegistTime: "",
		closeRegistDate: "",
		closeRegistTime: "",
		actStartDate: "",
		actStartTime: "",
		actEndDate: "",
		actEndTime: ""
	},
	competences: [],
	steps: []
};

const mutations = {
	[CREATE_ACTIVITY_DETAIL](stateData, data) {
		stateData.detailInput = {
			...stateData.detailInput,
			...data
		};
	},
	[CREATE_ACTIVITY_COMPETENCE](stateData, data) {
		stateData.competences = [
			...data
		];
	},
	[CREATE_ACTIVITY_STEP](stateData, data) {
		stateData.steps = [
			...data
		];
	}
};

const actions = {
	setDetailInput({ commit }, data) {
		commit(CREATE_ACTIVITY_DETAIL, data);
	},
	setCompetenceInput({ commit }, data) {
		commit(CREATE_ACTIVITY_COMPETENCE, data);
	},
	addStep({ commit, state: stateData }, data) {
		if (stateData.steps.includes(data)) {
			return;
		}

		const payload = [
			...stateData.steps,
			data
		];

		commit(CREATE_ACTIVITY_STEP, payload);
	},
	deleteStep({ commit, state: stateData }, data) {
		const payload = stateData.steps.filter((step) => step !== data);
		commit(CREATE_ACTIVITY_STEP, payload);
	},
	clearStep({ commit }) {
		commit(CREATE_ACTIVITY_STEP, []);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};