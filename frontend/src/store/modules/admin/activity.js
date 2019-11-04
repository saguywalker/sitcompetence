import { AdminActivity } from "@/services";
import {
	LOAD_POST_ACTIVITIES,
	LOAD_SAVE_ACTIVITIES,
	LOAD_ACTIVITIES,
	LOAD_ACTIVITY,
	EDIT_ACTIVITY,
	DELETE_ACTIVITY,
	APPROVE_ACTIVITY
} from "../../mutationTypes";

const state = {
	pastActivities: [],
	postActivities: [],
	saveActivities: [],
	activities: [],
	responseData: {},
	activity: {}
};

const mutations = {
	[LOAD_POST_ACTIVITIES](stateData, data) {
		stateData.postActivities = [
			...data
		];
	},
	[LOAD_SAVE_ACTIVITIES](stateData, data) {
		stateData.saveActivities = [
			...data
		];
	},
	[LOAD_ACTIVITIES](stateData, data) {
		stateData.activities = [
			...data
		];
	},
	[LOAD_ACTIVITY](stateData, data) {
		stateData.activity = {
			...data
		};
	},
	[EDIT_ACTIVITY](stateData, data) {
		stateData.activity = {
			...data
		};
	},
	[DELETE_ACTIVITY](stateData) {
		stateData.activity = {};
	},
	[APPROVE_ACTIVITY](stateData, data) {
		stateData.responseData = {
			...data
		};
	}
};

const actions = {
	loadPostActivities({ commit }, data) {
		commit(LOAD_POST_ACTIVITIES, data);
	},
	loadSaveActivities({ commit }, data) {
		commit(LOAD_SAVE_ACTIVITIES, data);
	},
	async loadActivity({ commit }) {
		const response = await AdminActivity.getActivities();

		if (response.status === 200) {
			commit(LOAD_ACTIVITIES, response.data);
		}

		return response;
	},
	async loadActivityById({ commit }, id) {
		const response = await AdminActivity.getActivityById(id);

		if (response.status === 200) {
			commit(LOAD_ACTIVITY, response.data[0]);
		}

		return response;
	},
	async deleteActivityById({ commit }, id) {
		const response = await AdminActivity.deleteActivityById(id);

		if (response.status === 200) {
			commit(DELETE_ACTIVITY);
		}

		return response;
	},
	async submitApprove({ commit }, data) {
		const payload = data.approvedStudents.map((student) => {
			return {
				student_id: student.id,
				activity_id: data.activityId,
				approver: data.approver
			};
		});

		const	response = await AdminActivity.postApproveActivity(payload);
		if (response.status === 200) {
			commit(APPROVE_ACTIVITY, response.data);
		}

		return response;
	}
};

const getters = {
	postedActivities: (stateData) => {
		return stateData.activities.filter((activity) => activity.student_site);
	},
	savedActivities: (stateData) => {
		return stateData.activities.filter((activity) => !activity.student_site);
	},
	getActivityById: (stateData) => (id) => {
		return stateData.activities.find((activity) => activity.activity_id === id);
	},
	// eslint-disable-next-line no-shadow
	getApprovedActivitiesBySemester: (state, getters) => (semester) => {
		return getters.getApprovedActivities.filter((activity) => activity.semester === semester);
	},
	getApprovedActivities: (stateData) => {
		return stateData.activities.filter((activity) => activity.approved);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions,
	getters
};