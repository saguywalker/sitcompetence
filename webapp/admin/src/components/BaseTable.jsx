import React, { useState, useEffect } from "react";
import BaseImagePic from "components/BaseImagePic";
import BasePopup from "components/BasePopup";

function BaseTable() {
	const [selectRows, setSelectRow] = useState([]);

	function handleSelectRow(e) {
		const selectItem = parseInt(e.target.value, 10);
		// Find the object of the select data by id
		const selectData = tableData.find((item) => item.id === selectItem);

		if (!e.target.checked) {
			// Remove the value from the array by unchecked checkbox index
			const index = selectRows.findIndex((item) => item.id === selectItem);
			selectRows.splice(index, 1);
			setSelectRow([
				...selectRows
			]);
			return;
		}

		setSelectRow([
			...selectRows,
			selectData
		]);
	}

	useEffect(() => {
		setSelectRow(selectRows);
	}, [selectRows]);

	const [tableData] = useState([
		{
			id: 1,
			firstName: "John",
			lastName: "Wick",
			img: "https://bulma.io/images/placeholders/128x128.png",
			stdId: "59130500191",
			department: "DSI",
			badges: [
				{
					badgeImg: "https://bulma.io/images/placeholders/64x64.png",
					badgeCount: 2
				},
				{
					badgeImg: "https://bulma.io/images/placeholders/64x64.png",
					badgeCount: 3
				},
				{
					badgeImg: "https://bulma.io/images/placeholders/64x64.png",
					badgeCount: 1
				}
			]
		},
		{
			id: 2,
			firstName: "Uzumaki",
			lastName: "Naruto",
			img: "https://bulma.io/images/placeholders/128x128.png",
			stdId: "59130500555",
			department: "IT",
			badges: [
				{
					badgeImg: "https://bulma.io/images/placeholders/64x64.png",
					badgeCount: 2
				},
				{
					badgeImg: "https://bulma.io/images/placeholders/64x64.png",
					badgeCount: 2
				},
				{
					badgeImg: "https://bulma.io/images/placeholders/64x64.png",
					badgeCount: 2
				}
			]
		},
		{
			id: 0,
			firstName: "Tindanai",
			lastName: "Wongpipattanopas",
			img: "https://bulma.io/images/placeholders/128x128.png",
			stdId: "59130500210",
			department: "CS",
			badges: [
				{
					badgeImg: "https://bulma.io/images/placeholders/64x64.png",
					badgeCount: 1
				},
				{
					badgeImg: "https://bulma.io/images/placeholders/64x64.png",
					badgeCount: 2
				},
				{
					badgeImg: "https://bulma.io/images/placeholders/64x64.png",
					badgeCount: 4
				}
			]
		}
	]);


	return (
		<>
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
							<th>
								Student
							</th>
							<th>
								Gained Competence Badge
							</th>
						</tr>
					</thead>
					<tbody>
						{
							tableData.map((data, index) => (
								<tr key={`${data.id}${index}`}>
									<td>
										<input
											className="base-table-checkbox"
											type="checkbox"
											value={data.id}
											onChange={handleSelectRow}
										/>
									</td>
									<td>
										<div className="base-table-student">
											<BaseImagePic
												size="64"
												round={true}
												imgSrc={data.img}
											/>
											<div className="base-table-student-data">
												<h1 className="title is-6">
													{data.firstName} {data.lastName}
												</h1>
												<p className="subtitle is-7">
													<strong>{data.department}</strong>  {data.stdId}
												</p>
											</div>
										</div>
									</td>
									<td>
										<div className="base-table-badge">
											{
												data.badges.map((item, index) => (
													<div
														className="base-table-badge-item"
														key={`${item}${index}`}
													>
														<BaseImagePic
															size="32"
															round={true}
															imgSrc={item.badgeImg}
														/>
														<p className="base-table-badge-item-multiple">
															x {item.badgeCount}
														</p>
													</div>
												))
											}
										</div>
									</td>
								</tr>
							))
						}
					</tbody>
				</table>
			</section>
			<BasePopup items={selectRows}/>
		</>
	);
}

export default BaseTable;