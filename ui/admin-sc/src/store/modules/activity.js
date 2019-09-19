import {
	LOAD_POST_ACTIVITIES,
	LOAD_SAVE_ACTIVITIES
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
	saveActivities: [{
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
	}]
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
	}
};

const actions = {
	loadPostActivities({ commit }, data) {
		commit(LOAD_POST_ACTIVITIES, data);
	},
	loadSaveActivities({ commit }, data) {
		commit(LOAD_SAVE_ACTIVITIES, data);
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};