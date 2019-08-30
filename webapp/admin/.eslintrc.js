module.exports = {
	root: true,
  env: {
    browser: true,
    node: true
  },
	parserOptions: {
		parser: "babel-eslint"
	},
  extends: [
		"plugin:vue/recommended",
		"plugin:vue/base",
		"maqe"
	],
	plugins: [
		"vue"
	],
	overrides: [
		{
			"files": "**/*.vue",
			"rules": {
				"quotes": ["error", "double"],
				"vue/html-indent": [
					"error",
					"tab"
				]
			}
		},
		{
			"files": "**/*",
			"rules": {
				"linebreak-style": "off",
				"import/extensions": 0
			}
		}
	],
	settings: {
	}
};