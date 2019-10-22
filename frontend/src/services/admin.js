import axios from "axios";
import { getLoginToken } from "@/helpers";

const apiClient = axios.create({
	baseURL: process.env.VUE_APP_API_URL,
	headers: {
		"X-Session-Token": getLoginToken()
	}
});

const GiveBadge = {
	postGiveBadge(data) {
		return apiClient.post("/giveBadge", data);
	}
};

const AdminActivity = {
	postApproveActivity(data) {
		return apiClient.post("/approveActivity", data);
	},
	postCreateActivity(data) {
		return apiClient.post("/activity", data);
	},
	getActivities() {
		return apiClient.get("/activity");
	},
	getActivityById(id) {
		return apiClient.get(`/search/activity?activity_id=${id}`);
	},
	editActivityById(data) {
		return apiClient.put("/activity", data);
	},
	deleteActivityById(id) {
		return apiClient.delete(`/activity/${id}`);
	}
};

export {
	GiveBadge,
	Activity
};