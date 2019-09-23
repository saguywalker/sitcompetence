import { Activity } from "@/services";
import {
	CREATE_ACTIVITY_STEP,
	CREATE_ACTIVITY_DETAIL,
	CREATE_ACTIVITY_COMPETENCE,
	CREATE_ACTIVITY_SUBMIT
} from "../mutationTypes";

const state = {
	detailInput: {
		name: "",
		description: "",
		activityDate: ""
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
		stateData.detailInput = {};
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
		const payload = {
			activity_name: data.name,
			description: data.description,
			activity_date: data.activityDate,
			student_site: data.student_site,
			creator: data.creator,
			competences: stateData.competences
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