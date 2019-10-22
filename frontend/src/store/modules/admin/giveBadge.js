import { GiveBadge } from "@/services";
import {
	GIVE_BADGE_SELECT_STUDENT,
	GIVE_BADGE_STEP,
	GIVE_BADGE_SUBMIT,
	GIVE_BADGE_SELECT_BADGE,
	GIVE_BADGE_SUCCESS,
	GIVE_BADGE_CLEAR
} from "../../mutationTypes";

const state = {
	steps: [],
	selectedStudents: [],
	success: []
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
	[GIVE_BADGE_SUBMIT](stateData, data) {
		stateData.selectedStudents = [
			...data
		];
	},
	[GIVE_BADGE_SUBMIT](stateData) {
		stateData.selectedStudents = [];
	},
	[GIVE_BADGE_SUCCESS](stateData, data) {
		stateData.success = [
			...data
		];
	},
	[GIVE_BADGE_STEP](stateData, data) {
		stateData.steps = [
			...data
		];
	},
	[GIVE_BADGE_CLEAR](stateData) {
		stateData.selectedStudents = [];
		stateData.steps = [];
		stateData.success = [];
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
	async submitGiveBadge({ commit, state: stateData }, data) {
		const filterData = stateData.selectedStudents.map((st) => {
			const idBadges = st.badges.map((badge) => badge.competence_id);

			const separateBadge = idBadges.map((id) => {
				return {
					student_id: st.student_id,
					competence_id: id,
					giver: data.giver,
					semester: data.semester
				};
			});

			return separateBadge;
		});

		const payload = [];

		filterData.forEach((student) => {
			student.forEach((x) => {
				payload.push(x);
			});
		});

		const	response = await GiveBadge.postGiveBadge(payload);
		if (response.status === 200) {
			commit(GIVE_BADGE_SUCCESS, payload);
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

		commit(GIVE_BADGE_STEP, payload);
	},
	deleteStep({ commit, state: stateData }, data) {
		const payload = stateData.steps.filter((step) => step !== data);
		commit(GIVE_BADGE_STEP, payload);
	},
	clearStep({ commit }) {
		commit(GIVE_BADGE_CLEAR);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};