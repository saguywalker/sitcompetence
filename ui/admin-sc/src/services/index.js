import axios from "axios";
import { getLoginToken } from "@/helpers";

const apiClient = axios.create({
	baseURL: "http://localhost:3000/api",
	headers: {
		"X-Session-Token": getLoginToken
	}
});

const GiveBadge = {
	postGiveBadge(data) {
		return apiClient.post("/giveBadge", data);
	}
};

const Base = {
	login(data) {
		return apiClient.post("/login", data);
	},
	getAllBadges() {
		return apiClient.get("/competence");
	},
	getAllStudents() {
		return apiClient.get("/student");
	},
	getStudentsPage(page) {
		return apiClient.get(`/student?page=${page}`);
	},
	getAllStaffs() {
		return apiClient.get("/staffs");
	},
	getUserDetail() {
		return apiClient.get("/userDetail");
	}
};

const Activity = {
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

const Verify = {
	postVerifyTransaction(data) {
		return apiClient.post("/verify", data);
	}
};

export {
	Base,
	GiveBadge,
	Activity,
	Verify
};