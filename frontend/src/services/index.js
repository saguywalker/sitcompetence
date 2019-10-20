import axios from "axios";
import { getLoginToken } from "@/helpers";

const apiClient = axios.create({
	baseURL: "http://localhost:3000/api",
	headers: {
		"X-Session-Token": getLoginToken
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

export {
	Login,
	Base,
	Verify
};