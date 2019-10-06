import Vue from "vue";
import Router from "vue-router";
import store from "@/store";

Vue.use(Router);

const router = new Router({
	mode: "history",
	routes: [
		{
			path: "/login",
			name: "login",
			component: () => import(/* webpackChunkName: "login" */ "@/pages/Login.vue")
		},
		{
			path: "/user/:id/portfolio",
			name: "user_id-portfoilio",
			component: () => import(/* webpackChunkName: "user-portfoilio" */ "@/pages/UserPortfolio_id.vue")
		},
		{
			path: "/",
			component: () => import(/* webpackChunkName: "student-layout" */"@/layouts/StudentLayout.vue"),
			children: [
				{
					path: "/",
					redirect: { name: "dashboard" }
				},
				{
					name: "dashboard",
					path: "/dashboard",
					component: () => import(/* webpackChunkName: "dashboard" */"@/pages/Dashboard.vue")
				},
				{
					name: "activity",
					path: "/activity",
					component: () => import(/* webpackChunkName: "activity" */"@/pages/Activity.vue"),
					beforeEnter: async (to, from, next) => {
						try {
							router.app.$Progress.start();
							await store.dispatch("activity/loadActivities");
							next();
						} catch (err) {
							router.app.$Progress.fail();
							router.app.$bvToast.toast(`Fetching data problem: ${err.message}`, {
								title: "Fetching activity error",
								variant: "danger",
								autoHideDelay: 1500
							});
						} finally {
							router.app.$Progress.finish();
						}
					}
				},
				{
					name: "activity-detail_id",
					path: "/activity/detail/:id",
					component: () => import(/* webpackChunkName: "activity-detail" */"@/pages/ActivityDetail_id.vue")
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
			path: "/adminsitcompetence",
			beforeEnter() {
				location.href = "http://localhost:8080/admin/activity";
			}
		},
		{
			name: "error404",
			path: "*",
			component: () => import("@/pages/Error404.vue")
		}
	]
});

// router.beforeEach(() => {
// 	app.$Progress.start();
// 	// Ignore login and error page
// 	// const isLogin = sessionStorage.getItem("loginSession");
// 	// if (to.name !== "login" && to.name !== "error404" && !isLogin) {
// 	// 	next({
// 	// 		name: "login"
// 	// 	});
// 	// }
// 	// next();
// });

export default router;