const path = require("path");

if (process.env.NODE_ENV === "production") {
	process.env.VUE_APP_API_URL = "https://10.4.56.22:3000/api";
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