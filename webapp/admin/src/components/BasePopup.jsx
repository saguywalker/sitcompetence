import React from "react";
import { Link } from "@reach/router";
import { GIVE_BADGE_SELECT } from "router/routeName";
import BaseImagePic from "components/BaseImagePic";

function BasePopup({items}) {
	return (
		<div className="base-popup">
			<div className="base-popup-select">
				{
					items.map((item, index) => (
						<div
							className="base-popup-select-item"
							key={`${item}${index}`}
						>
							<BaseImagePic
								size="32"
								round={true}
								imgSrc={item.img}
							/>
							<p className="base-popup-select-item-id">
								{item.stdId}
							</p>
						</div>
					))
				}
			</div>
			<div className="base-popup-next">
				<Link to={GIVE_BADGE_SELECT}>
					<button className="button is-button-main">
						Next
					</button>
				</Link>
			</div>
		</div>
	);
}

export default BasePopup;