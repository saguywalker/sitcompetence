import React from "react";
import Sidebar from "components/Sidebar";
import Navbar from "components/Navbar";
// import RouterAdmin from "router/RouterAdmin";

function AdminMainLayout() {
	return (
		<div className="columns">
			<div className="column">
				<Sidebar />
				<Navbar />
			</div>
			{/* <div className="column is-narrow admin-sidebar">
			</div> */}
			{/* <div className="column">
				<RouterAdmin/>
			</div> */}
		</div>
	);
}

export default AdminMainLayout;