const path = require("path");
const readYaml = require("read-yaml");

process.env.VUE_APP_SKKEY = readYaml.sync("../config.yaml").SecretKey;

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