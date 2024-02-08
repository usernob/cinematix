<script lang="ts">
	import { Button } from 'flowbite-svelte';
	import type { PageData } from './$types';
	import { dataFilm } from '@/lib/stores/data';

	export let data: PageData;

	let harga = 0;

	let kursiDiPilih: number[] = [];

	const pilihKursi = (id: number) => {
		if (kursiDiPilih.includes(id)) {
			// remove the kursi from the array
			kursiDiPilih = kursiDiPilih.filter((kursi) => kursi !== id);
		} else {
			kursiDiPilih = [...kursiDiPilih, id];
		}
		calculateHarga();
	};

	const calculateHarga = () => {
		let baseHarga =
			$dataFilm?.penayangan.find((item) => item.id === Number(data.penayanganid))?.harga ?? 0;

		harga = baseHarga * kursiDiPilih.length;
	};
</script>

<div class="min-h-screen">
	<div class="container">
    <div class="overflow-x-auto">
		<div class="flex flex-col items-center justify-center gap-8 w-[30rem] md:w-[45rem] mx-auto">
			<div
				class="w-[96%] bg-gray-800 py-2 text-center text-2xl font-bold text-white dark:bg-slate-200 dark:text-gray-800"
			>
				Layar
			</div>
			<div class="grid w-full auto-cols-fr auto-rows-auto grid-cols-10 grid-rows-10 gap-2 md:gap-4">
				{#each data.dataKursi as kursi (kursi.id)} 
					{#if kursi.seat?.length > 0}
						<button
							class="cursor-not-allowed rounded-md bg-gray-700 py-4 text-white text-sm md:text-md"
							disabled
						>
							{kursi.nama}
						</button>
					{:else}
						<button
							class="rounded-md text-sm md:text-md py-4 text-white"
							class:bg-primary-400={kursiDiPilih.includes(kursi.id)}
							class:bg-primary-700={!kursiDiPilih.includes(kursi.id)}
							on:click={() => {pilihKursi(kursi.id)}}
						>
							{kursi.nama}
						</button>
					{/if}
				{/each}
			</div>
		</div>
   </div>
	</div>
	<div
		class="fixed bottom-0 w-full bg-slate-50 p-4 shadow-[0_35px_60px_15px_rgba(0,0,0,0.3)] dark:bg-gray-800"
	>
		<div class="container flex items-center justify-between">
			<div>
				<h4 class="text-sm font-light md:text-2xl">Total</h4>
				<h3 class="text-lg font-semibold md:text-3xl">
					{Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(harga)}
				</h3>
			</div>
			<Button href="/film/{data.filmid}/konfirmasi/{data.penayanganid}" size="lg">Next</Button>
		</div>
	</div>
</div>
