import React from "react";
import GiveBadgeSearchBar from "components/GiveBadgeBar";
import BaseTable from "components/BaseTable";
import BasePopup from "components/BasePopup";

function GiveBadgeMain() {
	return (
		<>
			<div className="section">
				<GiveBadgeSearchBar />
				<BaseTable />
			</div>
			<BasePopup />
		</>
	);
}

export default GiveBadgeMain;