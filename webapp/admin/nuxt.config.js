
export default {
	mode: "universal",
	/*
	** Headers of the page
	*/
	head: {
		title: process.env.npm_package_name || "",
		meta: [
			{ charset: "utf-8" },
			{ name: "viewport", content: "width=device-width, initial-scale=1" },
			{ hid: "description", name: "description", content: process.env.npm_package_description || "" }
		],
		link: [
			{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }
		]
	},
	/*
	** Customize the progress-bar color
	*/
	loading: "~/components/loading.vue",
	/*
	** Global CSS
	*/
	css: [
		{ src: "~/styles/global.scss", lang: "scss", rel: "preload" }
	],
	/*
	** Plugins to load before mounting the App
	*/
	plugins: [
	],
	/*
	** Nuxt.js dev-modules
	*/
	buildModules: [
		// Doc: https://github.com/nuxt-community/eslint-module
		"@nuxtjs/eslint-module",
		"@nuxt/typescript-build"
	],
	/*
	** Nuxt.js modules
	*/
	modules: [
		// Doc: https://axios.nuxtjs.org/usage
		"@nuxtjs/axios",
		"bootstrap-vue/nuxt"
	],
	bootstrapVue: {
		bootstrapCSS: false, // Or `css: false`
    bootstrapVueCSS: false, // Or `bvCSS: false`
		componentPlugins: [
      "ToastPlugin",
      "ModalPlugin"
    ],
		components: ["BContainer", "BRow", "BCol", "BFormInput", "BButton", "BTable", "BSpinner"]
  },
	/*
	** Axios module configuration
	** See https://axios.nuxtjs.org/options
	*/
	axios: {
	},
	/*
	** Build configuration
	*/
	build: {
		/*
		** You can extend webpack config here
		*/
	}
};
