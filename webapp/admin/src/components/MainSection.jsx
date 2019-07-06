import React from "react";
import RouterAdmin from "router/RouterAdmin";

function MainContent() {
	return (
		<div className="columns is-mobile">
			<div className="column test is-1"/>
			<div className="column test">
				<div className="container">
					<div className="section">
						<RouterAdmin/>
					</div>
				</div>
			</div>
		</div>
	);
}

export default MainContent;