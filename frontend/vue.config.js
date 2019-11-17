const path = require("path");
// process.env.VUE_APP_API_URL = "https://sitcompetence.ilab.sit.kmutt.ac.th:3000/api";

if (process.env.NODE_ENV === "production") {

	process.env.VUE_APP_API_URL = "https://sitcompetence.ilab.sit.kmutt.ac.th/api";
} else {
	process.env.VUE_APP_API_URL = "http://localhost:3000/api";
}

module.exports = {
	configureWebpack: {
		resolve: {
			alias: {
				"~": path.resolve(__dirname, "src"),
				"@": path.resolve(__dirname, "src")
			}
		}
	}
};
