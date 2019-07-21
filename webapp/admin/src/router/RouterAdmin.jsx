import React from "react";
import { Router } from "@reach/router";
import GiveBadgeLayout from "layouts/GiveBadgeLayout";
import GiveBadgeMain from "pages/give-badge/GiveBadgeMain";
import GiveBadgeSelectBadge from "pages/give-badge/GiveBadgeSelectBadge";
import GiveBadgeConfirm from "pages/give-badge/GiveBadgeConfirm";
import Activity from "pages/Activity";
import Dashboard from "pages/Dashboard";

function RouterAdmin() {
	return (
		<Router>
			<Dashboard path="/"/>
			<GiveBadgeLayout path="give-badge">
				<GiveBadgeMain path="/"/>
				<GiveBadgeSelectBadge path="select"/>
				<GiveBadgeConfirm path="confirm"/>
			</GiveBadgeLayout>
			<Activity path="activity"/>
		</Router>
	);
}

export default RouterAdmin;