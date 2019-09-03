import EventService from "@/services/EventService";

import {
	SET_EVENTS,
	SET_EVENT
} from "../mutationTypes";

const state = {
	events: [],
	event: {
		time: "",
		date: "",
		title: ""
	}
};

const mutations = {
	[SET_EVENTS](stateData, data) {
		stateData.events = data;
	},
	[SET_EVENT](stateData, data) {
		stateData.event = data;
	}
};

const actions = {
	async fetchEvents({ commit }) {
		const { data } = await EventService.getEvents();
		commit(SET_EVENTS, data);

		return data;
	},
	async fetchEvent({ commit }, id) {
		const { data } = await EventService.getEvent(id);
		commit(SET_EVENT, data);

		return data;
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};