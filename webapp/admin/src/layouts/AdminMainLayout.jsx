import React from "react";
import Sidebar from "components/Sidebar";
import RouterAdmin from "router/RouterAdmin";

function AdminMainLayout() {
	return (
		<div className="columns is-gapless is-mobile">
			<div className="column is-narrow admin-sidebar">
				<Sidebar />
			</div>
			<div className="column">
				<RouterAdmin/>
			</div>
		</div>
	);
}

export default AdminMainLayout;