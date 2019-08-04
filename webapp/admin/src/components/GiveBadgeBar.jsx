import React, { useState } from "react";
import Dropdown from "components/GiveBadgeDropdownSearchFilter";

function GiveBadgeSearchBar() {
	const [selectedFilters, setSelectedFilter] = useState([]);

	function handleOnSelect(e) {
		const selectItem = e.target.value;
		if (e.target.checked) {
			setSelectedFilter([
				...selectedFilters,
				selectItem
			]);
		} else {
			// Remove the value from the unchecked checkbox from the array
			const index = selectedFilters.indexOf(selectItem);
			selectedFilters.splice(index, 1);
			setSelectedFilter([
				...selectedFilters
			]);
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
					<div className="gb-bar-search level-item">
						<div className="field has-addons">
							<div className="control">
								<Dropdown
									items={["Year", "Month", "Semester"]}
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
						{
							selectedFilters.length > 0 &&
							<div className="gb-bar-filter-tag">
								{
									selectedFilters.map((filter, index) => (
										<div
											className="gb-bar-filter-tag-item tag is-my-blue"
											key={`${filter}${index}`}
										>
											{filter}
										</div>
									))
								}
							</div>
						}
					</div>
				</div>
			</nav>
		</>
	);
}

export default GiveBadgeSearchBar;