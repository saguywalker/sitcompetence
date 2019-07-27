import React from "react";

function BaseBadge({badgeName, size}) {
	function baseBadgeSize(size) {
		switch (size) {
		case "small":
			return "base-badge-sm";
		case "medium":
			return "base-badge-md";
		case "confirm":
			return "base-badge-confirm";
		default:
			return "base-badge";
		}
	}

	return (
		<div className={baseBadgeSize(size)}>
			<div className="wrapper">
				<img
					src="https://bulma.io/images/placeholders/128x128.png"
					alt="hi"
				/>
			</div>
			{
				badgeName &&
				<p className="badge-name">{badgeName}</p>
			}
		</div>
	);
}

export default BaseBadge;