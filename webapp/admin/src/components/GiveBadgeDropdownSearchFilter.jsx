import React, { useState } from "react";

function BaseDropdown({items, onSelect}) {
	const [dropdown, setDropdown] = useState(false);
	const [mouseOnItem, setMouseOnItem] = useState(false);

	function handleDropDownItemClick(e) {
		onSelect(e);
	}

	function handleDropdown() {
		setDropdown(!dropdown);
	}

	function handleMouseOnItem() {
		setMouseOnItem(true);
	}

	function handleMouseNotOnItem() {
		setMouseOnItem(false);
	}

	function handleClickOutside() {
		if (dropdown && !mouseOnItem) {
			setDropdown(false);
		}
	}

	return (
		<div className="gb-dropdown">
			<div className={dropdown ? "dropdown is-active" : "dropdown"}>
				<div>
					<div
						className="dropdown-trigger"
						onBlur={handleClickOutside}
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
					</div>
					<div
						className="dropdown-menu"
						onMouseOver={handleMouseOnItem}
						onMouseLeave={handleMouseNotOnItem}
						onBlur={handleClickOutside}
					>
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
				</div>
			</div>
		</div>
	);
}

export default BaseDropdown;