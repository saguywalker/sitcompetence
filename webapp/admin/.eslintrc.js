module.exports = {
	root: true,
  env: {
    browser: true,
    node: true
  },
  parser: "vue-eslint-parser",
    parserOptions: {
        "parser": "@typescript-eslint/parser"
	},
  extends: [
		"plugin:vue/recommended",
		"plugin:vue/base",
		"@nuxtjs/eslint-config-typescript",
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
				"linebreak-style": "off"
			}
		}
	]
};