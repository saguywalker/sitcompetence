export interface Event {
	time: string;
	date: string;
	title: string;
};

export interface EventState {
	events: Event[];
	event: Event
};