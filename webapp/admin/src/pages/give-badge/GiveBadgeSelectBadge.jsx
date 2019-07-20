import React from "react";
import GiveBadgeStudentDescription from "components/GiveBadgeStudentDescription";
import GiveBadgeSelection from "components/GiveBadgeSelection";
import { Link } from "@reach/router";
import {
	GIVE_BADGE_MAIN,
	GIVE_BADGE_CONFIRM
} from "router/routeName";

function GiveBadgeSelectBadge() {
	return (
		<div className="container">
			<div className="section">
				<h1 className="title">Select Badge</h1>
				<GiveBadgeStudentDescription />
				<GiveBadgeSelection />
			</div>
			<div className="gb-select-pagination-bar">
				<Link to={GIVE_BADGE_MAIN}>
					<button className="button is-button-main">Back</button>
				</Link>
				<Link to={GIVE_BADGE_CONFIRM}>
					<button className="button is-button-main">Next</button>
				</Link>
			</div>
		</div>
	);
}

export default GiveBadgeSelectBadge;