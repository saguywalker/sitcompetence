import axios from "axios";

const apiClient = axios.create({
	baseURL: "http://localhost:3000/api",
	withCredentials: false,
	headers: {
		Accept: "application/json",
		"Content-Type": "application/json"
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

const Verify = {
	postVerifyTransaction(data) {
		return apiClient.post("/verify", data);
	}
};

export {
	Base,
	GiveBadge,
	Verify
};