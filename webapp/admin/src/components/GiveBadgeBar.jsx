import React, { useState } from "react";
import Dropdown from "components/GiveBadgeDropdownSearchFilter";

function GiveBadgeSearchBar() {
	const [selectedFilter, setSelectedFilter] = useState([]);

	function handleOnSelect(e) {
		let tmp = selectedFilter;
		let index;
		if (e.target.checked) {
			// Add the value of the checkbox to array
			tmp.push(e.target.value);
			setSelectedFilter(tmp);
		} else {
		// Remove the value from the unchecked checkbox from the array
			index = tmp.indexOf(e.target.value);
			tmp.splice(index, 1);
		}
	}

	return (
		<>
			<nav className="level">
				<div className="level-left">
					<div className="level-item">
						<h1 className="title">Give Badge</h1>
					</div>
				</div>
				<div className="level-right">
					<div className="gb-main-search level-item">
						<div className="field has-addons">
							<div className="control">
								<Dropdown
									items={["a", "b", "c"]}
									onSelect={handleOnSelect}
								/>
							</div>
							<div className="control">
								<input type="text" className="input"/>
							</div>
							<div className="control">
								<button className="button">
									Search
								</button>
							</div>
						</div>
						<div className="gb-main-filter-tag">
							<div className="gb-main-filter-tag-item tag is-my-blue">
								Bar
								<button className="delete is-small" />
							</div>
						</div>
					</div>
				</div>
			</nav>
		</>
	);
}

export default GiveBadgeSearchBar;