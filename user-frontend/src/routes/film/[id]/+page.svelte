<script lang="ts">
	import type { Penayangan } from '@/lib/types/modelTypes.js';
	import Button from 'flowbite-svelte/Button.svelte';
	import Rating from 'flowbite-svelte/Rating.svelte';
	import Badge from 'flowbite-svelte/Badge.svelte';
	import { routeApi } from '@/lib/util.js';
	import type { PageData } from './$types';
	import type { ApiResponse } from '@/lib/types/apiResponse';
	import { dataFilm, type FilmData } from '@/lib/stores/data';

	export let data: PageData;

	const loadData = async () => {
		const res = await fetch(routeApi(`films/${data.id}`));
		const resdata: ApiResponse<FilmData> = await res.json();
		if (!res.ok) {
			throw new Error(resdata.message);
		}
		const firstData: FilmData = resdata.data;
		dataFilm.set(firstData);

		let jadwal: { date: string; time: { id: number; mulai: string; selesai: string }[] }[] = [];
		if (firstData) {
			firstData.penayangan.forEach((item: Penayangan) => {
				const dateobj: Date = new Date(item.tanggal);
				const localeDate = dateobj.toLocaleString('id-ID', {
					weekday: 'long',
					month: 'long',
					day: 'numeric',
				});

				const extractedItem = {
					id: item.id,
					mulai: new Date(item.mulai).toLocaleTimeString('en-US', {
						hour: '2-digit',
						minute: '2-digit',
						hour12: false,
					}),
					selesai: new Date(item.selesai).toLocaleTimeString('en-US', {
						hour: '2-digit',
						minute: '2-digit',
						hour12: false,
					}),
				};

				if (jadwal.length === 0) {
					return jadwal.push({ date: localeDate, time: [extractedItem] });
				}

				if (!jadwal.find((item) => item.date === localeDate)) {
					jadwal.push({ date: localeDate, time: [extractedItem] });
				} else {
					jadwal.find((item) => item.date === localeDate)?.time.push(extractedItem);
				}
			});
		}
		return { resdata: firstData, jadwal };
	};
</script>

<main class="min-h-screen w-full">
	{#await loadData()}
		<p>Loading...</p>
	{:then data}
		<div class="container flex flex-col gap-4 pt-10 md:flex-row md:gap-8">
			<div>
				<img
					src={routeApi(data.resdata?.poster_path)}
					alt={data.resdata?.title}
					class="aspect-[9.2/13] h-auto w-52 max-w-max rounded-lg object-cover shadow-lg md:w-80 lg:w-96"
				/>
			</div>
			<div class="md:max-w-md lg:max-w-lg">
				<h2 class="text-2xl font-bold md:text-4xl">{data.resdata?.title}</h2>
				<Rating count rating={Number(data.resdata?.rating.toFixed(2))} id="example-4">
					<span class="mx-1.5 h-1 w-1 rounded-full bg-gray-500 dark:bg-gray-400" />
					<p>imdb</p>
				</Rating>
				<h3 class="mt-4 text-lg font-semibold md:text-2xl">Sinopsis</h3>
				<p class="font-light md:font-medium">{data.resdata?.sinopsis}</p>
				{#if data.resdata?.genre}
					<h3 class="mb-2 mt-4 text-lg font-semibold md:text-2xl">Genre</h3>
					<div class="flex flex-wrap items-center justify-start gap-2">
						{#each data.resdata?.genre as genre}
							<Badge border large>{genre.nama}</Badge>
						{/each}
					</div>
				{/if}
			</div>
			<div class="flex-1">
				<h3 class="text-lg font-semibold md:text-2xl">Jadwal Tayang</h3>
				<div class="flex flex-col justify-center">
					{#each data.jadwal as jadwal}
						<div>
							<h3 class="text-md mb-2 mt-4 font-semibold md:text-xl">{jadwal.date}</h3>
							<div class="flex flex-wrap items-center justify-start gap-2">
								{#each jadwal.time as time}
									<Button href="/film/{data.resdata?.id}/pesan/{time.id}">{time.mulai}</Button>
								{/each}
							</div>
						</div>
					{/each}
				</div>
			</div>
		</div>
	{/await}
</main>
