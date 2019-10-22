import axios from "axios";
import { getLoginToken } from "@/helpers";

const apiClient = axios.create({
	baseURL: process.env.VUE_APP_API_URL,
	headers: {
		"X-Session-Token": getLoginToken()
	}
});

const Activity = {
	getActivities() {
		return apiClient.get("/activity?std=true");
	},
	getActivityById(id) {
		return apiClient.get(`/search/activity?activity_id=${id}`);
	},
	postJoinActivity(data) {
		return apiClient.post("/joinActivity", data);
	}
};

export {
	Activity
};