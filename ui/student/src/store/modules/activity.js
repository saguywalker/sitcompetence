import { Activity } from "@/services";
import {
	LOAD_ACTIVITIES,
	LOAD_ACTIVITY_ID,
	JOIN_ACTIVITY
} from "../mutationTypes";

const state = {
	activities: [],
	activity: {},
	join: false
};

const mutations = {
	[LOAD_ACTIVITIES](stateData, data) {
		stateData.activities = [
			...data
		];
	},
	[LOAD_ACTIVITY_ID](stateData, data) {
		stateData.activity = {
			...data
		};
	},
	[JOIN_ACTIVITY](stateData, data) {
		stateData.join = data;
	}
};

const actions = {
	async loadActivities({ commit, state: stateData }) {
		if (stateData.activities.length > 0) {
			return false;
		}

		const response = await Activity.getActivities();
		if (response.status === 200) {
			commit(LOAD_ACTIVITIES, response.data);
		}

		return response;
	},
	async loadActivityById({ commit }, id) {
		const response = await Activity.getActivityById(id);

		if (response.status === 200) {
			commit(LOAD_ACTIVITY_ID, response.data[0]);
		}

		return response;
	},
	async joinActivity({ commit }, data) {
		const response = await Activity.postJoinActivity(data);

		if (response.status === 200) {
			commit(JOIN_ACTIVITY, true);
		}

		return response;
	}
};

const getters = {
	getActivityById: (stateData) => (id) => {
		return stateData.activities.find((activity) => activity.activity_id === id);
	},
	// eslint-disable-next-line no-shadow
	getApprovedActivitiesBySemester: (state, getters) => (semester) => {
		return getters.getApprovedActivities.filter((activity) => activity.semester === semester);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions,
	getters
};