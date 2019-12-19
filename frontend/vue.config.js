const path = require("path");

if (process.env.NODE_ENV === "production") {
	process.env.VUE_APP_PROD_URL = "http://domainname:port";
	process.env.VUE_APP_API_URL = "http://domainname:port/api";
} else if (process.env.NODE_ENV === "localproduction") {
	process.env.VUE_APP_PROD_URL = `http://${process.env.VUE_APP_BASE_IP}`;
	process.env.VUE_APP_API_URL = `http://${process.env.VUE_APP_BASE_IP}:3000/api`;
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
