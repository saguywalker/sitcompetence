import EventService from "@/services/EventService";

export const state = () => ({
	events: [],
	event: {
		time: "",
		date: "",
		title: ""
	}
});

export const mutations = {
	SET_EVENTS(stateData, data) {
		stateData.events = data;
	},
	SET_EVENT(stateData, data) {
		stateData.event = data;
	}
};

export const actions = {
	async fetchEvents({ commit }) {
		const { data } = await EventService.getEvents();
		commit("SET_EVENTS", data);

		return data;
	},
	async fetchEvent({ commit }, id) {
		const { data } = await EventService.getEvent(id);
		commit("SET_EVENT", data);

		return data;
	}
};