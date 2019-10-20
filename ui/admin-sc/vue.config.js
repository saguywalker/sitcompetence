const path = require("path");
const SWPrecache = require('sw-precache-webpack-plugin');

module.exports = {
	publicPath: "admin",
	outputDir: path.resolve(__dirname, "dist", "admin"),
	configureWebpack: {
		plugins: [
      new SWPrecache({
				cacheId: "sit-competence",
				filepath: 'service-worker.js',
				staticFileGlobs: [
					"index.html",
					"manifest.json",
					"dist/admin/*.{js,css}"
				],
				stripPrefix: "/admin"
			})
    ],
		resolve: {
			alias: {
				"~": path.resolve(__dirname, "src"),
				"@": path.resolve(__dirname, "src")
			}
		}
	}
};