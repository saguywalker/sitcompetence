import axios from "axios";
import { getLoginToken } from "@/helpers";

const apiClient = axios.create({
	baseURL: process.env.VUE_APP_API_URL,
	headers: {
		"X-Session-Token": getLoginToken()
	}
});

const Login = {
	login(data) {
		return apiClient.post("/login", data);
	}
};

const Base = {
	getBadges() {
		return apiClient.get("/competence");
	},
	getBadgesByStudentId(id) {
		return apiClient.get(`/search/competence?student_id=${id}`);
	},
	getStudents() {
		return apiClient.get("/student");
	},
	getStudentsPage(page) {
		return apiClient.get(`/search/student?page=${page}`);
	},
	getStaffs() {
		return apiClient.get("/staff");
	},
	getUserDetail() {
		return apiClient.get("/userDetail");
	}
};

const Verify = {
	postVerifyTransaction(data) {
		return apiClient.post("/verify", data);
	}
};

// -------------- Admin -------------
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
// ------------------------------------

// ------------ Student ---------------
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
// -------------------------------------


export {
	Login,
	Base,
	Verify,
	GiveBadge,
	AdminActivity,
	Activity
};