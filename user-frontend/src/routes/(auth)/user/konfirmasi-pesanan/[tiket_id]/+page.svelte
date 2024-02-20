<script lang="ts">
	import Radio from 'flowbite-svelte/Radio.svelte';
	import { routeApi } from '@/lib/util';
	import type { PageData } from './$types';
	import Modal from 'flowbite-svelte/Modal.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import Alert from 'flowbite-svelte/Alert.svelte';
	import CheckCircleSolid from 'flowbite-svelte-icons/CheckCircleSolid.svelte';
	import InfoCircleSolid from 'flowbite-svelte-icons/InfoCircleSolid.svelte';
	import { goto } from '$app/navigation';
	import type { ApiResponse } from '@/lib/types/apiResponse';

	export let data: PageData;
	$: penayangan = data.dataFilm.penayangan[0];
	$: seats = penayangan.tiket[0].kursi;

	let popupModal: boolean = false;

	const handleSubmit = async () => {
		const { kursi, ...postTiket } = data.dataFilm.penayangan[0].tiket[0];
		const res = await fetch(routeApi('user/pesanan/update-pembayaran'), {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${data.token}`,
			},
			body: JSON.stringify(postTiket),
		});
		const json: ApiResponse<any> = await res.json();
		if (json.status === 'ok') {
			popupModal = true;
		}
	};

	const getExpTime = () => {
		const d1: Date = new Date(penayangan.tiket[0].created_at);
		const d2: Date = new Date(d1);
		d2.setMinutes(d1.getMinutes() + 30);
		return d2.toLocaleTimeString('en-US', { hour: 'numeric', minute: 'numeric', hour12: false });
	};
</script>

<div class="container">
	<div class="mx-auto flex w-fit flex-col items-center justify-center gap-8">
		<div class="flex flex-col items-center justify-center gap-8 md:flex-row md:items-start">
			<img
				src={routeApi(data.dataFilm.poster_path ?? '')}
				alt="film poster"
				class="aspect-[9.2/13] h-auto w-36 max-w-max rounded-lg object-cover shadow-lg md:w-56"
			/>
			<div>
				<h4 class="mt-4 text-lg font-semibold md:text-2xl">Detail Pesanan</h4>
				<div
					class="md:text-md grid auto-cols-fr grid-cols-2 gap-x-4 gap-y-2 text-sm text-gray-600 dark:text-gray-400"
				>
					<p>Judul</p>
					<p>{data.dataFilm.title}</p>

					<p>Total Harga</p>
					<p>
						{Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR' }).format(
							penayangan.harga * seats.length
						)}
					</p>

					<p>Audiotorium</p>
					<p>{penayangan.audiotorium_id}</p>

					<p>Tanggal</p>
					<p>
						{new Date(penayangan.tanggal).toLocaleDateString('id-ID', {
							weekday: 'long',
							year: 'numeric',
							month: 'long',
							day: 'numeric',
						})}
					</p>

					<p>Waktu</p>
					<p>
						{new Date(penayangan.mulai).toLocaleTimeString('en-US', {
							hour: '2-digit',
							minute: '2-digit',
							hour12: false,
						})}
						~
						{new Date(penayangan.selesai).toLocaleTimeString('en-US', {
							hour: '2-digit',
							minute: '2-digit',
							hour12: false,
						})}
					</p>

					<p>Nomor Kursi</p>
					<p class="max-w-40">{seats.map((val) => val.nama).join(', ')}</p>
				</div>
			</div>
		</div>
		<div class="w-full">
			<h2 class="mt-4 text-lg font-semibold md:text-2xl">Pilih Metode Pembayaran</h2>
			<div class="flex flex-col items-center justify-center gap-4">
				{#each data.listPembayaran as listPembayaran}
					<div
						class="flex w-full items-center justify-between rounded border border-gray-200 p-4 dark:border-gray-700"
					>
						<Radio name="bordered" class="w-full">{listPembayaran.name}</Radio>
						<img src={listPembayaran.src} alt={listPembayaran.name} class="w-16" />
					</div>
				{/each}
			</div>
			<Alert color="red" class="mt-4">
				<InfoCircleSolid slot="icon" class="h-4 w-4" />
				<p>harap segera melakukan pembayaran sebelum pukul {getExpTime()}</p>
				<p>Jika melebihi tengat waktu maka pesanan akan dibatalkan</p>
			</Alert>
			<Button on:click={handleSubmit} class="mt-4">Bayar</Button>
		</div>
	</div>
	<Modal bind:open={popupModal} size="xs" dismissable={false}>
		<div class="text-center">
			<CheckCircleSolid class="mx-auto mb-4 h-20 w-20 text-green-400 dark:text-green-200" />
			<h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">Pembayaran Berhasil</h3>
			<Button color="alternative" on:click={() => goto(`/user/tiket/${penayangan.tiket[0].id}`)}
				>Tutup</Button
			>
		</div>
	</Modal>
</div>
