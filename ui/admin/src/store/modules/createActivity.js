import { Activity } from "@/services";
import {
	CREATE_ACTIVITY_STEP,
	CREATE_ACTIVITY_DETAIL,
	CREATE_ACTIVITY_COMPETENCE,
	CREATE_ACTIVITY_SUBMIT
} from "../mutationTypes";

const state = {
	detailInput: {
		activity_name: "",
		description: "",
		activity_date: "",
		start_time: "",
		location: "",
		organizer: "",
		category: ""
	},
	competences: [],
	steps: []
};

const mutations = {
	[CREATE_ACTIVITY_DETAIL](stateData, data) {
		stateData.detailInput = {
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
	},
	[CREATE_ACTIVITY_SUBMIT](stateData) {
		stateData.steps = [];
		stateData.competences = [];
		stateData.detailInput = {
			activity_name: "",
			description: "",
			activity_date: "",
			start_time: "",
			location: "",
			organizer: "",
			category: ""
		};
	}
};

const actions = {
	setDetailInput({ commit }, data) {
		commit(CREATE_ACTIVITY_DETAIL, data);
	},
	setCompetenceInput({ commit }, data) {
		commit(CREATE_ACTIVITY_COMPETENCE, data);
	},
	async submitCreateActivity({ commit, state: stateData }, data) {
		const competenceIds = stateData.competences.map((com) => com.competence_id);

		const payload = {
			...data,
			competences_id: competenceIds
		};

		const	response = await Activity.postCreateActivity(payload);
		if (response.status === 200) {
			commit(CREATE_ACTIVITY_SUBMIT);
		}

		return response;
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