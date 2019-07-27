import React from "react";
import BaseImagePic from "components/BaseImagePic";
import GiveBadgePagination from "components/GiveBadgePagination";
import BaseBadge from "components/BaseBadge";
import {
	GIVE_BADGE_SELECT,
	GIVE_BADGE_SUCCESS
} from "router/routeName";

function GiveBadgeConfirm() {
	return (
		<div className="container">
			<div className="section">
				<h1 className="title">Confirmation</h1>
				<div className="gb-student-description">
					<figure>
						<BaseImagePic />
					</figure>
					<div className="gb-student-description-content">
						<h1 className="title is-5">Mr. Tindanai Wongpipattanopas</h1>
						<p className="subtitle is-6">
							59130500210
							<p>Computer Science</p>
						</p>
					</div>
					<div className="gb-student-description-badge">
						<BaseBadge
							badgeName="Het"
							size="confirm"
						/>
						<BaseBadge
							badgeName="Het"
							size="confirm"
						/>
						<BaseBadge
							badgeName="Het"
							size="confirm"
						/>
						<BaseBadge
							badgeName="Het"
							size="confirm"
						/>
						<BaseBadge
							badgeName="Het"
							size="confirm"
						/>
					</div>
				</div>
				<hr />
				<GiveBadgePagination
					nextPage={GIVE_BADGE_SUCCESS}
					prevPage={GIVE_BADGE_SELECT}
					nextText="Confirm"
					prevText="Back"
				/>
			</div>
		</div>
	);
}

export default GiveBadgeConfirm;