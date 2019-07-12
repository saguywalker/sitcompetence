import React from "react";
import { Link } from "@reach/router";
import { GIVE_BADGE_SELECT } from "router/routeName";

function BasePopup() {
	return (
		<div className="base-popup">
			<div className="base-popup-wrapper">
				<div>
					Select card componentsvfvfdvdfvfdvd
				</div>
				<Link to={GIVE_BADGE_SELECT}>
					<button className="button is-button-main base-popup-next-btn">Next</button>
				</Link>
			</div>
		</div>
	);
}

export default BasePopup;