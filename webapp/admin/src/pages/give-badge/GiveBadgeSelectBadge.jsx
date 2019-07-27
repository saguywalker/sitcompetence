import React from "react";
import GiveBadgeStudentDescription from "components/GiveBadgeStudentDescription";
import GiveBadgeSelectSelection from "components/GiveBadgeSelectSelection";
import GiveBadgePagination from "components/GiveBadgePagination";
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
				<GiveBadgeSelectSelection />
				<hr/>
				<GiveBadgeStudentDescription />
				<GiveBadgeSelectSelection />
				<hr/>
			</div>
			<GiveBadgePagination
				nextPage={GIVE_BADGE_CONFIRM}
				prevPage={GIVE_BADGE_MAIN}
				nextText="Next"
				prevText="Back"
			/>
		</div>
	);
}

export default GiveBadgeSelectBadge;