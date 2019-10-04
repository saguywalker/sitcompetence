import Vue from "vue";
import Router from "vue-router";
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
					name: "activity-detail",
					path: "activity/detail/:id",
					component: () => import("@/pages/activity/detail_id.vue")
				},
				{
					name: "activity-approve",
					path: "activity/approve/:id",
					component: () => import("@/pages/activity/approve_id.vue")
				},
				{
					name: "activity-approve-success",
					path: "activity/approve/:id/success",
					component: () => import("@/pages/activity/approve-success.vue")
				},
				{
					name: "past-activity",
					path: "activity/past",
					component: () => import("@/pages/activity/past-activity.vue")
				},
				{
					path: "activity/create",
					component: () => import("@/layouts/LayoutCreateEditActivity.vue"),
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
							path: "select-competence",
							component: () => import("@/pages/activity/create/select-competence.vue"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.competence,
								step: CREATE_ACTIVITY_STEP.competence
							}
						},
						{
							name: "create-activity-summary",
							path: "summary",
							component: () => import("@/pages/activity/create/summary.vue"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.summary,
								step: CREATE_ACTIVITY_STEP.summary
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
				{ // Use breadcrumb and step same as CREATE_ACTIVITY
					path: "activity/edit/:id",
					component: () => import("@/layouts/LayoutCreateEditActivity.vue"),
					children: [
						{
							name: "edit-activity",
							path: "/",
							component: () => import("@/pages/activity/edit"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.detail,
								step: EDIT_ACTIVITY_STEP.detail
							}
						},
						{
							name: "edit-activity-competence",
							path: "select-competence",
							component: () => import("@/pages/activity/edit/select-competence.vue"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.competence,
								step: EDIT_ACTIVITY_STEP.competence
							}
						},
						{
							name: "edit-activity-summary",
							path: "summary",
							component: () => import("@/pages/activity/edit/summary.vue"),
							meta: {
								breadcrumb: CREATE_ACTIVITY_BREADCRUMB.summary,
								step: EDIT_ACTIVITY_STEP.summary
							}
						},
						{
							name: "edit-activity-success",
							path: "success",
							component: () => import("@/pages/activity/edit/success.vue"),
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
				},
				{
					name: "verify-result",
					path: "verify/result",
					component: () => import("@/pages/verify/result.vue")
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