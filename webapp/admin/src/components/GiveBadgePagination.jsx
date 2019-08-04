import React from "react";
import { Link } from "@reach/router";

function GiveBadgePaginationBar({nextPage, prevPage, nextText, prevText}) {
	return (
		<div className="gb-pagination">
			<Link to={prevPage}>
				<button className="button">{prevText}</button>
			</Link>
			<Link to={nextPage}>
				<button className="button is-button-main">{nextText}</button>
			</Link>
		</div>
	);
}

export default GiveBadgePaginationBar;