import React from "react";
import { useHover, useWindowWidth } from "helpers/hooks";
import { Link } from "@reach/router";

function SidebarItem({iconName, itemName, pathName}) {
	const [hoverElementRef, isHoveringRef] = useHover();
	const width = useWindowWidth();

	return (
		<Link
			className="sidebar-link-item"
			to={pathName}
			ref={hoverElementRef}
		>
			<span className="icon has-text-sb-item">
				<i className={`fas fa-${iconName}`}></i>
			</span>
			{ (isHoveringRef && (width > 767)) && <span className="sidebar-link-item-tag tag is-sb-tag">{itemName}</span> }
		</Link>
	);
}

export default SidebarItem;