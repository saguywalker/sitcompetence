import React from "react";

function BaseTable() {
	return (
		<section className="section">
			<table className="table">
				<thead>
					<tr>
						<th>
							<input
								className="base-table-checkbox-all"
								type="checkbox"
							/>
						</th>
						<th>Student</th>
						<th>Gained badge</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>
							<input
								className="base-table-checkbox"
								type="checkbox"
							/>
						</td>
						<td>TODO: Student table data component, including img, name</td>
						<td>TODO: Student gain badge component, including badge pic and badge count per each</td>
					</tr>
					<tr>
						<td>
							<input
								className="base-table-checkbox"
								type="checkbox"
							/>
						</td>
						<td>Hii</td>
						<td>Hello, how are you guy doing ?</td>
					</tr>
				</tbody>
			</table>
		</section>
	);
}

export default BaseTable;