const path = require("path");
if (process.env.NODE_ENV === "production") {
	process.env.VUE_APP_PROD_URL = "https://sitcompetence.ilab.sit.kmutt.ac.th";
	process.env.VUE_APP_API_URL = "https://sitcompetence.ilab.sit.kmutt.ac.th/api";
} else {
	process.env.VUE_APP_PROD_URL = "http://localhost:8080";
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
