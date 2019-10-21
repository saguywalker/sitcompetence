import Vue from "vue";
import Router from "vue-router";
import store from "@/store";
import {
	GIVE_BADGE_BREADCRUMB,
	CREATE_ACTIVITY_BREADCRUMB,
	GIVE_BADGE_STEP,
	CREATE_ACTIVITY_STEP,
	EDIT_ACTIVITY_STEP
} from "@/constants";

Vue.use(Router);

const router = new Router({
	mode: "history",
	routes: [
		{
			path: "/login",
			name: "login",
			component: () => import(/* webpackChunkName: "login" */ "@/pages/login.vue")
		},
		{
			path: "/user/:id/portfolio",
			name: "user_id-portfoilio",
			component: () => import(/* webpackChunkName: "user-portfoilio" */ "@/pages/student/user/_id-portfolio.vue")
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
					component: () => import(/* webpackChunkName: "dashboard" */"@/pages/student/dashboard.vue")
				},
				{
					name: "activity",
					path: "/activity",
					component: () => import(/* webpackChunkName: "activity" */"@/pages/student/activity"),
					beforeEnter: async (to, from, next) => {
						try {
							router.app.$Progress.start();
							await store.dispatch("activity/loadActivities");
						} catch (err) {
							router.app.$Progress.fail();
							router.app.$bvToast.toast(`Fetching data problem: ${err.message}`, {
								title: "Fetching activity error",
								variant: "danger",
								autoHideDelay: 1500
							});
						} finally {
							router.app.$Progress.finish();
							next();
						}
					}
				},
				{
					name: "activity-detail_id",
					path: "/activity/detail/:id",
					component: () => import(/* webpackChunkName: "activity-detail" */"@/pages/student/activity/detail.vue")
				},
				{
					name: "portfolio",
					path: "/portfolio",
					component: () => import(/* webpackChunkName: "portfolio" */"@/pages/student/portfolio")
				}
			]
		},
		{
			path: "/admin",
			component: () => import("@/layouts/AdminLayoutDefault.vue"),
			children: [
				{
					name: "admin",
					path: "/",
					redirect: { name: "give-badge" }
				},
				{
					name: "admin-activity",
					path: "activity",
					component: () => import("@/pages/admin/activity")
				},
				{
					name: "activity-detail",
					path: "activity/detail/:id",
					component: () => import("@/pages/admin/activity/detail_id.vue")
				},
				{
					name: "activity-approve",
					path: "activity/approve/:id",
					component: () => import("@/pages/admin/activity/approve_id.vue")
				},
				{
					name: "activity-approve-success",
					path: "activity/approve/:id/success",
					component: () => import("@/pages/admin/activity/approve-success.vue")
				},
				{
					name: "past-activity",
					path: "activity/past",
					component: () => import("@/pages/admin/activity/past-activity.vue")
				},
				{
					path: "activity/create",
					component: () => import("@/layouts/AdminLayoutCreateEditActivity.vue"),
					children: [
						{
							name: "create-activity",
							path: "/",
							component: () => import("@/pages/admin/activity/create"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.detail,
								step: CREATE_ACTIVITY_STEP.detail
							}
						},
						{
							name: "create-activity-competence",
							path: "select-competence",
							component: () => import("@/pages/admin/activity/create/select-competence.vue"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.competence,
								step: CREATE_ACTIVITY_STEP.competence
							}
						},
						{
							name: "create-activity-summary",
							path: "summary",
							component: () => import("@/pages/admin/activity/create/summary.vue"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.summary,
								step: CREATE_ACTIVITY_STEP.summary
							}
						},
						{
							name: "create-activity-success",
							path: "success",
							component: () => import("@/pages/admin/activity/create/success.vue"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.success,
								step: CREATE_ACTIVITY_STEP.success
							}
						}
					]
				},
				{ // Use breadcrumb and step same as CREATE_ACTIVITY
					path: "activity/edit/:id",
					component: () => import("@/layouts/AdminLayoutCreateEditActivity.vue"),
					children: [
						{
							name: "edit-activity",
							path: "/",
							component: () => import("@/pages/admin/activity/edit"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.detail,
								step: EDIT_ACTIVITY_STEP.detail
							}
						},
						{
							name: "edit-activity-competence",
							path: "select-competence",
							component: () => import("@/pages/admin/activity/edit/select-competence.vue"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.competence,
								step: EDIT_ACTIVITY_STEP.competence
							}
						},
						{
							name: "edit-activity-summary",
							path: "summary",
							component: () => import("@/pages/admin/activity/edit/summary.vue"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.summary,
								step: EDIT_ACTIVITY_STEP.summary
							}
						},
						{
							name: "edit-activity-success",
							path: "success",
							component: () => import("@/pages/admin/activity/edit/success.vue"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.success,
								step: EDIT_ACTIVITY_STEP.success
							}
						}
					]
				},
				{
					name: "badge-setting",
					path: "badge-setting",
					component: () => import("@/pages/admin/badge-setting")
				},
				{
					path: "give-badge",
					component: () => import("@/layouts/AdminLayoutGiveBadge.vue"),
					children: [
						{
							name: "give-badge",
							path: "/",
							component: () => import("@/pages/admin/give-badge"),
							meta: {
								breadcrumb: GIVE_BADGE_BREADCRUMB.main,
								step: GIVE_BADGE_STEP.main
							}
						},
						{
							name: "give-badge-selection",
							path: "selection",
							component: () => import("@/pages/admin/give-badge/selection.vue"),
							meta: {
								breadcrumb: GIVE_BADGE_BREADCRUMB.selection,
								step: GIVE_BADGE_STEP.selection
							}
						},
						{
							name: "give-badge-confirmation",
							path: "confirmation",
							component: () => import("@/pages/admin/give-badge/confirmation.vue"),
							meta: {
								breadcrumb: GIVE_BADGE_BREADCRUMB.confirmation,
								step: GIVE_BADGE_STEP.confirmation
							}
						},
						{
							name: "give-badge-success",
							path: "success",
							component: () => import("@/pages/admin/give-badge/success.vue"),
							meta: {
								breadcrumb: GIVE_BADGE_BREADCRUMB.success
							}
						}
					]
				},
				{
					name: "setting",
					path: "setting",
					component: () => import("@/pages/admin/setting")
				},
				{
					name: "verify",
					path: "verify",
					component: () => import("@/pages/admin/verify")
				},
				{
					name: "verify-result",
					path: "verify/result",
					component: () => import("@/pages/admin/verify/result.vue")
				}
			]
		},
		{
			name: "error404",
			path: "*",
			component: () => import("@/pages/error/404.vue")
		}
	]
});

router.beforeEach((to, from, next) => {
	// Ignore login and error page
	const isLogin = sessionStorage.getItem("inlog");
	if (to.name !== "login" && to.name !== "error404" && !isLogin) {
		next({
			name: "login"
		});
	}
	next();
});

export default router;