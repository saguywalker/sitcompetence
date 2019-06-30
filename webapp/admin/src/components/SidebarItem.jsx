import React from "react";

function SidebarItem({iconName, itemName, isHover}) {

	return (
		<div className="sidebar-header-item">
			<div className="sidebar-link-item">
				<span className="icon has-text-sb-item">
					<i className={`fas fa-${iconName}`}></i>
				</span>
				{ isHover &&
					<span className="sidebar-link-item-tag tag is-sb-tag">
						{ itemName }
					</span>
				}
			</div>
		</div>
	);
}

export default SidebarItem;