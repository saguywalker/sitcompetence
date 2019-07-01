import React from "react";
import { useHover } from "../helpers/hooks";

function SidebarItem({iconName, itemName}) {
	const [hoverElementRef, isHoveringRef] = useHover();

	return (
		<div
			className="sidebar-link-item"
			ref={hoverElementRef}
		>
			<span className="icon has-text-sb-item">
				<i className={`fas fa-${iconName}`}></i>
			</span>
			{ isHoveringRef && <span className="sidebar-link-item-tag tag is-sb-tag">{itemName}</span> }
		</div>
	);
}

export default SidebarItem;