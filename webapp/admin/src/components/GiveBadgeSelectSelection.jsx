import React from "react";
import BaseBadge from "components/BaseBadge";

function GiveBadgeSelectSelection() {
	return (
		<div className="gb-selection">
			<div className="gb-selection-wrapper">
				<div className="gb-selection-item">
					<BaseBadge badgeName="Problem-Solving" />
				</div>
				<div className="gb-selection-item">
					<BaseBadge badgeName="Team-working" />
				</div>
				<div className="gb-selection-item">
					<BaseBadge badgeName="Supportive" />
				</div>
				<div className="gb-selection-item">
					<BaseBadge badgeName="Flexible" />
				</div>
				<div className="gb-selection-item">
					<BaseBadge badgeName="Ah Shit" />
				</div>
			</div>
		</div>
	);
}

export default GiveBadgeSelectSelection;