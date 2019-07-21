import React from "react";
import BaseBadge from "components/BaseBadge";

function GiveBadgeSelection() {
	return (
		<div className="gb-selection">
			<div className="gb-selection-wrapper">
				<div className="gb-selection-item">
					<BaseBadge />
				</div>
				<div className="gb-selection-item">
					<BaseBadge />
				</div>
				<div className="gb-selection-item">
					<BaseBadge />
				</div>
				<div className="gb-selection-item">
					<BaseBadge />
				</div>
				<div className="gb-selection-item">
					<BaseBadge />
				</div>
			</div>
		</div>
	);
}

export default GiveBadgeSelection;