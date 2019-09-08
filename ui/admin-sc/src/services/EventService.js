import axios from "axios";

const apiClient = axios.create({
	baseURL: "http://my-json-server.typicode.com/Code-Pop/real-world-nxt/",
	withCredentials: false,
	headers: {
		Accept: "application/json",
		"Content-Type": "applcation/json"
	}
});

export default {
	getEvents() {
		return apiClient.get("/events");
	},
	getEvent(id) {
		return apiClient.get(`/events/${id}`);
	}
};