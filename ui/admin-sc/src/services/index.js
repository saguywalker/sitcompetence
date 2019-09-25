import axios from "axios";

const apiClient = axios.create({
	baseURL: "http://localhost:3000/api",
	withCredentials: true,
	headers: {
		"Access-Control-Allow-Origin": "http://localhost:3000"
	}
});

const GiveBadge = {
	postGiveBadge(data) {
		return apiClient.post("/giveBadge", data);
	}
};

const Base = {
	getAllBadges() {
		return apiClient.get("/competences");
	},
	getAllStudents() {
		return apiClient.get("/students");
	}
};

const Activity = {
	postApproveActivity(data) {
		return apiClient.post("/approve", data);
	},
	postCreateActivity(data) {
		return apiClient.post("/activity", data);
	},
	getActivities() {
		return apiClient.get("/activity");
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