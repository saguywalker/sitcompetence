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
	async fetchEvents({ commit, dispatch }) {
		let response;

		try {
			response = await EventService.getEvents();
			commit(SET_EVENTS, response.data);
		} catch (err) {
			const notification = {
				title: "Fetch Event",
				message: `There was a problem fetching events: ${err.message}`,
				variant: "danger"
			};

			dispatch("notification/add", notification, { root: true });
		}
	},
	async fetchEvent({ commit, dispatch }, id) {
		let response;

		try {
			response = await EventService.getEvent(id);
			commit(SET_EVENT, response.data);
		} catch (err) {
			const notification = {
				title: "Fetch Event",
				message: `There was a problem fetching events: ${err.message}`,
				variant: "danger"
			};

			dispatch("notification/add", notification, { root: true });
		}
	}
};

export default {
	namespaced: true,
	state,
	mutations,
	actions
};