import Vue from "vue";
import Router from "vue-router";
import { GIVE_BADGE_BREADCRUMB } from "@/constants/breadcrumb";
import { GIVE_BADGE_STEP } from "@/constants/step";

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
					component: () => import("@/pages/index.vue")
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
					path: "hello",
					component: () => import("@/pages/hello"),
					children: [
						{
							name: "hello_id",
							path: "/:id",
							component: () => import("@/pages/hello/_id.vue")
						}
					]
				}
			]
		}
	]
});

export default router;