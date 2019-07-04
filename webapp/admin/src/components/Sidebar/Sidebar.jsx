import React from "react";
import SidebarItem from "./SidebarItem";
import logo from "assets/images/sitcom_logo.png";

function Sidebar() {
	return (
		<div className="sidebar">
			<div className="sidebar-fixed-container">
				<div className="sidebar-inner-content">
					<div className="sidebar-header-container">
						<div className="sidebar-header-wrapper">
							<div className="sidebar-header-primary-item">
								<img
									className="sidebar-sitcom-logo"
									src={logo}
									alt="SIT-Competence"
								/>
							</div>
							<div className="sidebar-header-item">
								<SidebarItem
									iconName="search"
									itemName="Search"
									pathName="/"
								/>
							</div>
						</div>
					</div>
					<div className="sidebar-main-container">
						<div className="sidebar-main-wrapper">
							<SidebarItem
								iconName="hand-holding-usd"
								itemName="Giving"
								pathName="/give"
							/>
							<SidebarItem
								iconName="cubes"
								itemName="Activity"
								pathName="/activity"
							/>
							<SidebarItem
								iconName="certificate"
								itemName="Badges"
								pathName="/activity"
							/>
						</div>
					</div>
					<div className="sidebar-footer-container">
						<SidebarItem
							iconName="cog"
							itemName="Setting"
							pathName="/activity"
						/>
						<SidebarItem
							iconName="sign-out-alt"
							itemName="Logout"
							pathName="/activity"
						/>
					</div>
				</div>
			</div>
		</div>
	);
}

export default Sidebar;