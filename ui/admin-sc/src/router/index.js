import Vue from "vue";
import Router from "vue-router";
import {
	GIVE_BADGE_BREADCRUMB,
	CREATE_ACTIVITY_BREADCRUMB
} from "@/constants/breadcrumb";
import {
	GIVE_BADGE_STEP,
	CREATE_ACTIVITY_STEP
} from "@/constants/step";

Vue.use(Router);

const router = new Router({
	mode: "history",
	routes: [
		{
			path: "/admin",
			component: () => import("@/layouts/LayoutAdminDefault.vue"),
			children: [
				{
					name: "admin",
					path: "/",
					redirect: { name: "dashboard" }
				},
				{
					name: "dashboard",
					path: "dashboard",
					component: () => import("@/pages/dashboard.vue")
				},
				{
					name: "activity",
					path: "activity",
					component: () => import("@/pages/activity")
				},
				{
					name: "activity_id",
					path: "activity/d/:id",
					component: () => import("@/pages/activity/activity_id.vue")
				},
				{
					path: "activity/create",
					component: () => import("@/layouts/LayoutCreateActivity.vue"),
					children: [
						{
							name: "create-activity",
							path: "/",
							component: () => import("@/pages/activity/create"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.detail,
								step: CREATE_ACTIVITY_STEP.detail
							}
						},
						{
							name: "create-activity-competence",
							path: "competence",
							component: () => import("@/pages/activity/create/competence.vue"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.competence,
								step: CREATE_ACTIVITY_STEP.competence
							}
						},
						{
							name: "create-activity-confirmation",
							path: "confirmation",
							component: () => import("@/pages/activity/create/confirmation.vue"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.confirmation,
								step: CREATE_ACTIVITY_STEP.confirmation
							}
						},
						{
							name: "create-activity-success",
							path: "success",
							component: () => import("@/pages/activity/create/success.vue"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.success,
								step: CREATE_ACTIVITY_STEP.success
							}
						}
					]
				},
				{
					name: "badge-setting",
					path: "badge-setting",
					component: () => import("@/pages/badge-setting")
				},
				{
					path: "give-badge",
					component: () => import("@/layouts/LayoutGiveBadge.vue"),
					children: [
						{
							name: "give-badge",
							path: "/",
							component: () => import("@/pages/give-badge"),
							meta: {
								breadcrumb: GIVE_BADGE_BREADCRUMB.main,
								step: GIVE_BADGE_STEP.main
							}
						},
						{
							name: "give-badge-selection",
							path: "selection",
							component: () => import("@/pages/give-badge/selection.vue"),
							meta: {
								breadcrumb: GIVE_BADGE_BREADCRUMB.selection,
								step: GIVE_BADGE_STEP.selection
							}
						},
						{
							name: "give-badge-confirmation",
							path: "confirmation",
							component: () => import("@/pages/give-badge/confirmation.vue"),
							meta: {
								breadcrumb: GIVE_BADGE_BREADCRUMB.confirmation,
								step: GIVE_BADGE_STEP.confirmation
							}
						},
						{
							name: "give-badge-success",
							path: "success",
							component: () => import("@/pages/give-badge/success.vue"),
							meta: {
								breadcrumb: GIVE_BADGE_BREADCRUMB.success
							}
						}
					]
				},
				{
					name: "verify",
					path: "verify",
					component: () => import("@/pages/verify")
				}
			]
		},
		{
			path: "*",
			name: "error404",
			component: () => import("@/pages/error/error404")
		}
	]
});

export default router;