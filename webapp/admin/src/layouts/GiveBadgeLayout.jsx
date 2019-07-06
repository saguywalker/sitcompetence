import React from "react";

function GiveBadge(props) {
	return (
		<div>
			<h1 className="title">Give Badge</h1>
			{props.children}
		</div>
	);
}

export default GiveBadge;