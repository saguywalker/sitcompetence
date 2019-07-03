import React from "react";
import { useHover } from "helpers/hooks";
import { Link } from "@reach/router";

function SidebarItem({iconName, itemName, pathName}) {
	const [hoverElementRef, isHoveringRef] = useHover();

	return (
		<Link
			className="sidebar-link-item"
			to={pathName}
			ref={hoverElementRef}
		>
			<span className="icon has-text-sb-item">
				<i className={`fas fa-${iconName}`}></i>
			</span>
			{ isHoveringRef && <span className="sidebar-link-item-tag tag is-sb-tag">{itemName}</span> }
		</Link>
	);
}

export default SidebarItem;