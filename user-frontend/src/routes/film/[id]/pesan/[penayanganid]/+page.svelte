<script lang="ts">
	import Modal from 'flowbite-svelte/Modal.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import ExclamationCircleOutline from 'flowbite-svelte-icons/ExclamationCircleOutline.svelte';
	import type { PageData } from './$types';
	import { dataFilm, seatPesanan } from '@/lib/stores/data';
	import type { Kursi, Tiket } from '@/lib/types/modelTypes';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { routeApi } from '@/lib/util';
	import type { ApiResponse } from '@/lib/types/apiResponse';

	export let data: PageData;

	let popupModal: boolean = false;
	let messageModal: string = '';

	let harga: number = 0;

	let kursiDiPilih: Kursi[] = [];

	let excludeKursiId: number[] = [];

	$: isKursiDipilhExists = (id: number) => {
		const kursi = kursiDiPilih.find((kursi) => kursi.id === id);
		return kursi !== undefined;
	};

	const findKursi = (id: number): Kursi | undefined => {
		return data.dataKursi.find((val) => val.id === id) as Kursi | undefined;
	};

	const checkStatusKursi = async (kursi_id: number) => {
		const response = await fetch(routeApi(`kursi/${data.penayanganid}/${kursi_id}`));
		return response.ok;
	};

	const pilihKursi = async (id: number) => {
		const status = await checkStatusKursi(id);
		if (!status) {
			popupModal = true;
			messageModal = 'Maaf kursi ini baru saja dipilih orang lain';
			return excludeKursi(id);
		}

		if (isKursiDipilhExists(id)) {
			// remove the kursi from the array
			kursiDiPilih = kursiDiPilih.filter((kursi) => kursi.id !== id);
			removeFromStore(id);
		} else {
			const newKursi = findKursi(id);
			if (!newKursi) return;
			kursiDiPilih = [...kursiDiPilih, newKursi];
			saveToStore(newKursi);
		}
		calculateHarga();
	};

	const calculateHarga = () => {
		let baseHarga =
			$dataFilm?.penayangan.find((item) => item.id === Number(data.penayanganid))?.harga ?? 0;
		console.log('hiii');
		harga = baseHarga * kursiDiPilih.length;
	};

	const saveToStore = (newSeat: Kursi) => {
		$seatPesanan = [...$seatPesanan, newSeat];
	};

	const removeFromStore = (id: number) => {
		$seatPesanan = $seatPesanan.filter((val) => val.id !== id);
	};

	const handleSubmit = async () => {
		if (kursiDiPilih.length === 0) {
			popupModal = true;
			messageModal = 'Harap pilih setidaknya satu kursi';
			return;
		}

		const postData = {
			penayangan_id: parseInt(data.penayanganid),
			kursis: kursiDiPilih,
			total_harga: harga,
			film_id: parseInt(data.filmid),
		};

		const res = await fetch(routeApi('user/pesanan/add'), {
			body: JSON.stringify(postData),
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${data.token}`,
			},
		});

		if (res.ok) {
			const resData: ApiResponse<Tiket> = await res.json();
			return goto('/user/konfirmasi-pesanan/' + resData.data.id);
		} else {
			popupModal = true;
			messageModal = 'Gagal memesan film';
		}
	};

	const excludeKursi = (id: number) => {
		excludeKursiId = [...excludeKursiId, id];
	};

	onMount(() => {
		data.dataKursi.forEach((val) => {
			if (val.tiket?.length > 0) {
				excludeKursi(val.id);
			}
		});
		calculateHarga();
	});

	const beforeUnload = (e: BeforeUnloadEvent) => {
		e.preventDefault();
		e.returnValue = '';
		return 'Data akan hilang jika refresh browser';
	};
</script>

<svelte:window on:beforeunload={beforeUnload} />
<div class="min-h-screen">
	<div class="container">
		<div class="overflow-x-auto">
			<div class="mx-auto flex w-[30rem] flex-col items-center justify-center gap-8 md:w-[45rem]">
				<div
					class="w-[96%] bg-gray-800 py-2 text-center text-2xl font-bold text-white dark:bg-slate-200 dark:text-gray-800"
				>
					Layar
				</div>
				<div
					class="grid w-full auto-cols-fr auto-rows-auto grid-cols-10 grid-rows-10 gap-2 md:gap-4"
				>
					{#each data.dataKursi as kursi (kursi.id)}
						<button
							class="rounded-md py-4 text-[0.675rem] md:text-md text-white"
							class:bg-gray-500={excludeKursiId.includes(kursi.id)}
							class:bg-primary-400={!excludeKursiId.includes(kursi.id) &&
								isKursiDipilhExists(kursi.id)}
							class:bg-primary-700={!excludeKursiId.includes(kursi.id) &&
								!isKursiDipilhExists(kursi.id)}
							class:cursor-not-allowed={excludeKursiId.includes(kursi.id)}
							disabled={excludeKursiId.includes(kursi.id)}
							on:click={() => {
								if (excludeKursiId.includes(kursi.id)) return;
								pilihKursi(kursi.id);
							}}
						>
							{kursi.nama}
						</button>
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
			<Button on:click={handleSubmit} size="lg">Next</Button>
		</div>
	</div>

	<Modal bind:open={popupModal} size="xs" autoclose>
		<div class="text-center">
			<ExclamationCircleOutline class="mx-auto mb-4 h-12 w-12 text-gray-400 dark:text-gray-200" />
			<h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">{messageModal}</h3>
			<Button color="alternative">Tutup</Button>
		</div>
	</Modal>
</div>
