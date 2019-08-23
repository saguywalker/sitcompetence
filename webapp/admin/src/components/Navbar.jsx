import React from "react";

function Navbar() {
	return (
		<nav className="navbar">
			<div className="navbar-hamburger">
				<span className="icon has-text-sb-item">
					<i className="fas fa-bars"></i>
				</span>
			</div>
			<div className="navbar-menu">
				<h2 className="navbar-menu-user">Tindanai W.</h2>
			</div>
		</nav>
	);
}

export default Navbar;