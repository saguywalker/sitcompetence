import axios, { AxiosResponse, AxiosInstance } from "axios";

const apiClient: AxiosInstance = axios.create({
	baseURL: "http://my-json-server.typicode.com/Code-Pop/real-world-nuxt/",
	withCredentials: false,
	headers: {
		Accept: "application/json",
		"Content-Type": "applcation/json"
	}
});

interface Event {
	time: string
	date: string
	title: string
}

export default {
	getEvents(): Promise<AxiosResponse<Event[]>> {
		return apiClient.get("/events");
	},
	getEvent(id: number): Promise<AxiosResponse<Event>> {
		return apiClient.get(`/events/${id}`);
	}
}