import Vue from "vue";
import Router from "vue-router";
import store from "@/store";
import {
	GIVE_BADGE_BREADCRUMB,
	CREATE_ACTIVITY_BREADCRUMB,
	GIVE_BADGE_STEP,
	CREATE_ACTIVITY_STEP,
	EDIT_ACTIVITY_STEP,
	STUDENT_ROUTE_NAMES
} from "@/constants";
import { getLoginUserRole, clearLoginState, isLoggedIn } from "@/helpers";

Vue.use(Router);

const router = new Router({
	mode: "history",
	routes: [
		{
			path: "/login",
			name: "login",
			component: () => import("@/pages/login.vue"),
			meta: {
				rule: "isPublic"
			},
			beforeEnter: (to, from, next) => {
				// Prevent user go back to Login page when already logged in
				const isLogin = isLoggedIn();
				const role = getLoginUserRole();

				if (isLogin && to.name === "login") {
					if (role === "inst_group") {
						next({ name: "admin" });
					} else {
						next({ name: "student" });
					}

					return;
				}

				clearLoginState();
				next();
			}
		},
		{
			path: "/viewProfile",
			name: "share-portfolio",
			component: () => import("@/pages/_id-portfolio.vue"),
			meta: {
				rule: "isPublic"
			}
		},
		{
			path: "/",
			component: () => import("@/layouts/StudentLayout.vue"),
			meta: {
				rule: "isStudent"
			},
			children: [
				{
					path: "/",
					name: "student",
					redirect: { name: "dashboard" },
					meta: {
						rule: "isStudent"
					}
				},
				{
					name: "dashboard",
					path: "/dashboard",
					component: () => import("@/pages/student/dashboard.vue"),
					meta: {
						rule: "isStudent"
					}
				},
				{
					name: "activity",
					path: "/activity",
					component: () => import("@/pages/student/activity"),
					meta: {
						rule: "isStudent"
					},
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
					component: () => import("@/pages/student/activity/detail.vue"),
					meta: {
						rule: "isStudent"
					}
				},
				{
					name: "portfolio",
					path: "/portfolio",
					component: () => import("@/pages/student/portfolio"),
					meta: {
						rule: "isStudent"
					}
				}
			]
		},
		{
			path: "/admin",
			component: () => import("@/layouts/AdminLayoutDefault.vue"),
			meta: {
				rule: "isAdmin"
			},
			children: [
				{
					name: "admin",
					path: "/",
					redirect: { name: "give-badge" },
					meta: {
						rule: "isAdmin"
					}
				},
				{
					name: "admin-activity",
					path: "activity",
					component: () => import("@/pages/admin/activity"),
					meta: {
						rule: "isAdmin"
					}
				},
				{
					name: "activity-detail",
					path: "activity/detail/:id",
					component: () => import("@/pages/admin/activity/detail_id.vue"),
					meta: {
						rule: "isAdmin"
					}
				},
				{
					name: "activity-approve",
					path: "activity/approve/:id",
					component: () => import("@/pages/admin/activity/approve_id.vue"),
					meta: {
						rule: "isAdmin"
					}
				},
				{
					name: "activity-approve-success",
					path: "activity/approve/:id/success",
					component: () => import("@/pages/admin/activity/approve-success.vue"),
					meta: {
						rule: "isAdmin"
					}
				},
				{
					name: "past-activity",
					path: "activity/past",
					component: () => import("@/pages/admin/activity/past-activity.vue"),
					meta: {
						rule: "isAdmin"
					}
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
								rule: "isAdmin",
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.detail,
								step: CREATE_ACTIVITY_STEP.detail
							}
						},
						{
							name: "create-activity-competence",
							path: "select-competence",
							component: () => import("@/pages/admin/activity/create/select-competence.vue"),
							meta: {
								rule: "isAdmin",
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.competence,
								step: CREATE_ACTIVITY_STEP.competence
							}
						},
						{
							name: "create-activity-summary",
							path: "summary",
							component: () => import("@/pages/admin/activity/create/summary.vue"),
							meta: {
								rule: "isAdmin",
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.summary,
								step: CREATE_ACTIVITY_STEP.summary
							}
						},
						{
							name: "create-activity-success",
							path: "success",
							component: () => import("@/pages/admin/activity/create/success.vue"),
							meta: {
								rule: "isAdmin",
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
								rule: "isAdmin",
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.detail,
								step: EDIT_ACTIVITY_STEP.detail
							}
						},
						{
							name: "edit-activity-competence",
							path: "select-competence",
							component: () => import("@/pages/admin/activity/edit/select-competence.vue"),
							meta: {
								rule: "isAdmin",
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.competence,
								step: EDIT_ACTIVITY_STEP.competence
							}
						},
						{
							name: "edit-activity-summary",
							path: "summary",
							component: () => import("@/pages/admin/activity/edit/summary.vue"),
							meta: {
								rule: "isAdmin",
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.summary,
								step: EDIT_ACTIVITY_STEP.summary
							}
						},
						{
							name: "edit-activity-success",
							path: "success",
							component: () => import("@/pages/admin/activity/edit/success.vue"),
							meta: {
								rule: "isAdmin",
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.success,
								step: EDIT_ACTIVITY_STEP.success
							}
						}
					]
				},
				{
					name: "badge-setting",
					path: "badge-setting",
					component: () => import("@/pages/admin/badge-setting"),
					meta: {
						rule: "isAdmin"
					}
				},
				{
					name: "user-setting",
					path: "user/setting",
					component: () => import("@/pages/admin/user/setting.vue"),
					meta: {
						rule: "isAdmin"
					}
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
								rule: "isAdmin",
								breadcrumb: GIVE_BADGE_BREADCRUMB.main,
								step: GIVE_BADGE_STEP.main
							}
						},
						{
							name: "give-badge-selection",
							path: "selection",
							component: () => import("@/pages/admin/give-badge/selection.vue"),
							meta: {
								rule: "isAdmin",
								breadcrumb: GIVE_BADGE_BREADCRUMB.selection,
								step: GIVE_BADGE_STEP.selection
							}
						},
						{
							name: "give-badge-confirmation",
							path: "confirmation",
							component: () => import("@/pages/admin/give-badge/confirmation.vue"),
							meta: {
								rule: "isAdmin",
								breadcrumb: GIVE_BADGE_BREADCRUMB.confirmation,
								step: GIVE_BADGE_STEP.confirmation
							}
						},
						{
							name: "give-badge-success",
							path: "success",
							component: () => import("@/pages/admin/give-badge/success.vue"),
							meta: {
								rule: "isAdmin",
								breadcrumb: GIVE_BADGE_BREADCRUMB.success
							}
						}
					]
				},
				{
					name: "verify",
					path: "verify",
					component: () => import("@/pages/admin/verify"),
					meta: {
						rule: "isAdmin"
					}
				},
				{
					name: "verify-result",
					path: "verify/result",
					component: () => import("@/pages/admin/verify/result.vue"),
					meta: {
						rule: "isAdmin"
					}
				}
			]
		},
		{
			name: "user-genkey",
			path: "/admin/user/genkey",
			component: () => import("@/pages/admin/user/genkey.vue"),
			meta: {
				rule: "isAdmin"
			}
		},
		{
			name: "error404",
			path: "*",
			component: () => import("@/pages/error/404.vue"),
			meta: {
				rule: "isPublic"
			}
		},
		{
			name: "unknown",
			path: "/unknown",
			component: () => import("@/pages/error/404.vue"),
			meta: {
				rule: "isPublic"
			}
		},
		{
			name: "error403",
			path: "/notfound",
			component: () => import("@/pages/error/403.vue"),
			meta: {
				rule: "isPublic"
			},
			beforeEnter: (to, from, next) => {
				// Force user to login again when he/she try to access without authentication
				clearLoginState();
				next();
			}
		}
	]
});

router.beforeEach((to, from, next) => {
	// Ignore login and error page
	const isLogin = isLoggedIn();
	const role = getLoginUserRole();

	if (to.name !== "login" && to.name !== "error404" && to.name !== "share-portfolio" && !isLogin) {
		next({ name: "login" });
	} else if (role === "inst_group" && STUDENT_ROUTE_NAMES.includes(to.name)) {
		next({ name: "admin" });
	} else {
		next();
	}
});

export default router;