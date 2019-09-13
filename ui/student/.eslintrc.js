module.exports = {
  root: true,
  env: {
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
      files: "**/*.vue",
      rules: {
        quotes: [
          "error",
          "double"
        ],
        "vue/html-indent": [
          "error",
          "tab"
        ],
        "import/extensions": [
          "error",
          "never",
          {
            "vue": "always"
          }
        ]
      }
    },
    {
      files: "**/*",
      rules: {
        "linebreak-style": "off",
        "import/extensions": 0,
        quotes: [
          "error",
          "double"
        ],
        indent: [
          "error",
          "tab"
        ]
      }
    },
    {
      files: [
        "**/__tests__/*.{j,t}s?(x)"
      ],
      env: {
        "mocha": true
      }
    }
  ]
}
