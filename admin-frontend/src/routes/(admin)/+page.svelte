<script lang="ts">
	import Chart from 'flowbite-svelte/Chart.svelte';
	import Card from 'flowbite-svelte/Card.svelte';
	import type { PageServerData } from './$types';
	import type { Report } from '@/lib/types/modelTypes';

	export let data: PageServerData;
	$: report = data.data;

	const groupByDate = (data: Report[]): Record<string, number> => {
		const returner: Record<string, number> = {};
		data.forEach((item) => {
			const localDate = new Date(item.created_at).toLocaleDateString('id-ID', {
				month: 'long',
				day: 'numeric',
			});
			if (returner[localDate]) {
				returner[localDate] += item.pendapatan;
			} else {
				returner[localDate] = item.pendapatan;
			}
		});
		return returner;
	};

	const Options = () => {
		const groupedData = groupByDate(report);

		return {
			chart: {
				height: '400px',
				maxWidth: '100%',
				type: 'area',
				fontFamily: 'Inter, sans-serif',
				dropShadow: {
					enabled: false,
				},
				toolbar: {
					show: false,
				},
			},
			tooltip: {
				enabled: true,
				x: {
					show: false,
				},
			},
			fill: {
				type: 'gradient',
				gradient: {
					opacityFrom: 0.55,
					opacityTo: 0,
					shade: '#1C64F2',
					gradientToColors: ['#1C64F2'],
				},
			},
			dataLabels: {
				enabled: false,
			},
			stroke: {
				width: 6,
			},
			grid: {
				show: false,
				strokeDashArray: 4,
				padding: {
					left: 2,
					right: 2,
					top: 0,
				},
			},
			series: [
				{
					name: 'Pendapatan',
					data: Object.values(groupedData),
					color: '#1A56DB',
				},
			],
			xaxis: {
				categories: Object.keys(groupedData),
				labels: {
					show: false,
				},
				axisBorder: {
					show: false,
				},
				axisTicks: {
					show: false,
				},
			},
			yaxis: {
				show: false,
			},
		};
	};
</script>

<Card class="w-full max-w-full md:w-1/2">
	<div class="flex justify-between">
		<div>
			<h5 class="pb-2 text-3xl font-bold leading-none text-gray-900 dark:text-white">
				{Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(
					Object.entries(groupByDate(report)).reduce((a, b) => a + b[1], 0)
				)}
			</h5>
			<p class="text-base font-normal text-gray-500 dark:text-gray-400">Pendapatan Minggu Ini</p>
		</div>
	</div>
	<Chart options={Options()} />
</Card>
