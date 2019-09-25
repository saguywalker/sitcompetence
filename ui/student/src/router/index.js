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
			component: () => import(/* webpackChunkName: "student-layout" */"@/layouts/StudentLayout.vue"),
			children: [
				{
					name: "account",
					path: "/account",
					component: () => import(/* webpackChunkName: "account" */"@/pages/Account.vue")
				},
				{
					name: "portfolio",
					path: "/portfolio",
					component: () => import(/* webpackChunkName: "portfolio" */"@/pages/Portfolio.vue")
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
