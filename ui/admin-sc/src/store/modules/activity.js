import { Activity } from "@/services";
import {
	LOAD_POST_ACTIVITIES,
	LOAD_SAVE_ACTIVITIES,
	LOAD_ACTIVITIES,
	LOAD_ACTIVITY,
	EDIT_ACTIVITY,
	DELETE_ACTIVITY,
	APPROVE_ACTIVITY
} from "../mutationTypes";

const state = {
	pastActivities: [
		{
			id: "2093",
			title: "Run",
			description: "Bab P' Toon"
		},
		{
			id: "321",
			title: "Forrest Gump",
			description: ""
		},
		{
			id: "8632",
			title: "Pokemon",
			description: "The best Pokemon Trainer in the world is Sasuke wtf..."
		},
		{
			id: "4433",
			title: "Jeab",
			description: ""
		},
		{
			id: "4321",
			title: "Prep",
			description: "Cheapest Flight"
		},
		{
			id: "6332",
			title: "Potato Corner",
			description: "French Fries"
		}
	],
	postActivities: [
		{
			id: "2093",
			title: "Run",
			description: "Bab P' Toon"
		},
		{
			id: "321",
			title: "Forrest Gump",
			description: ""
		},
		{
			id: "8632",
			title: "Pokemon",
			description: "The best Pokemon Trainer in the world is Sasuke wtf..."
		},
		{
			id: "4433",
			title: "Jeab",
			description: ""
		},
		{
			id: "4321",
			title: "Prep",
			description: "Cheapest Flight"
		},
		{
			id: "6332",
			title: "Potato Corner",
			description: "French Fries"
		}
	],
	saveActivities: [
		{
			id: "1823",
			title: "Charlie",
			description: "Nack Charlie"
		},
		{
			id: "9121",
			title: "Tame Impala",
			description: ""
		},
		{
			id: "5555",
			title: "Miyabi",
			description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Minima cupiditate laudantium ipsa est maxime quasi minus soluta laborum. Unde soluta natus itaque qui! Quod asperiores aliquid odio nihil, veniam in."
		}
	],
	activities: [
		{
			id: "2093",
			title: "Run",
			description: "Bab P' Toon"
		},
		{
			id: "321",
			title: "Forrest Gump",
			description: ""
		},
		{
			id: "8632",
			title: "Pokemon",
			description: "The best Pokemon Trainer in the world is Sasuke wtf..."
		},
		{
			id: "4433",
			title: "Jeab",
			description: ""
		},
		{
			id: "4321",
			title: "Prep",
			description: "Cheapest Flight"
		},
		{
			id: "6332",
			title: "Potato Corner",
			description: "French Fries"
		}
	],
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
		const response = await Activity.getActivities();

		if (response.status === 200) {
			commit(LOAD_ACTIVITIES, response.data);
		}

		return response;
	},
	async loadActivityById({ commit }, id) {
		const response = await Activity.getActivityById(id);

		if (response.status === 200) {
			commit(LOAD_ACTIVITY, response.data);
		}

		return response;
	},
	async deleteActivityById({ commit }, id) {
		const response = await Activity.deleteActivityById(id);

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

		const	response = await Activity.postApproveActivity(payload);
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
		return stateData.activities.find((activity) => activity.id === id);
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