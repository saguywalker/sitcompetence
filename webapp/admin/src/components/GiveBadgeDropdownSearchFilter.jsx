import React, { useState } from "react";

function BaseDropdown({items, onSelect}) {
	const [dropdown, setDropdown] = useState(false);

	function handleDropDownItemClick(e) {
		onSelect(e);
	}

	function handleDropdown() {
		setDropdown(!dropdown);
	}

	return (
		<div className="dropdown is-active">
			<div>
				<div
					className="dropdown-trigger"
					onBlur={() => setDropdown(false)}
				>
					<button
						className="button"
						onClick={handleDropdown}
					>
						<span>Filter</span>
						<span className="icon is-small">
							<i className="fas fa-angle-down" />
						</span>
					</button>
					{
						dropdown &&
						<div className="dropdown-menu">
							<div className="dropdown-content">
								{
									items.map((item, index) =>
										<div
											className="dropdown-item"
											key={`${item}${index}`}
										>
											<label className="checkbox">
												<input
													type="checkbox"
													value={item}
													onChange={handleDropDownItemClick}
												/>
												{item}
											</label>
										</div>
									)
								}
							</div>
						</div>
					}
				</div>
			</div>
		</div>
	);
}

export default BaseDropdown;