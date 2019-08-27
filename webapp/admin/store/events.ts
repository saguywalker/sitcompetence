import { ActionTree, MutationTree, ActionContext } from "vuex";
import EventService from "@/services/EventService";
import { EventState, Event } from "@/types/events";

export const state: () => EventState = () => ({
		events: [],
		event: {
			time: "",
			date: "",
			title: ""
		}
});

export const mutations: MutationTree<EventState> = {
	SET_EVENTS(stateData: EventState, data: Event[]) {
		stateData.events = data;
	},
	SET_EVENT(stateData: EventState, data: Event) {
		stateData.event = data;
	}
};

export const actions: ActionTree<EventState, any> = {
	async fetchEvents({ commit }: ActionContext<EventState, any>): Promise<any> {
		const { data } = await EventService.getEvents();
		commit("SET_EVENTS", data);

		return data;
	},
	async fetchEvent({ commit }: ActionContext<EventState, any>, id: number): Promise<any> {
		const { data } = await EventService.getEvent(id);
		commit("SET_EVENT", data);

		return data;
	}
};