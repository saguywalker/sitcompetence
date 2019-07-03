import React from "react";
import { Router } from "@reach/router";
import GiveBadge from "pages/GiveBadge";
import Activity from "pages/Activity";
import Dashboard from "pages/Dashboard";

function MainContent() {
	return (
		<div className="columns is-mobile">
			<div className="column test is-1"/>
			<div className="column test">
				<div className="container">
					<div className="section">
						<Router>
							<Dashboard path="/"/>
							<GiveBadge path="/give"/>
							<Activity path="/activity"/>
						</Router>
					</div>
				</div>
			</div>
		</div>
	);
}

export default MainContent;