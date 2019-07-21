import React from "react";
import SidebarItemLink from "components/SidebarItemLink";
import logo from "assets/images/sitcom_logo.png";
import {
	GIVE_BADGE_MAIN,
	ACTIVITY_MAIN
} from "router/routeName";

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
								<SidebarItemLink
									iconName="search"
									itemName="Search"
									pathName="/"
								/>
							</div>
						</div>
					</div>
					<div className="sidebar-main-container">
						<div className="sidebar-main-wrapper">
							<SidebarItemLink
								iconName="hand-holding-usd"
								itemName="Giving"
								pathName={GIVE_BADGE_MAIN}
							/>
							<SidebarItemLink
								iconName="cubes"
								itemName="Activity"
								pathName={ACTIVITY_MAIN}
							/>
							<SidebarItemLink
								iconName="certificate"
								itemName="Badges"
								pathName="/badges"
							/>
						</div>
					</div>
					<div className="sidebar-footer-container">
						<SidebarItemLink
							iconName="cog"
							itemName="Setting"
							pathName="/activity"
						/>
						<SidebarItemLink
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