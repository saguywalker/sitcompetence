import axios from "axios";
import router from "@/router";

axios.defaults.withCredentials = true;
// Handlers
const handleResponse = (response) => response;
const handleErrors = (error) => {
	if (error.response.status === 403) {
		router.push({ name: "error403" });
	} else if (error.response.status === 451) {
		router.push({ name: "user-genkey" });
	} else if (error.response.status === 404) {
		router.push({ name: "error404" });
	}
};

const apiClient = axios.create({
	baseURL: process.env.VUE_APP_API_URL
});

const apiClientLogin = axios.create({
	baseURL: process.env.VUE_APP_API_URL
});

// ------- Interceptors --------
apiClient.interceptors.response.use(handleResponse, handleErrors);
// -----------------------------

const Login = {
	login(data) {
		return apiClientLogin.post("/login", data);
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
	getStudentsBySearch(key, value) {
		return apiClient.get(`/search/student?${key}=${value}`);
	},
	getStudentsPage(page) {
		return apiClient.get(`/search/student?page=${page}`);
	},
	getStaffs() {
		return apiClient.get("/staff");
	},
	getUserDetail() {
		return apiClient.get("/userDetail");
	},
	postSetPublicKey(data) {
		return apiClient.post("/admin/setkey", data);
	}
};

const Portfolio = {
	getBadgeWithToken(key) {
		return apiClient.get(`/profile/${key}`);
	},
	getShareLink() {
		return apiClient.get("/student/shareProfile");
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
	Activity,
	Portfolio
};