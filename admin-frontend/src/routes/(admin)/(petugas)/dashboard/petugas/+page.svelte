<script lang="ts">
	import type { ApiResponse } from '@/lib/types/apiResponse';
	import type { Film, Kursi, Penayangan, Tiket } from '@/lib/types/modelTypes';
	import { routeApi } from '@/lib/util';
	import Button from 'flowbite-svelte/Button.svelte';
	import Input from 'flowbite-svelte/Input.svelte';
	import Label from 'flowbite-svelte/Label.svelte';
	import type { PageData } from '../../../$types';
	import ActionModal from '@/components/actionModal.svelte';
	import { goto, invalidateAll } from '$app/navigation';
	import CheckCircleSolid from 'flowbite-svelte-icons/CheckCircleSolid.svelte';

  export let data: PageData
	let tiket: string[];
	let disableButton: boolean = true;

	type PenayanganTiket = Film & {
		penayangan: Penayangan[] &
			{
				tiket: Tiket[] &
					{
						kursi: Kursi[];
					}[];
			}[];
	};

	let dataFilm: PenayanganTiket | null = null;

	$: penayangan = dataFilm?.penayangan[0];
	$: seats = penayangan?.tiket ? penayangan.tiket[0].kursi : null;

	const onInput = (e: Event) => {
		disableButton = true;
		const value = (e.target as HTMLInputElement).value;
		if (!value) {
			return;
		}
		const regexMatch = value.match(/^([0-9]*)\-([0-9]*)$$/);
		if (!regexMatch) {
			return;
		}

		disableButton = false;

		tiket = regexMatch;
	};

	const getTiket = async () => {
		const req = await fetch(
			routeApi('admin/tiket/search?') +
				new URLSearchParams({
					user_id: tiket[1],
					tiket_id: tiket[2],
				}),
      {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${data.token}`,
        }
      }
		);
		const res: ApiResponse<PenayanganTiket> = await req.json();
		if (res.status == 'ok') {
			dataFilm = res.data;
		} else {
      dataFilm = null
    }
	};

  let modalState: boolean = false;
  let modalMessage: string = '';
  let modalSuccess: boolean = false;


  const checkIn = async () => {
    const req = await fetch(
      routeApi(`admin/tiket/checkin/${penayangan?.tiket[0].id}`),
      {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${data.token}`,
        }
      }
    )
    const res: ApiResponse<null> = await req.json();
    modalState = true;
    if (res.status == 'ok') {
      modalSuccess = true;
      modalMessage = 'Check In Berhasil';
    } else {
      modalSuccess = false;
      modalMessage = 'Check In Gagal';
    }
  }


</script>

<div>
	<h1
		class="mb-4 text-2xl font-bold leading-none tracking-tight text-gray-900 dark:text-white md:text-3xl"
	>
		Check In Tiket
	</h1>
	<Label class="text-md mb-2 flex flex-col items-center gap-4 font-medium md:flex-row">
		<Input type="text" placeholder="Tiket ID" on:input={onInput} />
		<Button bind:disabled={disableButton} on:click={getTiket}>Search</Button>
	</Label>
	{#if dataFilm && penayangan && seats}
    <div class="flex flex-col items-center justify-center gap-4 mt-8">
		<div class="flex flex-col items-center justify-center gap-8 md:flex-row md:items-start">
			<img
				src={routeApi(dataFilm.poster_path ?? '')}
				alt="film poster"
				class="aspect-[9.2/13] h-auto w-36 max-w-max rounded-lg object-cover shadow-lg md:w-56"
			/>
			<div>
				<h4 class="mt-4 text-lg font-semibold md:text-2xl">Detail Pesanan</h4>
				<div
					class="text-md grid auto-cols-fr grid-cols-2 gap-x-4 gap-y-2 text-gray-600 dark:text-gray-400"
				>
					<p>Judul</p>
					<p>{dataFilm.title}</p>

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
						{new Date(penayangan.mulai).toLocaleDateString('id-ID', {
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
					<p>{seats.map((val) => val.nama).join(', ')}</p>
				</div>
			</div>
		</div>
      {#if penayangan.tiket[0].signed_by}
      <div class="flex justify-center items-center text-green-400 dark:text-green-200 gap-4">
      <CheckCircleSolid class="mx-auto h-8 w-8 " />
      <h1 class="text-xl font-bold">Tiket Sudah Di Check In</h1>
      </div>
      {:else}
      <Button on:click={checkIn}>Check In</Button>
      {/if}
    </div>
	{/if}
</div>
	<ActionModal
		bind:modalState
		bind:success={modalSuccess}
		bind:message={modalMessage}
		onConfirm={async () => {
			await invalidateAll();
			await getTiket();
		}}
	/>

