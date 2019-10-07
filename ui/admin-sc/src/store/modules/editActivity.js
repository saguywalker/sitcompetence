import { Activity } from "@/services";
import {
	EDIT_ACTIVITY_STEP,
	EDIT_ACTIVITY_DETAIL,
	EDIT_ACTIVITY_COMPETENCE,
	EDIT_ACTIVITY_SUBMIT
} from "../mutationTypes";

const state = {
	detailInput: {
		activity_id: "",
		activity_name: "",
		description: "",
		activity_date: "",
		start_time: "",
		student_site: "",
		location: "",
		organizer: "",
		category: "",
		creator: "",
		semester: "",
		competences: ""
	},
	competences: [],
	steps: []
};

const mutations = {
	[EDIT_ACTIVITY_DETAIL](stateData, data) {
		stateData.detailInput = {
			...data
		};
	},
	[EDIT_ACTIVITY_COMPETENCE](stateData, data) {
		stateData.competences = [
			...data
		];
	},
	[EDIT_ACTIVITY_STEP](stateData, data) {
		stateData.steps = [
			...data
		];
	},
	[EDIT_ACTIVITY_SUBMIT](stateData) {
		stateData.steps = [];
		stateData.competences = [];
		stateData.detailInput = {
			activity_id: "",
			activity_name: "",
			description: "",
			activity_date: "",
			start_time: "",
			student_site: "",
			location: "",
			organizer: "",
			category: "",
			creator: "",
			semester: "",
			competences: ""
		};
	}
};

const actions = {
	setDetailInput({ commit }, data) {
		commit(EDIT_ACTIVITY_DETAIL, data);
	},
	setCompetenceInput({ commit }, data) {
		commit(EDIT_ACTIVITY_COMPETENCE, data);
	},
	async loadActivityById({ commit }, id) {
		const	response = await Activity.getActivityById(id);
		if (response.status === 200) {
			commit(EDIT_ACTIVITY_DETAIL, response.data[0]);
		}

		return response;
	},
	async submitEditActivity({ commit, state: stateData }, data) {
		const competenceIds = stateData.competences.map((com) => com.id);

		const payload = {
			...data,
			competences: competenceIds
		};

		const	response = await Activity.editActivityById(payload);
		if (response.status === 200) {
			commit(EDIT_ACTIVITY_SUBMIT);
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

		commit(EDIT_ACTIVITY_STEP, payload);
	},
	deleteStep({ commit, state: stateData }, data) {
		const payload = stateData.steps.filter((step) => step !== data);
		commit(EDIT_ACTIVITY_STEP, payload);
	},
	clearStep({ commit }) {
		commit(EDIT_ACTIVITY_STEP, []);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};