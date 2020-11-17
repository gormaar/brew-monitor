import React from 'react';
import { ResponsiveLine } from "@nivo/line";

type GraphProps = {
    //data: number;
	legend: string;
}

const Graph: React.FC<GraphProps> = (props) => {

	const data = [
		{
			id: "Respiration",
			data: [
				{ x: 1, y: 30 },
				{ x: 2, y: 31 },
				{ x: 3, y: 30 },
				{ x: 4, y: 33 },
				{ x: 5, y: 50 },
				{ x: 6, y: 40 },
				{ x: 7, y: 39 },
				{ x: 8, y: 45 },
				{ x: 9, y: 52 },
			],
		},
	];

		return (
				<ResponsiveLine
						data={data}
						margin={{ top: 50, right: 110, bottom: 50, left: 60 }}
						xScale={{ type: "point" }}
						yScale={{
							type: "linear",
							min: "auto",
							max: "auto",
							stacked: true,
							reverse: false,
						}}
						axisTop={null}
						axisRight={null}
						axisBottom={{
							orient: "bottom",
							tickSize: 1,
							tickPadding: 5,
							tickRotation: 0,
							legend: "Time",
							legendOffset: 36,
							legendPosition: "middle",
						}}
						axisLeft={{
							orient: "left",
							tickSize: 1,
							tickPadding: 5,
							tickRotation: 0,
							legend: props.legend,
							legendOffset: -40,
							legendPosition: "middle",
						}}
						colors={{ scheme: "nivo" }}
						pointSize={10}
						pointColor={{ theme: "background" }}
						pointBorderColor={{ from: "serieColor" }}
						pointBorderWidth={2}
						pointLabel="y"
						pointLabelYOffset={-12}
						useMesh={true}
						legends={[
							{
								anchor: "bottom-right",
								direction: "column",
								justify: false,
								translateX: 100,
								translateY: 0,
								itemsSpacing: 0,
								itemDirection: "left-to-right",
								itemWidth: 80,
								itemHeight: 20,
								itemOpacity: 0.75,
								symbolSize: 12,
								symbolShape: "circle",
								symbolBorderColor: "rgba(0, 0, 0, .5)",
								effects: [
									{
										on: "hover",
										style: {
											itemBackground: "rgba(0, 0, 0, .03)",
											itemOpacity: 1,
										},
									},
								],
							},
						]}
					/>
	)      
}

export default Graph;