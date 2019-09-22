import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

export default new Router({
	mode: "history",
	routes: [
		{
			path: "/login",
			name: "login",
			component: () => import(/* webpackChunkName: "login" */ "@/pages/Login.vue")
		},
		{
			path: "/main",
			name: "main",
			component: () => import(/* webpackChunkName: "main" */ "@/pages/Main.vue")
		}
	]
});
