import React from "react";
import Sidebar from "components/Sidebar/Sidebar";
import MainContent from "components/MainContent";

function AdminMainLayout() {
	return (
		<>
			<Sidebar/>
			<MainContent/>
		</>
	);
}

export default AdminMainLayout;