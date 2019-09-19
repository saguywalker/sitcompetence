const path = require("path");

module.exports = {
	outputDir: path.resolve(__dirname, "dist", "admin"),
	publicPath: "/admin",
	configureWebpack: {
		resolve: {
			alias: {
				"~": path.resolve(__dirname, "src"),
				"@": path.resolve(__dirname, "src")
			}
		}
	}
};