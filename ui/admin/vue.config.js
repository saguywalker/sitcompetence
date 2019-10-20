const path = require("path");

module.exports = {
	publicPath: "admin",
	outputDir: path.resolve(__dirname, "dist", "admin"),
	configureWebpack: {
		resolve: {
			alias: {
				"~": path.resolve(__dirname, "src"),
				"@": path.resolve(__dirname, "src")
			}
		}
	}
};