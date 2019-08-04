import React from "react";
import GiveBadgeSearchBar from "components/GiveBadgeBar";
import BaseTable from "components/BaseTable";

function GiveBadgeMain() {
	return (
		<>
			<div className="section">
				<GiveBadgeSearchBar />
				<BaseTable />
			</div>
		</>
	);
}

export default GiveBadgeMain;