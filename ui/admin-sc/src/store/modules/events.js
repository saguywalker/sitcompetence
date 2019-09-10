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

		return response;
	},
	async fetchEvent({ commit, dispatch, getters }, id) {
		let response;
		const event = getters.getEventById(id);

		if (event) {
			commit(SET_EVENT, event);
			return event;
		}

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

		return response.data;
	}
};

const getters = {
	getEventById: (stateData) => (id) => (
		stateData.events.find((event) => event.id === parseInt(id, 10))
	)
};

export default {
	namespaced: true,
	state,
	mutations,
	actions,
	getters
};