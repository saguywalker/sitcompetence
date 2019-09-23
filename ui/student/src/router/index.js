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
			path: "/",
			component: () => import(/* webpackChunkName: "studentLayout" */"@/layouts/StudentLayout.vue"),
			children: [
				{
					name: "dashboard",
					path: "/",
					component: () => import(/* webpackChunkName: "studentLayout" */"@/pages/Dashboard.vue")
				}
			]
		},
		{
			name: "admin",
			path: "/admin",
			beforeEnter() {
				location.href = "http://github.com";
			}
		}
	]
});
