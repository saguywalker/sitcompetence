import React from "react";
import SidebarItem from "../components/SidebarItem";
import logo from "../assets/images/sitcom_logo.png";

function Sidebar() {
	return (
		<div className="sidebar">
			<div className="sidebar-fixed-container">
				<div className="sidebar-inner-content">
					<div className="sidebar-header-container">
						<div className="sidebar-header-wrapper">
							<div className="sidebar-header-primary-item">
								<img className="sidebar-sitcom-logo" src={logo} alt="SIT-Competence"/>
							</div>
							<SidebarItem
								iconName="search"
								itemName="Search"
							/>
						</div>
					</div>
					<div className="sidebar-main-container">
						<div className="sidebar-main-wrapper">
							<SidebarItem
								iconName="hand-holding-usd"
								itemName="Giving"
							/>
							<SidebarItem
								iconName="cubes"
								itemName="Activity"
							/>
							<SidebarItem
								iconName="certificate"
								itemName="Badges"
							/>
						</div>
					</div>
					<div className="sidebar-footer-container">
						<SidebarItem
							iconName="cog"
							itemName="Setting"
						/>
						<SidebarItem
							iconName="sign-out-alt"
							itemName="Logout"
						/>
					</div>
				</div>
			</div>
		</div>
	);
}

export default Sidebar;