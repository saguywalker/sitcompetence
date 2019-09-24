import { Activity } from "@/services";
import {
	LOAD_POST_ACTIVITIES,
	LOAD_SAVE_ACTIVITIES,
	LOAD_ACTIVITY
} from "../mutationTypes";

const state = {
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
			activity_id: "1823",
			title: "Charlie",
			description: "Nack Charlie"
		},
		{
			activity_id: "9121",
			title: "Tame Impala",
			description: ""
		},
		{
			activity_id: "5555",
			title: "Miyabi",
			description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Minima cupactivity_iditate laudantium ipsa est maxime quasi minus soluta laborum. Unde soluta natus itaque qui! Quod asperiores aliquid odio nihil, veniam in."
		},
		{
			activity_id: "2093",
			title: "Run",
			description: "Bab P' Toon"
		},
		{
			activity_id: "321",
			title: "Forrest Gump",
			description: ""
		},
		{
			activity_id: "8632",
			title: "Pokemon",
			description: "The best Pokemon Trainer in the world is Sasuke wtf..."
		},
		{
			activity_id: "4433",
			title: "Jeab",
			description: ""
		},
		{
			activity_id: "4321",
			title: "Prep",
			description: "Cheapest Flight"
		},
		{
			activity_id: "6332",
			title: "Potato Corner",
			description: "French Fries"
		}
	]
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
	[LOAD_ACTIVITY](stateData, data) {
		stateData.activities = [
			...data
		];
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
			commit(LOAD_ACTIVITY, response.data);
		}
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
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions,
	getters
};